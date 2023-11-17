package models

import "github.com/google/uuid"

type TLine struct {
	BaseDBModel
	Code                string    `json:"code" gorm:"column:code"`                                   //线编号
	Title               string    `json:"title" gorm:"column:title"`                                 //线名称
	BoxBegin            int       `json:"box_begin" gorm:"column:box_begin"`                         //盒子的开始编号
	Boxes               int       `json:"boxes" gorm:"column:boxes"`                                 //线的总box数量
	Ip                  string    `json:"ip" gorm:"column:ip"`                                       //ip地址（开始2位）
	RuleFloors          int       `json:"rule_floors" gorm:"column:rule_floors"`                     //每个盒子多少层
	RuleMiners          int       `json:"rule_miners" gorm:"column:rule_miners"`                     //每层多少个miners
	RuleBoxColumns      int       `json:"rule_box_columns" gorm:"column:rule_box_columns"`           //每个盒子包括多少列，默认1列，有的3列
	RuleUserName        string    `json:"rule_username" gorm:"column:rule_username"`                 //每个盒子多少层
	RulePassword        string    `json:"rule_password" gorm:"column:rule_password"`                 //每层多少个miners
	ActualTotalMiners   int       `json:"actual_total_miners" gorm:"column:actual_total_miners"`     //实际总机器数量
	ActualEmptyPlace    string    `json:"actual_empty_place" gorm:"column:actual_empty_place"`       //实际的空机位
	ActualEmptyMiners   int       `json:"actual_empty_miners" gorm:"column:actual_empty_miners"`     //其中，空机器数量
	ActualNormalMiners  int       `json:"actual_normal_miners" gorm:"column:actual_normal_miners"`   //所有正常的机器
	ActualOfflineMiners int       `json:"actual_offline_miners" gorm:"column:actual_offline_miners"` //所有离线的机器
	IsLift              *bool     `json:"islift" gorm:"column:islift"`                               //盒子高度是否需要升降机
	IpStart             int       `json:"ip_start" gorm:"column:ip_start"`                           //ip从第几位开始，默认从1，也有从6或者25的                        //
	MinrBrandId         uuid.UUID `json:"miner_brandid" gorm:"column:miner_brandid"`                 //矿机品牌
} //t_line

type TBoxes struct {
	BaseDBModel
	LineId        uuid.UUID `json:"lineid" gorm:"column:lineid"`                 //所在货架id
	Ip_Start      string    `json:"ip_start" gorm:"column:ip_start"`             //开始ip
	Ip_End        string    `json:"ip_end" gorm:"column:ip_end"`                 //结束ip
	Islift        *bool     `json:"islift" gorm:"column:islift"`                 //是否需要升降机
	TotalPlace    int       `json:"total_place" gorm:"column:total_place"`       //实际的总机位
	EmptyPlace    int       `json:"empty_place" gorm:"column:empty_place"`       //实际的空机位
	TotalMiners   int       `json:"total_miners" gorm:"column:total_miners"`     //实际机器数量
	EmptyMiners   int       `json:"empty_miners" gorm:"column:empty_miners"`     //空壳器数量
	NormalMiners  int       `json:"normal_miners" gorm:"column:normal_miners"`   //所有正常的机器
	OfflineMiners int       `json:"offline_miners" gorm:"column:offline_miners"` //所有离线的机器
} //t_boxes

type TMiner struct {
	BaseDBModel
	BoxId           uuid.UUID   `json:"boxid" gorm:"column:boxid"`                       //所在盒子id
	Ip              string      `json:"ip" gorm:"type:inet;column:ip"`                   //IP地址
	Status          MinerStatus `json:"status" gorm:"column:status"`                     //状态
	Brand           MinerBrand  `json:"brand" gorm:"column:brand"`                       //矿机品牌
	Username        string      `json:"username" gorm:"column:username"`                 //用户名
	Password        string      `json:"password" gorm:"column:password"`                 //密码
	DefaultUsername string      `json:"default_username" gorm:"column:default_username"` //默认用户名
	DefaultPassword string      `json:"default_password" gorm:"column:default_password"` //默认密码
	ChildBox        int         `json:"child_box" gorm:"column:child_box"`               //如果盒子有分几栏，记录所在栏id
	Cols            int         `json:"col" gorm:"column:cols"`                          //所在行
	Rows            int         `json:"row" gorm:"column:rows"`                          //所在列

} //t_miner
