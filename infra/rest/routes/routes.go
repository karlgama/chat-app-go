package routes

import "github.com/gin-gonic/gin"

func SetupRoutes() {
	r := gin.Default()
	HandleUserRequest(r)
	r.Run(":8080")
}
