package testdb

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	id       int
	username string
	password string
}

var sqldata map[interface{}]interface{}
var db *sql.DB

func TestMysql() {
	initDB()
	testInsert()
	testQuery()
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
func initDB() {
	var err error
	db, err = sql.Open("mysql", "ppseaer:ppseaer@ppsea.com@/test?charset=utf8")
	checkErr(err)
}
func testInsert() {
	fmt.Println("testInsert-------------------------------------------")
	//插入数据
	stmt, err := db.Prepare("INSERT user SET username=?,password=?")
	checkErr(err)
	res, err := stmt.Exec("xiaowei", "xiaowei")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)
}
func testQuery() {
	fmt.Println("testQuery-------------------------------------------")
	var u User
	//查询数据
	rows, err := db.Query("SELECT * FROM user")
	checkErr(err)

	fmt.Println(rows.Columns())
	userinfo := make(map[interface{}]interface{})
	for rows.Next() {
		err := rows.Scan(&u.id, &u.username, &u.password)
		checkErr(err)
		userinfo[u.id] = u
	}
	fmt.Println(userinfo)
}
