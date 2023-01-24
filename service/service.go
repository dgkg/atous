package service

import "github.com/gin-gonic/gin"

func New(r *gin.Engine) {
	r.POST("/users", createUser)
	r.GET("/users", getListUsers)
	r.GET("/users/:id/say-hi", sayHiUser)
	r.GET("/users/:id", getUser)
	r.DELETE("/users/:id", deleteUser)
	r.PATCH("/users/:id", updateUser)
}
