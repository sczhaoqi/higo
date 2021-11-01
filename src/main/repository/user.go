package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"higo/src/main/entity"
	"net/http"
	"strconv"
)

// DB 定义全局变量存放 DB 结构体
var db *sqlx.DB

//SetHandler 该函数将新增或者修改用户信息
func SetHandler(w http.ResponseWriter, r *http.Request) {
	// 以下三个变量为从浏览器获得的数据
	name := r.FormValue("name") // r.FormValue 接收来自浏览器的信息,返回的是 string
	age := r.FormValue("age")
	gender := r.FormValue("gender")

	if name == "" {
		w.Write([]byte("name 不能为空"))
		return
	}

	// 将浏览器输入的字符串转换为整型
	ageInt, err := strconv.Atoi(age)
	if err != nil {
		w.Write([]byte("age 转换失败"))
		return
	}

	genderInt, err := strconv.Atoi(gender)
	if err != nil {
		w.Write([]byte("gender 转换失败"))
		return
	}

	// 插入用户信息
	insertSQL := "INSERT INTO test_user(name, age, gender) values(?, ?, ?)"
	// 执行 sql 语句
	_, err = db.Exec(insertSQL, name, ageInt, genderInt)
	if err != nil {
		w.Write([]byte("插入数据库失败"))
		return
	}
	w.Write([]byte("插入数据库成功"))
}

//ListHandler 用于显示浏览器的输出
func ListHandler() ([]entity.User, error) {
	initDB()
	var users []entity.User
	// 从数据库中查询用户信息
	querySQL := "SELECT * FROM test_user"
	// 执行查询语句
	err := db.Select(&users, querySQL)
	if err != nil {
		return nil, err
	}
	return users, nil
}

//DelHandler 删除用户信息
func DelHandler(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("name")
	if name == "" {
		w.Write([]byte("name 不能为空"))
		return
	}
	// 删除 SQL 语句
	deleteSQL := "DELETE FROM test_user WHERE name=?"
	// 执行 SQL 语句
	_, err := db.Exec(deleteSQL, name)
	if err != nil {
		w.Write([]byte("数据库删除失败"))
		return
	}
	w.Write([]byte("数据库删除成功"))
}

func initDB() error {
	var err error
	dsn := "root:root@tcp(192.168.49.2:30336)/test"
	// 使用 mysql 驱动连接数据库
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败：", err)
	}

	return nil
}
