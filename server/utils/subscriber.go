/*
 * @Author: chunhua yang
 * @Date: 2023-10-22 21:25:23
 * @LastEditors: Max jsrj78@gmail.com
 * @LastEditTime: 2023-10-23 21:55:32
 * @FilePath: /minermanager/server/utils/subscriber.go
 * @Description:订阅者模式
 *
 * Copyright (c) 2023 by Golink LLC All Rights Reserved.
 */
package utils

import "sync"

type Subscriber func()

type PubSub struct {
	mtx         sync.RWMutex
	subscribers []Subscriber
}

func (ps *PubSub) Subscribe(target Subscriber) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	ps.subscribers = append(ps.subscribers, target)

}

func (ps *PubSub) Unsubscribe(target Subscriber) {
	ps.mtx.Lock()
	defer ps.mtx.Unlock()

	for i, s := range ps.subscribers {
		if &s == &target {
			// 使用切片技巧删除该订阅者
			ps.subscribers = append(ps.subscribers[:i], ps.subscribers[i+1:]...)
			return
		}
	}
}

func (ps *PubSub) Publish() {
	ps.mtx.RLock()
	defer ps.mtx.RUnlock()

	for _, s := range ps.subscribers {
		go s()
	}
}
