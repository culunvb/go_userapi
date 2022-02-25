package users

import (
	"net/http"
	"strconv"

	usersdomain "github.com/culunvb/go_userapi/domains/users"
	"github.com/culunvb/go_userapi/services"
	"github.com/culunvb/go_userapi/utils/errors"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	userId, userErr := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userErr != nil {
		err := errors.BadRequestError("User id should be number")
		c.JSON(err.Status, err)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status, getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

func CreateUser(c *gin.Context) {
	var user usersdomain.User

	if err := c.ShouldBindJSON(&user); err != nil {
		restErr := errors.BadRequestError("invalid json body")
		c.JSON(restErr.Status, restErr)
		//fmt.Print(err.Error())
		// TOdo ; return bad request to the caller
		return
	}

	result, saveErr := services.CreateUser(user)
	if saveErr != nil {
		// TOdo : handle user creation error
		c.JSON(saveErr.Status, saveErr)
		return
	}

	c.JSON(http.StatusCreated, result)
	// fmt.Println(user)
	// fmt.Println(string(bytes))
	// c.String(http.StatusOK, "Test OK")
}

func SearchUser(c *gin.Context) {

}
