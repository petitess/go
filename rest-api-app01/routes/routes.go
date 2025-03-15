package routes

import (
	"github.com/gin-gonic/gin"
	"site.org/abc/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	server.GET("/users", getUsers)
	server.GET("/users/:id", getUser)
	server.POST("/signup", createUser)
	server.PUT("/users/:id", updateUser)
	server.DELETE("users/:id", deleteUser)
	server.POST("/login", login)
	authenticated := server.Group("/")
	authenticated.Use(middlewares.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/events/:id", updateEvent)
	authenticated.DELETE("events/:id", deleteEvent)
	authenticated.POST("/events/:id/register", registerForEvent)
	authenticated.DELETE("/events/:id/register", cancelRegistration)
}
