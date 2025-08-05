package server

import (
	"context"
	"errors"
	"net/http"

	genericoptions "github.com/onexstack/onexstack/pkg/options"

	"example.com/miniblog/internal/pkg/log"
)

// HTTPServer 代表一个 HTTP 服务器.
type HTTPServer struct {
	srv *http.Server
}

// NewHTTPServer 创建一个新的 HTTP 服务器实例.
func NewHTTPServer(httpOptions *genericoptions.HTTPOptions, handler http.Handler) *HTTPServer {
	return &HTTPServer{
		srv: &http.Server{
			Addr:    httpOptions.Addr,
			Handler: handler,
		},
	}
}

// RunOrDie 启动 HTTP 服务器并在出错时记录致命错误.
func (s *HTTPServer) RunOrDie() {
	log.Infow("Start to listening the incoming requests", "protocol", protocolName(s.srv), "addr", s.srv.Addr)
	if err := s.srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalw("Failed to server HTTP(s) server", "err", err)
	}
}

// GracefulStop 优雅地关闭 HTTP 服务器.
// http.Server 类型的 Shutdown 方法的工作流程如下：​​
// 首先关闭所有已开启的监听器，然后关闭所有空闲连接，最后等待所有活跃连接进入空闲状态后终止服务。
// 如果传入的 ctx 在服务完成终止之前超时，则 Shutdown 方法会返回与 context 相关的错误。否则会返回由关闭服务监听器引发的其他错误。
// 当 Shutdown 方法被调用时，Serve、ListenAndServe 以及 ListenAndServeTLS 方法会立即返回 ErrServerClosed 错误。
// ErrServerClosed 错误被视为服务关闭时的正常行为。因此，如果 ListenAndServe 返回该错误，程序不会打印错误信息。
func (s *HTTPServer) GracefulStop(ctx context.Context) {
	log.Infow("Gracefully stop HTTP(s) server")
	if err := s.srv.Shutdown(ctx); err != nil {
		log.Errorw("HTTP(s) server forced to shutdown", "err", err)
	}
}
