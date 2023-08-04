package CampusServicesControllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetCampuses(c *gin.Context) {
	var campuses []Models.Campus
	if err := Database.DB.Find(&campuses).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	if len(campuses) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Campuses Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Fetch Campuses Success",
		"data":    campuses,
	})
}

func GetCampusById(c *gin.Context) {
	id := c.Param("id")
	var campus Models.Campus
	if err := Database.DB.First(&campus, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Campus with ID %s Not Found", id),
			})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
				"status":  "Failed",
				"message": "Internal Server Error",
				"error":   err.Error(),
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Fetch Campus with ID %s Success", id),
		"data":    campus,
	})
}

func CreateCampus(c *gin.Context) {
	var campus Models.Campus
	if err := c.ShouldBindJSON(&campus); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if campus.UniversityName == "" || campus.Location == "" || campus.Website == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if err := Database.DB.Create(&campus).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Create Campus Success",
		"data":    campus,
	})
}

func UpdateCampusById(c *gin.Context) {
	id := c.Param("id")
	var campus Models.Campus
	if err := c.ShouldBindJSON(&campus); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if campus.UniversityName == "" || campus.Location == "" || campus.Website == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if Database.DB.Model(&campus).Where("id = ?", id).Updates(&campus).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Campus with ID %s Not Found", id),
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Update Campus with ID %s Success", id),
		"data": Models.Campus{
			ID:             uint(idUint),
			UniversityName: campus.UniversityName,
			Location:       campus.Location,
			Website:        campus.Website,
		},
	})
}

func DeleteCampusById(c *gin.Context) {
	id := c.Param("id")
	var campus Models.Campus
	if Database.DB.Delete(&campus, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Campus with ID %s Not Found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Delete Campus with ID %s Success", id),
		"data":    map[string]interface{}{},
	})
}
