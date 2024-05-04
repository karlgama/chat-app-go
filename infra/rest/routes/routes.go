package routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	r := gin.Default()
	HandleUserRequests(r)
	HandleAuthRequests(r)
	r.Run(":8080")
}
