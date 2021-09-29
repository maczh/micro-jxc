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
	URI_GET_SKU    = "/sku/get"
	URI_LIST_SKU   = "/sku/list"
	URI_SAVE_SKU   = "/sku/save"
	URI_UPDATE_SKU = "/sku/update"
	URI_DEL_SKU    = "/sku/del"
	URI_UP_SKU     = "/sku/up"
	URI_DOWN_SKU   = "/sku/down"
)

func GetProductSku(shopId, skuId, skuGuid string) (*model.ProductSku, error) {
	params := make(map[string]string)
	if skuGuid != "" {
		params["skuGuid"] = skuGuid
	}
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSku model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSku)
		return &productSku, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductSku(shopId, keyword, skusMap, productId, categoryId string, status, page, size int) ([]model.ProductSku, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	if keyword != "" {
		params["keyword"] = keyword
	}
	if skusMap != "" {
		params["skusMap"] = skusMap
	}
	if productId != "" {
		params["productId"] = productId
	}
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if status > -1 {
		params["status"] = strconv.Itoa(status)
	}
	params["page"] = strconv.Itoa(page)
	params["size"] = strconv.Itoa(size)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_SKU, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSkuList []model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSkuList)
		return productSkuList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func SaveProductSku(shopId, productId, skusMap, skuId, skuGuid, skuName, priceList string) (*model.ProductSku, error) {
	params := make(map[string]string)
	if skuGuid != "" {
		params["skuGuid"] = skuGuid
	}
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	if skusMap != "" {
		params["specs"] = skusMap
	}
	if productId != "" {
		params["productId"] = productId
	}
	if skuName != "" {
		params["name"] = skuName
	}
	if priceList != "" {
		params["prices"] = priceList
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSku model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSku)
		return &productSku, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateProductSku(skuGuid, skuId, name, skuName, barCode, skusMap, priceList string, status int) (*model.ProductSku, error) {
	params := make(map[string]string)
	params["skuGuid"] = skuGuid
	if name != "" {
		params["name"] = name
	}
	if skuName != "" {
		params["skuName"] = skuName
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	if skusMap != "" {
		params["specs"] = skusMap
	}
	if barCode != "" {
		params["barCode"] = barCode
	}
	if status > -1 {
		params["status"] = strconv.Itoa(status)
	}
	if priceList != "" {
		params["prices"] = priceList
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSku model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSku)
		return &productSku, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteProductSku(skuGuid, shopId, skuId string) common.Result {
	params := make(map[string]string)
	if skuGuid != "" {
		params["skuGuid"] = skuGuid
	}
	if shopId != "" {
		params["shopId"] = shopId
	}
	if skuId != "" {
		params["skuId"] = skuId
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_SKU, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func IncrProductSkuSortNumber(skuGuid string) ([]model.ProductSku, error) {
	params := make(map[string]string)
	params["skuGuid"] = skuGuid
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DOWN_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DOWN_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSkuList []model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSkuList)
		return productSkuList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DecrProductSkuSortNumber(skuGuid string) ([]model.ProductSku, error) {
	params := make(map[string]string)
	params["skuGuid"] = skuGuid
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UP_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UP_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSkuList []model.ProductSku
		utils.FromJSON(utils.ToJSON(result.Data), &productSkuList)
		return productSkuList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}
