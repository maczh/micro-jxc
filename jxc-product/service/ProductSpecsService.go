package service

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	"ququ.im/jxc-base/model"
	"ququ.im/jxc-product/dao"
)

func ListProductSpecs(shopId, productId string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户账号不可为空")
	}
	if productId == "" {
		return *common.Error(-1, "商品编号不可为空")
	}
	productSpecsList := dao.ListProductSpecsByProductId(shopId, productId)
	if productSpecsList == nil {
		return *common.Error(-1, "未定义产品规格")
	}
	return *common.Success(productSpecsList)
}

func GetProductSpecs(id int, shopId, productId, specsName string) common.Result {
	var productSpecs *model.ProductSpecs
	if id > 0 {
		productSpecs = dao.GetProductSpecsById(id)
	} else if shopId != "" && productId != "" && specsName != "" {
		productSpecs = dao.GetProductSpecs(shopId, productId, specsName)
	} else {
		return *common.Error(-1, "参数不全")
	}
	if productSpecs == nil {
		return *common.Error(-1, "未定义产品规格")
	}
	return *common.Success(productSpecs)
}

func SaveProductSpecs(shopId, productId, specsName, specsValues string) common.Result {
	if shopId == "" {
		return *common.Error(-1, "商户编号不可为空")
	}
	if specsName == "" {
		return *common.Error(-1, "商品规格名称不可为空")
	}
	if specsValues == "" {
		return *common.Error(-1, "商品规格内容不可为空")
	}
	if productId == "" {
		return *common.Error(-1, "商品编号不可为空")
	}
	productSpecs := dao.GetProductSpecs(shopId, productId, specsName)
	if productSpecs != nil {
		return UpdateProductSpecs(productSpecs.Id, shopId, productId, specsName, specsValues)
	}
	productSpecs = new(model.ProductSpecs)
	productSpecs.ShopId = shopId
	productSpecs.Name = specsName
	productSpecs.Values = specsValues
	productSpecs.ProductId = productId
	productSpecs, err := dao.SaveProductSpecs(productSpecs)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSpecs(shopId, productId, specsName))
}

func UpdateProductSpecs(id int, shopId, productId, specsName, specsValues string) common.Result {
	productSpecs := dao.GetProductSpecsById(id)
	if productSpecs == nil {
		return *common.Error(-1, "无此商品规格记录")
	}
	productSpecs.ProductId = productId
	productSpecs.ShopId = shopId
	productSpecs.Name = specsName
	productSpecs.Values = specsValues
	productSpecs, err := dao.UpdateProductSpecs(productSpecs)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSpecsById(id))
}

func DeleteProductSpecs(id int) common.Result {
	err := dao.DeleteProductSpecs(id)
	if err != nil {
		return *common.Error(-1, "无此商品规格记录")
	}
	return *common.Success(nil)
}

func AddSpecsValue(id int, value string) common.Result {
	productSpecs := dao.GetProductSpecsById(id)
	if productSpecs == nil {
		return *common.Error(-1, "无此商品规格记录")
	}
	var values []string
	utils.FromJSON(productSpecs.Values, &values)
	exists := false
	for _, v := range values {
		if v == value {
			exists = true
		}
	}
	if !exists {
		values = append(values, value)
	}
	productSpecs.Values = utils.ToJSON(values)
	_, err := dao.UpdateProductSpecs(productSpecs)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSpecsById(id))
}

func RemoveSpecsValue(id int, value string) common.Result {
	productSpecs := dao.GetProductSpecsById(id)
	if productSpecs == nil {
		return *common.Error(-1, "无此商品规格记录")
	}
	var values []string
	utils.FromJSON(productSpecs.Values, &values)
	for k, v := range values {
		if v == value {
			values = append(values[:k], values[k+1:]...)
			break
		}
	}
	productSpecs.Values = utils.ToJSON(values)
	_, err := dao.UpdateProductSpecs(productSpecs)
	if err != nil {
		return *common.Error(-1, err.Error())
	}
	return *common.Success(dao.GetProductSpecsById(id))
}
