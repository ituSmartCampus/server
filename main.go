// @title Smart Campus API
// @version 1.0
// @description API for managing Smart Campus data and authentication
// @host localhost:3000
// @BasePath /api/v1
package main

import (
	"smartCampus/core/api/auth"
	"smartCampus/core/api/data"
	"smartCampus/core/utils"

	docs "smartCampus/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var router = gin.Default()

func main() {
	utils.CreateTables()
	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	authApi := router.Group("/api/auth")
	{
		authApi.POST("/signup", auth.SignupValidation, auth.Signup)
		authApi.POST("/signin", auth.SigninValidation, auth.Signin)
		authApi.POST("/checkemail/:email", auth.Checkemail)

	}
	dataApi := router.Group("/api/data")
	{
		dataApi.POST("/create", data.CreateDataMiddleware, data.Create)
		dataApi.GET("/list", data.List)
		dataApi.GET("/list/moisture", data.ListMoisture)
		dataApi.GET("/list/temperature", data.ListTemperature)
		dataApi.GET("/list/air", data.ListAir)
		dataApi.GET("/get/:id", data.Get)
	}

	// dashboardApi := router.Group("/api/dashboard")
	// {
	// 	dashboardApi.GET("")

	// }
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run("0.0.0.0:3131")
}
