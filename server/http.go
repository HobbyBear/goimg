package server

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var Engine = gin.New()

func Handle(method string, pattern string, handler gin.HandlerFunc) {
	Engine.Handle(method, pattern, handler)
}

// 启动 HTTP 服务
// start http server
func RunHttp(adrr string) {

	// 手工配置 http.Server 服务
	server := http.Server{
		Addr:              adrr,            // 监听地址和端口
		Handler:           Engine,          // Handle
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
		c.Header("Access-Control-Allow-Origin", c.Request.Header.Get("Origin")) // 可将将 * 替换为指定的域名
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")
		if c.Request.Method != http.MethodOptions {
			c.Next()
		} else {
			c.Header("Content-Type", "application/json")
			c.AbortWithStatus(http.StatusOK)
		}
	}
}
