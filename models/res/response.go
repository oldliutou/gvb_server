package res

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int         `json:"code"`
	Data any `json:"data"`
	Msg  string      `json:"msg"`
}
const (
	// 成功
	SUCCESS = 0
	ERROR = 7
)
func Result(code int, data any, msg string, c *gin.Context) {
	c.JSON(http.StatusOK,
		Response{Code: code, Data: data, Msg: msg},
	)
}

func OK(data any, msg string,c *gin.Context){
	Result(SUCCESS, data, msg,c)
}
func OKWithData(data any, c *gin.Context){
	Result(SUCCESS, data, "成功",c)
}

func OKWithMessage(msg string, c *gin.Context){
	Result(SUCCESS, map[string]interface{}{}, msg,c)
}

func Fail(data any, msg string,c *gin.Context){
	Result(ERROR, data, msg,c)
}

func FailWithMessage(msg string, c *gin.Context){
	Result(ERROR, map[string]interface{}{}, msg,c)
}
func FailWithCode(code ErrorCode, c *gin.Context){
	msg,ok := ErrorMap[code]
	if ok  {
		Result(int(code), map[string]interface{}{}, msg,c)
		return
	}
	Result(ERROR, map[string]interface{}{}, "未知错误",c)


}

