package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"ququ.im/common"
	"ququ.im/jxc-stock/aop"
	"ququ.im/jxc-stock/controller"
	_ "ququ.im/jxc-stock/docs"
)

/**
统一路由映射入口
*/
func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	engine := gin.Default()

	//设置接口日志
	engine.Use(aop.SetRequestLogger())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, *common.Error(-1, "请求的方法不存在"))
	})

	var result common.Result
	//添加所需的路由映射
	engine.Any("/stock/list", func(c *gin.Context) {
		result = controller.ListSkuStock(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/stock/inout", func(c *gin.Context) {
		result = controller.InOutSkuStock(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/stock/num", func(c *gin.Context) {
		result = controller.GetSkuStockNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/stock/get", func(c *gin.Context) {
		result = controller.GetSkuStock(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *common.Error(-1, "系统异常，请联系客服"))
}
