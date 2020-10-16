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
	return m.TableName
}

func (m *BaseModel) SetTableName(tableName string) {
	m.TableName = tableName
}

func (m *BaseModel) Db() *gorm.DB {
	return db.Mysql.Table(m.TableName)
}
