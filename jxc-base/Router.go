package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	"ququ.im/common"
	"ququ.im/jxc-base/aop"
	"ququ.im/jxc-base/controller"
	_ "ququ.im/jxc-base/docs"
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
	//添加跨域处理
	engine.Use(aop.Cors())

	//添加swagger支持
	engine.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//engine.GET("/docs/*any", swagger_skin.HandleReDoc)

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, *common.Error(-1, "请求的方法不存在"))
	})

	var result common.Result
	//添加所需的路由映射
	//出库类型
	engine.Any("/deliverytype/list", func(c *gin.Context) {
		result = controller.ListDeliveryTypeByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/deliverytype/save", func(c *gin.Context) {
		result = controller.SaveDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/deliverytype/update", func(c *gin.Context) {
		result = controller.UpdateDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/deliverytype/del", func(c *gin.Context) {
		result = controller.DeleteDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//入库类型
	engine.Any("/entrytype/list", func(c *gin.Context) {
		result = controller.ListEntryTypeByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entrytype/save", func(c *gin.Context) {
		result = controller.SaveEntryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entrytype/update", func(c *gin.Context) {
		result = controller.UpdateEntryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entrytype/del", func(c *gin.Context) {
		result = controller.DeleteEntryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//单位
	engine.Any("/unit/list", func(c *gin.Context) {
		result = controller.ListUnitByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/save", func(c *gin.Context) {
		result = controller.SaveUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/update", func(c *gin.Context) {
		result = controller.UpdateUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/del", func(c *gin.Context) {
		result = controller.DeleteUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//仓库
	engine.Any("/storage/list", func(c *gin.Context) {
		result = controller.ListShopStorageByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/save", func(c *gin.Context) {
		result = controller.SaveShopStorage(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/update", func(c *gin.Context) {
		result = controller.UpdateShopStorage(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/del", func(c *gin.Context) {
		result = controller.DeleteShopStorage(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/get", func(c *gin.Context) {
		result = controller.GetShopStorage(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/last", func(c *gin.Context) {
		result = controller.GetStorageLastOperator(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/up", func(c *gin.Context) {
		result = controller.DecrShopStorageSortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/storage/down", func(c *gin.Context) {
		result = controller.IncrShopStorageSortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *common.Error(-1, "系统异常，请联系客服"))
}
