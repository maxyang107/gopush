/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-21 13:55:40
 * @LastEditTime: 2021-12-22 12:38:07
 * @LastEditors: liutq
 * @Reference:
 */
package model

import (
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/maxyang107/gopush/utils"
)

func Test_user(t *testing.T) {
	// db, err := gorm.Open("mysql", "root:root12345@tcp(127.0.0.1:3306)/gopush?charset=utf8mb4&parseTime=true&loc=Local")
	utils.Dump(utils.DbConfObj.DbDsn)
	db, err := gorm.Open(utils.DbConfObj.DbType, utils.DbConfObj.DbDsn)
	if err != nil {

		fmt.Println(err)
	}

	//迁移
	// db.AutoMigrate(&Auth{})

	// table := PushMessage{
	// 	BaseModel: BaseModel{},
	// 	Pcid:      "5",
	// 	Ptype:     0,
	// 	Psource:   "123",
	// 	Pcontent:  "nihao",
	// 	Pstatus:   "",
	// 	Ptip:      "",
	// 	Pushat:    nil,
	// }
	// fmt.Println(table)
	// db.Create(&table)
	defer db.Close()
}
