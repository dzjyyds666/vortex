package middleware

import (
	"github.com/dzjyyds666/Allspark-go/logx"

	"github.com/labstack/echo/v4"
)

type VortexHttpMiddleware echo.MiddlewareFunc

// 打印响应的日志
func LogResponse() VortexHttpMiddleware {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			if nil != err {
				return err
			}
			// 打印响应的日志
			logx.Infof("response=>[ %d %s ]", c.Response().Status, c.Request().URL)
			return nil
		}
	}
}

// 打印请求的日志
func LogRequest() VortexHttpMiddleware {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 打印请求的日志
			logx.Infof("request=>[ %s %s ]", c.Request().Method, c.Request().URL)
			return next(c)
		}
	}
}
