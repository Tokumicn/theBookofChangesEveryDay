package main

import (
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/database"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/internal/handler"
	"github.com/Tokumicn/theBookofChangesEveryDay/mhxyHelper/pkg/logger"
	"github.com/gin-gonic/gin"
)

func main() {
	logger.NewLogger()
	// 初始化数据库连接
	_, err := database.InitDB()
	if err != nil {
		panic(err)
	}

	r := gin.Default()
	r.GET("/ping", handler.Ping)

	// 字典构建
	dictRouter := r.Group("/dict")
	dictRouter.POST("/build", handler.BuildDict)

	stuffRouter := r.Group("/stuff")
	// 物品构建
	stuffRouter.POST("/build", handler.BuildStuff)
	// 物品查询
	stuffRouter.POST("/query", handler.QueryStuff)

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
}
