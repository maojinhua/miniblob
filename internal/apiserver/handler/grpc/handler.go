// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package grpc

import (
	"example.com/miniblog/internal/apiserver/biz"
	apiv1 "example.com/miniblog/pkg/api/apiserver/v1"
)

// Handler 负责处理博客模块的请求.
type Handler struct {
	apiv1.UnimplementedMiniBlogServer

	// Handler 层代码通过调用 Biz 层的方法来完成业务逻辑的处理。所以 Handler 结构体中包含了 biz.IBiz 类型的字段 biz。
	biz biz.IBiz
}

// NewHandler 创建一个新的 Handler 实例.
func NewHandler(biz biz.IBiz) *Handler {
	return &Handler{
		biz: biz,
	}
}
