package CommunityServicesControllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetCommunities(c *gin.Context) {
	var communities []Models.Community
	if err := Database.DB.Find(&communities).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	if len(communities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Communities Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Fetch Communities Success",
		"data":    communities,
	})
}

func GetCommunityById(c *gin.Context) {
	id := c.Param("id")
	var community Models.Community
	if err := Database.DB.First(&community, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Community with ID %s Not Found", id),
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
		"message": fmt.Sprintf("Fetch Community with ID %s Success", id),
		"data":    community,
	})
}

func CreateCommunity(c *gin.Context) {
	var community Models.Community
	if err := c.ShouldBindJSON(&community); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if community.Name == "" || community.Description == "" || community.Members == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if err := Database.DB.Create(&community).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Create Community Success",
		"data":    community,
	})
}

func UpdateCommunityById(c *gin.Context) {
	id := c.Param("id")
	var community Models.Community
	if err := c.ShouldBindJSON(&community); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if community.Name == "" || community.Description == "" || community.Members == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if Database.DB.Model(&community).Where("id = ?", id).Updates(&community).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Community with ID %s Not Found", id),
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Update Community with ID %s Success", id),
		"data": Models.Community{
			ID:          uint(idUint),
			Name:        community.Name,
			Description: community.Description,
			Members:     community.Members,
		},
	})
}

func DeleteCommunityById(c *gin.Context) {
	id := c.Param("id")
	var community Models.Community
	if Database.DB.Delete(&community, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Community with ID %s Not Found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Delete Community with ID %s Success", id),
		"data":    map[string]interface{}{},
	})
}
