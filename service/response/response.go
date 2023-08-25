package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func Success(c *gin.Context, data interface{}, Message string) {
	c.JSON(http.StatusOK, Response{
		200,
		data,
		Message,
	})
}
func Error(c *gin.Context, data interface{}, Message string) {
	c.JSON(400, Response{
		400,
		data,
		Message,
	})
}
