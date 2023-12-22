package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Sayeed1-dotcom/users-api/domain/users"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"github.com/Sayeed1-dotcom/users-api/services"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err :=ioutil.ReadAll(c.Request.Body)
	if err !=nil{
		//TODO: Handle error
		return
	}
    if err:=json.Unmarshal(bytes, &user); err !=nil {
		fmt.Println(err.Error())
		//TODO: Handle json error
		return
	}
	
	_, saveErr := services.CreateUser(user)
	if saveErr != nil {
		//TODO: Handle json error
		return
	}
	c.String(http.StatusNotImplemented, "implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, "implement me!")
}
