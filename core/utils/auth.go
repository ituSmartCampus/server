package utils

import (
	"errors"
	"net/http"
	"smartCampus/core/models"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("smartCamp.82lsSum!!!..ii")

type JWTClaim struct {
	UserID uint
	jwt.StandardClaims
}

func ValidateToken(signedToken string) (user models.User, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("token expired")
		return
	} else {
		if models.DB.First(&user, "ID = ?", claims.UserID).Error != nil {
			return user, errors.New("user doesn't exists")
		}
		return user, nil
	}
}

func GenerateJWT(id uint) (tokenString string, err error) {
	expirationTime := time.Now().Add(time.Minute * 60 * 24 * 7)
	claims := &JWTClaim{
		UserID: id,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err = token.SignedString(jwtKey)
	return
}

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.GetHeader("Authorization") != "" {
			a := strings.Split(c.GetHeader("Authorization"), " ")
			if len(a) == 2 {
				user, err := ValidateToken(a[1])
				if err != nil {
					c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
				}
				c.Set("user", user)
				c.Next()
			} else {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			}
		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		}
	}
}
