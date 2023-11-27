/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 14:55:06
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 20:56:58
 * @FilePath: /minermanager/server/models/miner.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */

package models

import (
	"context"
	"database/sql/driver"
	"fmt"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"golang.org/x/time/rate"
)

const pingInterval = 60 * time.Second

type MinerStatus int

type HashBoardStatus int

type MinerBrand int

type FaultType int

const (
	MinerOnline  MinerStatus = 0
	MinerOffline MinerStatus = 1
	MinerEmpty   MinerStatus = 2
	MinerRepair  MinerStatus = 3

	BoardNormal     HashBoardStatus = 0
	BoardAbnormal   HashBoardStatus = 1
	BoardDisconnect HashBoardStatus = 2
	BoardEmpty      HashBoardStatus = 3
	BoardRepair     HashBoardStatus = 4

	FaultFans         FaultType = 0
	FaultHashBoard    FaultType = 1
	FaultPower        FaultType = 2
	FaultControlBoard FaultType = 3
	FaultNetwork      FaultType = 4
	FaultUnknown      FaultType = 5

	AntMinerBrand    MinerBrand = 0
	WhatsMinerBrand  MinerBrand = 1
	AvalonMinerBrand MinerBrand = 2
)

var limiter = rate.NewLimiter(10, 1) // 每秒10个请求，桶大小为1

func (ms MinerStatus) String() string {
	return [...]string{"online", "offline", "empty", "repair"}[ms]
}

func (ms MinerStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ms.String() + `"`), nil
}

// GORM custom data type for saving as string in DB
func (ms MinerStatus) Value() (driver.Value, error) {
	return ms.String(), nil
}

func (ms *MinerStatus) Scan(value interface{}) error {
	switch value.(string) {
	case "online":
		*ms = MinerOnline
	case "offline":
		*ms = MinerOffline
	case "empty":
		*ms = MinerEmpty
	case "repair":
		*ms = MinerRepair
	default:
		return fmt.Errorf("unsupported value for MinerStatus: %v", value)
	}
	return nil
}

func (ms HashBoardStatus) String() string {
	return [...]string{"normal", "abnormal", "empty", "disconnect", "repair"}[ms]
}

func (ms HashBoardStatus) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ms.String() + `"`), nil
}

// GORM custom data type for saving as string in DB
func (ms HashBoardStatus) Value() (driver.Value, error) {
	return ms.String(), nil
}

func (ms *HashBoardStatus) Scan(value interface{}) error {
	switch value.(string) {
	case "normal":
		*ms = BoardNormal
	case "abnormal":
		*ms = BoardAbnormal
	case "empty":
		*ms = BoardEmpty
	case "disconnect":
		*ms = BoardDisconnect
	case "repair":
		*ms = BoardRepair
	default:
		return fmt.Errorf("unsupported value for HashBoardStatus: %v", value)
	}
	return nil
}

func (ms FaultType) String() string {
	return [...]string{"fan", "hashboard", "controllerboard", "power", "network"}[ms]
}

func (ms FaultType) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ms.String() + `"`), nil
}

// GORM custom data type for saving as string in DB
func (ms FaultType) Value() (driver.Value, error) {
	return ms.String(), nil
}

func (ms *FaultType) Scan(value interface{}) error {
	switch value.(string) {
	case "fan":
		*ms = FaultFans
	case "hashboard":
		*ms = FaultHashBoard
	case "controllerboard":
		*ms = FaultControlBoard
	case "power":
		*ms = FaultPower
	case "network":
		*ms = FaultNetwork
	case "unknown":
		*ms = FaultUnknown
	default:
		return fmt.Errorf("unsupported value for FaultType: %v", value)
	}
	return nil
}

func (ms MinerBrand) String() string {
	return [...]string{"AntMiner", "WhatsMiner", "AvalonMiner"}[ms]
}

func (ms MinerBrand) MarshalJSON() ([]byte, error) {
	return []byte(`"` + ms.String() + `"`), nil
}

