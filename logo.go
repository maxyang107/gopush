/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-16 13:41:34
 * @LastEditTime: 2021-12-16 18:08:41
 * @LastEditors: liutq
 * @Reference:
 */
package main

import (
	"fmt"

	"github.com/maxyang107/gopush/utils"
)

var header = `

=================================推送中心启动成功==================================
`

var logo = `
                                                                                 
____  ___      ____  _   _ ____  _   _      ____ _____ _   _ _____ _____ ____  
/ ___|/ _ \    |  _ \| | | / ___|| | | |    / ___| ____| \ | |_   _| ____|  _ \ 
| |  _| | | |   | |_) | | | \___ \| |_| |   | |   |  _| |  \| | | | |  _| | |_) |
| |_| | |_| |   |  __/| |_| |___) |  _  |   | |___| |___| |\  | | | | |___|  _ < 
\____|\___/    |_|    \___/|____/|_| |_|    \____|_____|_| \_| |_| |_____|_| \_\

               稳     定     运     行     没      有     bug  																																						
`

var footer = `
==================================================================================
`

func FprintLogo() {
	fmt.Println(header)
	fmt.Println(logo)
	fmt.Println(footer)
	fmt.Println(fmt.Sprintf("当前监听地址: %s，   监听端口：%d", utils.ConfigObj.Address, utils.ConfigObj.Port))
}
