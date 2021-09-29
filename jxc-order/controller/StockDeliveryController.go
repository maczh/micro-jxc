package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	_ "ququ.im/jxc-base/model"
	"ququ.im/jxc-order/service"
	"strconv"
)

// ListStockDelivery	godoc
// @Summary		列出出库单接口
// @Description	列出出库单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	deliveryNo formData string false "出库单号"
// @Param	orderNo formData string false "外部销售单号"
// @Param	customerId formData string false "客户账号"
// @Param	storageId formData string false "商户仓库编号"
// @Param	startTime formData string false "开始日期，格式为 yyyy-MM-dd"
// @Param	endTime formData string false "结束日期，格式为yyyy-MM-dd"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} model.StockDelivery	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/delivery/list [post][get]
func ListStockDelivery(params map[string]string) common.Result {
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
	return service.ListStockDelivery(params["shopId"], params["deliveryNo"], params["skuId"], params["storageId"], params["orderNo"], params["customerId"], params["startTime"], params["endTime"], page, size)
}

// GetStockDelivery	godoc
// @Summary		获取特定出库单的单一货品记录接口
// @Description	获取特定出库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "出库单记录号"
// @Param	shopId formData string false "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	deliveryNo formData string false "出库单号"
// @Success 200 {object} model.StockDelivery	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/delivery/get [post][get]
func GetStockDelivery(params map[string]string) common.Result {
	id := 0
	if utils.Exists(params, "id") {
		id, _ = strconv.Atoi(params["id"])
	}
	if id == 0 {
		if !utils.Exists(params, "shopId") {
			return *common.Error(-1, "未传入商户账号参数")
		}
		if !utils.Exists(params, "deliveryNo") {
			return *common.Error(-1, "未传入出库单号参数")
		}
		if !utils.Exists(params, "skuId") {
			return *common.Error(-1, "未传入货品编号参数")
		}
	}
	return service.GetStockDelivery(params["shopId"], params["deliveryNo"], params["skuId"], id)
}

// SaveStockDelivery	godoc
// @Summary		保存出库单的单一货品记录接口
// @Description	保存出库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	deliveryNo formData string true "出库单号"
// @Param	type formData string true "出库类型"
// @Param	unit formData string true "货品单位"
// @Param	orderNo formData string false "外部销售单号"
// @Param	storageId formData string true "商户仓库编号"
// @Param	customerId formData string false "客户账号"
// @Param	customerName formData string false "客户名称"
// @Param	operator formData string true "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int true "出库数量，必须大于0"
// @Param	cost formData int false "成本均价，单位为分"
// @Param	price formData int false "出库价，单位为分"
// @Success 200 {object} model.StockDelivery	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/delivery/save [post][get]
func SaveStockDelivery(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	cost, price := 0, 0
	if !utils.Exists(params, "number") {
		return *common.Error(-1, "缺少货品数量参数")
	}
	number, _ := strconv.Atoi(params["number"])
	if utils.Exists(params, "cost") {
		cost, _ = strconv.Atoi(params["cost"])
	}
	if utils.Exists(params, "price") {
		price, _ = strconv.Atoi(params["price"])
	}
	return service.SaveStockDelivery(params["shopId"], params["deliveryNo"], params["type"], params["skuId"], params["unit"], params["orderNo"], params["storageId"], params["customerId"], params["customerName"], params["operator"], params["remark"], number, cost, price)
}

// UpdateStockDelivery	godoc
// @Summary		修改出库单的单一货品记录接口
// @Description	修改出库单的单一货品记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData string true "记录号"
// @Param	skuId formData string false "货品编码"
// @Param	unit formData string false "货品单位"
// @Param	orderNo formData string false "外部销售单号"
// @Param	storageId formData string false "商户仓库编号"
// @Param	customerId formData string false "客户账号"
// @Param	customerName formData string false "客户名称"
// @Param	operator formData string false "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int false "出库数量，若需修改则大于0"
// @Param	cost formData int false "成本均价，单位为分"
// @Param	price formData int false "出库价，单位为分"
// @Success 200 {object} model.StockDelivery	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/delivery/update [post][get]
func UpdateStockDelivery(params map[string]string) common.Result {
	number, cost, price := 0, 0, 0
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "缺少记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if utils.Exists(params, "number") {
		number, _ = strconv.Atoi(params["number"])
	}
	if utils.Exists(params, "cost") {
		cost, _ = strconv.Atoi(params["cost"])
	}
	if utils.Exists(params, "price") {
		price, _ = strconv.Atoi(params["price"])
	}
	return service.UpdateStockDelivery(id, params["skuId"], params["unit"], params["orderNo"], params["storageId"], params["customerId"], params["customerName"], params["operator"], params["remark"], number, cost, price)
}

// DeleteStockDelivery	godoc
// @Summary		删除出库单记录接口
// @Description	删除出库单记录接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "出库单记录号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/delivery/del [post][get]
func DeleteStockDelivery(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入出库单记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteStockDelivery(id)
}