// GORM custom data type for saving as string in DB
func (ms MinerBrand) Value() (driver.Value, error) {
	return ms.String(), nil
}

func (ms *MinerBrand) Scan(value interface{}) error {
	switch value.(string) {
	case "AntMiner":
		*ms = AntMinerBrand
	case "WhatsMiner":
		*ms = WhatsMinerBrand
	case "AvalonMiner":
		*ms = AvalonMinerBrand
	default:
		return fmt.Errorf("unsupported value for MinerBrand: %v", value)
	}
	return nil
}

type Fans struct {
	Id          int       `json:"id" gorm:"column:id"`                    //风扇
	Postion     int       `json:"postion" gorm:"column:postion"`          //风扇位置
	Speed       int       `json:"speed" gorm:"column:speed"`              //风扇转速
	IsNormal    bool      `json:"isNormal" gorm:"column:is_normal"`       //风扇是否正常
	InstallDate uuid.Time `json:"installDate" gorm:"column:install_date"` //风扇安装日期

	MinerId uuid.UUID `json:"minerId" gorm:"column:miner_id"` //矿机ID
}

type HashBoard struct {
	Id          uuid.UUID       `json:"id" gorm:"column:id"`                    //Hash板id
	Postion     int             `json:"postion" gorm:"column:postion"`          //Hash板位置
	Code        string          `json:"code" gorm:"column:code"`                //Hash板编号
	HashRate    int             `json:"hashRate" gorm:"column:hash_rate"`       //Hash板算力
	Temperature []int           `json:"temperature" gorm:"column:temperature"`  //Hash板温度
	Status      HashBoardStatus `json:"status" gorm:"column:status"`            //Hash板是否正常
	InstallDate uuid.Time       `json:"installDate" gorm:"column:install_date"` //Hash板安装日期

	MinerId uuid.UUID `json:"minerId" gorm:"column:miner_id"` //矿机ID
}

type Pools struct {
	Id        uuid.UUID `json:"id" gorm:"column:id"`               //矿池ID
	URL       string    `json:"url" gorm:"column:url"`             //矿池URL
	User      string    `json:"user" gorm:"column:user"`           //矿池用户名
	Status    string    `json:"status" gorm:"column:status"`       //矿池状态
	Priority  int       `json:"priority" gorm:"column:priority"`   //矿池优先级
	GetWorks  int       `json:"getworks" gorm:"column:getworks"`   //矿池获取任务数
	Accepted  int       `json:"accepted" gorm:"column:accepted"`   //矿池接受任务数
	Rejected  int       `json:"rejected" gorm:"column:rejected"`   //矿池拒绝任务数
	Discarded int       `json:"discarded" gorm:"column:discarded"` //矿池丢弃任务数
	Stale     int       `json:"stale" gorm:"column:stale"`         //矿池过期任务数
	Diff      string    `json:"diff" gorm:"column:diff"`           //矿池难度
	Diff1     int       `json:"diff1" gorm:"column:diff1"`         //矿池难度1
	DiffA     int       `json:"diffa" gorm:"column:diffa"`         //矿池难度A
	DiffR     int       `json:"diffr" gorm:"column:diffr"`         //矿池难度R
	DiffS     int       `json:"diffs" gorm:"column:diffs"`         //矿池难度S
	LSDiff    int       `json:"lsdiff" gorm:"column:lsdiff"`       //矿池最小难度
	LSTime    string    `json:"lstime" gorm:"column:lstime"`       //矿池最小难度时间

	MinerId uuid.UUID `json:"minerId" gorm:"column:miner_id"` //矿机ID
}

