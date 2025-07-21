package vortex

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dzjyyds666/Allspark-go/logx"
	"github.com/dzjyyds666/Allspark-go/protocol"
	"github.com/labstack/echo/v4"
)

type VortexHttpHandle func(ctx *Context) error

type vortexHttpRouter struct {
	path        string
	methods     []string // 请求方法
	apiDesc     string   // 路由的描述信息
	handle      VortexHttpHandle
	middlewares []VortexHttpMiddleware // 路由中间件
}

func (vh *vortexHttpRouter) WithApiDesc(desc string) *vortexHttpRouter {
	vh.apiDesc = desc
	return vh
}

func AppendHttpRouter(methods []string, path string, handle VortexHttpHandle, desc string, middlwares ...VortexHttpMiddleware) *vortexHttpRouter {
	middlwares = append(middlwares, logReqAndResp(), VerifyJwt())
	return &vortexHttpRouter{
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
	em := getEmByLang(lang, status.I18nKey)
	//构造数据
	resp := &protocol.VortexPb{
		Body: data,
		Head: protocol.Head{
			Ec:        status.RespCode,
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

// 预制路由
func prepareRouters(v *Vortex) []*vortexHttpRouter {
	return []*vortexHttpRouter{
		AppendHttpRouter(
			[]string{http.MethodPost, http.MethodGet},
			"/v1/ws",
			v.handleWebSocket,
			"处理websocket的handle",
			logReqAndResp(), VerifyJwt()),
		AppendHttpRouter(
			[]string{http.MethodGet},
			"/v1/checkalive",
			v.handleCheckAlive,
			"程序健康检查",
			logReqAndResp(), VerifyJwt()),
		AppendHttpRouter(
			[]string{http.MethodGet, http.MethodPost},
			"/v1/cmd",
			v.handleCmd,
			"cmd接口，使用统一的结构",
			logReqAndResp(), VerifyJwt(),
		),
	}
}

// 处理websocket的handle
func (v *Vortex) handleWebSocket(ctx *Context) error {
	return nil
}

// 统一处理http的handle，根据请求体中设置的cmd接口进行访问
func (v *Vortex) handleCmd(ctx *Context) error {
	var cmd protocol.VortexPb
	decoder := json.NewDecoder(ctx.Request().Body)
	if err := decoder.Decode(&cmd); nil != err {
		logx.Errorf("Vortex|handleCmd|Error|Params Invaild|%v", err)
		return HttpJsonResponse(ctx, Status{RespCode: 400}, nil)
	}
	return nil
}

// 健康检查
func (v *Vortex) handleCheckAlive(ctx *Context) error {
	return HttpJsonResponse(ctx, Status{RespCode: 200}, nil)
}
