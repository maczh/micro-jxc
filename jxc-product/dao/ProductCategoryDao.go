package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/jxc-base/model"
	"strconv"
)

const TABLE_PRODUCT_CATEGORY string = "product_category"

func ListProductCategoryByParentLevel(shopId, parent string, level int) []model.ProductCategory {
	var productCategoryList []model.ProductCategory
	err := cache.GetCache("category:parent", parent, &productCategoryList)
	if err == nil && len(productCategoryList) > 0 {
		return productCategoryList
	}
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("shop_id = ? AND level = ? AND parent_id = ?", shopId, level, parent).Order("sort_number ASC").Find(&productCategoryList)
	if len(productCategoryList) == 0 {
		return nil
	} else {
		cache.SetCache("category:parent", parent, productCategoryList)
		return productCategoryList
	}
}

func GetProductCategoryById(id int) *model.ProductCategory {
	var productCategory model.ProductCategory
	err := cache.GetCache("category:id", strconv.Itoa(id), &productCategory)
	if err == nil && productCategory.ShopId != "" {
		return &productCategory
	}
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("id = ?", id).First(&productCategory)
	if productCategory.Id == 0 {
		return nil
	} else {
		cache.SetCache("category:id", strconv.Itoa(id), productCategory)
		return &productCategory
	}
}

func GetProductCategory(shopId, categoryId string) *model.ProductCategory {
	var productCategory model.ProductCategory
	err := cache.GetCache("category:shop:"+shopId, categoryId, &productCategory)
	if err == nil && productCategory.ShopId != "" {
		return &productCategory
	}
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("shop_id = ? AND category_id = ?", shopId, categoryId).First(&productCategory)
	if productCategory.Id == 0 {
		return nil
	} else {
		cache.SetCache("category:shop:"+shopId, categoryId, productCategory)
		return &productCategory
	}
}

func GetProductCategoryByName(shopId, categoryName, parent string, level int) *model.ProductCategory {
	var productCategory model.ProductCategory
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("shop_id = ? AND level = ? AND parent_id = ? AND name = ?", shopId, level, parent, categoryName).First(&productCategory)
	if productCategory.Id == 0 {
		return nil
	} else {
		return &productCategory
	}
}

func ListProductCategoryTree(shopId, categoryId string) []model.ProductCategory {
	var productCategoryList []model.ProductCategory
	err := cache.GetCache("category:tree:"+shopId, categoryId, &productCategoryList)
	if err == nil && len(productCategoryList) > 0 {
		return productCategoryList
	}
	productCategory := GetProductCategory(shopId, categoryId)
	if productCategory == nil {
		return nil
	}
	productCategoryList = make([]model.ProductCategory, productCategory.Level)
	productCategoryList[productCategory.Level-1] = *productCategory
	for productCategory.ParentId != "" {
		productCategory = GetProductCategory(shopId, productCategory.ParentId)
		productCategoryList[productCategory.Level-1] = *productCategory
	}
	cache.SetCache("category:tree:"+shopId, categoryId, productCategoryList)
	return productCategoryList
}

func SaveProductCategory(productCategory *model.ProductCategory) (*model.ProductCategory, error) {
	if productCategory.SortNumber == 0 {
		productCategoryList := ListProductCategoryByParentLevel(productCategory.ShopId, productCategory.ParentId, productCategory.Level)
		if productCategoryList == nil {
			productCategory.SortNumber = 101
		} else {
			productCategory.SortNumber = productCategoryList[len(productCategoryList)-1].SortNumber + 1
		}
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_CATEGORY).Create(productCategory).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("category:parent", productCategory.ParentId)
	return productCategory, nil
}

func UpdateProductCategory(productCategory *model.ProductCategory) (*model.ProductCategory, error) {
	if productCategory.Id == 0 {
		return nil, errors.New("未指定分类编号")
	}
	data := make(map[string]interface{})
	if productCategory.ShopId != "" {
		data["shop_id"] = productCategory.ShopId
	}
	data["parent_id"] = productCategory.ParentId
	if productCategory.CategoryId != "" {
		data["category_id"] = productCategory.CategoryId
	}
	if productCategory.Name != "" {
		data["name"] = productCategory.Name
	}
	if productCategory.Level > 0 {
		data["level"] = productCategory.Level
	}
	if productCategory.SortNumber > 0 {
		data["sort_number"] = productCategory.SortNumber
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_CATEGORY).Where("id = ?", productCategory.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("category:parent", productCategory.ParentId)
	cache.DelCache("category:tree:"+productCategory.ShopId, productCategory.CategoryId)
	cache.DelCache("category:id", strconv.Itoa(productCategory.Id))
	cache.DelCache("category:shop:"+productCategory.ShopId, productCategory.CategoryId)
	return productCategory, nil
}

//排序向后移一位
func IncrProductCategorySortNumber(categoryId int) {
	productCategory := GetProductCategoryById(categoryId)
	if productCategory == nil {
		return
	}
	var productCategoryNext model.ProductCategory
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("shop_id = ? AND level = ? AND parent_id = ? AND sort_number > ?", productCategory.ShopId, productCategory.Level, productCategory.ParentId, productCategory.SortNumber).Order("sort_number ASC").First(&productCategoryNext)
	if productCategoryNext.Id == 0 {
		return
	}
	productCategory.SortNumber, productCategoryNext.SortNumber = productCategoryNext.SortNumber, productCategory.SortNumber
	UpdateProductCategory(productCategory)
	UpdateProductCategory(&productCategoryNext)
}

//排序向前移一位
func DecrProductCategorySortNumber(categoryId int) {
	productCategory := GetProductCategoryById(categoryId)
	if productCategory == nil {
		return
	}
	var productCategoryBefore model.ProductCategory
	config.Mysql.Table(TABLE_PRODUCT_CATEGORY).Where("shop_id = ? AND level = ? AND parent_id = ? AND sort_number < ?", productCategory.ShopId, productCategory.Level, productCategory.ParentId, productCategory.SortNumber).Order("sort_number DESC").First(&productCategoryBefore)
	if productCategoryBefore.Id == 0 {
		return
	}
	productCategory.SortNumber, productCategoryBefore.SortNumber = productCategoryBefore.SortNumber, productCategory.SortNumber
	UpdateProductCategory(productCategory)
	UpdateProductCategory(&productCategoryBefore)
}

func DeleteProductCategory(productCategoryId int) error {
	productCategory := GetProductCategoryById(productCategoryId)
	if productCategoryId == 0 {
		return errors.New("未指定商品分类编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_CATEGORY).Where("id = ?", productCategoryId).Delete(model.ProductCategory{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("category:parent", productCategory.ParentId)
	cache.DelCache("category:tree:"+productCategory.ShopId, productCategory.CategoryId)
	cache.DelCache("category:id", strconv.Itoa(productCategoryId))
	cache.DelCache("category:shop:"+productCategory.ShopId, productCategory.CategoryId)
	return nil
}