type Miner struct {
	Id    uuid.UUID  `json:"id" gorm:"column:id"`       //矿机ID
	Brand MinerBrand `json:"brand" gorm:"column:brand"` //矿机品牌
	Model string     `json:"model" gorm:"column:model"` //矿机类型
	SN    string     `json:"sn" gorm:"column:sn"`       //矿机序列号

	Status MinerStatus `json:"status" gorm:"column:status"` //矿机状态

	Ip         string `json:"ip" gorm:"column:ip"`                 //矿机IP
	MacAddress string `json:"macAddress" gorm:"column:macaddress"` //矿机MAC地址
	HostName   string `json:"hostName" gorm:"column:hostname"`     //矿机主机名
	Netmask    string `json:"netmask" gorm:"column:netmask"`       //矿机子网掩码
	Gateway    string `json:"gateway" gorm:"column:gateway"`       //矿机网关
	DNS        string `json:"dns" gorm:"column:dns"`               //矿机DNS

	UserName        string `json:"username" gorm:"column:username"`                //矿机用户名
	Password        string `json:"password" gorm:"column:password"`                //矿机密码
	DefaultUserName string `json:"defaultUsername" gorm:"column:default_username"` //矿机默认用户名
	DefaultPassword string `json:"defaultPassword" gorm:"column:default_password"` //矿机默认密码

	ControlBoardModel       string `json:"controlBoardModel" gorm:"column:controlboard_model"`            //控制板型号
	ControlBoardSN          string `json:"controlBoardSN" gorm:"column:controlboard_sn"`                  //控制板序列号
	ControlBoardInstallDate string `json:"controlBoardInstallDate" gorm:"column:controlboard_intalldate"` //控制板MAC地址
	FirmwareVersion         string `json:"firmwareVersion" gorm:"column:firmware_version"`                //固件版本

	PowerSupplyModel       string `json:"powerSupplyModel" gorm:"column:powersupply_model"`             //电源
	PowerSupplySN          string `json:"powerSupplySN" gorm:"column:powersupply_sn"`                   //电源
	PowerSupplyInstallDate string `json:"powerSupplyInstallDate" gorm:"column:powersupply_installdate"` //电源功率

	//Fans      []Fans      `json:"fans" gorm:"column:fans"`            //风扇
	//HashBoard []HashBoard `json:"hashBoard" gorm:"column:hash_board"` //Hash板
	//Pools     []Pools     `json:"pools" gorm:"column:pools"`          //矿池

	Behavior MinerBehavior `json:"behavior" gorm:"-"` //矿机行为

	StopCh        chan bool   `json:"stopCh" gorm:"-"`        //矿机停止通道
	RequestClient http.Client `json:"requestclient" gorm:"-"` //矿机Http请求客户端
}

type MinerFault struct {
	Id        uuid.UUID   `json:"id"`         //故障ID
	Ip        string      `json:"ip" `        //矿机IP
	FaultType FaultType   `json:"faultType" ` //故障类型
	Content   interface{} `json:"content"`    //故障内容
	Memo      string      `json:"memo"`       //备注
}

var instanceMinerStatus *MinerStatusManager
var onceMinerStatus sync.Once

// 实时矿机状态管理（只考虑在线离线2种状态，其他状态交给数据库处理）
type MinerStatusManager struct {
	Mux         sync.Mutex
	OnlineList  map[uuid.UUID]string
	OfflineList map[uuid.UUID]string
	FaultList   map[uuid.UUID]interface{}
}

// 矿机行为接口
type MinerBehavior interface {
	Run(manager *MinerStatusManager)
	Login() bool
	Normal() bool
	Sleep() bool

	// Ping() bool
	//Reboot()
	// Upgrade()
	// Reset()

	// GetLog()
	// GetHistoryLog()

	// GetFan()
	// GetHashBoard()
}

// 定义单一实例变量
var instanceMinerFactory *MinerFactory
var onceMinerFactory sync.Once

type MinerFactory struct {
	Miners             map[uuid.UUID]*Miner
	MinerStatusManager *MinerStatusManager
}

