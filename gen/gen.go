package gen

import (
	"generator/module"
	"generator/utils"
	"log"
	"os"
	"strings"
	"text/template"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var Test = utils.Config.Db
var TemplateData = utils.Config.TemplateData

func Run() {
	db, err := gorm.Open(Test.Type, Test.Url)
	if err != nil {
		log.Fatal("数据库连接错误")
	}
	db.SingularTable(true)

	var table_infos []module.TableInfo
	sql := "select COLUMN_NAME,DATA_TYPE,COLUMN_COMMENT from information_schema.COLUMNS where TABLE_SCHEMA = ? and TABLE_NAME = ?"
	if err = db.Raw(sql, TemplateData.DbName, TemplateData.TableName).Find(&table_infos).Error; err != nil {
		log.Fatal("查询错误")
	}

	//向模板数据中添加表字段数据
	TemplateData.StructList = table_infos
	//生成
	ByTemplateGeneratorFile(TemplateData)
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

//格式化数据类型
func FormatType(types string) string {
	switch types {
	case "boolean", "bool":
		return "bool"
	case "tinyint":
		return "int8"
	case "smallint", "year":
		return "int16"
	case "integer", "mediumint", "int":
		return "int32"
	case "bigint":
		return "int64"
	case "date", "timestamp without time zone", "timestamp with ime zone", "timestamp", "datatime", "time":
		return "time.Time"
	case "bytea", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "[]byte"
	case "varchar", "text", "json", "longtext":
		return "string"
	default:
		log.Println("未识别类型:", types)
		return types
	}
}

//按模板生成文件
func ByTemplateGeneratorFile(data module.TemplateData) bool {
	tmpl, err := template.New(data.TemplateFileName).
		Funcs(template.FuncMap{"Case2Camel": Case2Camel, "FormatType": FormatType}).
		ParseFiles(data.TemplatePath + data.TemplateFileName)
	if err != nil {
		log.Println("读取模板文件错误")
		return false
	}
	file, err := os.OpenFile(data.FileSavePath+data.FileSaveName, os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		log.Println("打开文件错误")
		return false
	}
	err = tmpl.Execute(file, data)
	if err != nil {
		log.Println("写入文件错误")
		return false
	}
	return true
}
