package controllers

import (
	"mmserver/models"
	"mmserver/services"

	"github.com/gin-gonic/gin"
)

var lineService = services.LineService{}

type SysController struct {
	BaseController
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
