package vortex

import (
	"context"

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
	respCode int64
	subCode  int64
	i18nKey  string
}
