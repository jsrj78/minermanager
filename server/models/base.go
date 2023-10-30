/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 15:00:19
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-24 22:42:58
 * @FilePath: /minermanager/server/models/base.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package models

import (
	"time"

	"github.com/google/uuid"
)


type BaseDBModel struct {
	CreateDate time.Time `json:"create_date" gorm:"column:create_date"` //创建日期
	CreateUser uuid.UUID `json:"create_user" gorm:"column:create_user"` //创建用户
	IsDelete   bool      `json:"is_delete" gorm:"column:is_delete"`     //是否删除
	DeleteDate time.Time `json:"delete_date" gorm:"column:delete_date"` //删除日期
	DeleteUser uuid.UUID `json:"delete_user" gorm:"column:delete_user"` //删除用户
	UpdateDate time.Time `json:"update_date" gorm:"column:update_date"` //更新日期
	UpdateUser uuid.UUID `json:"update_user" gorm:"column:update_user"` //更新用户
}
