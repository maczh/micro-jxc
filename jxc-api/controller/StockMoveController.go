package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-api/vo"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListStockMove	godoc
// @Summary		列出移库单接口
// @Description	列出移库单,直接按条件查出所有符合条件的移库单记录，可分页
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	moveNo formData string false "移库单号"
// @Param	fromStorageId formData string false "原仓库编码"
// @Param	toStorageId formData string false "目标仓库编码"
// @Param	startTime formData string false "开始日期，格式为 yyyy-MM-dd"
// @Param	endTime formData string false "结束日期，格式为yyyy-MM-dd"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} vo.StockMoveRowDetail	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
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
	return service.ListStockMove(params["shopId"], params["skuId"], params["moveNo"], params["fromStorageId"], params["toStorageId"], params["startTime"], params["endTime"], page, size)
}

// GetStockMoveRow	godoc
// @Summary		按移库单号和货品编码查询一条移库单记录接口
// @Description	按移库单号和货品编码查询一条移库单记录
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	moveNo formData string true "移库单号"
// @Param	skuId formData string true "货品编码"
// @Success 200 {object} vo.StockMoveRowDetail	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/get [post][get]
func GetStockMoveRow(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "moveNo") {
		return *common.Error(-1, "未传入移库单号参数")
	}
	if !utils.Exists(params, "skuId") {
		return *common.Error(-1, "未传入货品编码参数")
	}
	return service.GetStockMoveRow(params["shopId"], params["skuId"], params["moveNo"])
}

// MoveSkuStock	godoc
// @Summary		移库操作接口
// @Description	完成移库操作过程，从原库出库，到目标库入库，并且保存移库单
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	moveNo formData string false "移库单号"
// @Param	fromStorageId formData string false "原仓库编码"
// @Param	toStorageId formData string false "目标仓库编码"
// @Param	unit formData string true "货品单位"
// @Param	operator formData string true "操作员名称"
// @Param	remark formData string false "备注"
// @Param	number formData int true "移库数量，必须大于0"
// @Success 200 {object} model.StockMove	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/stock [post][get]
func MoveSkuStock(params map[string]string) common.Result {
	if !utils.Exists(params, "number") {
		return *common.Error(-1, "缺少货品数量参数")
	}
	number, _ := strconv.Atoi(params["number"])
	return service.MoveSkuStock(params["shopId"], params["skuId"], params["moveNo"], params["unit"], params["fromStorageId"], params["toStorageId"], params["operator"], params["remark"], number)
}

// DeleteStockMoveRow	godoc
// @Summary		删除移库单记录接口
// @Description	删除移库单记录,对应的移库单库存也会还原，从目标库转移到原库
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "移库单记录号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/move/del [post][get]
func DeleteStockMove(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "缺少移库单记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteStockMove(id)
}
