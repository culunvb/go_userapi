package routes

import "github.com/culunvb/go_userapi/controllers/users"

func UsersRoutes() {
	router.GET("/users/:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}
