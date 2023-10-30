/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 20:53:54
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:14:08
 * @FilePath: /minermanager/server/utils/config.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package utils

import (
	"encoding/json"
	"log"
	"os"
	"sync"

	"github.com/fsnotify/fsnotify"
)

const configFilename = "config.json"

type ApiServer struct {
	Host    string `json:"host"`
	Port    int    `json:"port"`
	Version string `json:"version"`
	Path    string `json:"path"`
}

type RedisServer struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Key  string `json:"key"`
}

type PostgresServer struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type Environment string

const (
	Production  Environment = "production"
	Testing     Environment = "testing"
	Development Environment = "development"
)

type Config struct {
	Mode    Environment `json:"mode"`
	IsDebug bool        `json:"isDebug"`
	// 服务器配置
	Api ApiServer `json:"apiServer"`
	// 数据库配置
	Postgres PostgresServer `json:"dbServer"`
	// Redis配置
	Redis RedisServer `json:"redisServer"`
}

var (
	AppConfig Config
	configMtx sync.RWMutex

	AppsConfigChanged = &PubSub{}
)

// 初始化配置文件(默认在项目的根目录下)
func InitConfig() error {
	err := loadConfig()
	if err != nil {
		return err
	}

	// 监听配置文件变化
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer watcher.Close()

	err = watcher.Add(configFilename)
	if err != nil {
		log.Fatal(err)
		return err
	}

	go watchConfig(watcher)

	return nil
}

// 读取配置文件
func loadConfig() error {

	// 获取当前目录
	currenDir, err := os.Getwd()
	if err != nil {
		//log.Fatalf("Failed to get current directory: %s", err)
		return err
	}

	data, err := os.ReadFile(currenDir + "/config.json")
	if err != nil {
		//log.Fatalf("Error reading config file: %v", err)
		return err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		//log.Fatalf("Error unmarshalling config data: %v", err)
		return err
	}

	configMtx.Lock()
	AppConfig = config
	configMtx.Unlock()

	return nil
}

// 监听配置文件变化
func watchConfig(watcher *fsnotify.Watcher) {
	for {
		func() {
			defer func() {
				if r := recover(); r != nil {
					log.Printf("Recovered from panic in watchConfig: %v", r)
					log.Println("Restarting watchConfig goroutine...")
				}
			}()

			for {
				select {
				case event, ok := <-watcher.Events:
					if !ok {
						return // 重新启动 goroutine
					}
					if event.Op&fsnotify.Write == fsnotify.Write {
						if err := loadConfig(); err != nil {
							log.Printf("Error reloading config: %v", err)
							return // 如果遇到错误，重新启动 goroutine
						}
						AppsConfigChanged.Publish()
					}
				case err, ok := <-watcher.Errors:
					if !ok {
						return // 重新启动 goroutine
					}
					log.Printf("Watcher error: %v", err)
				}
			}
		}()
	}
}
