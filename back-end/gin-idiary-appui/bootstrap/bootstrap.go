/*
 * @Author: liziwei01
 * @Date: 2022-03-03 16:04:06
 * @LastEditors: liziwei01
 * @LastEditTime: 2022-03-03 16:22:57
 * @Description: file content
 */

package bootstrap

import (
	"context"
	"log"
)

/**
 * @description: start APP
 * @param {*}
 * @return {*}
 */
func Init() {
	// parse app.toml
	config, err := ParserAppConfig(appConfPath)
	if err != nil {
		log.Fatal("ParserAppConfig failed")
	}
	app := NewApp(context.Background(), config)

	//  start APP
	app.Start()
}
