package service

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/barkimedes/go-deepcopy"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"atous/db"
	"atous/geo"
	"atous/hash"
	"atous/model"
)

type ServiceUser struct {
	db         *db.DB
	geo        geo.Geocoder
	jwtKeySign []byte
}

func initServiceUser(r *gin.Engine, db *db.DB, geocoder geo.Geocoder, jwtKeySign []byte) {
	su := &ServiceUser{
		db:         db,
		geo:        geocoder,
		jwtKeySign: jwtKeySign,
	}
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

	// empty passwords
	for i := 0; i < len(us); i++ {
		us[i].Password = nil
	}

	c.JSON(http.StatusOK, gin.H{"users": us})
}

// restrives the user from the request body
func (su *ServiceUser) get(c *gin.Context) {
	id := c.Param("id")
	//user, ok := db.UserList[id]
	user, err := su.db.GetUser(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}

	// empty password
	user.Password = nil

	c.JSON(http.StatusOK, gin.H{"user": user})
}

// deletes the user from the request body
func (su *ServiceUser) delete(c *gin.Context) {
	id := c.Param("id")
	err := su.db.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	c.JSON(http.StatusAccepted, nil)
}

// updates the user from the request body
func (su *ServiceUser) update(c *gin.Context) {
	// get the user id param
	id := c.Param("id")
	// get the user from the db
	user, err := su.db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// get the body of the request and map it to a map
	newUser := map[string]interface{}{}
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// update the user with the new values
	setStringFromMap(newUser, "first_name", &user.FirstName)
	setStringFromMap(newUser, "last_name", &user.LastName)
	setStringFromMap(newUser, "email", &user.Email)
	setIntFromMap(newUser, "age", &user.Age)
	if value, ok := newUser["role_type"]; ok {
		if v, ok := value.(string); ok {
			user.RoleType = model.ToRoleType(v)
		}
	}
	// update the user in the db
	err = su.db.UpdateUser(id, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// empty password.
	user.Password = nil
	// return the updated user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

// creates a new user from the request body
func (su *ServiceUser) create(c *gin.Context) {
	// get the user from the request body
	var user model.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// copy the address from the user
	var a *model.Address
	if user.Address != nil {
		v, _ := deepcopy.Anything(user.Address)
		a = v.(*model.Address)
		user.Address = nil
	}
	// save the user in the db
	err := su.db.CreateUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go SayHello(user.FirstName, &wg)

	// save the address in the db
	if a != nil {
		// add the user id to the address
		a.UUIDOwner = user.ID
		// geocode the address
		a.Longitude, a.Latitude, err = su.geo.Geocode(a.String())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// save the address in the db
		err = su.db.CreateAddress(a)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		// set the address in the user
		user.Address = a
	}

	// empty password.
	user.Password = nil
	wg.Wait()
	// return the created user
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func SayHello(str string, wg *sync.WaitGroup) {
	time.Sleep(5 * time.Second)
	fmt.Println("serviceUser: Hello ", str)
	wg.Done()
}

// sayHi returns a message from the user
func (su *ServiceUser) sayHi(c *gin.Context) {
	// get the user id param
	id := c.Param("id")
	// get the user from the db
	user, err := su.db.GetUser(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// return the message
	c.JSON(http.StatusOK, gin.H{"message": user.SayHi()})
}

func (su *ServiceUser) login(c *gin.Context) {
	// get the login and password from the request body
	var payload model.Login
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// get the corresponding user from the db with the email
	u, err := su.db.GetUserByEmail(payload.Email)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	// check the password
	if u.Password != nil && hash.Password(payload.Password) != *u.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Wrong password"})
		return
	}

	// Create the JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         u.ID,
		"role_type":  u.RoleType,
		"first_name": u.FirstName,
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString(su.jwtKeySign)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// return the token
	c.JSON(http.StatusOK, gin.H{"jwt": tokenString})
}
