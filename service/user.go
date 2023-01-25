package service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"atous/db"
	"atous/model"
)

type ServiceUser struct {
	db *db.DB
}

func initServiceUser(r *gin.Engine, db *db.DB) {
	su := &ServiceUser{db: db}
	r.POST("/users", su.create)
	r.GET("/users", su.getList)
	r.GET("/users/:id/say-hi", su.sayHi)
	r.GET("/users/:id", su.get)
	r.DELETE("/users/:id", su.delete)
	r.PATCH("/users/:id", su.update)
	// login route
	r.POST("/user-login", su.login)
}

func (su *ServiceUser) getList(c *gin.Context) {
	us, err := su.db.GetListUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"users": us})
}

// restrives the user from the request body
func (su *ServiceUser) get(c *gin.Context) {
	id := c.Param("id")
	//user, ok := db.UserList[id]
	user, err := su.db.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// deletes the user from the request body
func (su *ServiceUser) delete(c *gin.Context) {
	id := c.Param("id")
	err := su.db.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

// updates the user from the request body
func (su *ServiceUser) update(c *gin.Context) {
	id := c.Param("id")
	user, err := su.db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	newUser := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if value, ok := newUser["first_name"]; ok {
		if v, ok := value.(string); ok {
			user.FirstName = v
		}
	}
	if value, ok := newUser["last_name"]; ok {
		if v, ok := value.(string); ok {
			user.LastName = v
		}
	}
	if value, ok := newUser["age"]; ok {
		if v, ok := value.(int); ok {
			user.Age = v
		}
	}

	if value, ok := newUser["role_type"]; ok {
		if v, ok := value.(string); ok {
			user.RoleType = model.ToRoleType(v)
		}
	}

	err = su.db.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (su *ServiceUser) create(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user = *model.NewUser(user.Email, user.Password, &model.ConfigUser{
		Age:       user.Age,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	})

	err := su.db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func (su *ServiceUser) sayHi(c *gin.Context) {
	id := c.Param("id")
	user, err := su.db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": user.SayHi()})
}

func (su *ServiceUser) login(c *gin.Context) {
	var login model.Login
	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := su.db.GetUserByEmail(login.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	if u.Password != login.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	//c.JSON(http.StatusOK, gin.H{"message": "Login successful"})
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         u.ID,
		"role_type":  u.RoleType,
		"first_name": u.FirstName,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte("secret"))

	c.JSON(http.StatusOK, gin.H{"jwt": tokenString})
}
