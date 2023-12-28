package auth

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func SignupValidation(c *gin.Context) {

	var input SignupSchema
	if err := c.ShouldBindJSON(&input); err != nil {

		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "all required fields are empty"})
			c.Abort()
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("signUpInput", input)
	c.Next()
}

type SigninSchema struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=8,max=20"`
}

func SigninValidation(c *gin.Context) {
	var input SigninSchema
	if err := c.ShouldBindJSON(&input); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "all required fields are empty"})
			c.Abort()
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}
	c.Set("signInInput", input)
	c.Next()
}
