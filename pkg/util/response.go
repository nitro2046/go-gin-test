package util

import (
	"github.com/gin-gonic/gin"
	"goImooc/go-gin-test/pkg/e"
)

/*
   数据返回信息的model，格式如下
*/
type Response struct {
	Code int         //自定义编码
	Msg  string      //自定义消息
	Data interface{} //返回的数据
}

func ResponseWithJson(code int, data interface{}, c *gin.Context) {
	c.JSON(200, &Response{
		Code: code,
		Msg:  e.GetMsg(code),
		Data: data,
	})
}
