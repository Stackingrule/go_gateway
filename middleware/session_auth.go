package middleware

import (
	"errors"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SessionAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
<<<<<<< HEAD
		if name,ok:=session.Get("user").(string);!ok||name==""{
=======
		if name, ok := session.Get("user").(string); !ok || name == "" {
>>>>>>> 75546dc0f54b4f6d4ece0208a542fdec4d21faa0
			ResponseError(c, InternalErrorCode, errors.New("user not login"))
			c.Abort()
			return
		}
		c.Next()
	}
}
