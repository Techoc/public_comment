package res

import (
	"github.com/gin-gonic/gin"
	"github.com/techoc/public_comment/myerr"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := myerr.DecodeErr(err)

	//http.StatusOK这个值是200
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
