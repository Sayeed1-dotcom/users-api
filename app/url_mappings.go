package app

import (
	"github.com/Sayeed1-dotcom/users-api/controllers/ping"
	"github.com/Sayeed1-dotcom/users-api/controllers/users"
)

func mapUrls() {
	router.GET(relativePath:"/ping", ping.Ping)

	router.GET(relativePath:"/users:user_id", users.GetUser)
	router.POST(relativePath:"/users", users.CreateUser)
}