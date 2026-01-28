package core

import "github.com/gin-gonic/gin"

func ServerRun() *gin.Engine {
	g := gin.New()
	// 默认的不太够，使用第三方的来兼容处理
	// g.Use(gin.Logger())
	// g.Use(gin.Recovery())

	// TODO

	return g
}
