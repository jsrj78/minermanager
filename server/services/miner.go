/*
 * @Author: chunhua yang
 * @Date: 2023-10-27 20:58:34
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 00:14:11
 * @FilePath: /minermanager/server/services/miner.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package services

import (
	"fmt"
	"mmserver/utils"
	"strings"
	"time"

	"mmserver/models"

	pag "github.com/mmuflih/gorm-paginator"
)

type MinerService struct {
	BaseService
}

// StartMiner
func (ms *MinerService) LoadMinerAndStart() error {
	dbminer := []models.TMiner{}
	//1. 获取在线和离线矿机
	if err := utils.DB.Where("status = ? or status = ? ", models.MinerOffline, models.MinerOnline).Find(&dbminer).Error; err == nil {

		miners := []*models.Miner{}

		for _, miner := range dbminer {
			m := models.Miner{
				Id:              miner.Id,
				Ip:              miner.Ip,
				Brand:           miner.Brand,
				UserName:        miner.Username,
				Password:        miner.Password,
				DefaultUserName: miner.DefaultUsername,
				DefaultPassword: miner.DefaultPassword,
				Status:          miner.Status,
			}
			miners = append(miners, &m)
		}

		f := models.NewMinerFactory()
		f.CreateMiners(miners)
		go func() {
			for {
				// 存储第3位相等的IP地址的和
				ipSum := make(map[string]int)

				//fmt.Println(ipSum)

				fmt.Println(f.MinerStatusManager.OfflineList)
				// 遍历 OfflineList
				for _, ip := range f.MinerStatusManager.OfflineList {
					// 使用 strings.Split 函数拆分IP地址
					parts := strings.Split(ip, ".")

					// 检查是否有足够的位数
					if len(parts) >= 3 {
						// 获取第3位并将其添加到和中
						thirdPart := parts[2]
						ipSum[thirdPart] += 1
					}
				}

				fmt.Println(ipSum)

				// 遍历和，检查是否大于5，如果是则发送警告
				for thirdPart, sum := range ipSum {
					if sum > 12 {
						fmt.Printf("警告：第3位为 %s 的IP地址总和超过10：%d\n", thirdPart, sum)
						// 在这里添加发送警告的代码
					}
				}

				// 1分钟循环一篇
				time.Sleep(time.Second * 60)
			}

		}()

	} else {
		utils.Log.Error(err)
		return err
	}

	// //读取online和offline的矿机
	// go func(msm *models.MinerStatusManager) {
	// 	for {

	// 		msm.Mux.Lock()
	// 		for id, value := range msm.OnlineList {
	// 			fmt.Println("start to read online and offline miner")
	// 			fmt.Printf("ID: %s, Value: %s\n", id, value)
	// 		}

	// 		for id, value := range msm.OfflineList {
	// 			fmt.Printf("ID: %s, Value: %s\n", id, value)
	// 		}

	// 		msm.Mux.Unlock() // 解锁访问

	// 		time.Sleep(time.Second * 3) // 可能需
	// 	}

	// }(f.MinerStatusManager)

	return nil

}

/**************************以下是数据库操作部分*************************/

// CreateMiner
func (ms *MinerService) CreateMiner(miner models.TMiner) (*models.TMiner, error) {

	if err := utils.DB.Create(&miner).Error; err != nil {

		return &miner, err
	}

	return &miner, nil
}

// UpdateMiner
func (ms *MinerService) UpdateMiner(miner models.TMiner) (*models.TMiner, error) {

	if err := utils.DB.Model(&miner).Updates(&miner).Error; err != nil {

		return &miner, err
	}

	return &miner, nil
}

// GetAllMinersPaginator
func (ms *MinerService) GetAllMinersPaginator(page int, pagesize int, order interface{}, query interface{}, args ...interface{}) *pag.Paginator {

	miner := []models.TMiner{}

	result := pag.Make(&pag.Config{
		DB:   utils.DB.Where(query, args...).Find(&miner).Order(order),
		Page: page,
		Size: pagesize,
	}, &miner)

	return result
}
