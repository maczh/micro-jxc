package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListProductSupplier	godoc
// @Summary		列出商品的所有供应商接口
// @Description	列出商品的所有供应商接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Success 200 {object} model.ProductSupplier	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/supplier/list [post][get]
func ListProductSupplier(params map[string]string) common.Result {
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.ListProductSupplier(params["shopId"], params["productId"])
}

// GetProductSupplier	godoc
// @Summary		查询商品的供应商接口
// @Description	查询商品的供应商接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "商品供应商记录号"
// @Param	shopId formData string false "商户账号"
// @Param	productId formData string false "商品编号"
// @Param	supplierId formData string false "商品供应商账号"
// @Success 200 {object} model.ProductSupplier	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/supplier/get [post][get]
func GetProductSupplier(params map[string]string) common.Result {
	id := 0
	if utils.Exists(params, "id") {
		id, _ = strconv.Atoi(params["id"])
	}
	return service.GetProductSupplier(id, params["shopId"], params["productId"], params["supplierId"])
}

// SaveProductSupplier	godoc
// @Summary		保存供应商接口
// @Description	保存供应商接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	productId formData string true "商品编号"
// @Param	shopId formData string true "商户账号"
// @Param	supplierId formData string true "商品供应商账号"
// @Param	supplierName formData string true "商品供应商名称"
// @Success 200 {object} model.ProductSupplier	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/supplier/save [post][get]
func SaveProductSupplier(params map[string]string) common.Result {
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号")
	}
	if !utils.Exists(params, "supplierId") {
		return *common.Error(-1, "未传入商品供应商账号")
	}
	if !utils.Exists(params, "supplierName") {
		return *common.Error(-1, "未传入商品供应商名称")
	}
	return service.SaveProductSupplier(params["shopId"], params["productId"], params["supplierId"], params["supplierName"])
}

// UpdateProductSupplier	godoc
// @Summary		修改商品供应商接口
// @Description	修改商品供应商接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品供应商id"
// @Param	productId formData string false "商品编号"
// @Param	shopId formData string false "商户账号"
// @Param	supplierId formData string false "商品供应商账号"
// @Param	supplierName formData string false "商品供应商名称"
// @Success 200 {object} model.ProductSupplier	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/supplier/update [post][get]
func UpdateProductSupplier(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品供应商id")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.UpdateProductSupplier(id, params["shopId"], params["supplierId"], params["supplierName"], params["productId"])
}

// DeleteProductSupplier	godoc
// @Summary		删除商品供应商接口
// @Description	删除商品供应商接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品供应商编号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/supplier/del [post][get]
func DeleteProductSupplier(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品供应商编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteProductSupplier(id)
}
