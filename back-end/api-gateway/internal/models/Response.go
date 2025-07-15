package models

import (
	"github.com/gin-gonic/gin"
)

type ResponseDataType struct {
	StatusCode int    `json:"StatusCode"`
	Data       any    `json:"Data"`
	Message    string `json:"Message"`
}

func Response(c *gin.Context, httpCode int, data any) {
	if httpCode >= 200 && httpCode <= 299 {
		c.JSON(httpCode, ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Successful",
		})
	} else if httpCode >= 400 && httpCode <= 499 {
		c.JSON(httpCode, ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Failure",
		})
	} else if httpCode >= 500 && httpCode <= 599 {
		c.JSON(httpCode, ResponseDataType{
			StatusCode: httpCode,
			Data:       data,
			Message:    "Server Error !",
		})
	}
}
