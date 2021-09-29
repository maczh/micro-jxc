package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-api/vo"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListStockEntryRow	godoc
// @Summary		列出入库单接口
// @Description	列出入库单,直接按条件查出所有符合条件的入库单记录，可分页
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	entryNo formData string false "入库单号"
// @Param	orderNo formData string false "外部销售单号"
// @Param	supplierId formData string false "客户账号"
// @Param	storageId formData string false "商户仓库编号"
// @Param	startTime formData string false "开始日期，格式为 yyyy-MM-dd"
// @Param	endTime formData string false "结束日期，格式为yyyy-MM-dd"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} vo.StockEntryRowDetail	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/list [post][get]
func ListStockEntryRow(params map[string]string) common.Result {
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
	return service.ListStockEntryRow(params["shopId"], params["entryNo"], params["skuId"], params["storageId"], params["orderNo"], params["supplierId"], params["startTime"], params["endTime"], page, size)
}

// GetStockEntryNote	godoc
// @Summary		按入库单号查询整个入库单接口
// @Description	按入库单号查询整个入库单，其中包含公共参数与多条货品入库记录
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	entryNo formData string true "入库单号"
// @Success 200 {object} vo.StockEntryNote	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/note/get [post][get]
func GetStockEntryNote(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "entryNo") {
		return *common.Error(-1, "未传入入库单号参数")
	}
	return service.GetStockEntryNote(params["shopId"], params["entryNo"])
}

// SaveOneStockEntryRow	godoc
// @Summary		按单一货品入库操作接口
// @Description	保存入库单的单一货品记录,商品自动增加到库存
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	entryNo formData string true "入库单号"
// @Param	type formData string true "入库类型"
// @Param	unit formData string true "货品单位"
// @Param	orderNo formData string false "外部销售单号"
// @Param	storageId formData string true "商户仓库编号"
// @Param	supplierId formData string false "客户账号"
// @Param	operator formData string true "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int true "入库数量，必须大于0"
// @Param	price formData int false "入库价，单位为分"
// @Success 200 {object} model.StockEntry	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/row/save [post][get]
func SaveOneStockEntryRow(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	price := 0
	if !utils.Exists(params, "number") {
		return *common.Error(-1, "缺少货品数量参数")
	}
	number, _ := strconv.Atoi(params["number"])
	if utils.Exists(params, "price") {
		price, _ = strconv.Atoi(params["price"])
	}
	stockEntry, err := service.SaveOneStockEntryRow(params["shopId"], params["entryNo"], params["type"], params["skuId"], params["unit"], params["orderNo"], params["storageId"], params["supplierId"], params["operator"], params["remark"], number, price)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(stockEntry)
}

// SaveStockEntryNote	godoc
// @Summary		按完整入库单一次性入库操作接口
// @Description	保存一个完整入库单，包含多条货品记录，并且自动增加相应货品的库存
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	entryNo formData string false "入库单号，不传时默认自动生成一个单号"
// @Param	type formData string true "入库类型"
// @Param	orderNo formData string false "外部销售单号"
// @Param	operator formData string true "操作员名称"
// @Param	stockEntryRowList formData string true "入库货品清单列表，JSON数组格式，必须包含以下字段:sku_id,unit,price,storage_id,supplier_id,number,remark字段"
// @Success 200 {object} model.StockEntry	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/note/save [post][get]
func SaveStockEntryNote(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.SaveStockEntryNote(params["shopId"], params["entryNo"], params["type"], params["orderNo"], params["operator"], params["stockEntryRowList"])
}

// UpdateStockEntryRow	godoc
// @Summary		修改入库单的单一货品记录接口
// @Description	修改入库单的单一货品记录,如果修改了数量、单位，则会自动调整库存数量，若修改了价格，则会自动重新计算成本价
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData string true "记录号"
// @Param	unit formData string false "货品单位"
// @Param	orderNo formData string false "外部销售单号"
// @Param	storageId formData string false "商户仓库编号"
// @Param	supplierId formData string false "客户账号"
// @Param	operator formData string false "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int false "入库数量，若需修改则大于0"
// @Param	price formData int false "入库价，单位为分"
// @Success 200 {object} model.StockEntry	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/update [post][get]
func UpdateStockEntryRow(params map[string]string) common.Result {
	number, price := 0, 0
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "缺少记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	if utils.Exists(params, "number") {
		number, _ = strconv.Atoi(params["number"])
	}
	if utils.Exists(params, "price") {
		price, _ = strconv.Atoi(params["price"])
	}
	return service.UpdateStockEntryRow(id, params["unit"], params["orderNo"], params["storageId"], params["supplierId"], params["operator"], params["remark"], number, price)
}

// DeleteStockEntryRow	godoc
// @Summary		删除入库单记录接口
// @Description	删除入库单记录,对应的入库单库存也会被扣除
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "入库单记录号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/entry/del [post][get]
func DeleteStockEntryRow(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "缺少入库单记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteStockEntryRow(id)
}
