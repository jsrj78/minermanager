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
	"mmserver/utils"

	"mmserver/models"

	pag "github.com/mmuflih/gorm-paginator"
)

type MinerService struct {
	BaseService
}

// StartMiner
func (ms *MinerService) LoadMinerAndStart() error {
	miners := []*models.Miner{}
	//1. 获取所有矿机
	if err := utils.DB.Where("status = ? or status = ? ", models.MinerOffline, models.MinerOnline).Find(&miners).Error; err == nil {
		f := models.NewMinerFactory()
		f.CreateMiners(miners)
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
