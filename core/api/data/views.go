package data

import (
	"net/http"
	"smartCampus/core/models"

	"github.com/gin-gonic/gin"
)

type ListSuccessResponse struct {
	Message string        `json:"message"`
	Datas   []models.Data `json:"datas,omitempty"`
}
type GetSuccessResponse struct {
	Message string      `json:"message"`
	Data    models.Data `json:"data,omitempty"`
}
type CreateSuccessResponse struct {
	Message string `json:"message"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

// List godoc
// @Summary Get a list of data records
// @Description Retrieve a list of data records from the database
// @Tags Data
// @Produce json
// @Success 200 {object} ListSuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/data/list [get]
func List(c *gin.Context) {

	var datas []models.Data
	if err := models.DB.Model(&models.Data{}).Find(&datas).Error; err != nil {
		c.JSON(400, ErrorResponse{Message: err.Error()})
		return
	}
	if len(datas) == 0 {
		c.JSON(400, ErrorResponse{Message: "data yok"})
		return
	}
	c.JSON(http.StatusOK, ListSuccessResponse{Message: "success", Datas: datas})
}

// Get godoc
// @Summary Get a single data record by ID
// @Description Retrieve a data record from the database based on its ID
// @Tags Data
// @Produce json
// @Param id path string true "Data ID"
// @Success 200 {object} GetSuccessResponse
// @Failure 400 {object} ErrorResponse
// @Router /api/data/get/{id} [get]
func Get(c *gin.Context) {

	id := c.Param("id")
	data := models.Data{}
	if models.DB.First(&data, "ID = ?", id).Error != nil {
		c.JSON(400, ErrorResponse{Message: "sensor verileri cekilemedi"})
		return
	}
	c.JSON(http.StatusOK, GetSuccessResponse{Message: "success", Data: data})

}

type CreateDataSchema struct {
	MeasuredValue string `json:"measuredValue" binding:"required"`
	SensorType    string `json:"sensorType" binding:"required"`
	Time          string `json:"time" binding:"required"`
	Location      string `json:"location" binding:"required"`
	Email         string `json:"email" binding:"required,email"`
}

// Create godoc
// @Summary Create multiple data records
// @Description Create multiple data records in the database
// @Tags Data
// @Accept json
// @Produce json
// @Param CreateDataInput body []CreateDataSchema true "Array of data records to create"
// @Success 200 {object} CreateSuccessResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/data/create [post]
func Create(c *gin.Context) {
	data := c.MustGet("CreateDataInput").([]CreateDataSchema)

	batchSize := 100

	result := models.DB.CreateInBatches(&data, batchSize)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Message: "Failed to create records in the database"})
		return
	}
	c.JSON(http.StatusOK, CreateSuccessResponse{Message: "success"})
}
