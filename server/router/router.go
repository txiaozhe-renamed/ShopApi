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
 *     Modify: 2017/07/19         Yang Zhengtian   添加返回收获地址
 *     Modify: 2017/07/20         Yang Zhengtain    添加修改密码
 */

package router

import (
	"github.com/labstack/echo"

	"ShopApi/handler"
)

func InitRouter(server *echo.Echo) {
	if server == nil {
		panic("[InitRouter], server couldn't be nil")
	}

	server.POST("/api/v1/user/create", handler.Create)
	server.POST("/api/v1/user/login", handler.Login)
	server.GET("/api/v1/user/logout", handler.Logout)
	server.POST("/api/v1/user/changemobilepass",handler.ChangeMobilePassword)
	server.POST("/api/v1/user/changeinfo", handler.ChangeUserInfo)
	server.POST("/api/v1/user/changepass",handler.ChangeMobilePassword,handler.MustLogin)
	server.POST("/api/vl/user/changephone",handler.Changephone)
	server.GET("/api/v1/user/getInfo", handler.GetInfo, handler.MustLogin)


	server.POST("/api/v1/contact/add", handler.AddAddress, handler.MustLogin)
	server.POST("/api/vl/contact/alter",handler.Alter)
	server.POST("/api/v1/contact/change",handler.ChangeAddress)
	server.POST("/api/v1/contact/getaddress", handler.GetAddress, handler.MustLogin)

	server.POST("/api/v1/products/create",handler.CreateProduct)//创建商品
	server.POST("/api/v1/product/getinfo",handler.GetProInfo,handler.MustLogin)
	server.POST("/api/v1/products/changestatus", handler.ChangeProStatus)
	server.POST("/api/v1/products/getlist", handler.GetProductList)
	server.POST("/api/v1/products/changecate",handler.ChangeCategories)

	server.POST("/api/v1/orders/get", handler.GetOrders, handler.MustLogin)
	server.POST("/api/v1/orders/create", handler.CreateOrder, handler.MustLogin)
	server.POST("/api/v1/orders/getone", handler.GetOneOrder)
	server.POST("/api/v1/orders/changestatus",handler.ChangeStatus)
	server.POST("/api/v1/orders/get", handler.GetOrders, handler.MustLogin)

	server.POST("/api/vl/categories/get", handler.GetCategories)
	server.POST("/api/v1/categories/create",handler.CreateCategories)

	server.POST("/api/v1/carts/delete", handler.Cartsdel, handler.MustLogin)
	server.POST("/api/vl/carts/altercartpro",handler.AlterCartPro)
	server.POST("/api/vl/carts/cartsput",handler.CartsPutIn)
	server.GET("/api/v1/carts/browse", handler.BrowseCart, handler.MustLogin)
}
