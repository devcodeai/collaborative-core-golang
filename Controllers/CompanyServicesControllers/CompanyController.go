package CompanyServicesControllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/devcodeai/collaborative-core-golang/Database"
	"github.com/devcodeai/collaborative-core-golang/Models"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func GetCompanies(c *gin.Context) {
	var companies []Models.Company
	if err := Database.DB.Find(&companies).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}
	if len(companies) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Companies Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Fetch Companies Success",
		"data":    companies,
	})
}

func GetCompanyById(c *gin.Context) {
	id := c.Param("id")
	var company Models.Company
	if err := Database.DB.First(&company, id).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": fmt.Sprintf("Company with ID %s Not Found", id),
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
		"message": fmt.Sprintf("Fetch Company with ID %s Success", id),
		"data":    company,
	})
}

func CreateCompany(c *gin.Context) {
	var company Models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if company.Name == "" || company.Address == "" || company.Email == "" || company.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if err := Database.DB.Create(&company).Error; err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "Failed",
			"message": "Internal Server Error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Create Company Success",
		"data":    company,
	})
}

func UpdateCompanyById(c *gin.Context) {
	id := c.Param("id")
	var company Models.Company
	if err := c.ShouldBindJSON(&company); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if company.Name == "" || company.Address == "" || company.Email == "" || company.Phone == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Failed",
			"message": "Bad Request",
		})
		return
	}
	if Database.DB.Model(&company).Where("id = ?", id).Updates(&company).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Company with ID %s Not Found", id),
		})
		return
	}

	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		fmt.Println("Error:", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Update Company with ID %s Success", id),
		"data": Models.Company{
			ID:      uint(idUint),
			Name:    company.Name,
			Address: company.Address,
			Email:   company.Email,
			Phone:   company.Phone,
		},
	})
}

func DeleteCompanyById(c *gin.Context) {
	id := c.Param("id")
	var company Models.Company
	if Database.DB.Delete(&company, id).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": fmt.Sprintf("Company with ID %s Not Found", id),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": fmt.Sprintf("Delete Company with ID %s Success", id),
		"data":    map[string]interface{}{},
	})
}
