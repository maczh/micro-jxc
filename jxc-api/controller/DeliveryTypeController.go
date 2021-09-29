package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListDeliveryTypeByShopId	godoc
// @Summary		列出商户的所有出库类型接口
// @Description	列出商户的所有出库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Success 200 {object} model.DeliveryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/type/delivery/list [post][get]
func ListDeliveryTypeByShopId(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.ListDeliveryTypeByShopId(params["shopId"])
}

// SaveDeliveryType	godoc
// @Summary		保存出库类型接口
// @Description	保存出库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string false "商户号"
// @Param	deliveryType formData string true "出库类型"
// @Success 200 {object} model.DeliveryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/type/delivery/save [post][get]
func SaveDeliveryType(params map[string]string) common.Result {
	if !utils.Exists(params, "deliveryType") {
		return *common.Error(-1, "未传入出库类型参数")
	}
	return service.SaveDeliveryType(params["shopId"], params["deliveryType"])
}

// UpdateDeliveryType	godoc
// @Summary		更新出库类型接口
// @Description	更新出库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "出库类型编号"
// @Param	deliveryType formData string true "出库类型"
// @Success 200 {object} model.DeliveryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/type/delivery/update [post][get]
func UpdateDeliveryType(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入出库类型编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if !utils.Exists(params, "deliveryType") {
		return *common.Error(-1, "未传入出库类型名称参数")
	}
	return service.UpdateDeliveryType(id, params["deliveryType"])
}

// DeleteDeliveryType	godoc
// @Summary		删除出库类型接口
// @Description	删除出库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "出库类型编号"
// @Success 200 {object} model.DeliveryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/type/delivery/del [post][get]
func DeleteDeliveryType(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入出库类型编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteDeliveryType(id)
}
