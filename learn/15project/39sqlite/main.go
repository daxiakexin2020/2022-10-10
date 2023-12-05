package main

import (
	"database/sql"

	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	fmt.Println("卧槽")
	return

	db, err := sql.Open("sqlite3", "./foo.db")

	checkErr(err)

	//插入数据

	stmt, err := db.Prepare("INSERT INTO userinfo(username, departname, created) values(?,?,?)")

	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")

	checkErr(err)

	id, err := res.LastInsertId()

	checkErr(err)

	fmt.Println(id)

	//更新数据

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")

	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)

	checkErr(err)

	affect, err := res.RowsAffected()

	checkErr(err)

	fmt.Println(affect)

	//查询数据

	rows, err := db.Query("SELECT * FROM userinfo")
	fmt.Println("data:", rows)
	checkErr(err)

	//删除数据

	stmt, err = db.Prepare("delete from userinfo where uid=?")

	checkErr(err)

	res, err = stmt.Exec(id)

	checkErr(err)

	affect, err = res.RowsAffected()

	checkErr(err)

	fmt.Println(affect)

	db.Close()

}

func checkErr(err error) {

	if err != nil {

		panic(err)

	}

}
