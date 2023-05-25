package http

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Start() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"status": "up",
		})
	})
	addr := os.Getenv("ANOTHER_NAS_TOOLS_ADDR")
	if addr == "" {
		addr = ":8080"
	}
	r.Run(addr)
}
