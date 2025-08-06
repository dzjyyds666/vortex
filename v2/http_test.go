package vortex

import (
	"strings"
	"testing"

	"github.com/dzjyyds666/Allspark-go/conv"
	"github.com/dzjyyds666/Allspark-go/jwtx"
	"github.com/smartystreets/goconvey/convey"
)

func Test_ParseJwtToken(t *testing.T) {
	convey.Convey("ParseJwtToken", t, func() {
		token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVzIjoxNzU0NDA5OTUwLCJ1aWQiOiJhYXJvbiJ9.63e-SoAsLK9WkngTFnbQEJtMbROPg6ASw-NSaiOIxIU"
		// 处理Bearer前缀
		if after, ok := strings.CutPrefix(token, "Bearer "); ok {
			token = after
		}
		claims, err := jwtx.ParseToken("123456", token)
		if nil != err {
			// 使用console的key试试
			claims, err = jwtx.ParseToken(consoleSecretKey, token)
			convey.So(err, convey.ShouldBeNil)
		}
		t.Logf("clamis:%s", conv.ToJsonWithoutError(claims))
	})
}
