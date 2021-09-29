package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	_ "ququ.im/jxc-base/model"
	"ququ.im/jxc-order/service"
	"strconv"
)

// ListStockMove	godoc
// @Summary		列出移库单接口
// @Description	列出移库单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	moveNo formData string false "移库单号"
// @Param	fromStorageId formData string false "移出仓库编号"
// @Param	toStorageId formData string false "移入仓库编号"
// @Param	startTime formData string false "开始日期，格式为 yyyy-MM-dd"
// @Param	endTime formData string false "结束日期，格式为yyyy-MM-dd"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} model.StockMove	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/list [post][get]
func ListStockMove(params map[string]string) common.Result {
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
	return service.ListStockMove(params["shopId"], params["moveNo"], params["skuId"], params["fromStorageId"], params["toStorageId"], params["startTime"], params["endTime"], page, size)
}

// GetStockMove	godoc
// @Summary		获取特定移库单的单一货品记录接口
// @Description	获取特定移库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "移库单记录号"
// @Param	shopId formData string false "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	moveNo formData string false "移库单号"
// @Success 200 {object} model.StockMove	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/get [post][get]
func GetStockMove(params map[string]string) common.Result {
	id := 0
	if utils.Exists(params, "id") {
		id, _ = strconv.Atoi(params["id"])
	}
	return service.GetStockMove(id, params["shopId"], params["moveNo"], params["skuId"])
}

// SaveStockMove	godoc
// @Summary		保存移库单的单一货品记录接口
// @Description	保存移库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	moveNo formData string true "移库单号"
// @Param	unit formData string true "货品单位"
// @Param	fromStorageId formData string true "移出仓库编号"
// @Param	toStorageId formData string true "移入仓库编号"
// @Param	operator formData string true "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int true "移库数量，必须大于0"
// @Success 200 {object} model.StockMove	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/save [post][get]
func SaveStockMove(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "number") {
		return *common.Error(-1, "缺少货品数量参数")
	}
	number, _ := strconv.Atoi(params["number"])
	return service.SaveStockMove(params["shopId"], params["moveNo"], params["skuId"], params["unit"], params["fromStorageId"], params["toStorageId"], params["operator"], params["remark"], number)
}

// UpdateStockMove	godoc
// @Summary		修改移库单的单一货品记录接口
// @Description	修改移库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData string true "记录号"
// @Param	skuId formData string false "货品编码"
// @Param	unit formData string false "货品单位"
// @Param	fromStorageId formData string false "移出仓库编号"
// @Param	toStorageId formData string false "移入仓库编号"
// @Param	operator formData string false "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int false "移库数量，若需修改则大于0"
// @Success 200 {object} model.StockMove	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/update [post][get]
func UpdateStockMove(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	number := 0
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "缺少记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if utils.Exists(params, "number") {
		number, _ = strconv.Atoi(params["number"])
	}
	return service.UpdateStockMove(id, params["skuId"], params["unit"], params["fromStorageId"], params["toStorageId"], params["operator"], params["remark"], number)
}

// DeleteStockMove	godoc
// @Summary		删除移库单接记录口
// @Description	删除移库单记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "移库单记录号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/del [post][get]
func DeleteStockMove(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入移库单记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteStockMove(id)
}
