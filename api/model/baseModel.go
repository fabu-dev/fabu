package model

import (
	"fabu.dev/api/pkg/db"
	"github.com/jinzhu/gorm"
)

type IModel interface {
	GetTableName() string
	Find() *gorm.DB
}

type BaseModel struct {
	TableName string
}

func (m *BaseModel) GetTableName() string {
	return ""
}

func (m *BaseModel) SetTableName(tableName string) {
	m.TableName = tableName
}

func (m *BaseModel) Find() *gorm.DB {
	return db.Mysql
}
