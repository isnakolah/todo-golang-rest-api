package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnakolah/todoAPI/config"
	"github.com/isnakolah/todoAPI/model"
)

func RegisterEndpoint(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userCheck model.User
	config.GetDB().First(&userCheck, "username = ?", user.Username)

	if userCheck.ID > 0 {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	config.GetDB().Save(&user)

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}
