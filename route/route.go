package route

import (
	"github.com/laixhe/goimg/server"
	"github.com/laixhe/goimg/uphand"
	"net/http"
)

// 注册访问路由
func InitRoute() {

	server.Engine.Use(server.CORS())

	// 路由处理绑定
	server.Handle(http.MethodGet, "/getImg/:imgid", uphand.Get)

	server.Handle(http.MethodGet, "/", uphand.Get)

	server.Handle(http.MethodPost, "/", uphand.Post)

	// 获取图片信息
	server.Handle(http.MethodGet, "/info", uphand.Info)

	// 测试上传
	server.Handle(http.MethodGet, "/test", uphand.Test)
	// 获取状态码
	server.Handle(http.MethodGet, "/statuscode", uphand.StatusCode)
}
