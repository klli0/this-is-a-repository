/*
 * @Author: liziwei01
 * @Date: 2022-03-03 15:33:30
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-04-22 20:51:29
 * @Description: main
 */
package main

import (
	"gin-idiary-appui/bootstrap"
	"gin-idiary-appui/httpapi"
	"log"
)

func main() {
	app, err := bootstrap.Setup()
	if err != nil {
		log.Fatalln(err)
	}
	// 注册接口路由
	httpapi.InitRouters(app.Handler)

	app.Start()
}
