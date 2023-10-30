/*
 * @Author: chunhua yang
 * @Date: 2023-10-23 22:29:34
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:38:16
 * @FilePath: /minermanager/server/models/user.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package models

type TUser struct {
	BaseDBModel

	UserName string `json:"userName" gorm:"column:user_name"` //用户名
	Password string `json:"password" gorm:"column:password"`  //密码
} //t_user
