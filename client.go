// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
	"github.com/maxyang107/gopush/utils"
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     CheckOriginT,
}

func CheckOriginT(r *http.Request) bool {
	return true
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	ClientId string

	Groups map[string]*Group

	// Buffered channel of outbound messages.
	send chan []byte
}

/**
 * @description: 方法描述：readPump将消息从websocket连接泵到集线器。应用程序在每个连接的goroutine协程中运行readPump。应用程序
 *  			 通过执行所有读取器，确保一个连接上最多有一个读取器从这个goroutine中读取。
 * @Author: maxyang
 * @return {*}
 */
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(int64(utils.ConfigObj.MaxMessageSize))
	c.conn.SetReadDeadline(time.Now().Add(utils.ConfigObj.PongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(utils.ConfigObj.PongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				utils.GoLog(err)
			}
			break
		}

		//心跳
		if string(message) == "@heart" {
			message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
			c.hub.broadcast <- message
		} else { //业务逻辑
			msg := make(map[string]interface{})
			err1 := json.Unmarshal(message, &msg)

			if err1 != nil || (msg["option"] == nil || msg["option"] == "") {
				c.packageAndSendMsg([]byte("非法连接请求，已断开链接"))
				return
			}

			if msg["option"].(string) != "bindUserId" && c.ClientId == "" {
				c.packageAndSendMsg([]byte("请完成用户与服务端绑定后再试"))
				return
			}
			switch msg["option"].(string) {
			case "bindUserId":
				c.ClientId = msg["userId"].(string)

				c.hub.onlinereg <- c

			case "joingroup":
				group, err := c.hub.groups[msg["groupName"].(string)]
				if !err {
					c.packageAndSendMsg([]byte("分组不存在，加入失败"))

				} else {
					group.AddGroup(c)
				}
			case "groupchat":
				content := msg["message"].(string)
				group, err := c.hub.groups[msg["groupName"].(string)]

				if err {
					group.GroupBroadCast(content)
				} else {
					c.packageAndSendMsg([]byte("分组不存在"))
				}

			}

		}
	}
}

/**
 * @description: 方法描述：writePump将消息从集线器泵到websocket连接。每个连接都会启动一个运行writePump的goroutine。应用程序确保一个连接最多有一个写入器执行这个goroutine的所有写操作。
 * @Author: maxyang
 * @return {*}
 */
func (c *Client) writePump() {
	ticker := time.NewTicker(utils.ConfigObj.PingPeriod)
	defer func() {
		utils.Dump("客户端连接断开，客户端ID：" + c.ClientId)
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(utils.ConfigObj.WriteWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)

			if err != nil {
				utils.GoLog(err)
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}
			if err := w.Close(); err != nil {
				utils.GoLog(err)
				return
			}

		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(utils.ConfigObj.WriteWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				utils.GoLog(err)
				return
			}
		}
	}
}

/**
 * @description: 方法描述：serveWs处理来自对等端的websocket请求。
 * @Author: maxyang
 * @return {*}
 * @param {*Hub} hub
 * @param {http.ResponseWriter} w
 * @param {*http.Request} r
 */
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		utils.GoLog(err)
		return
	}
	client := &Client{hub: hub, conn: conn, Groups: make(map[string]*Group), send: make(chan []byte, 256)}
	client.hub.register <- client

	client.send <- []byte("请绑定用户id")
	go client.writePump()
	go client.readPump()

	utils.Dump("客户端连接成功，当前客户端：" + conn.RemoteAddr().String())
}

func (c *Client) packageAndSendMsg(msg []byte) {
	message := bytes.TrimSpace(bytes.Replace(msg, newline, space, -1))
	c.send <- message
}
