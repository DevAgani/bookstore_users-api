package user

import (
	"github.com/DevAgani/bookstore_users-api/domain/users"
	"github.com/DevAgani/bookstore_users-api/services"
	"github.com/DevAgani/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetUser(c *gin.Context){
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr != nil{
		err := errors.NewBadRequestError("user id should be a number")
		c.JSON(err.Status, err)
	}

	user, getErr := services.GetUser(userId)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK,user)
}
func CreateUser(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user);err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil{
		// TODO Handle user creation error
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated,result)
}
func SearchUser(c *gin.Context){
	c.JSON(http.StatusNotImplemented,"implement me")
}