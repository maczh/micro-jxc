package dao

import (
	"errors"
	"ququ.im/common/cache"
	"ququ.im/common/config"
	"ququ.im/common/logs"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"strconv"
)

const TABLE_PRODUCT_INFO string = "product_info"

func ListProductInfoByShopId(shopId, keyword, categoryId string, page, size int) []model.ProductInfo {
	var productInfoList []model.ProductInfo
	if keyword == "" {
		err := cache.GetCache("product:list:"+shopId, categoryId, &productInfoList)
		if err == nil && len(productInfoList) > 0 {
			if page > 0 && size > 0 {
				return productInfoList[(page-1)*size : page*size]
			} else {
				return productInfoList
			}
		}
	}
	if page > 0 && size > 0 {
		config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND IF(? != '',category_id = ?,1=1) AND IF(? != '',concat(name,pinyin_full,pinyin_first) like ?,1=1)", shopId, categoryId, categoryId, keyword, "%"+keyword+"%").Order("product_id").Limit(size).Offset((page - 1) * size).Find(&productInfoList)
	} else {
		config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND IF(? != '',category_id = ?,1=1) AND IF(? != '',concat(name,pinyin_full,pinyin_first) like ?,1=1)", shopId, categoryId, categoryId, keyword, "%"+keyword+"%").Order("product_id").Find(&productInfoList)
		if keyword == "" && len(productInfoList) > 0 {
			cache.SetCache("product:list:"+shopId, categoryId, productInfoList)
		}
	}
	if len(productInfoList) == 0 {
		return nil
	} else {
		return productInfoList
	}
}

func CountProductInfoByShopId(shopId, keyword, categoryId string) int {
	count := 0
	config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND IF(? != '',category_id = ?,1=1) AND IF(? != '',concat(name,pinyin_full,pinyin_first) like ?,1=1)", shopId, categoryId, categoryId, keyword, "%"+keyword+"%").Count(&count)
	return count
}

func GetProductInfoById(id int) *model.ProductInfo {
	var productInfo model.ProductInfo
	err := cache.GetCache("product:id", strconv.Itoa(id), &productInfo)
	if err == nil && productInfo.ShopId != "" {
		return &productInfo
	}
	config.Mysql.Table(TABLE_PRODUCT_INFO).Where("id = ?", id).First(&productInfo)
	if productInfo.Id == 0 {
		return nil
	} else {
		cache.SetCache("product:id", strconv.Itoa(id), productInfo)
		return &productInfo
	}
}

func GetProductInfo(shopId, productId string) *model.ProductInfo {
	var productInfo model.ProductInfo
	err := cache.GetCache("product:shop:"+shopId, productId, &productInfo)
	if err == nil && productInfo.ShopId != "" {
		return &productInfo
	}
	config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND product_id = ?", shopId, productId).First(&productInfo)
	if productInfo.Id == 0 {
		return nil
	} else {
		cache.SetCache("product:shop:"+shopId, productId, productInfo)
		return &productInfo
	}
}

func GetProductInfoByBarCode(shopId, barCode string) *model.ProductInfo {
	var productInfo model.ProductInfo
	config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND bar_code = ?", shopId, barCode).First(&productInfo)
	if productInfo.Id == 0 {
		return nil
	} else {
		return &productInfo
	}
}

func GetProductInfoByName(shopId, productName string) *model.ProductInfo {
	var productInfo model.ProductInfo
	config.Mysql.Table(TABLE_PRODUCT_INFO).Where("shop_id = ? AND name = ?", shopId, productName).First(&productInfo)
	if productInfo.Id == 0 {
		return nil
	} else {
		return &productInfo
	}
}

func SaveProductInfo(productInfo *model.ProductInfo) (*model.ProductInfo, error) {
	if productInfo.Id > 0 {
		d := GetProductInfoById(productInfo.Id)
		if d != nil {
			return d, nil
		}
	} else {
		d := GetProductInfo(productInfo.ShopId, productInfo.ProductId)
		if d != nil {
			return d, nil
		}
	}
	if productInfo.PyFull == "" {
		productInfo.PyFull = utils.ToPinYin(productInfo.Name, true, false)
	}
	if productInfo.PyFirst == "" {
		productInfo.PyFirst = utils.ToPinYin(productInfo.Name, false, false)
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_INFO).Create(productInfo).Error
	if err != nil {
		logs.Error("插入数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("product:list:"+productInfo.ShopId, productInfo.CategoryId)
	return GetProductInfoById(productInfo.Id), nil
}

func UpdateProductInfo(productInfo *model.ProductInfo) (*model.ProductInfo, error) {
	if productInfo.Id == 0 {
		return nil, errors.New("未指定商品编号")
	}
	data := make(map[string]interface{})
	if productInfo.ShopId != "" {
		data["shop_id"] = productInfo.ShopId
	}
	if productInfo.ProductId != "" {
		p := GetProductInfo(productInfo.ShopId, productInfo.ProductId)
		if p != nil && p.Id != productInfo.Id {
			return nil, errors.New("商品自定义编号已存在，不可修改")
		}
		data["product_id"] = productInfo.ProductId
	}
	if productInfo.Name != "" {
		data["name"] = productInfo.Name
		data["pinyin_full"] = utils.ToPinYin(productInfo.Name, true, false)
		data["pinyin_first"] = utils.ToPinYin(productInfo.Name, false, false)
	}
	if productInfo.CategoryId != "" {
		data["category_id"] = productInfo.CategoryId
	}
	if productInfo.BaseUnit != "" {
		data["base_unit"] = productInfo.BaseUnit
	}
	if productInfo.BarCode != "" {
		data["bar_code"] = productInfo.BarCode
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_INFO).Where("id = ?", productInfo.Id).Update(data).Error
	if err != nil {
		logs.Error("更新数据错误:{}", err.Error())
		tx.Rollback()
		return nil, err
	}
	tx.Commit()
	cache.DelCache("product:list:"+productInfo.ShopId, productInfo.CategoryId)
	cache.DelCache("product:id", strconv.Itoa(productInfo.Id))
	cache.DelCache("product:shop:"+productInfo.ShopId, productInfo.ProductId)
	return productInfo, nil
}

func DeleteProductInfo(productInfoId int) error {
	productInfo := GetProductInfoById(productInfoId)
	if productInfoId == 0 {
		return errors.New("未指定商品编号")
	}
	tx := config.Mysql.Begin()
	err := tx.Table(TABLE_PRODUCT_INFO).Where("id = ?", productInfoId).Delete(model.ProductInfo{}).Error
	if err != nil {
		logs.Error("删除数据错误:{}", err.Error())
		tx.Rollback()
		return err
	}
	tx.Commit()
	cache.DelCache("product:list:"+productInfo.ShopId, productInfo.CategoryId)
	cache.DelCache("product:id", strconv.Itoa(productInfoId))
	cache.DelCache("product:shop:"+productInfo.ShopId, productInfo.ProductId)
	return nil
}
