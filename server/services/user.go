/*
 * @Author: chunhua yang
 * @Date: 2023-10-23 22:35:27
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 22:39:12
 * @FilePath: /minermanager/server/services/user.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package services

import (
	"mmserver/models"
	"mmserver/utils"

	pag "github.com/mmuflih/gorm-paginator"
)

type UserService struct {
	BaseService
}

// CreateUser
func (u *UserService) CreateUser(user models.TUser) (*models.TUser, error) {

	if err := utils.DB.Create(&user).Error; err != nil {

		return &user, err
	}

	return &user, nil
}

// UpdateUser
func (u *UserService) UpdateUser(user models.TUser) (*models.TUser, error) {

	if err := utils.DB.Model(&user).Updates(&user).Error; err != nil {

		return &user, err
	}

	return &user, nil
}

// GetAllUsers
func (u *UserService) GetAllUsersPaginator(page int, pagesize int, order interface{}, query interface{}, args ...interface{}) *pag.Paginator {

	user := []models.TUser{}

	result := pag.Make(&pag.Config{
		DB:   utils.DB.Where(query, args...).Find(&user).Order(order),
		Page: page,
		Size: pagesize,
	}, &user)

	return result
}
