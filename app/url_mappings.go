package app

import (
	"github.com/DevAgani/bookstore_users-api/controllers/health"
	"github.com/DevAgani/bookstore_users-api/controllers/user"
)

func mapUrls()  {
	router.GET("/health", health.HealthStatus)

	router.GET("/user/:user_id", user.GetUser)
	router.GET("/users/search", user.SearchUser)
	router.POST("/users", user.CreateUser)
}
