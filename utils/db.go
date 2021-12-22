/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-21 16:11:19
 * @LastEditTime: 2021-12-22 12:43:22
 * @LastEditors: liutq
 * @Reference:
 */
package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type DbConfig struct {
	DbHost string // 数据库连接

	DbPort int //数据库端口

	DbName string //数据库名称

	DbUser string //用户名称

	DbPwd string //用户密码

	DbCharset string //字符集

	DbParseTime string

	DbLoc string

	DbPrefix string //表前缀

	DbType string //数据库类型

	DbDsn string

	DbTimeout int
}

var DbConfObj *DbConfig

func init() {
	DbConfObj = &DbConfig{
		DbHost:      "127.0.0.1",
		DbPort:      3306,
		DbName:      "gopush",
		DbCharset:   "utf8mb4",
		DbParseTime: "true",
		DbLoc:       "Local",
		DbType:      "mysql",
		DbDsn:       "",
		DbTimeout:   10,
	}

	data, err := ioutil.ReadFile("db.json")

	if err == nil {
		err = json.Unmarshal(data, &DbConfObj)
		if err != nil {
			fmt.Println("配置文件加载错误：", err)
			panic(err)
		}
	}

	DbConfObj.DbDsn = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=%s&loc=%s&timeout=%ds",
		DbConfObj.DbUser,
		DbConfObj.DbPwd,
		DbConfObj.DbHost,
		DbConfObj.DbPort,
		DbConfObj.DbName,
		DbConfObj.DbCharset,
		DbConfObj.DbParseTime,
		DbConfObj.DbLoc,
		DbConfObj.DbTimeout)
}
