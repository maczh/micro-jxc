package nacos

import (
	"errors"
	"github.com/mitchellh/mapstructure"
	"ququ.im/common"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
)

const (
	PRODUCT_SERVICE = "jxc-product"
	BASE_SERVICE    = "jxc-base"
	URI_GET_SKU     = "/sku/get"
	URI_GET_UNIT    = "/unit/get"
	URI_GET_STORAGE = "/storage/get"
)

func GetProductSku(skuGuid, shopId, skuId string) (*model.ProductSku, error) {
	params := make(map[string]string)
	params["skuGuid"] = skuGuid
	params["shopId"] = shopId
	params["skuId"] = skuId
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_SKU, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_SKU, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productSku model.ProductSku
		mapstructure.Decode(result.Data, &productSku)
		return &productSku, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func GetProductUnit(shopId, productId, unit, baseUnit string) (*model.ProductUnit, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["productId"] = productId
	params["unit"] = unit
	params["baseUnit"] = baseUnit
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_UNIT, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_UNIT, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productUnit model.ProductUnit
		mapstructure.Decode(result.Data, &productUnit)
		return &productUnit, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func GetStorage(shopId, storageId string) (*model.ShopStorage, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["storageId"] = storageId
	res, err := utils.CallNacos(BASE_SERVICE, URI_GET_STORAGE, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", BASE_SERVICE, URI_GET_STORAGE, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var shopStorage model.ShopStorage
		mapstructure.Decode(result.Data, &shopStorage)
		return &shopStorage, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}
