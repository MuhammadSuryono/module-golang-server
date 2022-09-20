package Exception

import (
	"fmt"
	"github.com/MuhammadSuryono/module-golang-server/http/response"
	"github.com/gin-gonic/gin"
)

var ErrorMessage interface{}

func GetError() {
	ErrorMessage = recover()
}

func GetErrorJson(c *gin.Context) {
	err := recover()
	if err != nil {
		c.JSON(response.INTERNAL_SERVER_ERROR_CODE, response.FailureResponse(
			response.INTERNAL_SERVER_ERROR_STATUS,
			fmt.Sprintf("Error: %v", err),
			nil))
		return
	}
}
