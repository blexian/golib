package gin

import (
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"net/http"
)

func addCommonRoutes(g *gin.Engine) {
	// health
	g.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "ok")
	})
}

func addControllerRoute(g *gin.Engine) {
	v1 := g.Group("/api/v1")
	// 配置中间件
	v1.Use(
		gzip.Gzip(gzip.DefaultCompression),
	)
	v1.GET("/user", func(c *gin.Context) {})
}

func getRouter() {
	// 设置模式
	gin.SetMode(gin.ReleaseMode)
	// 初始化
	g := gin.New()
	// 启动Context特性
	g.ContextWithFallback = true
	// 配置中间件
	g.Use(
		gin.Recovery(),
	)
	// health
	addCommonRoutes(g)
	// controller
	addControllerRoute(g)
}
