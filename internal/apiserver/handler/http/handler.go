// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package http

import (
	"example.com/miniblog/internal/apiserver/biz"
	"example.com/miniblog/internal/apiserver/pkg/validation"
)

// Handler 处理博客模块的请求.
type Handler struct {
	biz biz.IBiz
	val *validation.Validator
}

// NewHandler 创建新的 Handler 实例.
func NewHandler(biz biz.IBiz, val *validation.Validator) *Handler {
	return &Handler{
		biz: biz,
		val: val,
	}
}
