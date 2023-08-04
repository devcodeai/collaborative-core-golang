package CampusServicesControllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Models"
)

func GetMajorsByCampusId(c *gin.Context) {
	campus_id := c.Query("campus_id")
	var majors []Models.Major
	if err := Database.DB.Where("campus_id = ?", campus_id).Find(&majors).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	if len(majors) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Majors with Campus ID %s Not Found", campus_id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Fetch Majors with Campus ID %s Success", campus_id),
		"data":    majors,
	})
}

func GetMajorById(c *gin.Context) {
	id := c.Param("id")
	var major Models.Major
	if err := Database.DB.First(&major, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Major with ID %s Not Found", id),
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
		"message": fmt.Sprintf("Fetch Major with ID %s Success", id),
		"data":    major,
	})
}

func CreateMajorByCampusId(c *gin.Context) {
	var major Models.Major
	if err := c.ShouldBindJSON(&major); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if major.Name == "" || major.CampusID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}

	var campus Models.Campus
	campus_id := major.CampusID
	if err := Database.DB.First(&campus, campus_id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Create Major Failed, Campus with ID %v Not Found", campus_id),
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
	if err := Database.DB.Create(&major).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Create Major Success",
		"data":    major,
	})
}

func UpdateMajorById(c *gin.Context) {
	id := c.Param("id")
	var major Models.Major
	if err := c.ShouldBindJSON(&major); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if major.Name == "" || major.CampusID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}

	var campus Models.Campus
	campus_id := major.CampusID
	if err := Database.DB.First(&campus, campus_id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Update Major Failed, Campus with ID %v Not Found", campus_id),
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
	if Database.DB.Model(&major).Where("id = ?", id).Updates(&major).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Major with ID %s Not Found", id),
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Update Major with ID %s Success", id),
		"data": Models.Major{
			ID:       uint(idUint),
			Name:     major.Name,
			CampusID: major.CampusID,
		},
	})
}

func DeleteMajorById(c *gin.Context) {
	id := c.Param("id")
	var major Models.Major
	if Database.DB.Delete(&major, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Major with ID %s Not Found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Delete Major with ID %s Success", id),
		"data":    map[string]interface{}{},
	})
}
