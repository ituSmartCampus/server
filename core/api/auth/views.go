package auth

import (
	"net/http"
	"smartCampus/core/models"

	"github.com/gin-gonic/gin"
	"github.com/hlandau/passlib"
)

type ErrorResponse struct {
	Message string `json:"message"`
}

// Signup godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signUpInput body SignupSchema true "User registration input"
// @Success 200 {string} string "success"
// @Failure 400 {object} ErrorResponse "user exists"
// @Router /auth/signup [post]
func Signup(c *gin.Context) {

	input, _ := c.Get("signUpInput")

	data := input.(SignupSchema)

	pass, _ := passlib.Hash(data.Password)
	user := models.User{Email: data.Email, Password: pass}
	if models.DB.First(&user, "Email = ?", data.Email).Error == nil {

		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "user exists"})
		return
	}

	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}

// Signin godoc
// @Summary Authenticate user
// @Description Authenticate user with email and password
// @Tags Authentication
// @Accept json
// @Produce json
// @Param signInInput body SigninSchema true "User authentication input"
// @Success 200 {string} string "success"
// @Failure 400 {object} ErrorResponse "user not found"
// @Failure 400 {object} ErrorResponse "incorrect credentials"
// @Router /auth/signin [post]
func Signin(c *gin.Context) {

	input, _ := c.Get("signInInput")
	data := input.(SigninSchema)

	user := models.User{}
	if models.DB.First(&user, "Email = ?", data.Email).Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "user not found"})
		return

	}
	if user.CheckPassword(data.Password) != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "incorrect credentials"})
	}
	// token, err := utils.GenerateJWT(user.ID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{
	// 		"message": "token can not be created!",
	// 	})
	// 	return
	// }

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})

}
