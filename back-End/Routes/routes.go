package routes

import (
	"example.com/middlewares"
	"github.com/gin-gonic/gin"
)
func Register(server *gin.Engine){
	server.Use(middlewares.CORSMiddleware())
	
	server.GET("/events", getEvents)
	server.GET("/events/:id", getEvent)
	
    server.POST("/events", middlewares.Authenticate, createEvent)
	server.PUT("/events/:id", middlewares.Authenticate,updateEvent)
	server.DELETE("/events/:id", middlewares.Authenticate,deleteEvent)

	server.GET("/personel", getPersonels)
	server.GET("/personel/:id", getPersonel)

	server.POST("/personel", middlewares.Authenticate, AddPersonel)
	server.PUT("/personel/:id", middlewares.Authenticate, updatePersonel)
	server.DELETE("/personel/:id", middlewares.Authenticate, deletePersonel)

	// server.POST("/events/:id/register", middlewares.Authenticate, registerForEvent)
	// server.DELETE("/events/:id/register", middlewares.Authenticate, cancelRegistration)


	server.POST("/signup", signup)
	server.POST("/login", login)
	server.GET("/users", getUsers)
	// server.OPTIONS("login", middlewares.CORSMiddleware(),login)
}
