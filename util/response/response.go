package response
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type socketResponse struct {
	Code int  `json:"code"`
	Action string `json:"action"`
	NextAction string `json:"nextaction"`
	Msg string `json:"msg"`
	Data  interface{} `json:"data"`
}
type Response struct{
	Code int  `json:"code"`
	Msg string `json:"msg"`
	Data  interface{} `json:"data"`
}


func Resp(c *gin.Context,code int,msg string,data interface{}){
	message:= GetErrorMessage(code,msg)
	jsonMap := Response{
		Code:code,
		Msg:message,
		Data:data,
	}
	c.JSON(http.StatusOK,jsonMap)
	return
}



const (
	ERROR   = 7
	SUCCESS = 0
)

func Result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		data,
	})
}

func Ok(c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func OkWithMessage(message string, c *gin.Context) {
	Result(SUCCESS, map[string]interface{}{}, message, c)
}

func OkWithData(data interface{}, c *gin.Context) {
	Result(SUCCESS, data, "操作成功", c)
}

func OkDetailed(data interface{}, message string, c *gin.Context) {
	Result(SUCCESS, data, message, c)
}

func Fail(c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func FailWithMessage(message string, c *gin.Context) {
	Result(ERROR, map[string]interface{}{}, message, c)
}

func FailWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	Result(code, data, message, c)
}
