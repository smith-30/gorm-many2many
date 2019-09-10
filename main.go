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

	return r
}
