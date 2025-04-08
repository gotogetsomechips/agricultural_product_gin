package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// 数据库配置
const (
	DBUsername = "root"
	DBPassword = "123456"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "traceability"
)

// GetDB 获取数据库连接
func GetDB() *sql.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		DBUsername, DBPassword, DBHost, DBPort, DBName)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("连接数据库失败:", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("数据库Ping失败:", err)
	}

	return db
}
