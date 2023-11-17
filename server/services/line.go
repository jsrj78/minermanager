package services

import (
	"fmt"
	"mmserver/models"
	"mmserver/utils"
	"strconv"

	"github.com/google/uuid"
)

type LineService struct {
	BaseService
}

// 按照规则添加矿机货架
func (this *LineService) AddShelfByRule(line models.TLine, minerbrand models.TBrand) error {

	// 货架编号，货架名称，货架是否需要lift，每个货架中有几列box，每个box有多少层，每层有多少个格子，ip从哪里开始，机器品牌，每个盒子有多少排
	tx := utils.DB.Begin()

	ipStart := line.IpStart - 1

	lineID := uuid.New()
	line.Id = lineID

	//1.loop box
	for k := 0; k < line.Boxes; k++ {

		totalMiners := 0

		boxID := uuid.New()
		box := models.TBoxes{}
		box.Id = boxID
		box.LineId = lineID

		if line.RuleBoxColumns > 1 {
			incColBox := 0
			for col := 1; col <= line.RuleBoxColumns; col++ {
				incBox:=0
				for i := 1; i <= line.RuleFloors; i++ {
					for j := 1; j <= line.RuleMiners; j++ {
						miner := models.TMiner{}
						miner.Id = uuid.New()
						miner.BoxId = boxID
						ip := line.Ip + strconv.Itoa(k+line.BoxBegin) + "." + strconv.Itoa(ipStart+incBox+j+incColBox)
						fmt.Println(ip)
						miner.Ip = ip
						miner.Id = uuid.New()
						miner.Status = models.MinerOnline
						miner.Password = line.RulePassword
						miner.Username = line.RuleUserName
						miner.DefaultPassword = minerbrand.Password
						miner.DefaultUsername = minerbrand.UserName
						miner.Brand = models.AntMinerBrand
						miner.ChildBox = col
						miner.Cols = j
						miner.Rows = i

						totalMiners++

						//create miner
						err := tx.Create(&miner).Error
						if err != nil {
							tx.Rollback()
							return err
						}
					}

					incBox = i * line.RuleMiners

				}
				incColBox = col * line.RuleMiners * line.RuleFloors
			}

		} else {

			incBox := 0
			for i := 1; i <= line.RuleFloors; i++ {

				for j := 1; j <= line.RuleMiners; j++ {
					miner := models.TMiner{}
					miner.Id = uuid.New()
					miner.BoxId = boxID
					ip := line.Ip + strconv.Itoa(k+line.BoxBegin) + "." + strconv.Itoa(ipStart+incBox+j)
					
					miner.Ip = ip
					miner.Id = uuid.New()
					miner.Status = models.MinerOnline
					miner.Password = line.RulePassword
					miner.Username = line.RuleUserName
					miner.DefaultPassword = minerbrand.Password
					miner.DefaultUsername = minerbrand.UserName
					miner.Brand = models.AntMinerBrand
					miner.ChildBox = 1
					miner.Cols = j
					miner.Rows = i

					totalMiners++

					//create miner
					err := tx.Create(&miner).Error
					if err != nil {
						tx.Rollback()
						return err
					}
				}

				incBox = i * line.RuleMiners

			}
		}

		//create box
		box.Islift = line.IsLift
		box.TotalMiners = totalMiners
		box.TotalPlace = totalMiners
		err := tx.Create(&box).Error
		if err != nil {
			tx.Rollback()
			return err
		}

	}

	//create line
	err := tx.Create(&line).Error
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	return nil
}

func (this *LineService) CreateLine(line models.TLine) (*models.TLine, error) {

	if err := utils.DB.Create(&line).Error; err != nil {

		return nil, err
	}
	return &line, nil

}

func (this *LineService) UpdateLine(line models.TLine) (*models.TLine, error) {

	if err := utils.DB.Save(&line).Error; err != nil {

		return nil, err
	}
	return &line, nil
}

func (this *LineService) DeleteLine(line models.TLine) error {

	if err := utils.DB.Delete(&line).Error; err != nil {

		return err
	}
	return nil
}

func (this *LineService) CreateBox(box models.TBoxes) (*models.TBoxes, error) {

	if err := utils.DB.Create(&box).Error; err != nil {

		return nil, err
	}
	return &box, nil

}

func (this *LineService) UpdateBox(box models.TBoxes) (*models.TBoxes, error) {

	if err := utils.DB.Save(&box).Error; err != nil {

		return nil, err
	}
	return &box, nil
}

func (this *LineService) DeleteBox(box models.TBoxes) error {

	if err := utils.DB.Delete(&box).Error; err != nil {

		return err
	}
	return nil
}

func (this *LineService) CreateMiner(miner models.TMiner) (*models.TMiner, error) {

	if err := utils.DB.Create(&miner).Error; err != nil {

		return nil, err
	}
	return &miner, nil

}

func (this *LineService) UpdateMiner(miner models.TMiner) (*models.TMiner, error) {

	if err := utils.DB.Save(&miner).Error; err != nil {

		return nil, err
	}
	return &miner, nil
}

func (this *LineService) DeleteMiner(miner models.TMiner) error {

	if err := utils.DB.Delete(&miner).Error; err != nil {

		return err
	}
	return nil
}
