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
 *     Initial: 2017/07/18        Li Zebang
 *     Modify : 2017/07/20        Yu Yi
 *     Modify : 2017/07/20        Yang Zhengtian
 */

package models

import (
	"time"

	"ShopApi/orm"
)

type ContactServiceProvider struct {
}

var ContactService *ContactServiceProvider = &ContactServiceProvider{}

type Contact struct {
	ID        uint64    `sql:"auto_increment; primary_key;" json:"id"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	Name      string    `json:"name"`
	Phone     string    `json:"phone"`
	Province  string    `json:"province"`
	City      string    `json:"city"`
	Street    string    `json:"street"`
	Address   string    `json:"address"`
	Created   time.Time `json:"created"`
	IsDefault uint8     `gorm:"column:isdefault" json:"isdefault" `
	Page       uint64    `json:"page" validate:"required,numeric"`
	PageSize   uint64    `json:"pagesize" validate:"required,numeric"`
}

type OrmContact struct {
	ID        uint64    `json:"id" validate:"required,numeric"`
	UserID    uint64    `gorm:"column:userid" json:"userid"`
	Name      string    `json:"name" validate:"required,alphanum,min=6,max=100"`
	Phone     string    `json:"phone" validate:"required,numeric,min=6,max=20"`
	Province  string    `json:"province" validate:"required,alphanum,min=6,max=100"`
	City      string    `json:"city" validate:"required,alphanum,min=6,max=100"`
	Street    string    `json:"street" validate:"required,alphanum,min=6,max=100"`
	Address   string    `json:"address" validate:"required,alphanum,min=6,max=200"`
	Created   time.Time `json:"created"`
	IsDefault uint8     `json:"isdefault" validate:"required,numeric"`
	Page       uint64    `json:"page" validate:"required,numeric"`
	PageSize   uint64    `json:"pagesize" validate:"required,numeric"`
}

type AddressGet struct {
	Province string `json:"province"`
	City     string `json:"city"`
	Street   string `json:"street"`
	Address  string `json:"address"`
}

type ChangeAddress struct {
	ID       uint64  `json:"id" validate:"numeric"`
	Name     *string `json:"name" validate:"required, alphaunicode, min=2,max=18"`
	Phone    *string `json:"phone" validate:"required, alphanum, min=6,max=30"`
	Province *string `json:"province" validate:"required, alphaunicode, min=2,max=30"`
	City     *string `json:"city" validate:"required, alphaunicode, min=2,max=30"`
	Street   *string `json:"street" validate:"required, alphaunicode, min=2,max=30"`
	Address  *string `json:"address" validate:"required, alphaunicode, min=2,max=30"`
}

func (Contact) TableName() string {
	return "contact"
}

func (csp *ContactServiceProvider) AddAddress(ormContact *OrmContact) error {
	ormContact.Created = time.Now()

	contact := &Contact{
		UserID:    ormContact.UserID,
		Name:      ormContact.Name,
		Phone:     ormContact.Phone,
		Province:  ormContact.Province,
		City:      ormContact.City,
		Street:    ormContact.Street,
		Address:   ormContact.Address,
		Created:   time.Now(),
		IsDefault: ormContact.IsDefault,
	}

	db := orm.Conn

	return db.Create(contact).Error
}

func (csp *ContactServiceProvider) ChangeAddress(addr OrmContact) error {
	var (
		con Contact
	)

	changeMap := map[string]interface{}{
		"name":     addr.Name,
		"phone":    addr.Phone,
		"province": addr.Province,
		"city":     addr.City,
		"street":   addr.Street,
		"address":  addr.Address,
	}

	db := orm.Conn
	err := db.Model(&con).Where("id = ?", addr.ID).Update(changeMap).Limit(1).Error

	return err
}

func (csp *ContactServiceProvider) FindAddressId(ID uint64) error {
	var (
		con Contact
	)

	db := orm.Conn
	err := db.Where("id = ?", ID).First(&con).Error

	return err
}

func (csp *ContactServiceProvider) GetAddressByUerId(userId uint64, pageStart, pageEnd uint64) ([]AddressGet, error) {
	var (
		list   Contact
		getAdd []AddressGet
	)

	db := orm.Conn
	err := db.Where("userid =? ", userId).Find(&list).Error
	if err != nil {
		return getAdd, err
	}
	sql := "SELECT * FROM contact WHERE userid = ? LIMIT ?, ? LOCK IN SHARE MODE"

	rows, err := db.Raw(sql, userId, pageStart, pageEnd).Rows()
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		db.ScanRows(rows, &list)
			add := AddressGet{
				Province: list.Province,
				City:     list.City,
				Street:   list.Street,
				Address:  list.Address,
			}
			getAdd = append(getAdd, add)
		}

	return getAdd, nil
}

func (csp *ContactServiceProvider) AlterDefault(id uint64) error {
	var (
		s   Contact
		con Contact
	)
	db := orm.Conn
	err := db.Where("id=?", id).Find(&s).Error

	updater := map[string]interface{}{"isdefault": s.IsDefault^1}

	err = db.Model(&con).Where("id=?", id).Update(updater).Limit(1).Error

		return err
}
