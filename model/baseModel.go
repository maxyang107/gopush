/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-21 10:54:22
 * @LastEditTime: 2021-12-21 17:53:12
 * @LastEditors: liutq
 * @Reference:
 */
package model

import "time"

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
