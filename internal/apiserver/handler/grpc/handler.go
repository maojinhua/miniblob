package grpc

import (
	// 重命名为 apiv1 为了和其他包的 v1 包做区分
	apiv1 "example.com/miniblog/pkg/api/apiserver/v1"
)

// Handler 负责处理博客模块的请求
type Handler struct {
	// 内嵌 UnimplementedMiniBlogServer，为了提供默认实现，确保未实现的方法返回“未实现”错误
	// 详细介绍 https://github.com/onexstack/miniblog/blob/feature/s09/docs/book/unimplemented.md
	apiv1.UnimplementedMiniBlogServer
}

// NewHandler 创建一个新的 Handler 实例
func NewHandler() *Handler {
	return &Handler{}
}
