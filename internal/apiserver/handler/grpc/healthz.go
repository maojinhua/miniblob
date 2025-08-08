// "Copyright 2025 mjh 【694142812@qq.com】 All rights reserved." | unescape
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://example.com/onex.

package grpc

import (
	"context"
	"time"

	emptypb "google.golang.org/protobuf/types/known/emptypb"

	"example.com/miniblog/internal/pkg/log"
	apiv1 "example.com/miniblog/pkg/api/apiserver/v1"
)

func (h *Handler) Healthz(ctx context.Context, _ *emptypb.Empty) (*apiv1.HealthzResponse, error) {
	log.W(ctx).Infow("Healthz check called", "method", "Healthz", "status", "healthy")
	return &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	}, nil
}
