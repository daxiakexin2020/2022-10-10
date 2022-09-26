package model

import "gorm.io/gorm"

var TestDB *gorm.DB

func SetDB(db *gorm.DB) {
	TestDB = db
}
func Get() *gorm.DB {
	return TestDB
}
