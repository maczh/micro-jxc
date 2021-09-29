package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListProductUnit	godoc
// @Summary		列出商品的所有单位换算规则接口
// @Description	列出商品的所有单位换算规则接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Success 200 {object} model.ProductUnit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/scale/list [post][get]
func ListProductUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	return service.ListProductUnit(params["shopId"], params["productId"])
}

// GetProductUnit	godoc
// @Summary		查询商品的单位换算规则接口
// @Description	查询商品的单位换算规则接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Param	unit formData string true "商品辅助单位，较大的单位"
// @Param	baseUnit formData string true "商品基础单位，最小单位"
// @Success 200 {object} model.ProductUnit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/scale/get [post][get]
func GetProductUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "unit") {
		return *common.Error(-1, "未传入商品辅助单位")
	}
	if !utils.Exists(params, "baseUnit") {
		return *common.Error(-1, "未传入商品基础单位")
	}
	return service.GetProductUnit(0, params["shopId"], params["productId"], params["unit"], params["baseUnit"])
}

// SaveProductUnit	godoc
// @Summary		保存单位换算规则接口
// @Description	保存单位换算规则接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Param	scale formData int true "商品单位换算比例值: 1辅助单位=scale个基础单位"
// @Param	unit formData string true "商品辅助单位，较大的单位"
// @Param	baseUnit formData string true "商品基础单位，最小单位"
// @Success 200 {object} model.ProductUnit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/scale/save [post][get]
func SaveProductUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "unit") {
		return *common.Error(-1, "未传入商品辅助单位")
	}
	if !utils.Exists(params, "baseUnit") {
		return *common.Error(-1, "未传入商品基础单位")
	}
	if !utils.Exists(params, "scale") {
		return *common.Error(-1, "未传入商品单位换算比例值: 1辅助单位=scale个基础单位")
	}
	scale, _ := strconv.Atoi(params["scale"])
	return service.SaveProductUnit(params["shopId"], params["productId"], params["unit"], params["baseUnit"], scale)
}

// UpdateProductUnit	godoc
// @Summary		修改商品单位换算规则接口
// @Description	修改商品单位换算规则接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	productId formData string true "商品编号"
// @Param	scale formData int true "商品单位换算比例值: 1辅助单位=scale个基础单位"
// @Param	unit formData string true "商品辅助单位，较大的单位"
// @Param	baseUnit formData string true "商品基础单位，最小单位"
// @Success 200 {object} model.ProductUnit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/scale/update [post][get]
func UpdateProductUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "productId") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	if !utils.Exists(params, "unit") {
		return *common.Error(-1, "未传入商品辅助单位")
	}
	if !utils.Exists(params, "baseUnit") {
		return *common.Error(-1, "未传入商品基础单位")
	}
	if !utils.Exists(params, "scale") {
		return *common.Error(-1, "未传入商品单位换算比例值: 1辅助单位=scale个基础单位")
	}
	scale, _ := strconv.Atoi(params["scale"])
	return service.UpdateProductUnit(params["shopId"], params["productId"], params["unit"], params["baseUnit"], scale)
}

// DeleteProductUnit	godoc
// @Summary		删除商品单位换算规则接口
// @Description	删除商品单位换算规则接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品单位换算编号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/scale/del [post][get]
func DeleteProductUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteProductUnit(id)
}
