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

func NewResponse(c *gin.Context) *JsonMsgResponse {
	return &JsonMsgResponse{Ctx: c}
}

func (r *JsonMsgResponse) Success(data interface{}) {
	res := JsonMsgResult{}
	res.Code = SUCCESS_CODE
	res.Message = SUCCESS_MSG
	res.Data = data
	r.Ctx.JSON(SUCCESS_CODE, res)
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
	r.Ctx.JSON(SUCCESS_CODE, res)
}
