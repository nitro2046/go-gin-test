package routers

import (
	"github.com/gin-gonic/gin"
	"goImooc/go-gin-test/pkg/setting"
	v1 "goImooc/go-gin-test/routers/v1"
)

/*
   初始化路由
*/
func InitRouter() *gin.Engine {
	r := gin.New()        //创建gin框架路由实例
	r.Use(gin.Logger())   //使用gin框架中的打印中间件
	r.Use(gin.Recovery()) //使用gin框架中的恢复中间件，可以从任何恐慌中恢复，如果有，则写入500

	gin.SetMode(setting.ServerSetting.RunMode) //设置运行模式，debug或release,如果放在gin.New或者gin.Default之后，还是会打印一些信息的。放之前则不会

	apiv1 := r.Group("/api/v1") //路由分组，apiv1代表v1版本的路由组

	{
		apiv1.GET("version", v1.GetAppVersionTest) //app版本升级
		//登录/注册接口
	}

	apiv1.POST("auth", v1.AuthStore)

	return r
}
