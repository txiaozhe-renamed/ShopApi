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
 *     Initial: 2017/07/19        Yusan Kurban
 */

package general

const (
	// User Type
	PhoneUser  = 0xff
	WechatUser = 0xfe

	// User Status
	UserActive   = 0xf0
	UserInactive = 0xf1

	// Login session
	SessionUserID = "userid"

	// sex
	Man   = 0x1
	Woman = 0x2

	//Address  Default
	Default   = 1
	Undefault = 0

	//Products Status
	ProductOnsale = 0xe0 // 224
	ProductUnsale = 0xe1 // 225
	ProductNotInCart = 1

	//Categories Status
	CategoriesOnuse = 0xa0 // 160
	CategoriesUnuse = 0xa1 // 161

	// Order Status
	OrderUnfinished = 0xef // 239
	OrderFinished   = 0xee // 238
	OrderGetAll		= 0xed // 237, Not order status
	OrderCanceled   = 0xec // 236
)
