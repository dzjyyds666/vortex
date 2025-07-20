package vortex

import (
	"context"

	"github.com/labstack/echo/v4"
)

type vortexContext struct {
	ctx          context.Context // 请求的上下文
	echo.Context                 // echo的上下文
}

// GetContext 获取底层的context.Context
func (vc *vortexContext) GetContext() context.Context {
	return vc.Request().Context()
}

// WithValue 在context中设置键值对
func (vc *vortexContext) WithValue(key, value interface{}) {
	vc.ctx = context.WithValue(vc.ctx, key, value)
}

// Value 从context中获取值
func (vc *vortexContext) Value(key interface{}) interface{} {
	return vc.ctx.Value(key)
}

type Status struct {
	respCode int64
	subCode  int64
	i18nKey  string
}
