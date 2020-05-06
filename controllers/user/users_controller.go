package user

import (
	"fmt"
	"github.com/DevAgani/bookstore_users-api/domain/users"
	"github.com/DevAgani/bookstore_users-api/services"
	"github.com/DevAgani/bookstore_users-api/utils/errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)


func getUserId(userIdParam string) (int64, *errors.RestErr){
	userId, userErr := strconv.ParseInt(userIdParam,10,64)
	if userErr != nil{
		fmt.Println(userErr)
		return 0, errors.NewBadRequestError("user id should be a number")
	}
	return userId,nil
}
func Get(c *gin.Context){
	userId,idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}

	user, getErr := services.GetUser(userId)
	if getErr !=nil{
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK,user)
}
func Create(c *gin.Context){
	var user users.User
	if err := c.ShouldBindJSON(&user);err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result, saveError := services.CreateUser(user)
	if saveError != nil{
		// TODO Handle user creation error
		c.JSON(saveError.Status, saveError)
		return
	}
	c.JSON(http.StatusCreated,result)
}
func Update(c *gin.Context)  {
	userId,idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}
	var user users.User
	if err := c.ShouldBindJSON(&user);err != nil{
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	user.Id = userId
	isPartial := c.Request.Method == http.MethodPatch
	result, err := services.UpdateUser(isPartial,user)
	if err != nil{
		c.JSON(err.Status,err)
		return
	}
	c.JSON(http.StatusOK,result)
}
func Delete(c *gin.Context)  {
	userId,idErr := getUserId(c.Param("user_id"))
	if idErr != nil{
		c.JSON(idErr.Status, idErr)
		return
	}
	if err := services.DeleteUser(userId);err != nil{
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, map[string]string{"status":"deleted"})
}
func Search(c *gin.Context){
	c.JSON(http.StatusNotImplemented,"implement me")
}