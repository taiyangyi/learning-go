package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	// 配置数据库连接信息
	dsn := "root:123321@tcp(127.0.0.1:3306)/learn_mysql_course"

	// 初始化数据库连接
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// 配置连接池
	db.SetMaxOpenConns(25)        // 最大打开连接数
	db.SetMaxIdleConns(25)        // 最大空闲连接数 0 表示无限制
	db.SetConnMaxLifetime(5 * 60) // 连接最大生存时间

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("数据库连接成功！")

}
