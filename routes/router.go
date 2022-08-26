package routes

import (
	"criu/app"
	"criu/cmd"
	"criu/middleware"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/dump", app.Dump)
	r.POST("/restore", app.Restore)
	r.POST("/startCmd", app.StartPageServer)
	_ = r.Run(":" + cmd.Port)
}
