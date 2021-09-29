package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-api/vo"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListSkuStock	godoc
// @Summary		列出商品/货品的库存接口
// @Description	列出商品/货品的库存
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品编码"
// @Param	name formData string false "搜索关键字，从商品名称中查询"
// @Param	storageId formData string false "商户仓库编号"
// @Param	productId formData string false "商品编号"
// @Param	page formData int false "分页页号参数，第几页，若要分页，页号>=1，否则为0"
// @Param	size formData int false "分页大小参数"
// @Success 200 {object} vo.SkuStockNote	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/stock/list [post][get]
func ListSkuStock(params map[string]string) common.Result {
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
	return service.ListSkuStock(params["shopId"], params["skuId"], params["name"], params["productId"], params["storageId"], page, size)
}

// GetSkuStock	godoc
// @Summary		获取货品在特定仓库中的库存信息
// @Description	获取货品在特定仓库中的库存信息
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	storageId formData string true "商户仓库编号"
// @Success 200 {object} vo.SkuStockNote	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/stock/get [post][get]
func GetSkuStock(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "skuId") {
		return *common.Error(-1, "未传入货品编码参数")
	}
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "缺少仓库编码参数")
	}
	return service.GetSkuStockDetail(params["shopId"], params["skuId"], params["storageId"])
}

// InOutSkuStock	godoc
// @Summary		货品入库/出库接口
// @Description	货品入库/出库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string true "货品编码"
// @Param	unit formData string true "入库/出库的货品单位"
// @Param	number formData int true "入库/出库的货品数量"
// @Param	storageId formData string true "商户仓库编号"
// @Param	cost formData int false "本次入库的货品强制设定加权均价，不传时由系统计算"
// @Param	price formData int false "本次入库的货品进货价，出库时不传"
// @Success 200 {object} model.SkuStock	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/stock/inout [post][get]
func InOutSkuStock(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "skuId") {
		return *common.Error(-1, "未传入商品编码参数")
	}
	if !utils.Exists(params, "number") {
		return *common.Error(-1, "缺少数量参数，入库为正整数，出库为负整数")
	}
	number, _ := strconv.Atoi(params["number"])
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "缺少商户仓库编码参数")
	}
	if number > 0 && !utils.Exists(params, "price") {
		return *common.Error(-1, "必须传入最新入库价格，单位为分")
	}
	price, _ := strconv.Atoi(params["price"])
	cost := 0
	if utils.Exists(params, "cost") {
		cost, _ = strconv.Atoi(params["cost"])
	}
	return service.InOutSkuStock(params["shopId"], params["skuId"], params["unit"], params["storageId"], number, cost, price)
}

// GetSkuStockNumber	godoc
// @Summary		获取指定货品/商品的库存量
// @Description	获取指定货品/商品的库存量
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	skuId formData string false "货品自定义编号"
// @Param	storageId formData string false "商户仓库编号"
// @Param	productId formData string false "商品编号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/stock/num [post][get]
func GetSkuStockNumber(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.GetSkuStockNumber(params["shopId"], params["skuId"], params["productId"], params["storageId"])
}
