// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

package store

//go:generate mockgen -destination mock_store.go -package store example.com/miniblog/internal/miniblog/store IStore,UserStore,PostStore,ConcretePostStore

import (
	"context"
	"github.com/google/wire"
	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"
	"sync"
)


// ProviderSet 是一个 Wire 的 Provider 集合，用于声明依赖注入的规则.
// 包含 NewStore 构造函数，用于生成 datastore 实例.
// wire.Bind 用于将接口 IStore 与具体实现 *datastore 绑定，
// 从而在依赖 IStore 的地方，能够自动注入 *datastore 实例.
var ProviderSet = wire.NewSet(NewStore, wire.Bind(new(IStore), new(*datastore)))

var (
	once sync.Once
	// 全局变量，方便其它包直接调用已初始化好的 datastore 实例.
	S *datastore
)

// IStore 定义了 Store 层需要实现的方法.
type IStore interface {
	// 返回 Store 层的 *gorm.DB 实例(也就是 S 变量)，在少数场景下会被用到.
	DB(ctx context.Context, wheres ...where.Where) *gorm.DB
	TX(ctx context.Context, fn func(ctx context.Context) error) error

	// IStore 接口中的 User() 方法和 Post() 方法，分别返回 User 资源的 Store 层方法和 Post 资源的 Store 层方法。
	// 这种设计方式其实是软件开发中的抽象工厂模式。
	// 通过使用不同的方法，创建不同的对象，对象之间相互独立，以此提高代码的可维护性。
	// 另外使用这种方法，也能有效提高资源接口的标准化。标准化的资源接口更易理解和使用。​
	User() UserStore
	Post() PostStore
}

// transactionKey 用于在 context.Context 中存储事务上下文的键.
type transactionKey struct{}

// datastore 是 IStore 的具体实现.
type datastore struct {
	core *gorm.DB

	// 可以根据需要添加其他数据库实例
	// fake *gorm.DB
}

// 确保 datastore 实现了 IStore 接口.
var _ IStore = (*datastore)(nil)

// NewStore 创建一个 IStore 类型的实例.
// 上述代码使用了 sync.Once来确保实例只被初始化一次。
// 实例创建完成后，将其赋值给包级变量 S，以便通过 store.S.User().Create()来调用 Store 层接口。
// 在 Go 的最佳实践中，建议尽量减少使用包级变量，因为包级变量的状态通常难以感知，会增加维护的复杂度。
// 然而，这并非绝对规则，可以根据实际需要选择是否设置包级变量来简化开发。
func NewStore(db *gorm.DB) *datastore {
	// 确保 S 只被初始化一次
	once.Do(func() {
		S = &datastore{db}
	})

	return S
}

// DB 根据传入的条件（wheres）对数据库实例进行筛选.
// 如果未传入任何条件，则返回上下文中的数据库实例（事务实例或核心数据库实例）.
func (store *datastore) DB(ctx context.Context, wheres ...where.Where) *gorm.DB {
	db := store.core
	// 从上下文中提取事务实例
	if tx, ok := ctx.Value(transactionKey{}).(*gorm.DB); ok {
		db = tx
	}

	// 遍历所有传入的条件并逐一叠加到数据库查询对象上
	for _, whr := range wheres {
		db = whr.Where(db)
	}
	return db
}

// TX 返回一个新的事务实例.
// TX 方法将 *gorm.DB 类型的实例注入到 context 中
// nolint: fatcontext
func (store *datastore) TX(ctx context.Context, fn func(ctx context.Context) error) error {
	return store.core.WithContext(ctx).Transaction(
		func(tx *gorm.DB) error {
			ctx = context.WithValue(ctx, transactionKey{}, tx)
			return fn(ctx)
		},
	)
}

// Users 返回一个实现了 UserStore 接口的实例.
func (store *datastore) User() UserStore {
	return newUserStore(store)
}

// Posts 返回一个实现了 PostStore 接口的实例.
func (store *datastore) Post() PostStore {
	return newPostStore(store)
}
