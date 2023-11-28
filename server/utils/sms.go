/*
 * @Author: chunhua yang
 * @Date: 2023-11-27 22:18:59
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-11-27 22:23:17
 * @FilePath: /minermanager/server/utils/sms.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */

package utils

import (
	"fmt"
	"strings"

	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/api/v2010"
)

// 短信接口，先虚拟一个接口，方便调用（后期再进行具体实现）
func SendSMS(receiver string, text string) (bool, error) {
	clientParam := twilio.ClientParams{
		Username:   AppConfig.SMSGateway.AccountsID,
		Password:   AppConfig.SMSGateway.Token,
		AccountSid: "",
	}
	client := twilio.NewRestClientWithParams(clientParam)

	params := &openapi.CreateMessageParams{}
	params.SetTo(receiver)
	params.SetFrom(AppConfig.SMSGateway.PhoneNumber)
	params.SetBody(text)

	_, err := client.Api.CreateMessage(params)
	if err != nil {
		return false, err
	} else {
		return true, nil
	}
}

// 短信发送到多个用户（多个用户之间用逗号分割，）
func SendSMSMuliteRecever(receiver string, text string) {
	user := strings.Split(receiver, ",")
	fmt.Println(user)
}
