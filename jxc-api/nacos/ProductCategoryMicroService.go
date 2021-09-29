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
	URI_GET_CATEGORY    = "/category/get"
	URI_LIST_CATEGORY   = "/category/list"
	URI_SAVE_CATEGORY   = "/category/save"
	URI_UPDATE_CATEGORY = "/category/update"
	URI_DEL_CATEGORY    = "/category/del"
	URI_UP_CATEGORY     = "/category/up"
	URI_DOWN_CATEGORY   = "/category/down"
)

func GetProductCategory(id int, shopId, categoryId string) (*model.ProductCategory, error) {
	params := make(map[string]string)
	if shopId != "" {
		params["shopId"] = shopId
	}
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if id != 0 {
		params["id"] = strconv.Itoa(id)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_GET_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_GET_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategory model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategory)
		return &productCategory, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func ListProductCategory(shopId, categoryId, parent string, level int) ([]model.ProductCategory, error) {
	params := make(map[string]string)
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if shopId != "" {
		params["shopId"] = shopId
	}
	if parent != "" {
		params["parent"] = parent
	}
	if level != 0 {
		params["level"] = strconv.Itoa(level)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_LIST_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_LIST_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategoryList []model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategoryList)
		return productCategoryList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func SaveProductCategory(shopId, categoryId, categoryName, parent string, level int) (*model.ProductCategory, error) {
	params := make(map[string]string)
	params["shopId"] = shopId
	params["name"] = categoryName
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	params["parent"] = parent
	params["level"] = strconv.Itoa(level)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_SAVE_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_SAVE_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategory model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategory)
		return &productCategory, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func UpdateProductCategory(id int, categoryId, categoryName, parent string, level int) (*model.ProductCategory, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	if categoryId != "" {
		params["categoryId"] = categoryId
	}
	if categoryName != "" {
		params["name"] = categoryName
	}
	if parent != "" {
		params["parent"] = parent
	}
	if level != 0 {
		params["level"] = strconv.Itoa(level)
	}
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UPDATE_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UPDATE_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategory model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategory)
		return &productCategory, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func DeleteProductCategory(id int) common.Result {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DEL_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DEL_CATEGORY, err.Error())
		return *common.Error(-1, "微服务请求失败")
	}
	var result common.Result
	utils.FromJSON(res, &result)
	return result
}

func DecrProductCategorySortNumber(id int) ([]model.ProductCategory, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_UP_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_UP_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategoryList []model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategoryList)
		return productCategoryList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}

func IncrProductCategorySortNumber(id int) ([]model.ProductCategory, error) {
	params := make(map[string]string)
	params["id"] = strconv.Itoa(id)
	res, err := utils.CallNacos(PRODUCT_SERVICE, URI_DOWN_CATEGORY, params)
	if err != nil {
		logs.Error("微服务{}{}调用异常:{}", PRODUCT_SERVICE, URI_DOWN_CATEGORY, err.Error())
		return nil, err
	}
	var result common.Result
	utils.FromJSON(res, &result)
	if result.Status == 1 {
		var productCategoryList []model.ProductCategory
		utils.FromJSON(utils.ToJSON(result.Data), &productCategoryList)
		return productCategoryList, nil
	} else {
		return nil, errors.New(result.Msg)
	}
}
