package routes

import (
	"github.com/culunvb/go_userapi/controllers/ping"
)

func PingRoutes() {
	router.GET("/ping", ping.Ping)

}
