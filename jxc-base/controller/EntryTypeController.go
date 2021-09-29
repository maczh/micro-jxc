package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/service"
	"strconv"
)

// ListEntryTypeByShopId	godoc
// @Summary		列出商户的所有入库类型接口
// @Description	列出商户的所有入库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Success 200 {object} model.EntryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entrytype/list [post][get]
func ListEntryTypeByShopId(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.ListEntryTypeByShopId(params["shopId"])
}

// SaveEntryType	godoc
// @Summary		保存入库类型接口
// @Description	保存入库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string false "商户号"
// @Param	entryType formData string true "入库类型"
// @Success 200 {object} model.EntryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entrytype/save [post][get]
func SaveEntryType(params map[string]string) common.Result {
	if !utils.Exists(params, "entryType") {
		return *common.Error(-1, "未传入入库类型参数")
	}
	return service.SaveEntryType(params["shopId"], params["entryType"])
}

// UpdateEntryType	godoc
// @Summary		更新入库类型接口
// @Description	更新入库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "入库类型编号"
// @Param	entryType formData string true "入库类型"
// @Success 200 {object} model.EntryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entrytype/update [post][get]
func UpdateEntryType(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入入库类型编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if !utils.Exists(params, "entryType") {
		return *common.Error(-1, "未传入入库类型名称参数")
	}
	return service.UpdateEntryType(id, params["entryType"])
}

// DeleteEntryType	godoc
// @Summary		删除入库类型接口
// @Description	删除入库类型接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "入库类型编号"
// @Success 200 {object} model.EntryType	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entrytype/del [post][get]
func DeleteEntryType(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入入库类型编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteEntryType(id)
}
