package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"ququ.im/common"
	"ququ.im/jxc-product/aop"
	"ququ.im/jxc-product/controller"
	_ "ququ.im/jxc-product/docs"
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
	//单位换算规则
	engine.Any("/unit/list", func(c *gin.Context) {
		result = controller.ListProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/save", func(c *gin.Context) {
		result = controller.SaveProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/update", func(c *gin.Context) {
		result = controller.UpdateProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/del", func(c *gin.Context) {
		result = controller.DeleteProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/get", func(c *gin.Context) {
		result = controller.GetProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//商品分类
	engine.Any("/category/list", func(c *gin.Context) {
		result = controller.ListProductCategory(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/save", func(c *gin.Context) {
		result = controller.SaveProductCategory(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/update", func(c *gin.Context) {
		result = controller.UpdateProductCategory(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/del", func(c *gin.Context) {
		result = controller.DeleteProductCategory(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/get", func(c *gin.Context) {
		result = controller.GetProductCategory(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/up", func(c *gin.Context) {
		result = controller.DecrProductCategorySortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/category/down", func(c *gin.Context) {
		result = controller.IncrProductCategorySortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//商品
	engine.Any("/product/list", func(c *gin.Context) {
		result = controller.ListProductInfo(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/product/get", func(c *gin.Context) {
		result = controller.GetProductInfo(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/product/save", func(c *gin.Context) {
		result = controller.SaveProductInfo(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/product/update", func(c *gin.Context) {
		result = controller.UpdateProductInfo(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/product/del", func(c *gin.Context) {
		result = controller.DeleteProductInfo(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//商品规格
	engine.Any("/specs/list", func(c *gin.Context) {
		result = controller.ListProductSpecs(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/save", func(c *gin.Context) {
		result = controller.SaveProductSpecs(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/update", func(c *gin.Context) {
		result = controller.UpdateProductSpecs(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/del", func(c *gin.Context) {
		result = controller.DeleteProductSpecs(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/get", func(c *gin.Context) {
		result = controller.GetProductSpecs(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/add", func(c *gin.Context) {
		result = controller.AddSpecsValue(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/specs/remove", func(c *gin.Context) {
		result = controller.RemoveSpecsValue(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//供应商
	engine.Any("/supplier/list", func(c *gin.Context) {
		result = controller.ListProductSupplier(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/supplier/get", func(c *gin.Context) {
		result = controller.GetProductSupplier(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/supplier/save", func(c *gin.Context) {
		result = controller.SaveProductSupplier(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/supplier/update", func(c *gin.Context) {
		result = controller.UpdateProductSupplier(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/supplier/del", func(c *gin.Context) {
		result = controller.DeleteProductSupplier(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//货品
	engine.Any("/sku/list", func(c *gin.Context) {
		result = controller.ListProductSku(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/save", func(c *gin.Context) {
		result = controller.SaveProductSku(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/update", func(c *gin.Context) {
		result = controller.UpdateProductSku(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/del", func(c *gin.Context) {
		result = controller.DeleteProductSku(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/get", func(c *gin.Context) {
		result = controller.GetProductSku(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/up", func(c *gin.Context) {
		result = controller.DecrProductSkuSortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/sku/down", func(c *gin.Context) {
		result = controller.IncrProductSkuSortNumber(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})

	return engine
}

func recoveryHandler(c *gin.Context, err interface{}) {
	c.JSON(http.StatusOK, *common.Error(-1, "系统异常，请联系客服"))
}
