package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	_ "ququ.im/jxc-base/model"
	"ququ.im/jxc-product/service"
	"strconv"
)

// ListProductSpecs	godoc
// @Summary		列出商品的所有规格指标接口
// @Description	列出商品的所有规格指标接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品自定义编码"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/list [post][get]
func ListProductSpecs(params map[string]string) common.Result {
	return service.ListProductSpecs(params["shopId"], params["productId"])
}

// GetProductSpecs	godoc
// @Summary		查询商品的规格指标接口
// @Description	查询商品的规格指标接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "商品规格编号"
// @Param	shopId formData string false "商户账号"
// @Param	productId formData string false "商品编号"
// @Param	specs formData string false "商品规格名"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/get [post][get]
func GetProductSpecs(params map[string]string) common.Result {
	id, _ := strconv.Atoi(params["id"])
	return service.GetProductSpecs(id, params["shopId"], params["productId"], params["specs"])
}

// SaveProductSpecs	godoc
// @Summary		保存规格指标接口
// @Description	保存规格指标接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Param	specs formData string true "商品规格名"
// @Param	values formData string true "商品规格值列表，JSON数组格式"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/save [post][get]
func SaveProductSpecs(params map[string]string) common.Result {
	return service.SaveProductSpecs(params["shopId"], params["productId"], params["specs"], params["values"])
}

// UpdateProductSpecs	godoc
// @Summary		修改商品规格指标接口
// @Description	修改商品规格指标接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品规格编号"
// @Param	shopId formData string false "商户账号"
// @Param	productId formData string false "商品编号"
// @Param	specs formData string false "商品规格名"
// @Param	values formData string false "商品规格值列表，JSON数组格式"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/update [post][get]
func UpdateProductSpecs(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品规格编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.UpdateProductSpecs(id, params["shopId"], params["productId"], params["specs"], params["values"])
}

// DeleteProductSpecs	godoc
// @Summary		删除商品规格指标接口
// @Description	删除商品规格指标接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品规格编号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/del [post][get]
func DeleteProductSpecs(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品规格编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteProductSpecs(id)
}

// AddSpecsValue	godoc
// @Summary		添加商品规格值接口
// @Description	添加商品规格值接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品规格编号"
// @Param	value formData string true "单个商品规格值"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/add [post][get]
func AddSpecsValue(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品规格编号参数")
	}
	if !utils.Exists(params, "value") {
		return *common.Error(-1, "未传入商品规格值参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.AddSpecsValue(id, params["value"])
}

// RemoveSpecsValue	godoc
// @Summary		移除商品规格值接口
// @Description	移除商品规格值接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品规格编号"
// @Param	value formData string true "要移除的单个商品规格值"
// @Success 200 {object} model.ProductSpecs	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/specs/remove [post][get]
func RemoveSpecsValue(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品规格编号参数")
	}
	if !utils.Exists(params, "value") {
		return *common.Error(-1, "未传入商品规格值参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.RemoveSpecsValue(id, params["value"])
}
