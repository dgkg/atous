package middleware

import (
	"atous/model"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func AccessLevel(signKey []byte, authRoles []model.Role) gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the auth header
		authValue := c.GetHeader("Authorization")
		if authValue == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// check if the auth header is valid
		if !strings.Contains(authValue, "Bearer ") {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		if len(authValue) < 100 {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		// get the jwt value
		jwtValue := authValue[7:]

		//claims := make(jwt.MapClaims)
		token, err := jwt.Parse(jwtValue, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return signKey, nil
		})
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
			return
		}

		// infered type is jwt.MapClaims
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			log.Println("middleware: ", claims["id"], claims["role_type"], claims["first_name"])
			if role, ok := claims["role_type"]; ok {
				fmt.Println(role)
				for _, r := range authRoles {
					if r.String() == role {
						c.Next()
						return
					}
				}
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		} else {
			c.AbortWithError(http.StatusUnauthorized, err)
			fmt.Println(err)
		}
	}
}
