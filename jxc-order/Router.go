package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"ququ.im/common"
	"ququ.im/jxc-order/aop"
	"ququ.im/jxc-order/controller"
	_ "ququ.im/jxc-order/docs"
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
	//入库
	engine.Any("/entry/list", func(c *gin.Context) {
		result = controller.ListStockEntry(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/get", func(c *gin.Context) {
		result = controller.GetStockEntry(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/save", func(c *gin.Context) {
		result = controller.SaveStockEntry(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/update", func(c *gin.Context) {
		result = controller.UpdateStockEntry(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/del", func(c *gin.Context) {
		result = controller.DeleteStockEntry(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//出库
	engine.Any("/delivery/list", func(c *gin.Context) {
		result = controller.ListStockDelivery(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/get", func(c *gin.Context) {
		result = controller.GetStockDelivery(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/save", func(c *gin.Context) {
		result = controller.SaveStockDelivery(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/update", func(c *gin.Context) {
		result = controller.UpdateStockDelivery(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/del", func(c *gin.Context) {
		result = controller.DeleteStockDelivery(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//移库
	engine.Any("/move/list", func(c *gin.Context) {
		result = controller.ListStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/get", func(c *gin.Context) {
		result = controller.GetStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/save", func(c *gin.Context) {
		result = controller.SaveStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/update", func(c *gin.Context) {
		result = controller.UpdateStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/del", func(c *gin.Context) {
		result = controller.DeleteStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *common.Error(-1, "系统异常，请联系客服"))
}
