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
 *     Initial: 2017/07/19       Li Zebang
 *     Modify : 2017/07/20       Yu Yi
 *     Modify : 2017/07/20       Yang Zhengtian
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

func AddAddress(c echo.Context) error {
	var (
		err     error
		contact models.OrmContact
	)

	if err = c.Bind(&contact); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	contact.UserID = session.Get(general.SessionUserID).(uint64)

	err = models.ContactService.AddAddress(&contact)
	if err != nil {
		log.Logger.Error("Mysql error in add address:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func ChangeAddress(c echo.Context) error {
	var (
		err  error
		addr models.OrmContact
	)

	if err = c.Bind(&addr); err != nil {
		log.Logger.Error("Bind change with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.ContactService.FindAddressId(addr.ID)
	if err != nil {

		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("Address id not find", err)

			return general.NewErrorWithMessage(errcode.ErrNotFound, err.Error())
		}

		log.Logger.Error("Mysql error", err)

		return general.NewErrorWithMessage(errcode.ErrNotFound, err.Error())
	}

	err = models.ContactService.ChangeAddress(addr)
	if err != nil {
		log.Logger.Error("Change address with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, nil)
}

func GetAddress(c echo.Context) error {
	var (
		err   	 error
		userId   uint64
		address  models.OrmContact
		list     []models.AddressGet
	)

	if err = c.Bind(&address); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	session := utility.GlobalSessions.SessionStart(c.Response().Writer, c.Request())
	s := session.Get(general.SessionUserID)
	userId = s.(uint64)

	pageStart, pageEnd := utility.Paging(address.Page, address.PageSize)
	list, err = models.ContactService.GetAddressByUerId(userId, pageStart, pageEnd)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Logger.Error("Id not find:", err)

			return general.NewErrorWithMessage(errcode.ErrNotFound, err.Error())
		}
		log.Logger.Error("Mysql err", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}

	return c.JSON(errcode.ErrSucceed, list)
}

func Alter(c echo.Context) error {
	var (
		err error
		m   models.Contact
	)

	if err = c.Bind(&m); err != nil {
		log.Logger.Error("Bind with error:", err)

		return general.NewErrorWithMessage(errcode.ErrInvalidParams, err.Error())
	}

	err = models.ContactService.AlterDefault(m.ID)
	if err != nil {
		log.Logger.Error("Alter Default with error:", err)

		return general.NewErrorWithMessage(errcode.ErrMysql, err.Error())
	}
	return c.JSON(errcode.ErrSucceed, nil)
}
