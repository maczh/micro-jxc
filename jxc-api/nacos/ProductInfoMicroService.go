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
	PRODUCT_SERVICE    = "jxc-product"
	URI_GET_PRODUCT    = "/product/get"
	URI_LIST_PRODUCT   = "/product/list"
	URI_SAVE_PRODUCT   = "/product/save"
	URI_UPDATE_PRODUCT = "/product/update"
	URI_DEL_PRODUCT    = "/product/del"
)

func GetProductInfo(id int, shopId, productId, barCode string) (*model.ProductInfo, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if productId != "" {
		params["productId"] = productId
	}
	if barCode != "" {
		params["barCode"] = barCode
	}
	if id > 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_PRODUCT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_PRODUCT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productInfo model.ProductInfo
		utils.FromJSON(utils.ToJSON(result.Data), &productInfo)
		return &productInfo, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductInfo(shopId, categoryId, keyword string, page, size int) ([]model.ProductInfo, *common.ResultPage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if keyword != "" {
		params["keyword"] = keyword
	}
	if page > 0 {
		params["page"] = strconv.Itoa(page)
	}
	if size > 0 {
		params["size"] = strconv.Itoa(size)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_PRODUCT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_PRODUCT, err.Error())
		return nil, nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productInfoList []model.ProductInfo
		utils.FromJSON(utils.ToJSON(result.Data), &productInfoList)
		return productInfoList, result.Page, nil
	} else {
		return nil, nil, errors.New(result.Msg)
	}
}

func SaveProductInfo(shopId, categoryId, productName, productId, barCode, baseUnit string) common.Result {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["name"] = productName
	params["categoryId"] = categoryId
	params["baseUnit"] = baseUnit
	if productId != "" {
		params["productId"] = productId
	}
	if barCode != "" {
		params["barCode"] = barCode
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_PRODUCT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_PRODUCT, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func UpdateProductInfo(id int, categoryId, productName, productId, baseUnit, barCode string) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if productName != "" {
		params["name"] = productName
	}
	if productId != "" {
		params["productId"] = productId
	}
	if baseUnit != "" {
		params["baseUnit"] = baseUnit
	}
	if barCode != "" {
		params["barCode"] = barCode
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_PRODUCT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_PRODUCT, err.Error())
		return *common.Error(-1, err.Error())
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func DeleteProductInfo(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_PRODUCT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_PRODUCT, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}
