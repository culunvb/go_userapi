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
	// router.SetTrustedProxies([]string{"192.168.1.2"})
	router.Run(":8080")
}
