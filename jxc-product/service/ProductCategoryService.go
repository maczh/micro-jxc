package service

import (
	"ququ.im/common"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-base/util"
	"ququ.im/jxc-product/dao"
)

func ListProductCategory(shopId, categoryId, parent string, level int) common.Result {
	var productCategoryList []model.ProductCategory
	if shopId == "" {
		return *common.Error(-1, "商户号或分类编号不能为空")
	}
	if categoryId != "" {
		productCategoryList = dao.ListProductCategoryTree(shopId, categoryId)
	} else {
		if level < 1 {
			return *common.Error(-1, "分类层级或分类编号不能为空")
		}
		productCategoryList = dao.ListProductCategoryByParentLevel(shopId, parent, level)
	}
	if productCategoryList == nil {
		return *common.Error(-1, "未定义产品分类")
	}
	return *common.Success(productCategoryList)
}

func GetProductCategory(id int, shopId, categoryId string) common.Result {
	var productCategory *model.ProductCategory
	if id > 0 {
		productCategory = dao.GetProductCategoryById(id)
	} else if shopId != "" && categoryId != "" {
		productCategory = dao.GetProductCategory(shopId, categoryId)
	} else {
		return *common.Error(-1, "参数不全")
	}
	if productCategory == nil {
		return *common.Error(-1, "未定义此产品分类")
	}
	return *common.Success(productCategory)
}

func SaveProductCategory(shopId, categoryName, categoryId, parent string, level int) common.Result {
	productCategory := dao.GetProductCategoryByName(shopId, categoryName, parent, level)
	if productCategory != nil {
		return *common.Success(productCategory)
	}
	productCategory = dao.GetProductCategory(shopId, categoryId)
	if productCategory != nil {
		return *common.Success(productCategory)
	}
	productCategory = new(model.ProductCategory)
	productCategory.ShopId = shopId
	if categoryId == "" {
		productCategory.CategoryId = util.GenerateId("FL", "category", shopId, 6)
	}
	productCategory.Name = categoryName
	productCategory.Level = level
	productCategory.ParentId = parent
	productCategory, err := dao.SaveProductCategory(productCategory)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategory)
}

func UpdateProductCategory(id, level int, parent, categoryName, categoryId string) common.Result {
	productCategory := dao.GetProductCategoryById(id)
	if productCategory == nil {
		return *common.Error(-1, "无此商品分类ID")
	}
	if categoryName != "" {
		productCategory.Name = categoryName
	}
	if categoryId != "" {
		productCategory.CategoryId = categoryId
	}
	if parent != "" {
		productCategory.ParentId = parent
	}
	if level > 0 {
		productCategory.Level = level
	}
	productCategory, err := dao.UpdateProductCategory(productCategory)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(productCategory)
}

func IncrProductCategorySortNumber(id int) common.Result {
	productCategory := dao.GetProductCategoryById(id)
	dao.IncrProductCategorySortNumber(id)
	return ListProductCategory(productCategory.ShopId, "", productCategory.ParentId, productCategory.Level)
}

func DecrProductCategorySortNumber(id int) common.Result {
	productCategory := dao.GetProductCategoryById(id)
	dao.DecrProductCategorySortNumber(id)
	return ListProductCategory(productCategory.ShopId, "", productCategory.ParentId, productCategory.Level)
}

func DeleteProductCategory(id int) common.Result {
	err := dao.DeleteProductCategory(id)
	if err != nil {
		return *common.Error(-1, "无此商品分类记录")
	}
	return *common.Success(nil)
}
