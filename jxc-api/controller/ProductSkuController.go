package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListProductSku	godoc
// @Summary		列出商品的所有货品接口
// @Description	按自由组合方式与关键字模糊查询方式列出货品
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	keyword formData string false "搜索关键字，从商品名称、全拼、拼音首字母"
// @Param	specs formData string false "商品规格，JSON格式"
// @Param	categoryId formData int false "商品分类编号"
// @Param	productId formData int false "商品编号"
// @Param	status formData int false "货品上架状态"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/list [post][get]
func ListProductSku(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "缺少商户账号参数")
	}
	status := -1
	if utils.Exists(params, "status") {
		status, _ = strconv.Atoi(params["status"])
	}
	page, size := 0, 0
	if utils.Exists(params, "page") {
		page, _ = strconv.Atoi(params["page"])
	}
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}
	return service.ListProductSku(params["shopId"], params["keyword"], params["specs"], params["productId"], params["categoryId"], status, page, size)
}

// GetProductSku	godoc
// @Summary		查询货品接口
// @Description	查询货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	skuGuid formData string false "货品全局编码"
// @Param	shopId formData string false "商户账号"
// @Param	skuId formData string false "货品自定义编码"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/get [post][get]
func GetProductSku(params map[string]string) common.Result {
	return service.GetProductSku(params["skuGuid"], params["shopId"], params["skuId"])
}

// SaveProductSku	godoc
// @Summary		保存货品接口
// @Description	保存货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Param	specs formData string true "货品规格，JSON格式，{<规格名>:<格式值>}"
// @Param	skuId formData string false "货品自定义编码，不传则自动随机生成"
// @Param	skuGuid formData string false "货品全局唯一编码，默认不传，随机生成32位码"
// @Param	name formData string false "货品名称，默认为商品名+规格值"
// @Param	prices formData string false "货品报价表，JSON格式"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/save [post][get]
func SaveProductSku(params map[string]string) common.Result {
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "specs") {
		return *common.Error(-1, "未传入货品规格，JSON格式")
	}
	return service.SaveProductSku(params["shopId"], params["productId"], params["specs"], params["skuId"], params["skuGuid"], params["name"], params["prices"])
}

// UpdateProductSku	godoc
// @Summary		修改货品接口
// @Description	修改货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	skuGuid formData string true "货品全局唯一编码"
// @Param	skuId formData string false "货品自定义编码，不传则自动随机生成"
// @Param	name formData string false "商品名称"
// @Param	skuName formData string false "货品名称"
// @Param	barCode formData string false "货品条码"
// @Param	specs formData string true "货品规格，JSON格式，{<规格名>:<格式值>}"
// @Param	prices formData string false "货品报价表，JSON格式"
// @Param	status formData int false "货品状态，0-禁售 1-可售"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/update [post][get]
func UpdateProductSku(params map[string]string) common.Result {
	status := -1
	if utils.Exists(params, "status") {
		status, _ = strconv.Atoi(params["status"])
	}
	if utils.Exists(params, "skuGuid") {
		return *common.Error(-1, "未传入货品全局唯一编码")
	}
	return service.UpdateProductSku(params["skuGuid"], params["skuId"], params["name"], params["skuName"], params["barCode"], params["specs"], params["prices"], status)
}

// DeleteProductSku	godoc
// @Summary		删除货品接口
// @Description	删除货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	skuGuid formData string false "货品全局编码"
// @Param	shopId formData string false "商户账号"
// @Param	skuId formData string false "货品自定义编码"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/del [post][get]
func DeleteProductSku(params map[string]string) common.Result {
	return service.DeleteProductSku(params["skuGuid"], params["shopId"], params["skuId"])
}

// DeleteProductSku	godoc
// @Summary		下移货品接口
// @Description	下移货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	skuGuid formData string true "货品全局编码"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/down [post][get]
func IncrProductSkuSortNumber(params map[string]string) common.Result {
	return service.IncrProductSkuSortNumber(params["skuGuid"])
}

// DeleteProductSku	godoc
// @Summary		上移货品接口
// @Description	上移货品接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	skuGuid formData string true "货品全局编码"
// @Success 200 {object} model.ProductSku	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/sku/up [post][get]
func DecrProductSkuSortNumber(params map[string]string) common.Result {
	return service.DecrProductSkuSortNumber(params["skuGuid"])
}
