package db

import (
	"chip_database/conf"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"log"
)

type client struct {
	db *gorm.DB
}

func NewClient(conf *conf.SqliteDatabaseConfig) (*client, error) {
	path := conf.Path + conf.DB
	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}})
	if err != nil {
		log.Fatalf("connect sqlite err:%s", err)
	}
	return &client{db: db}, nil
}
