package app

import (
	"github.com/DevAgani/bookstore_users-api/controllers/health"
	"github.com/DevAgani/bookstore_users-api/controllers/user"
)

func mapUrls()  {
	router.GET("/health", health.HealthStatus)

	router.POST("/users", user.Create)
	router.GET("/user/:user_id", user.Get)
	router.PUT("/user/:user_id", user.Update)
	router.PATCH("/user/:user_id", user.Update)
	router.DELETE("/user/:user_id",user.Delete)
	router.GET("/internal/users/search", user.Search)

}
