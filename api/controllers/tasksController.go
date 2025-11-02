package controllers

import (
	"fmt"
	"net/http"

	"github.com/Jojojojodr/taskflow/api/database"
	"github.com/Jojojojodr/taskflow/api/database/models"
	"github.com/Jojojojodr/taskflow/api/internal"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func GetTasks(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	secretToken := internal.Env("SECRET_TOKEN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretToken), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token subject"})
		return
	}

	userUUID, err := uuid.Parse(sub)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id in token"})
		return
	}

	var tasks []models.Task
	result := database.DB.Where("user_id = ?", userUUID).Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve tasks",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"tasks": tasks,
		"count": len(tasks),
	})
}

func CreateTask(c *gin.Context) {
	var task models.Task

	if err := c.ShouldBindJSON(&task); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Not authenticated"})
		return
	}

	secretToken := internal.Env("SECRET_TOKEN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretToken), nil
	})
	if err != nil || !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
		return
	}

	sub, ok := claims["sub"].(string)
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token subject"})
		return
	}

	userUUID, err := uuid.Parse(sub)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user id in token"})
		return
	}

	task.UserID = userUUID
	task.ID = uuid.New()

	result := database.DB.Create(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to create task",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Task created successfully",
		"task":    task,
	})
}

func GetTask(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task ID is required",
		})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid task ID format",
			"details": err.Error(),
		})
		return
	}

	var task models.Task
	result := database.DB.First(&task, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve task",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"task": task,
	})
}

func UpdateTask(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task ID is required",
		})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid task ID format",
			"details": err.Error(),
		})
		return
	}

	var existingTask models.Task
	result := database.DB.First(&existingTask, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to find task",
			"details": result.Error.Error(),
		})
		return
	}

	var updateData models.Task
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	updateData.ID = existingTask.ID

	result = database.DB.Model(&existingTask).Updates(updateData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to update task",
			"details": result.Error.Error(),
		})
		return
	}

	result = database.DB.First(&existingTask, "id = ?", id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve updated task",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Task updated successfully",
		"task old": existingTask,
		"task new": updateData,
	})
}

func DeleteTask(c *gin.Context) {
	idStr := c.Query("id")
	if idStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Task ID is required",
		})
		return
	}

	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid task ID format",
			"details": err.Error(),
		})
		return
	}

	var task models.Task
	result := database.DB.First(&task, "id = ?", id)

	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{
				"error": "Task not found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to find task",
			"details": result.Error.Error(),
		})
		return
	}

	result = database.DB.Delete(&task)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to delete task",
			"details": result.Error.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":      "Task deleted successfully",
		"deleted_task": task,
	})
}
