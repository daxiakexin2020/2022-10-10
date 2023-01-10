package session

import (
	"database/sql"
	"geeorm/dialect"
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Name string `geeorm:"PRIMARY KEY"`
	Age  int
}

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func NewSession() *Session {
	TestDB, _ = sql.Open("sqlite3", "../gee.db")
	return New(TestDB, TestDial)
}
