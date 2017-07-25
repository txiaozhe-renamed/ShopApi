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
 *     Initial: 2017/07/21       Li Zebang
 *     Modify : 2017/07/21       Zhang Zizhao 添加创建订单
 *	   Modify : 2017/07/21       Ai Hao       订单状态更改
 */

package handler

import (
	"errors"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"ShopApi/general"
	"ShopApi/general/errcode"
	"ShopApi/log"
	"ShopApi/models"
	"ShopApi/utility"
)

type ChangStatus struct {
	ID     uint64 `json:"id"`
	Status uint8  `json:"status"`
}

func CreateOrder(c echo.Context) error {
	var (
		order models.RegisterOrder
		err   error
	)

	if err = c.Bind(&order); err != nil {
		log.Logger.Error("Create crash with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	sess := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	numberID := sess.Get(general.SessionUserID).(uint64)

	err = models.OrderService.CreateOrder(numberID, order)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("Product not found:", err)

			return general.NewErrorWithMessage(errcode.ErrMysqlfound, err.Error())
		}
		log.Logger.Error("Mysql error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func GetOrders(c echo.Context) error {
	var (
		err    error
		orm    models.OrmOrders
		orders *[]models.Orders
	)

	if err = c.Bind(&orm); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	if orm.Status != general.OrderUnfinished && orm.Status != general.OrderFinished && orm.Status != general.OrderGetAll {
		err = errors.New("Invalid Orders Status")

		log.Logger.Error("Error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidOrdersStatus, err.Error())
	}

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	userID := session.Get(general.SessionUserID).(uint64)

	pageStart, pageEnd := utility.Paging(orm.Page, orm.PageSize)

	orders, err = models.OrderService.GetOrders(userID, orm.Status, pageStart, pageEnd)
	if err != nil {
		log.Logger.Error("Mysql error in GetOrders Function:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	if len(*orders) == 0 {
		err = errors.New("Orders Not Found")

		log.Logger.Error("Error:", err)

		return general.NewErrorWithMessage(errcode.ErrOrdersNotFound, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, orders)
}

func GetOneOrder(c echo.Context) error {
	var (
		err    error
		order  *models.OrmOrders
		OutPut *models.OrmOrders
	)

	if err = c.Bind(&order); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	UserID := session.Get(general.SessionUserID).(uint64)

	OutPut, err = models.OrderService.GetOneOrder(order.ID, UserID)

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("Find order with error:", err)

			return general.NewErrorWithMessage(errcode.ErrInformation, err.Error())
		}

		log.Logger.Error("Get Order with error:", err)

		return general.NewErrorWithMessage(errcode.ErrOrdersNotFound, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, OutPut)
}

func ChangeStatus(c echo.Context) error {
	var (
		err error
		st  ChangStatus
	)

	if err = c.Bind(&st); err != nil {
		log.Logger.Error("Input order status with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	if st.Status != general.OrderFinished && st.Status != general.OrderUnfinished && st.Status != general.OrderCanceled {
		err = errors.New("Status unExistence")
		log.Logger.Error("", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}
	err = models.OrderService.ChangeStatus(st.ID, st.Status)
	if err != nil {
		log.Logger.Error("Change status with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}
