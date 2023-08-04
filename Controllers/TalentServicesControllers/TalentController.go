package TalentServicesControllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetTalents(c *gin.Context) {
	var talents []Models.Talent
	if err := Database.DB.Find(&talents).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	if len(talents) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Talents Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Fetch Talents Success",
		"data":    talents,
	})
}

func GetTalentById(c *gin.Context) {
	id := c.Param("id")
	var talent Models.Talent
	if err := Database.DB.First(&talent, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Talent with ID %s Not Found", id),
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
		"message": fmt.Sprintf("Fetch Talent with ID %s Success", id),
		"data":    talent,
	})
}

func CreateTalent(c *gin.Context) {
	var talent Models.Talent
	if err := c.ShouldBindJSON(&talent); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if talent.Name == "" || talent.Email == "" || talent.Skills == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if err := Database.DB.Create(&talent).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Create Talent Success",
		"data":    talent,
	})
}

func UpdateTalentById(c *gin.Context) {
	id := c.Param("id")
	var talent Models.Talent
	if err := c.ShouldBindJSON(&talent); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if talent.Name == "" || talent.Email == "" || talent.Skills == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if Database.DB.Model(&talent).Where("id = ?", id).Updates(&talent).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Talent with ID %s Not Found", id),
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Update Talent with ID %s Success", id),
		"data": Models.Talent{
			ID:     uint(idUint),
			Name:   talent.Name,
			Email:  talent.Email,
			Skills: talent.Skills,
		},
	})
}

func DeleteTalentById(c *gin.Context) {
	id := c.Param("id")
	var talent Models.Talent
	if Database.DB.Delete(&talent, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Talent with ID %s Not Found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Delete Talent with ID %s Success", id),
		"data":    map[string]interface{}{},
	})
}
