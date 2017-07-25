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
 *	   Modify : 2017/07/21		 Ai Hao       订单状态更改
 *	   Modify : 2017/07/21		 Zhang Zizhao 创建订单
 */

package models

import (
	"time"

	"ShopApi/general"
	"ShopApi/orm"
	"fmt"
)

type OrderServiceProvider struct {
}

var OrderService *OrderServiceProvider = &OrderServiceProvider{}

// todo：参数检查 结构
type Orders struct {
	ID         uint64    `sql:"auto_increment;primary_key;" json:"id"`
	UserID     uint64    `gorm:"column:userid" json:"userid"`
	TotalPrice float64   `gorm:"column:totalprice"json:"totalprice"`
	Payment    float64   `json:"payment"`
	Freight    float64   `json:"freight"`
	Remark     string    `json:"remark"`
	Discount   uint8     `json:"discount"`
	Size       string    `json:"size"`
	Color      string    `json:"color"`
	Status     uint8     `json:"status"`
	Created    time.Time `json:"created"`
	PayWay     uint8     `gorm:"column:payway"json:"payway"`
}

type OrmOrders struct {
	ID         uint64    `json:"id" validate:"required,numeric"`
	UserID     uint64    `json:"userid" validate:"required,numeric"`
	TotalPrice float64   `json:"totalprice" validate:"required,numeric"`
	Payment    float64   `json:"payment" validate:"required,numeric"`
	Freight    float64   `json:"freight" validate:"required,numeric"`
	Remark     string    `json:"remark" validate:"alphanum"`
	Discount   uint8     `json:"discount" validate:"numeric"`
	Size       string    `json:"size" validate:"required,alphanum"`
	Color      string    `json:"color" validate:"required,alphanum"`
	Status     uint8     `json:"status" validate:"required,numeric"`
	Created    time.Time `json:"created"`
	PayWay     uint8     `json:"payway" validate:"required,numeric"`
	Page       uint64    `json:"page" validate:"required,numeric"`
	PageSize   uint64    `json:"pagesize" validate:"required,numeric"`
}

type RegisterOrder struct {
	Name       string  `json:"productname"`
	TotalPrice float64 `json:"totalprice"`
	Payment    float64 `json:"payment"`
	Freight    float64 `json:"freight"`
	Remark     string  `json:"remark"`
	Discount   uint8   `json:"discount"`
	Size       string  `json:"size"`
	Color      string  `json:"color"`
	Payway     uint8   `json:"payway"`
}

func (Orders) TableName() string {
	return "orders"
}

// todo：命名
func (osp *OrderServiceProvider) CreateOrder(numberID uint64, o RegisterOrder) error {
	var (
		pro Product
		err error
	)

	db := orm.Conn
	err = db.Where("name = ? AND size = ? AND color = ?", o.Name, o.Size, o.Color).Find(&pro).Error
	if err != nil {
		return err
	}

	order := Orders{
		UserID:     numberID,
		TotalPrice: o.TotalPrice,
		Payment:    o.Payment,
		Freight:    o.Freight,
		Remark:     o.Remark,
		Discount:   o.Discount,
		Size:       o.Size,
		Color:      o.Color,
		Status:     general.OrderFinished,
		Created:    time.Now(),
		PayWay:     o.Payway,
	}

	err = db.Create(&order).Error
	if err != nil {
		return err
	}

	return nil
}

func (osp *OrderServiceProvider) GetOrders(userID uint64, status uint8, pageStart, pageEnd uint64) (*[]Orders, error) {
	var (
		order  Orders
		orders []Orders
	)

	db := orm.Conn

	if status == general.OrderUnfinished || status == general.OrderFinished {
		sql := fmt.Sprintf("SELECT * FROM orders WHERE userid = ? AND status = ? LIMIT %d, %d LOCK IN SHARE MODE", pageStart, pageEnd)

		rows, err := db.Raw(sql, userID, status).Rows()
		defer rows.Close()
		if err != nil {
			return nil, err
		}

		for rows.Next() {
			db.ScanRows(rows, &order)
			orders = append(orders, order)
		}

		return &orders, nil
	}

	sql := fmt.Sprintf("SELECT * FROM orders WHERE userid = ? LIMIT %d, %d LOCK IN SHARE MODE", pageStart, pageEnd)

	rows, err := db.Raw(sql, userID).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		db.ScanRows(rows, &order)
		orders = append(orders, order)
	}

	return &orders, nil
}


func (osp *OrderServiceProvider) GetOneOrder(ID uint64, UserID uint64) (*OrmOrders, error) {
	var (
		err      error
		order    Orders
		getOrder *OrmOrders = &OrmOrders{}
	)

	db := orm.Conn
	err = db.Where("userid = ? AND id = ?", UserID, ID).Find(&order).Error
	if err != nil {
		return getOrder, err
	}

	*getOrder = OrmOrders{
		TotalPrice: order.TotalPrice,
		Payment:    order.Payment,
		Freight:    order.Freight,
		Discount:   order.Discount,
		Size:       order.Size,
		Color:      order.Color,
		Status:     order.Status,
		Created:    order.Created,
	}

	return getOrder, nil
}

func (osp *OrderServiceProvider) ChangeStatus(id uint64, status uint8) error {
	cha := Orders{
		Status: status,
	}

	updater := map[string]interface{}{"status": status}
	db := orm.Conn

	err := db.Model(&cha).Where("id=?", id).Update(updater).Limit(1).Error
	return err
}
