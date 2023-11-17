package models

import "github.com/google/uuid"

type TBrand struct {
	BaseDBModel
	Code     string `json:"code" gorm:"column:code"`         //厂家编号
	Title    string `json:"title" gorm:"column:title"`       //厂家名称
	IOrder   int    `json:"iorder" gorm:"column:iorder"`     //序号
	UserName string `json:"username" gorm:"column:username"` //默认用户名
	Password string `json:"password" gorm:"column:password"` //默认密码
} //t_brand

type TModel struct {
	BaseDBModel
	BrandID  uuid.UUID `json:"brandid" gorm:"column:brandid"`   //厂家id
	Code     string    `json:"code" gorm:"column:code"`         //型号编号
	Title    string    `json:"title" gorm:"column:title"`       //型号名称
	User     string    `json:"user" gorm:"column:user"`         //默认矿机用户名
	Password string    `json:"password" gorm:"column:password"` //默认矿机密码
	IOrder   int       `json:"iorder" gorm:"column:iorder"`     //排序
} //t_model
