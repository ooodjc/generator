package service

import (
	"time"
)

type {{ .TableName | Case2Camel }} struct {
	{{range .StructList}}
	{{ .ColumnName | Case2Camel }}	{{ .DataType | FormatType }}	`gorm:"column:{{.ColumnName}}"`
	{{ end }}
}

type {{ .TableName | Case2Camel }}Service struct {}

//查询所有
func ({{ .TableName | Case2Camel }}Service) QueryAll() (item []{{ .TableName | Case2Camel }}, err error) {
	if err = {{ .DbName | Case2Camel }}.Find(&item).Error; err != nil {
		return
	}
	return
}

//新增
func ({{ .TableName | Case2Camel }}Service) Inseter(info {{ .TableName | Case2Camel }}) bool {
	if err := {{ .DbName | Case2Camel }}.Create(info).Error; err != nil {
		return false
	}
	return true
}

//按id更新
func ({{ .TableName | Case2Camel }}Service) UpdateById(info map[string]interface{}) bool {
	if err := {{ .DbName | Case2Camel }}.Model({{ .TableName | Case2Camel }}{}).Where("id = ?", info["Id"]).Updates(info).Error; err != nil {
		return false
	}
	return true
}

//按id删除
func ({{ .TableName | Case2Camel }}Service) DeleteById(id int) bool {
	if err := {{ .DbName | Case2Camel }}.Where("id = ?", id).Delete({{ .TableName | Case2Camel }}{}).Error; err != nil {
		return false
	}
	return true
}