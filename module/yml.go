package module

type AppConfig struct {
	Db
	TemplateData
}

type Db struct {
	Type string
	Url  string
}

type TemplateData struct {
	StructList       []TableInfo //结构体类型列表
	TemplatePath     string      //模板文件路径
	TemplateFileName string      //模板文件名称
	FileSavePath     string      //生成文件路径
	FileSaveName     string      //生成文件名称
	TableNameList    []string    //表名列表
	TableName        string      //当前循环表名
	DbName           string      //数据库名
}

type TableInfo struct {
	ColumnName    string `gorm:"column:COLUMN_NAME"`
	DataType      string `gorm:"column:DATA_TYPE"`
	ColumnComment string `gorm:"column:COLUMN_COMMENT"`
}
