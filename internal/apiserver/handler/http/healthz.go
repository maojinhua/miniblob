package http

import (
	"time"

	"github.com/gin-gonic/gin"

	apiv1 "example.com/miniblog/pkg/api/apiserver/v1"
)

// Healthz 服务健康检查.
// 在企业级应用开发中，一个合格的 Web 服务需要提供健康检查接口，以供其他服务探测 Web 服务的健康状态。
func (h *Handler) Healthz(c *gin.Context) {
	// 返回 JSON 响应
	c.JSON(200, &apiv1.HealthzResponse{
		Status:    apiv1.ServiceStatus_Healthy,
		Timestamp: time.Now().Format(time.DateTime),
	})
}
