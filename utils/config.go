/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-16 13:45:22
 * @LastEditTime: 2021-12-16 18:33:05
 * @LastEditors: liutq
 * @Reference:
 */

package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"time"
)

type Config struct {
	Name    string //服务名称
	Address string //服务端ip

	Port int // 服务端口号

	WriteWait time.Duration //允许向对等端写入消息的时间

	PongWait time.Duration //允许接受客户端pong最长等待时间。

	PingPeriod time.Duration //向客户端发送心跳的最长等待时间，必须比pongwait小

	MaxMessageSize int64 //允许传输的最长消息长度

	Debug bool //是否开启debug模式

	LogName string //日志名称
}

var ConfigObj *Config

func init() {
	ConfigObj = &Config{
		Name:           "GO PUSH SERVER",
		Address:        "0.0.0.0",
		Port:           8088,
		WriteWait:      10,
		PongWait:       60,
		PingPeriod:     54,
		MaxMessageSize: 1024,
		Debug:          true,
		LogName:        "gopush",
	}
	data, err := ioutil.ReadFile("conf.json")

	if err == nil {
		err = json.Unmarshal(data, &ConfigObj)
		if err != nil {
			fmt.Println("配置文件加载错误：", err)
			panic(err)
		}
	}
	ConfigObj.WriteWait = ConfigObj.WriteWait * time.Second

	ConfigObj.PongWait = ConfigObj.PongWait * time.Second

	ConfigObj.PingPeriod = ConfigObj.PongWait * 9 / 10

}
