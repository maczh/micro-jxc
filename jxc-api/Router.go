package main

import (
	"github.com/ekyoung/gin-nice-recovery"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"net/http"
	"ququ.im/common"
	"ququ.im/jxc-api/aop"
	"ququ.im/jxc-api/controller"
	_ "ququ.im/jxc-api/docs"
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

	//处理全局异常
	engine.Use(nice.Recovery(recoveryHandler))

	//设置404返回的内容
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, *common.Error(-1, "请求的方法不存在"))
	})

	var result common.Result
	//添加所需的路由映射
	//出库类型
	engine.Any("/type/delivery/list", func(c *gin.Context) {
		result = controller.ListDeliveryTypeByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/delivery/save", func(c *gin.Context) {
		result = controller.SaveDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/delivery/update", func(c *gin.Context) {
		result = controller.UpdateDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/delivery/del", func(c *gin.Context) {
		result = controller.DeleteDeliveryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//入库类型
	engine.Any("/type/entry/list", func(c *gin.Context) {
		result = controller.ListEntryTypeByShopId(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/entry/save", func(c *gin.Context) {
		result = controller.SaveEntryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/entry/update", func(c *gin.Context) {
		result = controller.UpdateEntryType(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/type/entry/del", func(c *gin.Context) {
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
	//单位换算规则
	engine.Any("/unit/scale/list", func(c *gin.Context) {
		result = controller.ListProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/scale/save", func(c *gin.Context) {
		result = controller.SaveProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/scale/update", func(c *gin.Context) {
		result = controller.UpdateProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/scale/del", func(c *gin.Context) {
		result = controller.DeleteProductUnit(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/unit/scale/get", func(c *gin.Context) {
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
	//库存
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
	//入库单
	engine.Any("/entry/list", func(c *gin.Context) {
		result = controller.ListStockEntryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/note/get", func(c *gin.Context) {
		result = controller.GetStockEntryNote(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/row/save", func(c *gin.Context) {
		result = controller.SaveOneStockEntryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/note/save", func(c *gin.Context) {
		result = controller.SaveStockEntryNote(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/update", func(c *gin.Context) {
		result = controller.UpdateStockEntryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/entry/del", func(c *gin.Context) {
		result = controller.DeleteStockEntryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//出库单
	engine.Any("/delivery/list", func(c *gin.Context) {
		result = controller.ListStockDeliveryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/note/get", func(c *gin.Context) {
		result = controller.GetStockDeliveryNote(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/row/save", func(c *gin.Context) {
		result = controller.SaveOneStockDeliveryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/note/save", func(c *gin.Context) {
		result = controller.SaveStockDeliveryNote(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/update", func(c *gin.Context) {
		result = controller.UpdateStockDeliveryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/delivery/del", func(c *gin.Context) {
		result = controller.DeleteStockDeliveryRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	//移库
	engine.Any("/move/list", func(c *gin.Context) {
		result = controller.ListStockMove(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/get", func(c *gin.Context) {
		result = controller.GetStockMoveRow(c.GetParamMap())
		c.JSON(http.StatusOK, result)
	})
	engine.Any("/move/stock", func(c *gin.Context) {
		result = controller.MoveSkuStock(c.GetParamMap())
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
