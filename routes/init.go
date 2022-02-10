package routes

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	PingRoutes()
	UsersRoutes()
	router.Run(":8080")
}
