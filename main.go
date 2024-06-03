package main

import (
	"CallTheRoll/api/router"
	"CallTheRoll/api/services"
	"CallTheRoll/config"
	"CallTheRoll/logger"

	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func main() {
	// 加载配置
	config.LoadConfig()

	// 初始化日志
	if err := logger.LoggerInit(); err != nil {
		log.Fatalf("初始化日志失败: %v", err)
	}

	// 初始化数据库
	initDB()

	// 清空数据库
	if err := services.ClearDatabase(); err != nil {
		log.Fatalf("清空数据库失败: %v", err)
	}

	// 初始化路由
	r := router.InitRouter()

	// 启动服务器
	addr := config.G().Server.Addr()
	if err := r.Run(addr); err != nil {
		log.Fatalf("启动服务失败: %v", err)
	}
}

func initDB() {
	db, err := sql.Open("sqlite3", "./students.db")
	if err != nil {
		log.Fatalf("打开数据库失败: %v", err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE IF NOT EXISTS students (
		name TEXT,
		number TEXT,
		status TEXT
	);
	`
	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Fatalf("建表失败: %v", err)
	}
}
