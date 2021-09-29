package controller

import (
	"ququ.im/common"
	"ququ.im/common/utils"
	_ "ququ.im/jxc-base/model"
	"ququ.im/jxc-product/service"
	"strconv"
)

// ListProductCategory	godoc
// @Summary		列出商品的所有商品分类接口
// @Description	按菜单层级方式或路径方式列出商品分类
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string false "商户账号"
// @Param	categoryId formData string false "商品分类号，按编号查完整分类路径"
// @Param	parent formData string false "上级分类编号"
// @Param	level formData int false "分类层级"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/list [post][get]
func ListProductCategory(params map[string]string) common.Result {
	level := 0
	if utils.Exists(params, "level") {
		level, _ = strconv.Atoi(params["level"])
	}
	return service.ListProductCategory(params["shopId"], params["categoryId"], params["parent"], level)
}

// GetProductCategory	godoc
// @Summary		查询商品分类接口
// @Description	查询商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int false "商品分类记录号"
// @Param	shopId formData string false "商户账号"
// @Param	categoryId formData string false "商品分类自定义编号"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/get [post][get]
func GetProductCategory(params map[string]string) common.Result {
	id := 0
	if utils.Exists(params, "id") {
		id, _ = strconv.Atoi(params["id"])
	}
	return service.GetProductCategory(id, params["shopId"], params["categoryId"])
}

// SaveProductCategory	godoc
// @Summary		保存商品分类接口
// @Description	保存商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	shopId formData string true "商户账户"
// @Param	name formData string true "商品分类名称"
// @Param	categoryId formData string false "商品分类自定义编号,默认自动生成"
// @Param	parent formData string true "上级分类编号，第一级无上级分类的填0"
// @Param	level formData int true "商品分类层级，从第1级向下"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/save [post][get]
func SaveProductCategory(params map[string]string) common.Result {
	if !utils.Exists(params, "parent") {
		return *common.Error(-1, "未传入商品上级分类编号参数")
	}
	if !utils.Exists(params, "level") {
		return *common.Error(-1, "未传入商品分类层级，从1级开始向下")
	}
	if !utils.Exists(params, "shopId") {
		return *common.Error(-1, "未传入商户账号")
	}
	if !utils.Exists(params, "name") {
		return *common.Error(-1, "未传入商品分类名称")
	}
	level, _ := strconv.Atoi(params["level"])
	return service.SaveProductCategory(params["shopId"], params["name"], params["categoryId"], params["parent"], level)
}

// UpdateProductCategory	godoc
// @Summary		修改商品分类接口
// @Description	修改商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品分类编号"
// @Param	name formData string true "商品分类名称"
// @Param	categoryId formData string false "商品分类自定义编号"
// @Param	parent formData string false "上级分类编号，第一级无上级分类的填0"
// @Param	level formData int false "商品分类层级，从第1级向下"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/update [post][get]
func UpdateProductCategory(params map[string]string) common.Result {
	id, level := 0, 0
	if utils.Exists(params, "id") {
		id, _ = strconv.Atoi(params["id"])
	}
	if utils.Exists(params, "level") {
		level, _ = strconv.Atoi(params["level"])
	}
	return service.UpdateProductCategory(id, level, params["parent"], params["name"], params["categoryId"])
}

// DeleteProductCategory	godoc
// @Summary		删除商品分类接口
// @Description	删除商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品分类记录号"
// @Success 200 {object} common.Result	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/del [post][get]
func DeleteProductCategory(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品分类编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DeleteProductCategory(id)
}

// DeleteProductCategory	godoc
// @Summary		下移商品分类接口
// @Description	下移商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品分类记录号"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/down [post][get]
func IncrProductCategorySortNumber(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品分类编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.IncrProductCategorySortNumber(id)
}

// DeleteProductCategory	godoc
// @Summary		上移商品分类接口
// @Description	上移商品分类接口
// @Accept	x-www-form-urlencoded
// @Produce json
// @Param	id formData int true "商品分类记录号"
// @Success 200 {object} model.ProductCategory	"{\"status\":1,<br>\"msg\":\"成功\",<br>\"data\":{},<br>\"page\":null}"
// @Router	/category/up [post][get]
func DecrProductCategorySortNumber(params map[string]string) common.Result {
	if !utils.Exists(params, "id") {
		return *common.Error(-1, "未传入商品分类编号参数")
	}
	id, _ := strconv.Atoi(params["id"])
	return service.DecrProductCategorySortNumber(id)
}
