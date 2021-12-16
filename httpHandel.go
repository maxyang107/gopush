/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-14 14:42:05
 * @LastEditTime: 2021-12-16 18:51:29
 * @LastEditors: liutq
 * @Reference:
 */
package main

import (
	"fmt"
	"net/http"
)

func ServeHome(w http.ResponseWriter, r *http.Request) {

	httpmsg := r.URL.String()

	httpChan <- httpmsg

	fmt.Println("发送了一个chan信息")

}
