package models

import (
	"go-training/go-http/conf"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	*gorm.DB
}

func NewDB() (*DB, error) {
	db, err := gorm.Open(conf.Database, conf.DbName)
	if err != nil {
		return nil, err
	}
	return &DB{db}, nil
}

func (db *DB) InitSchema() {
	db.AutoMigrate(&User{})
}
