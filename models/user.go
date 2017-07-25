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
 *     Modify: 2017/07/21         Xu Haosheng    更改用户信息
 *     Modify: 2017/07/20	      Zhang Zizhao   登录检查
 *     Modify: 2017/07/21         Yang Zhengtian 添加判断用户是否存在和修改密码
 */

package models

import (
	"time"

	"ShopApi/general"
	"ShopApi/orm"
	"ShopApi/utility"
)

type UserServiceProvider struct {
}

var UserService *UserServiceProvider = &UserServiceProvider{}

type User struct {
	UserID   uint64    `sql:"auto_increment;primary_key;" gorm:"column:id" json:"userid"`
	OpenID   string    `gorm:"column:openid" json:"openid"`
	Name     string    `json:"name"`
	Password string    `json:"password"`
	Status   uint16    `json:"status"`
	Type     uint16    `json:"type"`
	Created  time.Time `json:"created"`
}

type UserInfo struct {
	UserID   uint64 `sql:"primary_key" gorm:"column:userid" json:"userid"`
	Avatar   string `json:"avatar"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Sex      uint8  `json:"sex"`
}

//todo：连接前端
type ConUsers struct {
	UserID   uint64    `gorm:"column:id" json:"userid"`
	OpenID   string    `gorm:"column:openid" json:"openid"`
	Name     string    `json:"name"`
	Status   uint16    `json:"status"`
	Type     uint16    `json:"type"`
	Created  time.Time `json:"created"`
	Avatar   string    `json:"avatar"`
	Nickname string    `json:"nickname"`
	Email    string    `json:"email"`
	Phone    string    `json:"phone"`
	Sex      uint8     `json:"sex"`
	Pass    *string `json:"pass" validate:"required,alphanum,min=6,max=30"`
	NewPass *string `json:"newpass" validate:"required,alphanum,min=6,max=30"`
}

func (User) TableName() string {
	return "users"
}

func (UserInfo) TableName() string {
	return "userinfo"
}

func (us *UserServiceProvider) Create(name, pass *string) error {
	hashedPass, err := utility.GenerateHash(*pass)
	if err != nil {
		return err
	}

	u := User{
		Name:     *name,
		Password: string(hashedPass),
		Status:   general.UserActive,
		Type:     general.PhoneUser,
		Created:  time.Now(),
	}

	db := orm.Conn

	tx := db.Begin()
	defer func() {
		if err != nil {
			err = tx.Rollback().Error
		} else {
			err = tx.Commit().Error
		}
	}()

	err = tx.Create(&u).Error
	if err != nil {
		return err
	}

	info := UserInfo{
		UserID: u.UserID,
		Phone:  *name,
		Sex:    general.Man,
	}

	err = tx.Create(&info).Error
	if err != nil {
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

// todo: 代码风格
func (us *UserServiceProvider) Login(name, pass *string) (bool, uint64, error) {
	var (
		u   User
		err error
	)

	db := orm.Conn
	err = db.Where("name = ?", *name).First(&u).Error
	if err != nil {

		return false, 0, err
	}

	if !utility.CompareHash([]byte(u.Password), *pass) {

		return false, 0, nil
	}

	return true, u.UserID, nil
}

func (us *UserServiceProvider) GetInfo(UserID uint64) (*UserInfo, error) {

	var (
		err error
		ui  *UserInfo = &UserInfo{}
	)

	db := orm.Conn
	err = db.Where("userid = ?", UserID).First(&ui).Error
	if err != nil {
		return ui, err
	}

	return ui, nil
}

func (us *UserServiceProvider) ChangePhone(UserID uint64, Phone string) error {
	var (
		err error
		con Contact
	)

	change := map[string]interface{}{"phone": Phone}

	db := orm.Conn
	err = db.Model(&con).Where("id=?", UserID).Update(change).Limit(1).Error

		return err
}

func (us *UserServiceProvider) GetUerPassword(id uint64) (string, error) {
	var (
		user User
		err  error
	)

	db := orm.Conn
	err = db.Where("id = ?", id).First(&user).Error

	return user.Password, err
}

func (us *UserServiceProvider) ChangeMobilePassword(newPass *string, id uint64) error {
	var (
		user User
		err  error
	)

	db := orm.Conn
	hashPass, err := utility.GenerateHash(*newPass)
	if err != nil {
		return err
	}

	updater := map[string]interface{}{"password": hashPass}
	err = db.Model(&user).Where("id =? ", id).Update(updater).Limit(1).Error

	return err
}

func (us *UserServiceProvider) ChangeUserInfo(info *UserInfo, userID uint64) error {
	var (
		con Contact
		nil uint8 = 0
	)

	changMap := map[string]interface{}{
		"nickname": info.Nickname,
		"email"       : info.Email,
		"phone"     : info.Phone,
		"sex"           : info.Sex,
		"avatar"      : info.Avatar,
	}

	db := orm.Conn

	for userInfo, value := range changMap {
		if value =="" || value == nil {
			delete(changMap, userInfo)
		}
	}
	err := db.Model(&con).Where("userid = ?", userID).Updates(changMap).Limit(1).Error

	return err
}
