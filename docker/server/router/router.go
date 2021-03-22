package router

import (
	// "server/api/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()

	if gin.IsDebugging() {
		r.Use(cors.New(cors.Config{
			AllowOrigins: []string{"*"},
			AllowMethods: []string{"GET, POST, PUT, PATCH, DELETE"},
			AllowHeaders: []string{"*"},
		}))
	}
	setUp(r)
	return r
}

func setUp(r *gin.Engine) {
	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello")
	}
}
