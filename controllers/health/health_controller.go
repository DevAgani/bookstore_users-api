package health

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthStatus(c *gin.Context){
	c.JSON(http.StatusOK,"I am Health :D")
}