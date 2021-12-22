/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-14 14:42:05
 * @LastEditTime: 2021-12-22 13:58:58
 * @LastEditors: liutq
 * @Reference:
 */
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-playground/validator/v10"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/maxyang107/gopush/model"
	"github.com/maxyang107/gopush/utils"
)

var validate *validator.Validate

func ServeHome(w http.ResponseWriter, r *http.Request) {

	if r.Method != "GET" {
		w.Write([]byte("只支持GET请求"))
		return
	}

	sourceCode := r.URL.Query().Get("sourceCode")

	if sourceCode == "" {

		w.Write([]byte("没有权限访问"))
		return
	}

	db, err := gorm.Open(utils.DbConfObj.DbType, utils.DbConfObj.DbDsn)

	if err != nil {
		utils.GoLog(err)
		return
	}
	defer db.Close()

	result := db.First(&model.Auth{}, "source_code = ?", sourceCode)

	if result.RowsAffected == 0 {
		utils.GoLog(result.Error)
		w.Write([]byte("没有权限访问"))
		return
	}

	httpmsg := r.URL.String()

	httpChan <- httpmsg

	fmt.Println("发送了一个chan信息")
}

/**
 * @description: 方法描述：权限管理
 * @Author: maxyang
 * @return {*}
 */

func manager(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Write([]byte("只支持POST raw json格式的请求"))
		return
	}
	var authMsg model.Auth
	body, _ := ioutil.ReadAll(r.Body)

	json.Unmarshal(body, &authMsg)
	validate = validator.New()
	err := validate.Struct(authMsg)

	if err != nil {
		validationErrors := err.Error()
		w.Write([]byte(validationErrors))
		return
	}

	db, err := gorm.Open(utils.DbConfObj.DbType, utils.DbConfObj.DbDsn)

	if err != nil {
		utils.GoLog(err)
		return
	}
	defer db.Close()
	db.Create(&authMsg)
}
