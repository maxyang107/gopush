/*
 * @Description:http消息推送日志
 * @Author: maxyang
 * @Date: 2021-12-21 10:55:53
 * @LastEditTime: 2021-12-21 16:10:53
 * @LastEditors: liutq
 * @Reference:
 */
package model

import "time"

type PushMessage struct {
	BaseModel
	Pcid     string `gorm:"type:varchar(64);unique_index;not null;"` //推送客户端id
	Ptype    int    `gorm:"type:tinyint(3);not null;"`               //推送消息类型
	Psource  string `gorm:"type:text;not null;"`                     //推送源
	Pcontent string `gorm:"type:text;not null;"`                     //推送消息
	Pstatus  string `gorm:"type:tinyint(3);not null;default:0;"`     //推送状态
	Ptip     string `gorm:"type:text;"`                              //推送日志
	Pushat   *time.Time
}

func (PushMessage) TableName() string {
	return "push_message"
}
