package httpserver

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewHttpServer() *gin.Engine {
	router := gin.Default()

	// Allow CORs Policy
	// TODO: Set the appropriate permission
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
		MaxAge: 12 * time.Hour,
	}))

	healthCheckGroup := router.Group("healthcheck")
	healthCheckGroup.GET("/", healthCheckHandler)
	return router
}
