package app

import (
	"github.com/DevAgani/bookstore_users-api/logger"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)
func StartApplication()  {
	mapUrls()

	logger.Info("Application is starting ...")
	router.Run(":8080")
	
}
