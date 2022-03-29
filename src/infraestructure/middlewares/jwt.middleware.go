package middlewares

import (
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
			c.String(http.StatusBadRequest, "Not bearer token found")
			c.Abort()
			return
		}
		tokenString := authHeader[len(BEARER_SCHEMA):]
		token, _ := jwtService.ValidateToken(tokenString)
		// fmt.Println("tk", tokenString, token, err, authHeader)
		if token.Valid {
			err := jwtService.ValidateRole(token, roles)
			if err != nil {
				c.String(http.StatusUnauthorized, "Invalid token"+err.Error())
				c.Abort()
				return
			} else {
				err, id := jwtService.GetId(token)
				if err != nil {
					c.String(http.StatusUnauthorized, "Invalid token")
					return
				}
				c.Set("id", id)
				c.Next()
				return
			}
		} else {
			c.String(http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}
	}
}
