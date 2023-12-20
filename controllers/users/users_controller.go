package users

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/Sayeed1-dotcom/users-api/domain/users"
)

func CreateUser(c *gin.Context) {
	var user users.User
	fmt.Println(user)
	bytes, err :=ioutil.ReadAll(c.Request.Body)
	fmt.Println(err)
	fmt.Println(string(bytes))
	c.String(http.StatusNotImplemented, format:"implement me!")
}

func GetUser(c *gin.Context) {
	c.String(http.StatusNotImplemented, format:"implement me!")
}
