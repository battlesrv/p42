package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
)

// ReadRequest ..
func ReadRequest(c *gin.Context, out interface{}) error {
	b, err := c.GetRawData()
	// defer c.Request.Body.Close()

	if err != nil {
		return err
	}

	if err = json.Unmarshal(b, &out); err != nil {
		return err
	}

	validate := validator.New()
	return validate.Struct(out)
}

// ResponseError ..
func ResponseError(c *gin.Context, err error) {
	c.JSON(http.StatusOK, gin.H{
		"ok": false,
		"result": map[string]interface{}{
			"message": err,
		},
	})
	c.Next()
}

// ResponseSuccess ..
func ResponseSuccess(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"ok":     true,
		"result": data,
	})
	c.Next()
}
