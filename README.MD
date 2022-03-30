# gf - Generic Functions For Go 1.18+

![Build And Test](https://github.com/tiancheng92/gf/actions/workflows/gf.yml/badge.svg)

### 注意：部分方法仅可在Go版本大于等于1.18时使用！

#### 什么是gf？
* gf 是一个容纳一些比较常用的函数的库，主要用于简化 Go 编程。

#### gf包含哪些功能？
* 字符串
  * 高性能的string与[]byte之间的转换方法
  * 字符串Gzip压缩、解压
  * 高性能的字符串拼接
* 数组
  * 去重
  * 判断包含关系
  * 是否相等
  * 交集
  * 并集
  * 差集
  * 过滤
  * 反转
  * 排序（插入、冒泡、快速、堆）只需实现less方法即可
* 浮点数
  * 保留指定位小数
* URL
  * 格式化URL