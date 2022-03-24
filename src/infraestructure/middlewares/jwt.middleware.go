package middlewares

import (
	"golang-gingonic-hex-architecture/src/infraestructure/response"
	"golang-gingonic-hex-architecture/src/infraestructure/utils/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

//Roles
const (
	ADMINISTRATOR        = "ADMINISTRATOR"
	TECHNICAL_WORKER     = "TECHNICAL_WORKER"
	LEGAL_REPRESENTATIVE = "LEGAL_REPRESENTATIVE"
	PUBLICATION_WRITER   = "PUBLICATION_WRITER"
	LEATHER              = "LEATHER"
)

func JWTMIddleware(jwtService jwt.JWTService, roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BEARER_SCHEMA = "Bearer "
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || len(authHeader) <= 8 {
			response.SendError(c, "Not bearer token found", "Not bearer token found", http.StatusUnauthorized)
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, err := jwtService.ValidateToken(tokenString)
		// fmt.Println("tk", tokenString, token, err, authHeader)
		if token.Valid {
			err := jwtService.ValidateRole(token, roles)
			if err != nil {
				response.SendError(c, err.Error(), "Invalid token", http.StatusUnauthorized)
				c.Abort()
				return
			} else {
				err, id := jwtService.GetId(token)
				if err != nil {
					response.SendError(c, err.Error(), "Invalid token", http.StatusUnauthorized)
					return
				}
				c.Set("id", id)
				c.Next()
				return
			}
		} else {
			response.SendError(c, err.Error(), "Invalid token", http.StatusUnauthorized)
			c.Abort()
			return
		}
	}
}
