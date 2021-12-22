/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-14 14:53:00
 * @LastEditTime: 2021-12-21 17:50:27
 * @LastEditors: liutq
 * @Reference:
 */
package main

import (
	"fmt"
	"sync"
)

type Group struct {
	Name    string
	Clients map[string]*Client
	mapLuck sync.RWMutex
}

func newGroup(name string, hub *Hub) *Group {
	group := &Group{
		Name:    name,
		Clients: make(map[string]*Client),
	}

	hub.groups[group.Name] = group

	return group
}

func (this *Group) AddGroup(client *Client) {
	this.mapLuck.Lock()
	defer this.mapLuck.Unlock()

	this.Clients[client.ClientId] = client

	client.Groups[this.Name] = this
	client.send <- []byte("您已加入" + this.Name)
	fmt.Println("用户ID：" + client.ClientId + "成功加入房间：" + this.Name)
}

/**
 * @description: 方法描述：离开组
 * @Author: maxyang
 * @return {*}
 * @param {*Client} client
 */
func (this *Group) LeftGroup(client *Client) {
	_, err := this.Clients[client.ClientId]

	if err {
		return
	}

	this.mapLuck.Lock()
	defer this.mapLuck.Unlock()

	delete(this.Clients, client.ClientId)

	fmt.Println("用户ID：" + client.ClientId + "已离开房间：" + this.Name)
}

/**
 * @description: 方法描述：组内广播
 * @Author: maxyang
 * @return {*}
 * @param {string} msg
 */
func (this *Group) GroupBroadCast(msg string) {
	for _, cli := range this.Clients {
		fmt.Println(cli.ClientId)
		cli.send <- []byte(msg)
	}
}
