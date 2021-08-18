package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/viniciusmazon/go-jwt/database"
	"github.com/viniciusmazon/go-jwt/models"
	"github.com/viniciusmazon/go-jwt/services"
)

func Login(c *gin.Context) {
	db := database.GetDatabase()

	var p models.Login

	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(400, gin.H{
			"error": "cannot bind JSON: " + err.Error(),
		})

		return
	}

	var user models.User
	dbError := db.Where("email = ?", p.Email).First(&user).Error
	if dbError != nil {
		c.JSON(400, gin.H{
			"error": "cannot find user",
		})

		return
	}

	if user.Password != services.SHA256Encoder(p.Password) {
		c.JSON(401, gin.H{
			"error": "invalid credentials",
		})

		return
	}

	token, err := services.NewJWTService().GenerateToken(user.ID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})

		return
	}

	c.JSON(200, gin.H{
		"token": token,
	})
}
