package v1

import (
	"github.com/gin-gonic/gin"
	"goImooc/go-gin-test/pkg/e"
)

//app更新接口
func GetAppVersionTest(c *gin.Context) {

	c.JSON(e.SUCCESS, gin.H{
		"Code": e.SUCCESS,
		"Msg":  e.GetMsg(e.SUCCESS),
		"Data": "返回数据成功",
	})
}
