// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"encoding/json"

	"github.com/maxyang107/gopush/utils"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	onlineClients map[string]*Client

	// Inbound messages from the clients.
	broadcast chan []byte

	groupbroadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	onlinereg chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	groups map[string]*Group
}

func newHub() *Hub {
	return &Hub{
		broadcast:      make(chan []byte),
		groupbroadcast: make(chan []byte),
		register:       make(chan *Client),
		onlinereg:      make(chan *Client),
		unregister:     make(chan *Client),
		clients:        make(map[*Client]bool),
		groups:         make(map[string]*Group),
		onlineClients:  make(map[string]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				//用户下线退出所有组
				for _, group := range client.Groups {
					group.mapLuck.Lock()
					delete(group.Clients, client.ClientId)
					group.mapLuck.Unlock()
				}
				delete(h.onlineClients, client.ClientId)
				delete(h.clients, client)
				close(client.send)
			}
		case client := <-h.onlinereg:
			_, err := h.onlineClients[client.ClientId]
			if !err {
				h.onlineClients[client.ClientId] = client
			} else {
				client.send <- []byte("当前客户端id被占用")
			}
			client.send <- []byte("绑定成功")
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		case message := <-h.groupbroadcast:
			msg := make(map[string]interface{})
			err := json.Unmarshal(message, &msg)
			if err != nil {
				utils.Dump("message 反序列化失败")
			} else {
				groupId := msg["groupId"].(string)

				clients := h.groups[groupId].Clients
				for _, client := range clients {
					select {
					case client.send <- []byte(msg["message"].(string)):
					default:
						close(client.send)
						delete(h.clients, client)
					}
				}
			}
		}
	}
}

func (h *Hub) GetOnlineClientByName(name string) bool {
	client, _ := h.onlineClients[name]
	if client != nil {
		return true
	}
	return false
}

func (h *Hub) getGroupList() {
	for k, _ := range h.groups {
		utils.Dump(k)
	}
}

func (h *Hub) getGroupMemberList(name string) {
	for k, _ := range h.groups[name].Clients {
		utils.Dump(k)
	}
}
