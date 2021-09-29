CREATE DATABASE `jxc` /*!40100 DEFAULT CHARACTER SET utf8 */

CREATE TABLE `delivery_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(10) DEFAULT NULL COMMENT '出库类型',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  PRIMARY KEY (`id`),
  KEY `delivery_type_shop_id_index` (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='出库类型表'

CREATE TABLE `entry_type` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `type` varchar(10) DEFAULT NULL COMMENT '入库类型',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  PRIMARY KEY (`id`),
  KEY `entry_type_shop_id_index` (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='入库类型表'

CREATE TABLE `product_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `parent_id` varchar(10) DEFAULT NULL COMMENT '上级分类id',
  `name` varchar(50) DEFAULT NULL COMMENT '分类名称',
  `level` int(11) DEFAULT NULL COMMENT '分类层级',
  `sort_number` int(11) DEFAULT NULL COMMENT '前端排序值',
  `category_id` varchar(10) DEFAULT NULL COMMENT '商户自定义类别编号',
  PRIMARY KEY (`id`),
  KEY `product_category_sort_number_index` (`sort_number`),
  KEY `product_category_shop_id_category_id_index` (`shop_id`,`category_id`),
  KEY `product_category_shop_id_parent_id_level_index` (`shop_id`,`parent_id`,`level`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品分类表'

CREATE TABLE `product_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `name` varchar(256) DEFAULT NULL COMMENT '商品名称',
  `category_id` varchar(10) DEFAULT NULL COMMENT '分类id',
  `base_unit` varchar(10) DEFAULT NULL COMMENT '基础单位',
  `bar_code` varchar(20) DEFAULT NULL COMMENT '条码',
  `pinyin_full` varchar(512) DEFAULT NULL COMMENT '商品名称拼音全称,小写',
  `pinyin_first` varchar(50) DEFAULT NULL COMMENT '商品名称拼音首字母大写',
  `product_id` varchar(10) DEFAULT NULL COMMENT '自定义商品编号',
  PRIMARY KEY (`id`),
  KEY `product_info_shop_id_bar_code_index` (`shop_id`,`bar_code`),
  KEY `product_info_shop_id_category_id_index` (`shop_id`,`category_id`),
  KEY `product_info_shop_id_name_pinyin_first_pinyin_full_index` (`shop_id`,`name`,`pinyin_first`,`pinyin_full`),
  KEY `product_info_shop_id_product_id_index` (`shop_id`,`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品基础信息表'

CREATE TABLE `product_specs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` varchar(10) DEFAULT NULL COMMENT '商品id',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `name` varchar(100) DEFAULT NULL COMMENT '规格名称',
  `values` varchar(512) DEFAULT NULL COMMENT '规格预设值，多个，JSON数组',
  PRIMARY KEY (`id`),
  KEY `product_specs_shop_id_product_id_index` (`shop_id`,`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品规格表'

CREATE TABLE `product_supplier` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `product_id` varchar(10) DEFAULT NULL COMMENT '商品编号',
  `supplier_id` varchar(10) DEFAULT NULL COMMENT '供货商商户id',
  `supplier_name` varchar(100) DEFAULT NULL COMMENT '供货商名称',
  PRIMARY KEY (`id`),
  KEY `product_supplier_shop_id_product_id_index` (`shop_id`,`product_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品供货商表'

CREATE TABLE `product_unit` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `product_id` varchar(10) DEFAULT NULL COMMENT '商品id',
  `unit` varchar(10) DEFAULT NULL COMMENT '辅助单位',
  `base_unit` varchar(10) DEFAULT NULL COMMENT '基础单位',
  `scale` int(11) DEFAULT NULL COMMENT '比例',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户账号',
  PRIMARY KEY (`id`),
  KEY `product_unit_product_id_shop_id_index` (`product_id`,`shop_id`),
  KEY `product_unit_shop_id_product_id_unit_base_unit_index` (`shop_id`,`product_id`,`unit`,`base_unit`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商品多单位换算表'

CREATE TABLE `shop_storage` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `name` varchar(50) DEFAULT NULL COMMENT '仓库名称',
  `sort_number` int(11) DEFAULT NULL COMMENT '前端排序值',
  `storage_id` varchar(10) DEFAULT NULL COMMENT '自定义仓库编号',
  `remark` varchar(50) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `shop_storage_shop_id_sort_number_index` (`shop_id`,`sort_number`),
  KEY `shop_storage_shop_id_sn_sort_number_index` (`shop_id`,`storage_id`,`sort_number`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='商家仓库表'

CREATE TABLE `sku_stock` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `sku_id` varchar(32) DEFAULT NULL COMMENT '货品全局唯一id',
  `product_id` varchar(10) DEFAULT NULL COMMENT '商品id',
  `name` varchar(100) DEFAULT NULL COMMENT '商品名称',
  `stocks` int(11) DEFAULT NULL COMMENT '库存量',
  `base_unit` varchar(10) DEFAULT NULL COMMENT '基础单位',
  `storage_id` varchar(10) DEFAULT NULL COMMENT '仓库号',
  `cost_price` int(11) DEFAULT NULL COMMENT '加权平均成本，单位为分/件',
  `last_price` int(11) DEFAULT NULL COMMENT '最新成本价，单位为分/件',
  `update_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '最后更新时间',
  PRIMARY KEY (`id`),
  KEY `sku_stock_shop_id_product_id_index` (`shop_id`,`product_id`),
  KEY `sku_stock_shop_id_sku_guid_index` (`shop_id`,`sku_id`),
  KEY `sku_stock_shop_id_storage_id_index` (`shop_id`,`storage_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='货品库存表'

CREATE TABLE `stock_delivery` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `delivery_no` varchar(32) DEFAULT NULL COMMENT '出库单号',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `type` varchar(10) DEFAULT NULL COMMENT '出库类型',
  `sku_id` varchar(32) DEFAULT NULL COMMENT '货品全局唯一id',
  `unit` varchar(10) DEFAULT NULL COMMENT '计量单位',
  `order_no` varchar(32) DEFAULT NULL COMMENT '外部采购单号',
  `cost` int(11) DEFAULT NULL COMMENT '成本价',
  `price` int(11) DEFAULT NULL COMMENT '进货价',
  `storage_id` varchar(10) DEFAULT NULL COMMENT '仓库id',
  `multi_storage` varchar(256) DEFAULT NULL COMMENT '多仓库出货详情，JSON数组格式',
  `delivery_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '出库时间',
  `number` int(10) DEFAULT NULL COMMENT '出库件数',
  `customer_id` varchar(10) DEFAULT NULL COMMENT '客户shopid',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `operator` varchar(50) DEFAULT NULL COMMENT '操作员',
  `customer_name` varchar(50) DEFAULT NULL COMMENT '客户名称',
  PRIMARY KEY (`id`),
  KEY `stock_delivery_shop_id_delivery_no_index` (`shop_id`,`delivery_no`),
  KEY `stock_delivery_shop_id_order_no_index` (`shop_id`,`order_no`),
  KEY `stock_delivery_shop_id_sku_guid_delivery_time_index` (`shop_id`,`sku_id`,`delivery_time`),
  KEY `stock_delivery_shop_id_delivery_time_index` (`shop_id`,`delivery_time`),
  KEY `stock_delivery_shop_id_storage_id_delivery_time_index` (`shop_id`,`storage_id`,`delivery_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='出库单'

CREATE TABLE `stock_entry` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `entry_no` varchar(32) DEFAULT NULL COMMENT '入库单号',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `type` varchar(10) DEFAULT NULL COMMENT '入库类型',
  `sku_id` varchar(32) DEFAULT NULL COMMENT '货品全局唯一id',
  `unit` varchar(10) DEFAULT NULL COMMENT '计量单位',
  `order_no` varchar(32) DEFAULT NULL COMMENT '外部采购单号',
  `price` int(11) DEFAULT NULL COMMENT '进货价',
  `storage_id` varchar(10) DEFAULT NULL COMMENT '仓库id',
  `entry_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '入库时间',
  `number` int(11) DEFAULT NULL COMMENT '入库件数',
  `supplier_id` varchar(10) DEFAULT NULL COMMENT '供货商shopId',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `operator` varchar(50) DEFAULT NULL COMMENT '操作员',
  PRIMARY KEY (`id`),
  KEY `stock_entry_shop_id_entry_no_index` (`shop_id`,`entry_no`),
  KEY `stock_entry_shop_id_entry_time_index` (`shop_id`,`entry_time`),
  KEY `stock_entry_shop_id_order_no_index` (`shop_id`,`order_no`),
  KEY `stock_entry_shop_id_sku_guid_entry_time_index` (`shop_id`,`sku_id`,`entry_time`),
  KEY `stock_entry_shop_id_storage_id_entry_time_index` (`shop_id`,`storage_id`,`entry_time`),
  KEY `stock_entry_shop_id_type_entry_time_index` (`shop_id`,`type`,`entry_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='入库单'

CREATE TABLE `stock_move` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `move_no` varchar(32) DEFAULT NULL COMMENT '移库单号',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  `sku_id` varchar(32) DEFAULT NULL COMMENT '货品全局唯一id',
  `unit` varchar(10) DEFAULT NULL COMMENT '计量单位',
  `from_storage_id` varchar(10) DEFAULT NULL COMMENT '原仓库id',
  `to_storage_id` varchar(10) DEFAULT NULL COMMENT '新仓库id',
  `move_time` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '移库时间',
  `number` int(11) DEFAULT NULL COMMENT '移库件数',
  `remark` varchar(256) DEFAULT NULL COMMENT '备注',
  `operator` varchar(50) DEFAULT NULL COMMENT '操作员',
  PRIMARY KEY (`id`),
  KEY `stock_move_shop_id_from_storage_id_move_time_index` (`shop_id`,`from_storage_id`,`move_time`),
  KEY `stock_move_shop_id_move_time_index` (`shop_id`,`move_time`),
  KEY `stock_move_shop_id_sku_guid_move_time_index` (`shop_id`,`sku_id`,`move_time`),
  KEY `stock_move_shop_id_to_storage_id_move_time_index` (`shop_id`,`to_storage_id`,`move_time`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='移库单'

CREATE TABLE `units` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `unit` varchar(10) DEFAULT NULL COMMENT '单位名称',
  `shop_id` varchar(10) DEFAULT NULL COMMENT '商户id',
  PRIMARY KEY (`id`),
  KEY `units_shop_id_index` (`shop_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='单位表'

