package vortex

import (
	"context"
	"net/http"

	"github.com/dzjyyds666/vortex/v2/locale"
	"github.com/labstack/echo/v4"
)

type Context struct {
	echo.Context // echo的上下文
}

// GetContext 获取底层的context.Context
func (vc *Context) GetContext() context.Context {
	return vc.Request().Context()
}

// Value 从context中获取值
func (vc *Context) Value(key interface{}) interface{} {
	return vc.Request().Context().Value(key)
}

// 从请求中获取到session信息
func (vc *Context) GetSessionPayload() *Session {
	session := vc.Get(HttpHeaderEnum.Session.XString())
	if session == nil {
		return nil
	}
	return session.(*Session)
}

type Status struct {
	RespCode int64
	subCode  int64
	I18nKey  string
}

func (s Status) WithSubCode(code SubCode) Status {
	if code.SubCode != 0 {
		s.subCode = code.SubCode
	}
	s.I18nKey = code.I18nKey
	return s
}

type SubCode struct {
	SubCode int64
	I18nKey string
}

var Statuses = struct {
	Success        Status // 请求成功
	ParamsInvaild  Status // 参数错误
	PermissionDeny Status // 权限不足
	UnAuthorized   Status // 未授权或token无效
	InternalError  Status // 系统内部故障
}{
	Success:        Status{RespCode: 200, I18nKey: locale.K.CODE_FOR_SUCCESS},
	ParamsInvaild:  Status{RespCode: 400, I18nKey: locale.K.CODE_FOR_PARAMS_INVAILD},
	PermissionDeny: Status{RespCode: 403, I18nKey: locale.K.CODE_FOR_PERMISSION_DENY},
	UnAuthorized:   Status{RespCode: 401, I18nKey: locale.K.CODE_FOR_UNAUTHORIZED},
	InternalError:  Status{RespCode: 500, I18nKey: locale.K.CODE_FOR_INTERNAL_ERROR},
}

type HttpHeaderOption func(header http.Header)
