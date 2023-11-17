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
	"database/sql/driver"
	"fmt"
	"time"

	"github.com/google/uuid"
)

// CustomTime 是一个自定义的时间类型
type CustomTime struct {
	time.Time
}

// Scan 用于从数据库扫描时间值
func (ct *CustomTime) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	switch v := value.(type) {
	case time.Time:
		ct.Time = v
	case []uint8:
		parseTime, err := time.Parse("2006-01-02 15:04:05", string(v))
		if err != nil {
			return err
		}
		ct.Time = parseTime
	default:
		return fmt.Errorf("无法处理的时间类型")
	}
	return nil
}

// Value 用于获取时间值以插入到数据库中
func (ct CustomTime) Value() (driver.Value, error) {
	return ct.Time, nil
}

// MarshalJSON 用于自定义JSON输出格式
func (ct CustomTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, ct.Time.Format("2006-01-02 15:04:05"))), nil
}

type BaseDBModel struct {
	Id         uuid.UUID   `json:"id" gorm:"column:id;primary_key"`     //主键
	CreateDate CustomTime  `json:"create_date" gorm:"column:create_date"` //创建日期
	CreateUser uuid.UUID   `json:"create_user" gorm:"column:create_user"` //创建用户
	IsDelete   bool        `json:"is_delete" gorm:"column:is_delete"`     //是否删除
	DeleteDate CustomTime `json:"delete_date" gorm:"column:delete_date"` //删除日期
	DeleteUser uuid.UUID   `json:"delete_user" gorm:"column:delete_user"` //删除用户
	UpdateDate CustomTime `json:"update_date" gorm:"column:update_date"` //更新日期
	UpdateUser uuid.UUID   `json:"update_user" gorm:"column:update_user"` //更新用户
}
