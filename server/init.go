/*
 * MIT License
 *
 * Copyright (c) 2017 SmartestEE Inc.
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

/*
 * Revision History:
 *     Initial: 2017/07/18        Yusan Kurban
 */

package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo"

	"ShopApi/log"
	"ShopApi/orm"
	"ShopApi/server/router"

	"ShopApi/general"
)

var (
	server *echo.Echo
)

func startServer() {
	server = echo.New()

	server.HTTPErrorHandler = general.EchoRestfulErrorHandler
	server.Validator = general.NewEchoValidator()

	router.InitRouter(server)
	log.Logger.Debug("Router already init")
	log.Logger.Fatal(server.Start(configuration.address))
}

func init() {
	readConfiguration()
	initMysql()
	startServer()
}

func initMysql() {
	user := configuration.mysqlUser
	pass := configuration.mysqlPass
	url := configuration.mysqlHost
	port := configuration.mysqlPort
	sqlName := configuration.mysqlDb

	conf := fmt.Sprintf(user + ":" + pass + "@" + "tcp(" + url + port + ")/" + sqlName + "?charset=utf8&parseTime=True&loc=Local")

	orm.InitOrm(conf)
}
