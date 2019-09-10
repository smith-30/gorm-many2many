package main

import "github.com/gin-gonic/gin"

func main() {
	r := setupRouter()

	err := r.Run(":8080")
	if err != nil {
		err.Error()
	}
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	apiRouter := r.Group("/api")

	apiRouter.POST("/roles", CreateRole)
	apiRouter.POST("/users", CreateUser)
	apiRouter.PUT("/users/:user_id", UpdateUser)
	apiRouter.GET("/users/:user_id", GetUserRole)

	return r
}
