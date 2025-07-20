package vortex

type HttpHeader string

func (h HttpHeader) String() string {
	return string(h)
}

func (h HttpHeader) XString() string {
	return "X-Vortex-" + string(h)
}

var HttpHeaderEnum = struct {
	ContentType    HttpHeader // 内容类型
	ContentLength  HttpHeader // 内容长度
	AcceptLanguage HttpHeader // 接收语言
	Authorization  HttpHeader // 授权
}{
	ContentType:    "Content-Type",
	ContentLength:  "Content-Length",
	AcceptLanguage: "Accept-Language",
	Authorization:  "Authorization",
}
