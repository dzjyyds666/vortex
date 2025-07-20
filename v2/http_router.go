package vortex

import (
	"time"

	"github.com/dzjyyds666/Allspark-go/protocol"
	"github.com/labstack/echo/v4"
)

type VortexHttpHandle func(ctx *Context) error

type VortexHttpRouter struct {
	path        string
	methods     []string // 请求方法
	apiDesc     string   // 路由的描述信息
	handle      VortexHttpHandle
	middlewares []VortexHttpMiddleware // 路由中间件
}

func (vh *VortexHttpRouter) WithApiDesc(desc string) *VortexHttpRouter {
	vh.apiDesc = desc
	return vh
}

func AppendHttpRouter(methods []string, path string, handle VortexHttpHandle, desc string, middlwares ...VortexHttpMiddleware) *VortexHttpRouter {
	middlwares = append(middlwares, logReqAndResp(), VerifyJwt())
	return &VortexHttpRouter{
		methods:     methods,
		path:        path,
		handle:      handle,
		middlewares: middlwares,
		apiDesc:     desc,
	}
}

func HttpJsonResponse(c echo.Context, status Status, data interface{}) error {
	// 获取当前请求想要返回的语言类型
	lang := c.Request().Header.Get(HttpHeaderEnum.AcceptLanguage.String())
	em := getEmByLang(lang, status.i18nKey)
	//构造数据
	resp := &protocol.VortexPb{
		Body: data,
		Head: protocol.Head{
			Ec:        status.respCode,
			Em:        em,
			TimeStamp: time.Now().UnixMilli(),
			Type:      reqType.Http,
		},
	}

	if status.subCode != 0 {
		resp.Head.Ec = status.subCode
	}

	return c.JSON(int(resp.Head.Ec), resp)
}
