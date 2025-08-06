package vortex

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/dzjyyds666/Allspark-go/conv"
	"github.com/dzjyyds666/Allspark-go/logx"
	"github.com/dzjyyds666/vortex/v2/locale"
	"github.com/labstack/echo/v4"
)

var reqType = struct {
	Http      string
	WebSocket string
}{
	Http:      "http",
	WebSocket: "websocket",
}

var (
	emMap            = make(map[string]string)
	consoleSecretKey = "" // 控制台的密钥
	sercetKey        = "" // 服务的密钥
)

type Vortex struct {
	ctx           context.Context
	e             *echo.Echo          // echo 实例
	port          string              // 服务的端口号
	routers       []*VortexHttpRouter // http路由
	hiddleBanner  bool                // 是否隐藏banner
	hiddleRouters bool                // 是否打印路由信息

}

// 程序启动入口
func BootStrap(ctx context.Context, options ...Option) *Vortex {
	e := echo.New()
	e.HideBanner = true

	v := &Vortex{
		ctx: ctx,
		e:   e,
	}

	v.routers = prepareRouters(v)
	// 设置默认的18n
	WithI18n(locale.V)
	for _, opt := range options {
		opt(v)
	}

	return v
}

// Start 启动服务器
func (v *Vortex) Start() {
	if !v.hiddleRouters {
		// 打印路由信息
		v.printRouters()
	}
	if !v.hiddleBanner {
		// 打印banner
		v.showBanner()
	}

	err := v.e.Start(":" + v.port)
	if nil != err {
		panic(err)
	}
}

type Option func(*Vortex) *Vortex

// 设置服务的端口
func WithPort(port string) Option {
	return func(v *Vortex) *Vortex {
		v.port = port
		return v
	}
}

// 设置http的路由
func WithRouters(routers []*VortexHttpRouter) Option {
	return func(v *Vortex) *Vortex {
		for _, router := range routers {
			echoHandler := func(c echo.Context) error {
				return router.handle(&Context{
					Context: c,
				})
			}
			echoMws := make([]echo.MiddlewareFunc, 0)
			// 转化中间件为echo中允许的类型
			for _, middleware := range router.middlewares {
				echoMws = append(echoMws, echo.MiddlewareFunc(middleware))
			}
			v.e.Match(router.methods, router.path, echoHandler, echoMws...)
		}
		v.routers = append(v.routers, routers...)
		return v
	}
}

// 设置i18n配置
func WithI18n(i18n string) Option {
	return func(v *Vortex) *Vortex {
		// 输入的i18n json字符串反序列化到emMap
		tmp := make(map[string]string)
		err := json.Unmarshal([]byte(i18n), &tmp)
		if nil != err {
			panic(err)
		}
		for k, v := range emMap {
			tmp[k] = v
		}
		emMap = tmp
		return v
	}
}

// 隐藏banner
func WithHiddleBanner() Option {
	return func(v *Vortex) *Vortex {
		v.hiddleBanner = true
		return v
	}
}

// 隐藏路由信息
func WithHiddleRouters() Option {
	return func(v *Vortex) *Vortex {
		v.hiddleRouters = true
		return v
	}
}

// 打印出路由的详细信息
func (v *Vortex) printRouters() {
	for _, router := range v.routers {
		logx.Infof("methods=>[%s]\tapi=>[%s]\tpath=>[%s]\tdesc=>[%s]", conv.ToJsonWithoutError(router.methods), router.apiDesc, router.path, router.apiDesc)
	}
}

func (v *Vortex) showBanner() {
	logx.Infof(fmt.Sprintf(`
 __     __                 _                 
 \ \   / /   ___    _ __  | |_    ___  __  __
  \ \ / /   / _ \  | '__| | __|  / _ \ \ \/ /
   \ V /   | (_) | | |    | |_  |  __/  >  < 
    \_/     \___/  |_|     \__|  \___| /_/\_\
----------------------------------------------
 Vortex Server Start Success on %s
==============================================
`, v.port))
}

func WithJwtSecretKey(key string) Option {
	return func(v *Vortex) *Vortex {
		sercetKey = key
		return v
	}
}

func WithConsoleSecretKey(key string) Option {
	return func(v *Vortex) *Vortex {
		consoleSecretKey = key
		return v
	}
}

// 根据语言类型获取对应Em
func getEmByLang(i18nkey, lang string) string {

	//拼接对应key
	if len(lang) <= 0 || len(i18nkey) <= 0 {
		return ""
	}
	if emMap == nil {
		return ""
	}
	key := i18nkey + "." + lang
	return emMap[key]
}
