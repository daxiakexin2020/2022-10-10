package conf

import (
	"log"
	"sync"
)

var (
	sqliteDatabaseConfig *SqliteDatabaseConfig
	sqliteDatabaseOnce   sync.Once
)

type SqliteDatabaseConfig struct {
	DB   string `json:"db"`
	Path string `json:"path"`
}

func (c *SqliteDatabaseConfig) CName() string {
	return "sqlite_database"
}

func makeSqliteDatabaseConfig() {
	sqliteDatabaseOnce.Do(func() {
		c := &SqliteDatabaseConfig{}
		if err := generate(c.CName(), c); err != nil {
			log.Fatalf("read %s config err%v\n", c.CName(), err)
		}
		sqliteDatabaseConfig = c
		log.Printf("read %s config ok::::::::::::::::::::::::%+v\n", c.CName(), c)
	})
}

func GetSqliteDatabaseConfig() *SqliteDatabaseConfig {
	if sqliteDatabaseConfig == nil {
		makeSqliteDatabaseConfig()
	}
	return sqliteDatabaseConfig
}
