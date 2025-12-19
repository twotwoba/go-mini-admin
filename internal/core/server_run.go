package core

import "github.com/gin-gonic/gin"

func ServerRun() *gin.Engine {
	g := gin.New()
	return g
}
