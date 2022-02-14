package middlewares

import (
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JWTMIddleware(jwtService jwt.JWTService, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer"
		authHeader := c.GetHeader("Authorization")
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		if token.Valid {
			err := jwtService.ValidateRole(token, roles)
			if err != nil {
				response.SendError(c, err.Error(), "Invalid token", http.StatusUnauthorized)
				c.Abort()
			} else {
				id, _ := jwtService.GetId(token)
				c.Set("id", id)
				c.Next()
			}
		} else {
			response.SendError(c, err.Error(), "Invalid token", http.StatusUnauthorized)
			c.Abort()
		}
	}
}
