package controller

import (
	"net/http"

	jwtapple2 "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"github.com/isnakolah/todoAPI/config"
	"github.com/isnakolah/todoAPI/model"
)

func CreateTask(c *gin.Context) {
	claims := jwtapple2.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todo model.Task
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo.UserID = user.ID
	config.GetDB().Save(&todo)
	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully!", "task": todo})
}

func FetchAllTask(c *gin.Context) {
	claims := jwtapple2.ExtractClaims(c)

	var user model.User
	config.GetDB().Where("id = ?", claims[config.IdentityKey]).First(&user)

	if user.ID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todos []model.Task
	config.GetDB().Where("user_id = ?", user.ID).Order("created_at desc").Find(&todos)

	if len(todos) <= 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No tasks found", "data": todos})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func FetchSingleTask(c *gin.Context) {
	todoID := c.Param("id")

	if len(todoID) <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user id"})
		return
	}

	var todo model.Task
	config.GetDB().First(&todo, todoID)

	if todo.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{"message": "No todo found!"})
		return
	}

	c.JSON(http.StatusOK, todo)
}
