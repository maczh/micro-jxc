package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/service"
	"strconv"
)

// ListUnitByShopId	godoc
// @Summary		列出商户的所有单位接口
// @Description	列出商户的所有单位接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Success 200 {object} model.Unit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/list [post][get]
func ListUnitByShopId(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.ListUnitByShopId(params["shopId"])
}

// SaveUnit	godoc
// @Summary		保存单位接口
// @Description	保存单位接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string false "商户号"
// @Param	unit formData string true "单位"
// @Success 200 {object} model.Unit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/save [post][get]
func SaveUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "unit") {
		return *common.Error(-1, "未传入单位参数")
	}
	return service.SaveUnit(params["shopId"], params["unit"])
}

// UpdateUnit	godoc
// @Summary		更新单位接口
// @Description	更新单位接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "单位编号"
// @Param	unit formData string true "单位"
// @Success 200 {object} model.Unit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/update [post][get]
func UpdateUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入单位编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if !utils.Exists(params, "unit") {
		return *common.Error(-1, "未传入单位名称参数")
	}
	return service.UpdateUnit(id, params["unit"])
}

// DeleteUnit	godoc
// @Summary		删除单位接口
// @Description	删除单位接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "单位编号"
// @Success 200 {object} model.Unit	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/unit/del [post][get]
func DeleteUnit(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入单位编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteUnit(id)
}
