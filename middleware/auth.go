package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/muhammadsarimin/indocyber-api/models"
	log "github.com/sirupsen/logrus"
)

func BasicAuth() gin.HandlerFunc {
	return func(c *gin.Context) {

		// check basic auth
		user, pass, ok := c.Request.BasicAuth()
		if !ok || !isUserAuthenticated(user, pass) {

			log.Error("Unauthorized", fmt.Errorf("unauthorized"))
			c.JSON(401, models.Response{
				ResponseCode:    "401",
				ResponseMessage: "Unauthorized",
			})

			c.Abort()
			return
		}

		c.Next()
	}
}

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
	"user3": "password3",
	"user4": "password4",
}

func isUserAuthenticated(user, pass string) bool {

	if password, ok := users[user]; ok && password == pass {
		return true
	}

	return false
}
