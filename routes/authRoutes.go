package routes

import (
	"ecomerce/controller/authenticate"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(router *gin.Engine) {
	router.POST("/signup", authenticate.Signup)
	router.POST("/login", authenticate.Login)
}
