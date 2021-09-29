# micro-jxc
一个基于Go微服务的进销存系统API后端

## 模块组成
本小型进销存系统包括商品模块、库存模块、单据模块、公共基础模块和统一网关模块组成
子模块|模块功能
---|:---
jxc-base|公共基础模块，含数据模型和基础配置数据增删改查接口
jxc-product|产品模块，含商品定义、规格管理、Sku管理、供应商管理、单位管理接口
jxc-order|单据模块，含入库单、出库单、移库单等单据的管理接口
jxc-stock|库存模块，含入库、出库、取库存等接口
jxc-api|网关模块，统一对外接口API模块，业务逻辑整合

## 系统使用的数据库
- MySQL 5.7
- MongoDB 3.4
- Redis 4.0

## 微服务相关模块
- 配置中心 SpringCloud Config Server
- 注册中心 Nacos

## 说明
本系统只有后端API部分，前端是Android/iOS原生APP和微信小程序
