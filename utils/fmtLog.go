/*
 * @Description:
 * @Author: maxyang
 * @Date: 2021-12-16 17:53:05
 * @LastEditTime: 2021-12-22 11:57:24
 * @LastEditors: liutq
 * @Reference:
 */
package utils

import (
	"fmt"
	"log"
	"os"
)

func Dump(msg string) {

	if ConfigObj.Debug {
		log.Printf("info: %s", msg)
	}
}

func GoLog(err error) {
	if ConfigObj.Debug {
		log.Printf("error: %v", err)
	}
	writeLog(err)
}

func writeLog(err error) {
	logfile, err := os.OpenFile(fmt.Sprintf("log/%s.log", ConfigObj.LogName), os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		logfile, err = os.Create("log/%s.log")
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
	defer logfile.Close()
	logger := log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Llongfile)
	logger.Println(err)
}
