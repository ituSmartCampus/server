package data

import (
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateDataMiddleware(c *gin.Context) {
	var data []CreateDataSchema

	if err := c.BindJSON(&data); err != nil {
		if err == io.EOF {
			c.JSON(http.StatusBadRequest, gin.H{"error": "all required fields are empty"})
			c.Abort()
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.Set("CreateDataInput", data)
	c.Next()

}
