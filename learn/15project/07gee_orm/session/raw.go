package session

import (
	"database/sql"
	"geeorm/clause"
	"geeorm/dialect"
	"geeorm/log"
	"geeorm/schema"
	"strings"
)

/*
db *sql.DB，使用 sql.Open() 方法连接数据库成功之后返回的指针。
sql	、sqlValues 用来拼接 SQL 语句和 SQL 语句中占位符的对应值,用户调用 Raw() 方法即可改变这两个变量的值。
*/
type Session struct {
	db       *sql.DB
	dialect  dialect.Dialect //实现了一些特定的 SQL 语句的转换，比如查询表是否存在，每个驱动的写法不一样
	tx       *sql.Tx         //执行事务
	refTable *schema.Schema  //任意对象转化成数据库对象
	clause   clause.Clause   //格式化sql语句，增删改查
	sql      strings.Builder //真正执行的sql 盒子
	sqlVars  []interface{}   //sql对应的占位参数
}

type CommonDB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

var _ CommonDB = (*sql.DB)(nil)
var _ CommonDB = (*sql.Tx)(nil)

func New(db *sql.DB, dialect dialect.Dialect) *Session {
	return &Session{
		db:      db,
		dialect: dialect,
		sqlVars: make([]interface{}, 0),
	}
}

// 清空此builder、参数
func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() CommonDB {
	if s.tx != nil {
		return s.tx
	}
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sql.WriteString(" ")
	s.sqlVars = append(s.sqlVars, values...)
	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if result, err = s.DB().Exec(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	return s.db.QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	log.Info(s.sql.String(), s.sqlVars)
	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); err != nil {
		log.Error(err)
	}
	return
}
