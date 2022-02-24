package jwt

import (
	"fmt"
	"os"
	"sort"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type JWTService interface {
	GenerateToken(id, email string, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	ValidateRole(token *jwt.Token, roles []string) error
	GetId(token *jwt.Token) (error, string)
}

type authCustomClaims struct {
	Name string `json:"name"`
	Role string `json:"role"`
	jwt.StandardClaims
}

type jwtServices struct {
	secretKey string
}

//auth-jwt
func NewJWTAuthService() JWTService {
	return &jwtServices{
		secretKey: getSecretKey(),
	}
}

func getSecretKey() string {
	secret := os.Getenv("JWT_PASSWORD")
	if secret == "" {
		secret = "secret"
	}
	return secret
}

func (service *jwtServices) GenerateToken(id, email string, role string) string {
	claims := &authCustomClaims{
		email,
		role,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 48).Unix(),
			IssuedAt:  time.Now().Unix(),
			Id:        id,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	//encoded string
	t, err := token.SignedString([]byte(service.secretKey))
	if err != nil {
		panic(err)
	}
	return t
}

func (service *jwtServices) ValidateToken(encodedToken string) (*jwt.Token, error) {
	return jwt.Parse(encodedToken, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token %v", token.Header["alg"])
		}
		return []byte(service.secretKey), nil
	})
}

func (service *jwtServices) ValidateRole(token *jwt.Token, roles []string) (err error) {
	tokenClaims := token.Claims.(jwt.MapClaims)
	role, ok := tokenClaims["role"].(string)
	if !ok {
		err = fmt.Errorf("Error obtaining token role")
	}

	var hasPermission bool
	containRore(roles, role, &hasPermission)
	if !hasPermission {
		err = fmt.Errorf("This user doesnt have permission in this route")
	} else {
		err = nil
	}

	return err
}

func (service *jwtServices) GetId(token *jwt.Token) (err error, id string) {
	tokenClaims := token.Claims.(jwt.MapClaims)
	id, ok := tokenClaims["id"].(string)
	if !ok {
		err = fmt.Errorf("Error obtaining token role")
	}
	return err, id
}

func containRore(roles []string, role string, exist *bool) {
	i := sort.SearchStrings(roles, role)
	if i < len(roles) && roles[i] == role {
		*exist = true
	} else {
		*exist = false
	}
}
