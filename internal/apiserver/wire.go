// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

//go:build wireinject
// +build wireinject

package apiserver

import (
	"github.com/google/wire"

	"example.com/miniblog/internal/apiserver/biz"
	"example.com/miniblog/internal/apiserver/pkg/validation"
	"example.com/miniblog/internal/apiserver/store"
	ginmw "example.com/miniblog/internal/pkg/middleware/gin"
	"example.com/miniblog/internal/pkg/server"
	"example.com/miniblog/pkg/auth"
)

func InitializeWebServer(*Config) (server.Server, error) {
	wire.Build(
		wire.NewSet(NewWebServer, wire.FieldsOf(new(*Config), "ServerMode")),
		wire.Struct(new(ServerConfig), "*"), // * 表示注入全部字段
		wire.NewSet(store.ProviderSet, biz.ProviderSet),
		ProvideDB, // 提供数据库实例
		validation.ProviderSet,
		wire.NewSet(
			wire.Struct(new(UserRetriever), "*"),
			wire.Bind(new(ginmw.UserRetriever), new(*UserRetriever)),
		),
		auth.ProviderSet,
	)
	return nil, nil
}
