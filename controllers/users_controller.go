package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"projectPath/domain/users"
	"projectPath/services"
	"projectPath/utils/errors"
	"strconv"
)


func CreateUser(c *gin.Context)  {
	var user users.User
	if err := c.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		restErr := errors.NewBadRequestError("invalid json body")
		c.JSON(restErr.Status,restErr)
		return
	}
	result,saveErr := services.CreateUser(user)
	if saveErr != nil {
		c.JSON(saveErr.Status, saveErr)
		return
	}
	c.JSON(http.StatusCreated, result)
}

func GetUser(c *gin.Context)  {
	userId, userErr := strconv.ParseInt(c.Param("user_id"),10,64)
	if userErr!= nil {
		err:= errors.NewBadRequestError("userid should be number")
		c.JSON(err.Status,err)
		return
	}
	user, getErr := services.GetUser(userId)
	if getErr != nil {
		c.JSON(getErr.Status,getErr)
		return
	}
	c.JSON(http.StatusOK, user)
}

//func SearchUser(c *gin.Context)  {
//	c.String(http.StatusNotImplemented, "Not Implement")
//}