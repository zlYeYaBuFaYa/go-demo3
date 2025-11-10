package response

import (
	"github.com/gin-gonic/gin"
)

type Res struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func SuccessResponse(c *gin.Context, data interface{}) {
	if data == nil {
		c.JSON(200, Res{Code: 0, Message: "success"})
	} else {
		c.JSON(200, Res{Code: 0, Message: "success", Data: data})
	}
}

func FailResponse(c *gin.Context, msg string) {
	c.JSON(200, Res{Code: -1, Message: msg})
}
