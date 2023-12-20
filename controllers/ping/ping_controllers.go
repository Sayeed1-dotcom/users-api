package ping

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.String(http.StatusOK, format:"pong")

}