// 执行这个ping函数需要root或管理员权限
func (miner *Miner) ping() (bool, error) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", "3", miner.Ip)
	case "linux":
		cmd = exec.Command("ping", "-c", "3", miner.Ip)
	default:
		return false, fmt.Errorf("unsupported operating system")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		//fmt.Println("Command failed:", err)
		return false, err
	}

	//fmt.Println("Command output:", string(output))

	successful := strings.Contains(string(output), "Lost = 0") &&
		strings.Contains(string(output), "Received = 3")

	return successful, nil
}

func (miner *Miner) Login() bool {
	return true
}

func (miner *Miner) Reboot() {
	fmt.Println("base miner reboot")
}

func (miner *Miner) Run(manager *MinerStatusManager) {
	//fmt.Println("miner run", miner.Ip)
	go func(miner *Miner) {
		ticker := time.NewTicker(pingInterval)
		defer ticker.Stop()

		for {
			select {
			case <-ticker.C:
				//fmt.Println("start to ping", miner.Ip)
				go func() {
					_, err := miner.ping()
					if err != nil {
						//fmt.Println("ping fail", miner.Ip, err.Error())
						miner.Status = MinerOffline
					} else {
						miner.Status = MinerOnline
						//在线状态下，判断是否登录成功
						if miner.Behavior.Login() {
							if miner.Behavior.Normal() {
								//整机运行正常，但是也有可能缺少Hash板，或者部分Hash板故障

							} else {
								// 错误检查
							}
						} else {
							//错误监测

						}

					}

					// 处理结果离线在线结果
					manager.ChangeState(miner, miner.Status)
				}()

			case <-miner.StopCh:
				return
			}
		}
	}(miner)
}

// 矿机状态管理器
func NewMinerManager() *MinerStatusManager {
	onceMinerStatus.Do(func() {
		instanceMinerStatus = &MinerStatusManager{
			OnlineList:  make(map[uuid.UUID]string),
			OfflineList: make(map[uuid.UUID]string),
		}
	})
	return instanceMinerStatus
}

func (manager *MinerStatusManager) ChangeState(miner *Miner, status MinerStatus) {
	manager.Mux.Lock()
	defer manager.Mux.Unlock()

	if status == MinerOnline {
		// Remove from OfflineList
		delete(manager.OfflineList, miner.Id)
		//fmt.Println("miner online", miner.Ip)
		manager.OnlineList[miner.Id] = miner.Ip

		// for id, value := range manager.OnlineList {
		// 	fmt.Println("start to read online and offline miner")
		// 	fmt.Printf("ID: %s, Value: %s\n", id, value)
		// }
	} else {
		// Remove from OnlineList
		delete(manager.OnlineList, miner.Id)

		manager.OfflineList[miner.Id] = miner.Ip
	}
}

func NewMinerFactory() *MinerFactory {
	onceMinerFactory.Do(func() {
		instanceMinerFactory = &MinerFactory{
			Miners:             make(map[uuid.UUID]*Miner),
			MinerStatusManager: NewMinerManager(),
		}
	})
	return instanceMinerFactory
}

func (factory *MinerFactory) CreateMiners(miners []*Miner) {
	// 创建一个每秒产生3个令牌的限速器
	limiter := rate.NewLimiter(500, 1)

	for _, miner := range miners {

		err := limiter.Wait(context.Background())
		if err != nil {
			fmt.Println("Error waiting for token:", err)
			return
		}

		//fmt.Println("create miner", miner.Ip)
		factory.Miners[miner.Id] = miner

		switch miner.Brand {
		case AntMinerBrand:
			ant := NewAntMinerClient(miner)
			miner.Behavior = ant
		case AvalonMinerBrand:
			avalon := NewAvalonMinerClient(miner)
			miner.Behavior = avalon
		}

		miner.Run(factory.MinerStatusManager)
	}
}

func (factory *MinerFactory) GetMiner(id uuid.UUID) *Miner {
	return factory.Miners[id]
}

func (factory *MinerFactory) StopMiner(id uuid.UUID) {
	factory.Miners[id].StopCh <- true
}
