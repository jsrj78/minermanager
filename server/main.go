/*
 * @Author: chunhua yang
 * @Date: 2023-10-21 22:05:41
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 00:27:49
 * @FilePath: /minermanager/Server/main.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */

package main

import (
	"fmt"
	"log"
	"mmserver/routes"
	"mmserver/services"
	"mmserver/utils"
	"runtime"
	"time"

	"github.com/robfig/cron/v3"
)

func main() {
	utils.Log.OutputToConsole(true)

	//加载配置文件
	utils.InitConfig()

	//连接到数据库
	utils.DBConnect()

	//连接到redis

	//启动矿机
	ms := services.MinerService{}
	ms.LoadMinerAndStart()

	//启动定时任务
	go func() {
		// 从数据库中获取定时任务
		jobs := []utils.CronJob{
			{
				Expression: "0 0 0 * * *",
				Task:       nil,
			},
			{
				Expression: "0 0 3 * * *",
				Task:       nil,
			},
		}

		c := cron.New(cron.WithSeconds())

		// 添加每个任务到Cron调度器中
		for _, job := range jobs {
			c.AddFunc(job.Expression, job.Task)
		}

		c.Start() // 启动Cron调度器
	}()

	//内存监控
	go func() {
		ticker := time.NewTicker(time.Second * 5)
		defer ticker.Stop()

		for range ticker.C {
			var m runtime.MemStats
			runtime.ReadMemStats(&m)
			//fmt.Printf("Alloc = %v MiB", utils.BToMb(m.Alloc))
		}
	}()

	//启动http服务
	r := routes.InitRouter()

	err := r.Run(utils.AppConfig.Api.Host + ":" + fmt.Sprintf("%d", utils.AppConfig.Api.Port))
	if err != nil {
		log.Fatal(err)
	}
}
