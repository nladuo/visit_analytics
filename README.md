# visit_analytics
自用的网站访问量统计工具

## STATUS
正在编写中...

## 使用
### 下载安装
``` sh
git clone https://github.com/nladuo/visit_analytics.git
cd visit_analytics
make prepare 		# 下载库文件
```

### 创建数据库
``` sql
CREATE DATABASE `visit_analytics` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
```

### 配置数据库
打开db.go, 配置mysql。
``` go
const (
	DB_USER = "root"
	DB_PASS = "root"
	DB_HOST = "localhost"
	DB_PORT = "3306"
	DBNAME  = "vist_analytics"
)
```

### 安装运行
``` sh
make && ./visit_analytics
```

## 使用
添加`<script type="text/javascript" src="http://localhost:3000/analytics.js"></script>`到要统计的网页中。

## LICENSE
MIT
