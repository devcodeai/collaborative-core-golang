package Routes

import (
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": c.Request.URL.Path + " not found!",
		})
	})
	routerGroup := router.Group("/api")
	{
		routerGroup.GET("/", func(c *gin.Context) {
			println("Welcome to the Collaborative Core API (Go)!")
			c.JSON(200, gin.H{
				"status":  "Success",
				"message": "Welcome to the Collaborative Core API (Go)!",
			})
		})
		// CompanyServices
		// CampusServices
		// TalentServices
		// CommunityServices
	}
	return router
}
