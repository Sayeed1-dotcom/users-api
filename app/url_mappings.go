package app

import (
	"github.com/Sayeed1-dotcom/users-api/controllers/ping"
	"github.com/Sayeed1-dotcom/users-api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.GET("/users:user_id", users.GetUser)
	router.POST("/users", users.CreateUser)
}