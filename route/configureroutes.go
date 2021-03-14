package route

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/isnakolah/todoAPI/auth"
	"github.com/isnakolah/todoAPI/controller"
)

func SetupRoutes() *gin.Engine {
	router := gin.Default()
	authMiddleware, err := auth.SetupAuth()

	if err != nil {
		log.Fatal("JWT Error" + err.Error())
	}

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Welcome to my Todo App"})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/login", authMiddleware.LoginHandler)

		v1.POST("/register", controller.RegisterEndpoint)

		todo := v1.Group("/todo")
		{
			todo.POST("/create", authMiddleware.MiddlewareFunc(), controller.CreateTask)
			todo.POST("/all", authMiddleware.MiddlewareFunc(), controller.FetchAllTask)
			todo.POST("/get/:id", authMiddleware.MiddlewareFunc(), controller.FetchSingleTask)
			todo.POST("/update/:id", authMiddleware.MiddlewareFunc(), controller.UpdateTask)
			todo.POST("/delete/:id", authMiddleware.MiddlewareFunc(), controller.DeleteTask)
		}
	}

	authorization := router.Group("/auth")
	authorization.GET("/refresh_token", authMiddleware.RefreshHandler)

	return router
}
