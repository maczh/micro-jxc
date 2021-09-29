package nacos

import (
	"errors"
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"strconv"
)

const (
	URI_GET_SUPPLIER    = "/supplier/get"
	URI_LIST_SUPPLIER   = "/supplier/list"
	URI_SAVE_SUPPLIER   = "/supplier/save"
	URI_UPDATE_SUPPLIER = "/supplier/update"
	URI_DEL_SUPPLIER    = "/supplier/del"
)

func GetProductSupplier(id int, shopId, productId, supplierId string) (*model.ProductSupplier, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if supplierId != "" {
		params["supplierId"] = supplierId
	}
	if id != 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_SUPPLIER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_SUPPLIER, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSupplier model.ProductSupplier
		utils.FromJSON(utils.ToJSON(result.Data), &productSupplier)
		return &productSupplier, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductSupplier(shopId, productId string) ([]model.ProductSupplier, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_SUPPLIER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_SUPPLIER, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSupplierList []model.ProductSupplier
		utils.FromJSON(utils.ToJSON(result.Data), &productSupplierList)
		return productSupplierList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveProductSupplier(shopId, productId, supplierId, supplierName string) (*model.ProductSupplier, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["productId"] = productId
	params["supplierId"] = supplierId
	params["supplierName"] = supplierName
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_SUPPLIER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_SUPPLIER, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSupplier model.ProductSupplier
		utils.FromJSON(utils.ToJSON(result.Data), &productSupplier)
		return &productSupplier, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateProductSupplier(id int, shopId, productId, supplierId, supplierName string) (*model.ProductSupplier, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if supplierId != "" {
		params["supplierId"] = supplierId
	}
	if supplierName != "" {
		params["supplierName"] = supplierName
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_SUPPLIER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_SUPPLIER, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSupplier model.ProductSupplier
		utils.FromJSON(utils.ToJSON(result.Data), &productSupplier)
		return &productSupplier, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteProductSupplier(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_SUPPLIER, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_SUPPLIER, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
