package routers

import (
	"gin/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	//房源列表
	router.POST("/bdlist",controllers.BdlistPost)
	return router
}