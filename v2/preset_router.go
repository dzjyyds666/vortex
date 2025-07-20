package vortex

import "net/http"

// 预制路由

func prepareRouters(v Vortex) []*VortexHttpRouter {
	return []*VortexHttpRouter{
		AppendHttpRouter(
			[]string{http.MethodPost, http.MethodGet},
			"/v1/ws",
			v.handleWebSocket,
			"处理websocket的handle",
			logReqAndResp(), VerifyJwt()),
	}
}

// 处理websocket的handle
func (v *Vortex) handleWebSocket(ctx *Context) error {
	return nil
}

// 统一处理http的handle，根据请求体中设置的cmd接口进行访问
func (v *Vortex) handleCmd(ctx *Context) error {
	return nil
}
