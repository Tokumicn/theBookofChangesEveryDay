package main

import (
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/ping", handler.Ping)

	// 字典
	dictRouter := r.Group("/dict")
	dictRouter.POST("/buildByStr", handler.BuildDictByStr)

	// 物品
	productRouter := r.Group("/product")
	productRouter.POST("/buildByStr", handler.BuildProductByStr)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
