/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-04-25 00:20:22
 * @LastEditTime: 2021-12-16 18:36:33
 * @LastEditors: liutq
 * @Reference:
 */

package main

import (
	"flag"
	"fmt"
	"net/http"
	"net/url"

	"github.com/maxyang107/gopush/utils"
)

var httpChan = make(chan string, 10)

func main() {
	FprintLogo()
	flag.Parse()
	hub := newHub()
	go hub.run()
	go ListenHttpChan(hub)
	http.HandleFunc("/", ServeHome)
	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
	utils.Dump(utils.ConfigObj.Name + " 启动成功")
	err := http.ListenAndServe(fmt.Sprintf("%s:%d", utils.ConfigObj.Address, utils.ConfigObj.Port), nil)
	if err != nil {
		utils.GoLog(err)
	}
}

/**
 * @description: 方法描述：监听httpchan，分发任务
 * @Author: maxyang
 * @return {*}
 * @param {*Hub} hub
 */
func ListenHttpChan(hub *Hub) {
	for {
		msg := <-httpChan

		umsg, err := url.Parse(msg)

		if err != nil {
			panic(err)
		}

		param, _ := url.ParseQuery(umsg.RawQuery)

		switch param.Get("func") {
		//创建组
		case "createGroup":

			group := newGroup(param.Get("groupName"), hub)

			fmt.Println(group.Name)
			//组内广播
		case "groupBroadCast":
			if param.Get("message") == "" {
				fmt.Println("消息不能为空")

			} else {
				group := hub.groups[param.Get("groupName")]

				group.GroupBroadCast(param.Get("message"))
			}

			//获取组列表
		case "groupList":
			hub.getGroupList()
		case "groupMember":
			hub.getGroupMemberList(param.Get("groupName"))
		}
	}
}
