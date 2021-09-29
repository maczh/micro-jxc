package service

import (
	"ququ.im/common"
	"ququ.im/jxc-api/nacos"
)

func ListProductCategory(shopId, categoryId, parent string, level int) common.Result {
	productCategoryList, err := nacos.ListProductCategory(shopId, categoryId, parent, level)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productCategoryList == nil {
		return *common.Error(-1, "未定义产品分类")
	}
	return *common.Success(productCategoryList)
}

func GetProductCategory(id int, shopId, categoryId string) common.Result {
	productCategory, err := nacos.GetProductCategory(id, shopId, categoryId)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	if productCategory == nil {
		return *common.Error(-1, "未定义此产品分类")
	}
	return *common.Success(productCategory)
}

func SaveProductCategory(shopId, categoryName, categoryId, parent string, level int) common.Result {
	productCategory, err := nacos.SaveProductCategory(shopId, categoryId, categoryName, parent, level)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategory)
}

func UpdateProductCategory(id, level int, parent, categoryName, categoryId string) common.Result {
	productCategory, err := nacos.UpdateProductCategory(id, categoryId, categoryName, parent, level)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategory)
}

func IncrProductCategorySortNumber(id int) common.Result {
	productCategoryList, err := nacos.IncrProductCategorySortNumber(id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategoryList)
}

func DecrProductCategorySortNumber(id int) common.Result {
	productCategoryList, err := nacos.DecrProductCategorySortNumber(id)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategoryList)
}

func DeleteProductCategory(id int) common.Result {
	return nacos.DeleteProductCategory(id)
}
