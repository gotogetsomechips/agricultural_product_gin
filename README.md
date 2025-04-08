# agricultural_product_gin

#### 介绍

农产品溯源管理系统的后台



#### 软件架构

Gin



#### 安装教程

1. 克隆仓库：`https://github.com/gotogetsomechips/agricultural_product_gin.git`
2. 初始化数据库：在MySQL中创建一个数据库`traceability`，并执行数据库初始化脚本`traceability.sql`
3. 在**config**目录下的`config.go`文件，进行数据的相关配置（账号、密码）
4. `config.go`内容示例如下：

```
const (
	DBUsername = "root"
	DBPassword = "123456"
	DBHost     = "localhost"
	DBPort     = "3306"
	DBName     = "traceability"
)
```

5. go mod tidy
6. go run main.go
