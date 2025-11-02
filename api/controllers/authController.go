package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Jojojojodr/taskflow/api/database"
	"github.com/Jojojojodr/taskflow/api/database/models"
	"github.com/Jojojojodr/taskflow/api/internal"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func LoginUser(c *gin.Context) {
	var credentials struct {
		Name     string `json:"name"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid JSON data",
			"details": err.Error(),
		})
		return
	}

	user, err := getUserByName(credentials.Name)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No user found with the provided name",
		})
		return
	}

	if !internal.CheckPasswordHash(credentials.Password, user.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid username or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	})

	secretToken := internal.Env("SECRET_TOKEN")
	tokenString, err := token.SignedString([]byte(secretToken))
	if err != nil {
		log.Printf("Token creation failed: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"user":    user,
	})
}

func LogoutUser(c *gin.Context) {
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", "", -1, "/", "", false, true)

	c.JSON(http.StatusOK, gin.H{
		"message": "Logout successful",
	})
}

func ValidateUser(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "No authorization cookie found",
		})
		return
	}

	secretToken := internal.Env("SECRET_TOKEN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretToken), nil
	})

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token",
		})
		return
	}

	if !token.Valid {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Token is not valid",
		})
		return
	}

	// Extract user ID from token claims
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		userIDStr, ok := claims["sub"].(string)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		// Get user from database
		var user models.User
		result := database.DB.First(&user, "id = ?", userIDStr)
		if result.Error != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "User not found",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Validation successful",
			"user":    user,
		})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid token claims",
		})
	}
}
