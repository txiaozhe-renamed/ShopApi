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
 *     Initial: 2017/07/21     Zhu Yaqiang
 *     Modify : 2017/07/22     Xu Haosheng    添加购物车
 *     Modify : 2017/07/23     Wang Ke
 *     Modify : 2017/07/24     Ma Chao
 */

package handler

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"ShopApi/general"
	"ShopApi/general/errcode"
	"ShopApi/log"
	"ShopApi/models"
	"ShopApi/utility"
)

func CartsPutIn(c echo.Context) error {
	var (
		err   error
		carts models.ConCarts
	)

	if err = c.Bind(&carts); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	userID := session.Get(general.SessionUserID)
	id := userID.(uint64)

	err = models.CartsService.CreateInCarts(&carts, id)
	if err != nil {
		log.Logger.Error("Mysql error in add address:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func Cartsdel(c echo.Context) error {
	var (
		err  error
		cart models.ConCarts
	)

	if err = c.Bind(&cart); err != nil {
		log.Logger.Error("Analysis crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.CartsService.CartsDelete(cart.ID, cart.ProductID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("This product doesn't exist !", err)

			return general.NewErrorWithMessage(errcode.ErrInformation, err.Error())
		}

		log.Logger.Error("Delete product with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func AlterCartPro(c echo.Context) error {
	var (
		err     error
		cartpro models.ConCarts
	)

	if err = c.Bind(&cartpro); err != nil {
		log.Logger.Error("Get crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.CartsService.AlterCartPro(cartpro.ID, cartpro.Count)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("This product doesn't exist !", err)

			return general.NewErrorWithMessage(errcode.ErrInformation, err.Error())
		}

		log.Logger.Error("Alter product with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func BrowseCart(c echo.Context) error {
	var (
		err    error
		output []models.ConCarts
	)

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	userID := session.Get(general.SessionUserID).(uint64)

	output, err = models.CartsService.BrowseCart(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("Find order with error:", err)

			return general.NewErrorWithMessage(errcode.ErrInformation, err.Error())
		}

		log.Logger.Error("Get Order with error:", err)

		return general.NewErrorWithMessage(errcode.ErrOrdersNotFound, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, output)
}
