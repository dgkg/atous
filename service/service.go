package service

import (
	"atous/db"

	"github.com/gin-gonic/gin"
)

func New(r *gin.Engine, db *db.DB) {
	su := &ServiceUser{db: db}
	r.POST("/users", su.createUser)
	r.GET("/users", su.getListUsers)
	r.GET("/users/:id/say-hi", su.sayHiUser)
	r.GET("/users/:id", su.getUser)
	r.DELETE("/users/:id", su.deleteUser)
	r.PATCH("/users/:id", su.updateUser)
}
