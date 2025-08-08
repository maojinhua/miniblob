// "Copyright 2025 mjh 【694142812@qq.com】 All rights reserved." | unescape
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file. The original repo for
// this file is https://example.com/miniblog. The professional
// version of this repository is https://example.com/onex.

package errno

import (
	"net/http"

	"example.com/miniblog/internal/pkg/errorsx"
)

// ErrPostNotFound 表示未找到指定的博客.
var ErrPostNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.PostNotFound", Message: "Post not found."}
