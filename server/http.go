package server

import (
	"github.com/gin-gonic/gin"
	"github.com/laixhe/goimg/uphand"
	"log"
	"net/http"
	"time"
)

var engine = gin.New()

// 注册访问路由
func HandleFunc(pattern string, c uphand.Controller) {
	engine.Handle(http.MethodPost, pattern, c.ServeHTTP)
	engine.Handle(http.MethodGet, pattern, c.ServeHTTP)
}
func Handle(method string, pattern string, handler gin.HandlerFunc) {
	engine.Handle(method, pattern, handler)
}

// 启动 HTTP 服务
// start http server
func RunHttp(adrr string) {

	engine.Use(CORS())

	// 手工配置 http.Server 服务
	server := http.Server{
		Addr:              adrr,            // 监听地址和端口
		Handler:           engine,          // Handle
		ReadTimeout:       5 * time.Second, // 读超时
		WriteTimeout:      5 * time.Second, // 写超时
		ReadHeaderTimeout: 5 * time.Second,
		IdleTimeout:       5 * time.Second,
	}

	// 启动监听
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Header("Access-Control-Allow-Headers", "authorization,content-type,xiao-user-agent")
		if c.Request.Method != http.MethodOptions {
			c.Next()
		} else {
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
