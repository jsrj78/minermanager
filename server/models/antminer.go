/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 14:55:06
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 20:52:01
 * @FilePath: /minermanager/server/models/antminer.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */

package models

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/icholy/digest"
)

// 	Response Body:
// <?xml version="1.0" encoding="iso-8859-1"?>
// <!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN"
//          "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
// <html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en">
//  <head>
//   <title>401 - Unauthorized</title>
//  </head>
//  <body>
//   <h1>401 - Unauthorized</h1>
//  </body>
// </html>

var (
	URL_Get_SystemInfo string = "/cgi-bin/get_system_info.cgi" //获取Miner系统信息
	URL_Get_Pools      string = "/cgi-bin/pools.cgi"           //获取矿池信息
	URL_Get_Stats      string = "/cgi-bin/stats.cgi"           //获取算力版信息
	URL_Get_Summary    string = "cgi-bin/summary.cgi"          //Miner算力和
	URL_Get_Chart      string = "/cgi-bin/chart.cgi"           //算力分时图
	URL_Get_Log        string = "/cgi-bin/log.cgi"             // 得到当前日志
	URL_Get_HistoryLog string = "/cgi-bin/hlog.cgi"            //得到历史日志

	URL_Get_Blink string = "/cgi-bin/get_blink_status.cgi" //得到Miner定位指示
	URL_Set_Blink string = "/cgi-bin/blink.cgi"            //设置Miner定位

	URL_Set_Network string = "/cgi-bin/set_network_conf.cgi" //设置网络配置

	URL_Post_FirmWire       string = "/cgi-bin/upgrade.cgi"       //上传固件
	URL_Post_FirmWire_Clear string = "/cgi-bin/upgrade_clear.cgi" //上传固件(清楚原来配置)
)

type AntMiner struct {
	Miner
}

func NewAntMinerClient(m *Miner) *AntMiner {

	antMiner := &AntMiner{}
	antMiner.Miner = *m
	//fmt.Println("user,password", antMiner.UserName, antMiner.Password)
	antMiner.RequestClient = http.Client{Transport: &digest.Transport{
		Username: antMiner.UserName,
		Password: antMiner.Password,
	}}

	return antMiner
}

func (ant *AntMiner) Login() bool {
	//fmt.Println("kaishideng", ant.Ip)
	res, err := ant.RequestClient.Get("http://" + ant.Ip + URL_Get_SystemInfo)
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		//fmt.Println("HTTP响应状态不是OK:",ant.Ip, res.Status)

		return false
	}else{
		//fmt.Println("HTTP  OK:",ant.Ip, res.Status)
	}

	// 使用json.Unmarshal解码JSON响应并将其赋值给结构体
	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&ant.Miner); err != nil {
		fmt.Println("JSON解码失败:", err)
		return false
	}
	//fmt.Println(ant.Miner.Ip, ant.Miner.Gateway)

	return true
}

func (ant *AntMiner) Sleep() bool {
	return true
}
