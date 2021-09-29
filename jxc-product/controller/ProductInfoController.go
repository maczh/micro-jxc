package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	_ "ququ.im/jxc-base/model"
	"ququ.im/jxc-product/service"
	"strconv"
)

// ListProductInfo	godoc
// @Summary		列出商品信息接口
// @Description	列出商品信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	keyword formData string false "模糊搜索关键字：商品名称/全拼(小写)/拼音首字母(大写)"
// @Param	categoryId formData string false "商品分类编号"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} model.ProductInfo	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/product/list [post][get]
func ListProductInfo(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	page, size := 0, 0
	if utils.Exists(params, "page") {
		page, _ = strconv.Atoi(params["page"])
	}
	if utils.Exists(params, "size") {
		size, _ = strconv.Atoi(params["size"])
	}
	return service.ListProductInfo(params["shopId"], params["keyword"], params["categoryId"], page, size)
}

// GetProductInfo	godoc
// @Summary		查询商品信息接口
// @Description	查询商品信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "商品记录号"
// @Param	shopId formData string false "商户账号"
// @Param	productId formData string false "商品自定义编号"
// @Param	barCode formData string false "商品条码"
// @Success 200 {object} model.ProductInfo	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/product/get [post][get]
func GetProductInfo(params map[string]string) common.Result {
	id, _ := strconv.Atoi(params["id"])
	return service.GetProductInfo(id, params["shopId"], params["productId"], params["barCode"])
}

// SaveProductInfo	godoc
// @Summary		保存商品信息接口
// @Description	保存商品信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	categoryId formData string true "商品分类编号"
// @Param	name formData string true "商品名称"
// @Param	productId formData string false "商品自定义编号，默认系统自动生成"
// @Param	barCode formData string false "商品条码"
// @Param	baseUnit formData string true "商品基础单位，最小单位"
// @Success 200 {object} model.ProductInfo	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/product/save [post][get]
func SaveProductInfo(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "categoryId") {
		return *common.Error(-1, "缺少商品分类编码参数")
	}
	if !utils.Exists(params, "name") {
		return *common.Error(-1, "未传入商品名称参数")
	}
	if !utils.Exists(params, "baseUnit") {
		return *common.Error(-1, "未传入商品基础单位")
	}
	return service.SaveProductInfo(params["productId"], params["shopId"], params["name"], params["baseUnit"], params["barCode"], params["categoryId"])
}

// UpdateProductInfo	godoc
// @Summary		修改商品信息接口
// @Description	修改商品信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品编号"
// @Param	categoryId formData string false "商品分类编码"
// @Param	productId formData string false "商品自定义编号"
// @Param	name formData string false "商品名称"
// @Param	baseUnit formData string false "商品基础单位，最小单位"
// @Param	barCode formData string false "商品条码"
// @Success 200 {object} model.ProductInfo	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/product/update [post][get]
func UpdateProductInfo(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.UpdateProductInfo(id, params["productId"], params["categoryId"], params["name"], params["baseUnit"], params["barCode"])
}

// DeleteProductInfo	godoc
// @Summary		删除商品信息接口
// @Description	删除商品信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品编号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/product/del [post][get]
func DeleteProductInfo(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteProductInfo(id)
}
