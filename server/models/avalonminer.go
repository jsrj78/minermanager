/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 14:55:06
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-29 20:39:41
 * @FilePath: /minermanager/server/models/avalonminer.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */

package models

import (
	"fmt"
	"net/http"

	"github.com/icholy/digest"
)

type AvalonMiner struct {
	Miner
}

func NewAvalonMinerClient(m *Miner) *AvalonMiner {
	avalonMiner := &AvalonMiner{}
	avalonMiner.Miner = *m
	avalonMiner.RequestClient = http.Client{Transport: &digest.Transport{
		Username: avalonMiner.UserName,
		Password: avalonMiner.Password,
	}}

	return avalonMiner
}

func (avalon *AvalonMiner) Login() bool {

	//fmt.Println("avalon miner reboot")

	fmt.Println("avalon", avalon.Ip)

	return false
}

func (avalon *AvalonMiner) Normal() bool {

	//fmt.Println("avalon miner reboot")

	fmt.Println("avalon", avalon.Ip)

	return false
}

func (avalon *AvalonMiner) Sleep() bool {
	fmt.Println("avalon miner sleep", avalon.Ip)
	return true
}
