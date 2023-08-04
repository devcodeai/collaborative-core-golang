package Routes

import (
	"github.com/devcodeai/collaborative-core-golang/Controllers/CampusServicesControllers"
	"github.com/devcodeai/collaborative-core-golang/Controllers/CommunityServicesControllers"
	"github.com/devcodeai/collaborative-core-golang/Controllers/CompanyServicesControllers"
	"github.com/devcodeai/collaborative-core-golang/Controllers/TalentServicesControllers"

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
		// > Company Controller
		routerGroup.GET("/companies", CompanyServicesControllers.GetCompanies)
		routerGroup.GET("/companies/:id", CompanyServicesControllers.GetCompanyById)
		routerGroup.POST("/companies", CompanyServicesControllers.CreateCompany)
		routerGroup.PUT("/companies/:id", CompanyServicesControllers.UpdateCompanyById)
		routerGroup.DELETE("/companies/:id", CompanyServicesControllers.DeleteCompanyById)
		// > Product Controller
		routerGroup.GET("/products", CompanyServicesControllers.GetProductsByCompanyId)
		routerGroup.GET("/products/:id", CompanyServicesControllers.GetProductById)
		routerGroup.POST("/products", CompanyServicesControllers.CreateProductByCompanyId)
		routerGroup.PUT("/products/:id", CompanyServicesControllers.UpdateProductById)
		routerGroup.DELETE("/products/:id", CompanyServicesControllers.DeleteProductById)
		// CampusServices
		// > Campus Controller
		routerGroup.GET("/campuses", CampusServicesControllers.GetCampuses)
		routerGroup.GET("/campuses/:id", CampusServicesControllers.GetCampusById)
		routerGroup.POST("/campuses", CampusServicesControllers.CreateCampus)
		routerGroup.PUT("/campuses/:id", CampusServicesControllers.UpdateCampusById)
		routerGroup.DELETE("/campuses/:id", CampusServicesControllers.DeleteCampusById)
		// > Major Controller
		routerGroup.GET("/majors", CampusServicesControllers.GetMajorsByCampusId)
		routerGroup.GET("/majors/:id", CampusServicesControllers.GetMajorById)
		routerGroup.POST("/majors", CampusServicesControllers.CreateMajorByCampusId)
		routerGroup.PUT("/majors/:id", CampusServicesControllers.UpdateMajorById)
		routerGroup.DELETE("/majors/:id", CampusServicesControllers.DeleteMajorById)
		// TalentServices
		// > TalentController
		routerGroup.GET("/talents", TalentServicesControllers.GetTalents)
		routerGroup.GET("/talents/:id", TalentServicesControllers.GetTalentById)
		routerGroup.POST("/talents", TalentServicesControllers.CreateTalent)
		routerGroup.PUT("/talents/:id", TalentServicesControllers.UpdateTalentById)
		routerGroup.DELETE("/talents/:id", TalentServicesControllers.DeleteTalentById)
		// CommunityServices
		// > CommunityController
		routerGroup.GET("/talents", CommunityServicesControllers.GetCommunities)
		routerGroup.GET("/talents/:id", CommunityServicesControllers.GetCommunityById)
		routerGroup.POST("/talents", CommunityServicesControllers.CreateCommunity)
		routerGroup.PUT("/talents/:id", CommunityServicesControllers.UpdateCommunityById)
		routerGroup.DELETE("/talents/:id", CommunityServicesControllers.DeleteCommunityById)
	}
	return router
}
