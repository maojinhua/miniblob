// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package http

import (
	"github.com/gin-gonic/gin"
	"github.com/onexstack/onexstack/pkg/core"
)

// 在语法题函数 HandleJSONRequest，HandleQueryRequest，HandleUriRequest 调用了 HandleRequest 函数，再调用 HandleRequest 时，
// 传入了对应的 gin.Context 类型解析请求的方法（c.ShouldBindJSON, c.ShouldBindQuery, c.ShouldBindUri）。
// 在 HandlerRequest 函数中，会先调用 ReadRequest 从请求中解析出请求参数，所有的请求信息都保存在 *gin.Context 类型的变量 c 中。
// 解析完请求之后会调用传入的 Handler 方法，进行请求处理。
// 比如说 Login 函数，会调用 h.biz.UserV1().Login 方法，将解析的请求参数传入，进行业务处理，并返回结果。


// Login 用户登录并返回 JWT Token.
func (h *Handler) Login(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().Login)
}

// RefreshToken 刷新 JWT Token.
func (h *Handler) RefreshToken(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().RefreshToken)
}

// ChangeUserPassword 修改用户密码.
func (h *Handler) ChangePassword(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().ChangePassword)
}

// CreateUser 创建新用户.
func (h *Handler) CreateUser(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().Create)
}

// UpdateUser 更新用户信息.
func (h *Handler) UpdateUser(c *gin.Context) {
	core.HandleJSONRequest(c, h.biz.UserV1().Update)
}

// DeleteUser 删除用户.
func (h *Handler) DeleteUser(c *gin.Context) {
	core.HandleUriRequest(c, h.biz.UserV1().Delete)
}

// GetUser 获取用户信息.
func (h *Handler) GetUser(c *gin.Context) {
	core.HandleUriRequest(c, h.biz.UserV1().Get)
}

// ListUser 列出用户信息.
func (h *Handler) ListUser(c *gin.Context) {
	core.HandleQueryRequest(c, h.biz.UserV1().List)
}