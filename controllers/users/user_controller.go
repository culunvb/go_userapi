package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"res":     true,
		"message": "OK",
	})
}

func CreateUser(c *gin.Context) {

}

func SearchUser(c *gin.Context) {

}
