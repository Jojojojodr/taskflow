package controllers

import (
	"log"
	"net/http"

	"github.com/Jojojojodr/taskflow/api/database"
	"github.com/Jojojojodr/taskflow/api/database/models"
	"github.com/Jojojojodr/taskflow/api/internal"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetUser(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	var user models.User
	result := database.DB.First(&user, "id = ?", idStr)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "User not found",
			})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error":   "Failed to retrieve user",
				"details": result.Error.Error(),
			})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func CreateUser(c *gin.Context) {
	var newUser struct {
		Name     string `json:"name"`
		Email	string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	// Check if a user with the same name already exists
	var existing models.User
	check := database.DB.First(&existing, "name = ?", newUser.Name)
	if check.Error == nil {
		c.JSON(http.StatusConflict, gin.H{
			"error": "User with that name already exists",
		})
		return
	} else if check.Error != nil && check.Error != gorm.ErrRecordNotFound {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to check existing user",
			"details": check.Error.Error(),
		})
		return
	}
	
	var user models.User
	user.ID = uuid.New() 
	user.Name = newUser.Name
	user.Email = newUser.Email
	user.Password = internal.HashPassword(newUser.Password)

	result := database.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create user",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully",
		"user":    user,
	})
}

func UpdateUser(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	result := database.DB.Save(&user)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update user",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User updated successfully",
		"user":    user,
	})
}

func DeleteUser(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID is required",
		})
		return
	}

	result := database.DB.Delete(&models.User{}, "id = ?", idStr)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete user",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "User deleted successfully",
	})
}

func getUserByName(name string) (models.User, error) {
	var user models.User

	// First try with 'name' column (new schema)
	result := database.DB.First(&user, "name = ?", name)

	if result.Error != nil {
		log.Printf("Database query failed for 'name': %v", result.Error)
	}

	return user, result.Error
}
