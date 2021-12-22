/*
 * @Description: 权限
 * @Author: maxyang
 * @Date: 2021-12-21 16:34:27
 * @LastEditTime: 2021-12-21 17:50:09
 * @LastEditors: liutq
 * @Reference:
 */
package model

type Auth struct {
	BaseModel

	Source string `gorm:"type:varchar(100);unique_index;not null;" validate:"required,max=100"`

	SourceCode int `gorm:"type:char(6);unique_index;not null;" validate:"required,gte=10000,lte=999999"`

	SourceTip string `gorm:"type:varchar(100)" validate:"max=100"`

	SourceStatus int `gorm:"type:tinyint(3);not null;default:0;" validate:"required,max=3,number"`
}

func (Auth) TableName() string {
	return "auth"
}
