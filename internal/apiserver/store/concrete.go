// Copyright 2024 孔令飞 <colin404@foxmail.com>. All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://github.com/onexstack/onex.

// nolint: dupl
package store

import (
	"context"
	"errors"

	"github.com/onexstack/onexstack/pkg/store/where"
	"gorm.io/gorm"

	"example.com/miniblog/internal/apiserver/model"
	"example.com/miniblog/internal/pkg/errno"
	"example.com/miniblog/internal/pkg/log"
)

// ConcretePostStore 定义了 post 模块在 store 层所实现的方法.
// 这么做是为了保留之前实现的方法，用来和 genericstore.Store 做对比。实际上只要有 genericstore.Store 的方法就可以了。
// 外部要用 b.store.ConcretePost() 来使用
type ConcretePostStore interface {
	Create(ctx context.Context, obj *model.PostM) error
	Update(ctx context.Context, obj *model.PostM) error
	Delete(ctx context.Context, opts *where.Options) error
	Get(ctx context.Context, opts *where.Options) (*model.PostM, error)
	List(ctx context.Context, opts *where.Options) (int64, []*model.PostM, error)

	ConcretePostExpansion
}

// ConcretePostExpansion 定义了帖子操作的附加方法.
type ConcretePostExpansion interface{}

// concretePostStore 是 ConcretePostStore 接口的实现.
type concretePostStore struct {
	store *datastore
}

// 确保 concretePostStore 实现了 ConcretePostStore 接口.
var _ ConcretePostStore = (*concretePostStore)(nil)

// newConcretePostStore 创建 concretePostStore 的实例.
func newConcretePostStore(store *datastore) *concretePostStore {
	return &concretePostStore{store}
}

// Create 插入一条帖子记录.
func (s *concretePostStore) Create(ctx context.Context, obj *model.PostM) error {
	if err := s.store.DB(ctx).Create(&obj).Error; err != nil {
		log.Errorw("Failed to insert post into database", "err", err, "post", obj)
		return errno.ErrDBWrite.WithMessage("database create failed: %v", err)
	}

	return nil
}

// Update 更新帖子数据库记录.
func (s *concretePostStore) Update(ctx context.Context, obj *model.PostM) error {
	if err := s.store.DB(ctx).Save(obj).Error; err != nil {
		log.Errorw("Failed to update post in database", "err", err, "post", obj)
		return errno.ErrDBWrite.WithMessage("database update failed: %v", err)
	}

	return nil
}

// Delete 根据条件删除帖子记录.
func (s *concretePostStore) Delete(ctx context.Context, opts *where.Options) error {
	err := s.store.DB(ctx, opts).Delete(new(model.PostM)).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Errorw("Failed to delete post from database", "err", err, "conditions", opts)
		return errno.ErrDBWrite.WithMessage("database delete failed: %v", err)
	}

	return nil
}

// Get 根据条件查询帖子记录.
func (s *concretePostStore) Get(ctx context.Context, opts *where.Options) (*model.PostM, error) {
	var obj model.PostM
	if err := s.store.DB(ctx, opts).First(&obj).Error; err != nil {
		log.Errorw("Failed to retrieve post from database", "err", err, "conditions", opts)
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrPostNotFound
		}
		return nil, errno.ErrDBRead.WithMessage("%v", err.Error())
	}

	return &obj, nil
}

// List 返回帖子列表和总数.
// nolint: nonamedreturns
func (s *concretePostStore) List(ctx context.Context, opts *where.Options) (count int64, ret []*model.PostM, err error) {
	err = s.store.DB(ctx, opts).Order("id desc").Find(&ret).Offset(-1).Limit(-1).Count(&count).Error
	if err != nil {
		log.Errorw("Failed to list posts from database", "err", err, "conditions", opts)
		err = errno.ErrDBRead.WithMessage("%v", err.Error())
	}
	return
}
