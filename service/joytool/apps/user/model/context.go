package model

import (
	"fmt"
	"github.com/gin-gonic/gin"
	lkit_go "github.com/xlkness/lkit-go"
)

type MyContext struct {
	ctx *gin.Context
}

func (myCtx *MyContext) SetGinContext(ctx *gin.Context) {
	myCtx.ctx = ctx
}
func (myCtx *MyContext) GetGinContext() *gin.Context {
	return myCtx.ctx
}
func (myCtx *MyContext) ResponseParseParamsFieldFail(path string, uri string, body string, field string, value string, err error) {
	lkit_go.Errorf("[%v] [%v] [%v] [%v] [%v] ResponseParseParamsFieldFail error:%v", path, uri, body, field, value, err)
	myCtx.RespFailMessage(300, fmt.Sprintf("解析参数错误：%v", err))
}

func (myCtx *MyContext) RespSuccessJson(data interface{}) {
	myCtx.respJsonAny(200, 200, data)
}

func (myCtx *MyContext) RespSuccessJsonP(data interface{}) {
	myCtx.respJsonPAny(200, 200, data)
}

func (myCtx *MyContext) RespSuccessMessage(msg string) {
	myCtx.respJsonAny(200, 200, msg)
}
func (myCtx *MyContext) RespFailJson(errCode int, data map[string]interface{}) {
	myCtx.respJsonAny(300, errCode, data)
}

func (myCtx *MyContext) RespFailMessage(errCode int, msg string) {
	myCtx.respJsonAny(300, errCode, msg)
}

func (myCtx *MyContext) respJsonAny(code int, contentCode int, payload any) {
	message := "ok"
	if code != 200 {
		message = payload.(string)
		payload = nil
	}
	myCtx.GetGinContext().JSON(code, map[string]interface{}{
		"code":    contentCode,
		"message": message,
		"payload": payload,
	})
}

func (myCtx *MyContext) respJsonPAny(code int, contentCode int, payload any) {
	message := "ok"
	if code != 200 {
		message = payload.(string)
		payload = nil
	}
	myCtx.GetGinContext().JSONP(code, map[string]interface{}{
		"code":    contentCode,
		"message": message,
		"payload": payload,
	})
}
