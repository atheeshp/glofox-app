package classes

import "github.com/gin-gonic/gin"

func ClassRouter(router *gin.RouterGroup) {
	router.POST("/", ValidateCreateClass, CreateClass)
}
