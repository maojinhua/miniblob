// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://github.com/onexstack/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package http

// Handler 处理博客模块的请求.
type Handler struct{}

// NewHandler 创建新的 Handler 实例.
func NewHandler() *Handler {
	return &Handler{}
}
