package response

import (
	"github.com/gin-gonic/gin"
)

type JsonMsgResponse struct {
	Ctx *gin.Context
}

type JsonMsgResult struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}
type nilStruct struct{}

const SUCCESS_CODE = 200
const SUCCESS_MSG = "成功"
const ERROR_MSG = "错误"

const code200 = 200

// Response 更加通用的返回方法 以后不用直接使用gin的返回方法
func Response(c *gin.Context, data interface{}, err error) {
	if err != nil {
		// 尝试将错误断言为*RespError类型
		if respErr, ok := err.(*RespError); ok {
			// 如果断言成功，返回RespError的Code和消息
			c.JSON(respErr.Code, respErr)
			return
		}
		// 如果错误不是*RespError类型，返回通用错误
		c.JSON(code200, JsonMsgResult{
			Code:    COMMON_FAIL.Code,
			Message: COMMON_FAIL.Msg,
			Data:    err.Error(),
		})
		return
	}
	// 如果没有错误，正常返回
	c.JSON(code200, JsonMsgResult{
		Code:    SUCCESS_CODE,
		Message: SUCCESS_MSG,
		Data:    data,
	})
}

func NewResponse(c *gin.Context) *JsonMsgResponse {
	return &JsonMsgResponse{Ctx: c}
}

func (r *JsonMsgResponse) Success(data interface{}) {
	res := JsonMsgResult{}
	res.Code = SUCCESS_CODE
	res.Message = SUCCESS_MSG
	res.Data = data
	r.Ctx.JSON(code200, res)
}

func (r *JsonMsgResponse) Error(mc MsgCode) {
	r.error(mc.Code, mc.Msg)
}

func (r *JsonMsgResponse) error(code int, message string) {
	if message == "" {
		message = ERROR_MSG
	}
	res := JsonMsgResult{}
	res.Code = code
	res.Message = message
	res.Data = nilStruct{}
	r.Ctx.JSON(code200, res)
}
