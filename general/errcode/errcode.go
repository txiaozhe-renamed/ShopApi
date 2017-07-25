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
 *     Initial: 2017/05/14        Feng Yifei
 */

package errcode

const (
	ErrInformation         = 0x4
	ErrSucceed             = 0x0
	ErrInvalidParams       = 0x1
	ErrMysql               = 0x2
	ErrDelete              = 0x3 //用户登出错误
	ErrMysqlfound          = 0x4
	ErrNameFormat          = 0x5
	ErrGetsess             = 0x6
	ErrInvalidOrdersStatus = 0x7
	ErrOrdersNotFound      = 0x8
	NoOrder                = 0x11
	ErrCategoriesNotFound  = 0x9
	ErrNotFound            = 0xa
	ErrAccess              = 0xb
	ErrInvalidPhone        = 0xc
	ErrInput		 =0xd

	// 需要登录
	ErrLoginRequired    = 0x800
	ErrPermissionDenied = 0x801

	// 严重错误
	ErrNoConnection      = 0x1000
	ErrDBOperationFailed = 0x1001
)
