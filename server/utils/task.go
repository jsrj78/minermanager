/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 20:50:45
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-22 20:50:51
 * @FilePath: /minermanager/server/utils/task.go
 * @Description:
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package utils

type CronJob struct {
	Expression string
	Task       func()
}
