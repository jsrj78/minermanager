/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 23:57:21
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:24:54
 * @FilePath: /minermanager/server/utils/db.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package utils

import (
	"fmt"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

const maxRetries = 3

var (
	DB *gorm.DB

	dbMutex sync.RWMutex
)

func init() {
	// 订阅配置文件变化事件
	AppsConfigChanged.Subscribe(DBConnect)
}

func DBConnect() {
	dbMutex.Lock()
	defer dbMutex.Unlock()

	dsn := generateDSN(AppConfig.Postgres)
	fmt.Println(dsn)

	var err error

	// Configure a database connection pool
	dbConfig := &gorm.Config{
		Logger: logger.New(
			Log, // 使用封装好的日志库
			logger.Config{
				SlowThreshold: time.Second,
				LogLevel:      logger.Info,
				Colorful:      false,
			},
		),
		NamingStrategy: schema.NamingStrategy{
			//TablePrefix:   "t_",
			SingularTable: true,
			NoLowerCase:   false,
		},
	}

	// 关闭旧的数据库连接
	if DB != nil {
		sqlDB, _ := DB.DB()
		sqlDB.Close()
	}

	// 尝试连接到数据库，如果失败则重试
	for i := 0; i < maxRetries; i++ {
		DB, err = gorm.Open(postgres.Open(dsn), dbConfig)
		if err == nil {
			break
		}
		Log.Warn("Failed to connect to  database, retrying: ", err)
		time.Sleep(3 * time.Second) // 等待3秒再重试
	}
	if err != nil {
		Log.Error("Failed to connect to  database after retries: ", err)
		return
	}

	configureConnectionPool(DB)
}

func configureConnectionPool(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		Log.Error("Failed to get SQL database instance: ", err)
		return
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)
}

func generateDSN(config PostgresServer) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		config.Host, config.User, config.Password, config.DbName, config.Port)
}
