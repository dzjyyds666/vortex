package vortex

import (
	"github.com/dzjyyds666/Allspark-go/jwtx"
	"github.com/dzjyyds666/Allspark-go/logx"
	"github.com/golang-jwt/jwt/v5"

	"github.com/labstack/echo/v4"
)

type Session struct {
	Uid    string `json:"uid"`    // 用户id
	Sid    string `json:"sid"`    // session id
	Expire int64  `json:"expire"` // 过期时间
}

func (s *Session) AsJwtClaims() jwt.MapClaims {
	claims := make(jwt.MapClaims)
	claims["uid"] = s.Uid
	claims["sid"] = s.Sid
	claims["expire"] = s.Expire
	return claims
}

func (s *Session) Bind(claims jwt.MapClaims) *Session {
	s.Uid = claims["uid"].(string)
	s.Sid = claims["sid"].(string)
	s.Expire = int64(claims["expire"].(float64))
	return s
}

type VortexHttpMiddleware echo.MiddlewareFunc

// 打印请求和响应的日志
func logReqAndResp() VortexHttpMiddleware {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// 打印请求的日志
			logx.Infof("req ==>> [ %s %s ]", c.Request().Method, c.Request().URL)
			err := next(c)
			if nil != err {
				return err
			}
			// 打印响应的日志
			logx.Infof("rsp ==>> [ %d %s ]", c.Response().Status, c.Request().URL)
			return nil
		}
	}
}

// jwt解析
func VerifyJwt() VortexHttpMiddleware {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			token := c.Request().Header.Get(HttpHeaderEnum.Authorization.String())
			if len(token) > 0 {
				claims, err := jwtx.ParseToken(sercetKey, token)
				if nil != err {
					logx.Errorf("vortex|VerifyMw|VerifyJwt ParseToken err:%v", err)
					return err
				}
				session := new(Session).Bind(claims)
				// 还是交给业务侧判断是否根据过期时间拒绝该次请求
				c.Set(HttpHeaderEnum.Session.XString(), session)
			}
			return next(c)
		}
	}
}
