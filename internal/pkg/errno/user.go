package errno

import (
	"net/http"

	"example.com/miniblog/internal/pkg/errorsx"
)

var (
	// ErrUsernameInvalid 表示用户名不合法.
	ErrUsernameInvalid = &errorsx.ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "InvalidArgument.UsernameInvalid",
		Message: "Invalid username: Username must consist of letters, digits, and underscores only, and its length must be between 3 and 20 characters.",
	}

	// ErrPasswordInvalid 表示密码不合法.
	ErrPasswordInvalid = &errorsx.ErrorX{
		Code:    http.StatusBadRequest,
		Reason:  "InvalidArgument.PasswordInvalid",
		Message: "Password is incorrect.",
	}

	// ErrUserAlreadyExists 表示用户已存在.
	ErrUserAlreadyExists = &errorsx.ErrorX{Code: http.StatusBadRequest, Reason: "AlreadyExist.UserAlreadyExists", Message: "User already exists."}

	// ErrUserNotFound 表示未找到指定用户.
	ErrUserNotFound = &errorsx.ErrorX{Code: http.StatusNotFound, Reason: "NotFound.UserNotFound", Message: "User not found."}
)
