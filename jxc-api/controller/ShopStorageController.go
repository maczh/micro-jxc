package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-api/service"
	_ "ququ.im/jxc-base/model"
	"strconv"
)

// ListShopStorageByShopId	godoc
// @Summary		列出商户的所有仓库接口
// @Description	列出商户的所有仓库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/list [post][get]
func ListShopStorageByShopId(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	return service.ListShopStorageByShopId(params["shopId"])
}

// SaveShopStorage	godoc
// @Summary		保存仓库接口
// @Description	保存仓库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	name formData string true "仓库名称"
// @Param	storageId formData string false "自定义仓库编码，不传则自动生成"
// @Param	remark formData string false "备注"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/save [post][get]
func SaveShopStorage(params map[string]string) common.Result {
	if !utils.Exists(params, "name") {
		return *common.Error(-1, "未传入仓库名称参数")
	}
	return service.SaveShopStorage(params["shopId"], params["name"], params["storageId"], params["remark"])
}

// UpdateShopStorage	godoc
// @Summary		更新仓库接口
// @Description	更新仓库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "仓库记录序列号，若不传时则不可修改自定义仓库编码"
// @Param	shopId formData string false "商户账号"
// @Param	storageId formData string false "自定义仓库编码"
// @Param	name formData string false "仓库名称"
// @Param	remark formData string false "备注"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/update [post][get]
func UpdateShopStorage(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入仓库记录号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.UpdateShopStorage(id, params["shopId"], params["storageId"], params["name"], params["remark"])
}

// DeleteShopStorage	godoc
// @Summary		删除仓库接口
// @Description	删除仓库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	storageId formData string true "自定义仓库编码"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/del [post][get]
func DeleteShopStorage(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "未传入仓库编码参数")
	}
	return service.DeleteShopStorage(params["shopId"], params["storageId"])
}

// GetShopStorage	godoc
// @Summary		获取仓库接口
// @Description	获取仓库接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Param	storageId formData string false "仓库编号"
// @Param	name formData string false "仓库名称"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/get [post][get]
func GetShopStorage(params map[string]string) common.Result {
	return service.GetShopStorage(params["storageId"], params["shopId"], params["name"])
}

// DecrShopStorageSortNumber	godoc
// @Summary		仓库顺序号向上移一个位置接口
// @Description	仓库顺序号向上移一个位置接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Param	storageId formData string true "仓库编号"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/up [post][get]
func DecrShopStorageSortNumber(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "未传入仓库编码参数")
	}
	return service.DecrShopStorageSortNumber(params["shopId"], params["storageId"])
}

// IncrShopStorageSortNumber	godoc
// @Summary		仓库顺序号向下移一个位置接口
// @Description	仓库顺序号向下移一个位置接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户号"
// @Param	storageId formData string true "仓库编号"
// @Success 200 {object} model.ShopStorage	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/down [post][get]
func IncrShopStorageSortNumber(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "未传入仓库编码参数")
	}
	return service.IncrShopStorageSortNumber(params["shopId"], params["storageId"])
}

// GetStorageLastOperator	godoc
// @Summary		获取指定仓库最后修改信息接口
// @Description	获取指定仓库最后修改信息接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账号"
// @Param	storageId formData string true "仓库编号"
// @Success 200 {object} model.StockDelivery	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/storage/last [post][get]
func GetStorageLastOperator(params map[string]string) common.Result {
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号参数")
	}
	if !utils.Exists(params, "storageId") {
		return *common.Error(-1, "未传入仓库编号参数")
	}
	return service.GetStorageLastOperator(params["shopId"], params["storageId"])
}
