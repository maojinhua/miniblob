// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

// nolint: dupl
package store

import (
	"context"

	genericstore "github.com/onexstack/onexstack/pkg/store"
	"github.com/onexstack/onexstack/pkg/store/where"

	"example.com/miniblog/internal/apiserver/model"
)

// UserStore 定义了 user 模块在 store 层所实现的方法.
// UserStore 接口中的方法分为两大类：资源标准 CURD 方法和扩展方法。
// 通过将方法分类，可以进一步提高接口中方法的可维护性和代码的可读性。
// 将方法分为标准方法和扩展方法的开发技巧，在 Kuberentes client-go 项目中被大量使用，miniblog 的设计思路正是来源于 client-go 的设计。

// 在 Go 项目开发中，我习惯将资源标准方法中的方法按固定的接口顺序排序定义：Create、Update、Delete、Get、List。
// 在 miniblog 项目中，Store 层其他资源的方法定义、Biz 层的 UserBiz 接口、PostBiz 接口的方法定义以及 Handler 层的方法定义、
// Protobuf 文件中的服务接口定义、请求/返回参数定义等，均采用了一致的方法顺序。
// 通过将方法顺序标准化，可以在一定程度上提高代码的可阅读性和可维护性。
type UserStore interface {
	Create(ctx context.Context, obj *model.UserM) error
	Update(ctx context.Context, obj *model.UserM) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.UserM, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.UserM, error)

	UserExpansion
}

// UserExpansion 定义了用户操作的附加方法.
// nolint: iface
type UserExpansion interface{}

// userStore 是 UserStore 接口的实现.
type userStore struct {
	*genericstore.Store[model.UserM]
}

// 确保 userStore 实现了 UserStore 接口.
var _ UserStore = (*userStore)(nil)

// newUserStore 创建 userStore 的实例.
func newUserStore(store *datastore) *userStore {
	return &userStore{
		Store: genericstore.NewStore[model.UserM](store, NewLogger()),
	}
}
