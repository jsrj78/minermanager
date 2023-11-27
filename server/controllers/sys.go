package controllers

import (
	"fmt"
	"mmserver/models"
	"mmserver/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var lineService = services.LineService{}

type SysController struct {
	BaseController
}

func (this *SysController) Test(c *gin.Context) {
	fmt.Println("test")
	miner := models.TMiner{}
	miner.Ip = "10.5.5.1"
	miner.Brand = models.AntMinerBrand
	miner.Id = uuid.New()
	_, err := lineService.CreateMiner(miner)
	if err != nil {
		this.ResponseError(c, err.Error())
		return
	}

	this.ResponseSuccess(c, nil)
}

// 按照规则添加矿机货架
func (this *SysController) AddShelfByRule(c *gin.Context) {
	shelf := models.TLine{}

	if err := c.ShouldBindJSON(&shelf); err != nil {
		this.ResponseError(c, err.Error())
		return
	}

	// 货架编号，货架名称，货架是否需要lift，每个货架中有多少个box，每个box有多少层，每层有多少个格子，ip从哪里开始，机器品牌，每个盒子有多少排
	err := lineService.AddShelfByRule(shelf, models.TBrand{})
	if err != nil {
		this.ResponseError(c, err.Error())
		return
	}

	this.ResponseSuccess(c, nil)
}
