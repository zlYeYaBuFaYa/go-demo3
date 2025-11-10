# 图书管理后台 Go 项目文档

## 一、项目简介

本项目是一个基于 Golang 实现的图书管理后台 Demo。核心特点：

- 基于 Gin 实现 Restful API
- GORM + MySQL 持久化存储
- 配置与日志解耦，支持配置文件灵活定义
- 完善的分层架构，强可维护、方便扩展
- 实现 Book 实体的增删改查，带有统一接口响应结构
- 启动时自动生成 10 条真实中文图书测试数据 -- 已生成

---

## 二、项目结构说明

```
go_demo2/
├─ cmd/
│   └─ main.go           # 程序入口，调用注册/配置/路由
├─ internal/
│   ├─ app/v1              # 业务接口 handler 及路由注册
│   │   ├─ book_handler.go
│   │- service         #  定义业务逻辑,与store层交互
|   |    |- book.go    #  book相关
|   |
|   |-store             # 数据持久化逻辑,定义与数据库交互
|   |   |- book.go      # book模型的数据库操作
|
|   |-router
|   |   |-router.go  # 总路由入口
|   |   |- book.go   # book api路由   
|
| 
│   ├─ models/           # 所有领域模型
│   │   └─ book.go
│   └─ response/         # 统一API响应结构体
│       └─ response.go   #  定义SuccessResponse, FailResponse,封装c.Json,根据传递的data是否为空,setData.
|-config
|   |-config.go 配置文件
|
|-global
|  |-global.go 全局变量文件, 定义global.db,  global.log等变量,用于全局访问
|  |-variabal.go  全局变量
|
|
├─ pkg/
│   ├─ db/               # 数据库初始化与工具
│   │   └─ db.go
│   └─ logger/           # 日志初始化与工具
│       └─ logger.go
├─ etc/
│   └─ config.yaml       # 数据库、日志等配置
├─ go.mod、go.sum
└─ Readme.md(你当前看到的本文件)
```

---

## 三、功能模块及实现要点

### 1. 配置集中管理

- 所有数据库、日志、全局参数在 **etc/config.yaml** 中集中管理
- main.go 启动时自动加载，无硬编码安全隐患

### 2. 日志模块

- 使用 zap+zapcore+lumberjack
- 支持日志文件名、级别、大小、保存天数、最大数量、压缩等参数自定义

### 3. 数据库模块

- 支持最大连接: 100、最大空闲连接: 10、设置连接超时时间,空闲最大时间
- mysql 连接参数全部在 config.yaml 可灵活切换

### 4. Book 模型

- 字段包括 ID, Name, Author, Category, PublishedDate, Description, CreatedAt, UpdatedAt, Removed
- 位于 internal/models/book.go

### 5. API接口（RESTful）

- 接口路由定义在 internal/router目录中, router.go 为总入口, book.go为book接口路由定义
- /api/books POST/GET/PUT/DELETE，全部支持
- 业务 handler 定义在 internal/app/v1/book_handler.go，所有数据库调用、数据校验与响应解耦，示范最佳开发习惯
- handler 与模型彻底分离，便于后续测试与维护
- internal/service 定义业务逻辑,与store层交互, book.go为book相关
- internal/store 数据持久化逻辑,定义与数据库交互, book.go为book模型操作
- 调用response封装的SuccessResponse, FailResponse return返回结果

### 6. 统一响应结构

- internal/response/response.go 定义SuccessResponse, FailResponse,封装c.Json,根据传递的data是否为空,setData.
- 支持自动省略 data 字段，符合主流返回规则

### 7. 统一配置文件
- config/config.go 定义全局配置


### 8. 全局变量
- global/global.go 全局变量文件, 定义global.db,  global.log等变量,用于全局访问
- glbal/variabal.go 其他全局引用的常量或变量 根据实际情况定义在global, variabal文件中


### 9. 初始化测试数据 - 已创建表,跳过

- 创建book表
```
CREATE TABLE `books` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT '主键',
  `name` VARCHAR(100) NOT NULL COMMENT '书名',
  `author` VARCHAR(100) NOT NULL COMMENT '作者',
  `category` VARCHAR(50) NOT NULL COMMENT '类别',
  `published_date` DATE COMMENT '出版日期',
  `description` TEXT COMMENT '内容简介',
  `created_at` DATETIME COMMENT '创建时间',
  `updated_at` DATETIME COMMENT '更新时间',
  `removed` BOOLEAN DEFAULT FALSE COMMENT '逻辑删除'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
```

- 插入数据
```
INSERT INTO `books` (`name`, `author`, `category`, `published_date`, `description`, `created_at`, `updated_at`, `removed`) VALUES
('三体', '刘慈欣', '科幻', '2008-01-01', '外星文明与人类命运', NOW(), NOW(), 0),
('活着', '余华', '小说', '1993-01-01', '生命的坚韧和苦难', NOW(), NOW(), 0),
('解忧杂货店', '东野圭吾', '推理', '2013-03-01', '温情推理小说', NOW(), NOW(), 0),
('平凡的世界', '路遥', '文学', '1986-01-01', '奋斗与理想', NOW(), NOW(), 0),
('小王子', '圣埃克苏佩里', '童话', '1943-04-01', '成长寓言故事', NOW(), NOW(), 0),
('人类简史', '尤瓦尔·赫拉利', '历史', '2011-01-01', '从动物到上帝', NOW(), NOW(), 0),
('围城', '钱钟书', '小说', '1947-05-01', '婚姻爱情讽刺', NOW(), NOW(), 0),
('嫌疑人X的献身', '东野圭吾', '推理', '2005-08-01', '天才之间的较量', NOW(), NOW(), 0),
('西游记', '吴承恩', '古典', '1592-01-01', '中国古典神魔小说', NOW(), NOW(), 0),
('百年孤独', '加西亚·马尔克斯', '魔幻', '1967-01-01', '马孔多家族传奇', NOW(), NOW(), 0);
```

---

## 四、数据库配置示例（etc/config.yaml）

```yaml
db:
  user: root
  password: 123456
  host: 127.0.0.1
  port: 3306
  name: demo
  charset: utf8mb4
  max_open_conns: 50
  max_idle_conns: 10
  conn_max_lifetime: 3600

log:
  file: app.log
  level: info
  max_size: 100    # 单个文件最大100MB
  max_age: 31      # 日志保存天数
  max_backups: 31  # 最多保存31个文件
  compress: true   # 打开log压缩
```

---

## 五、运行方式

1. 启动数据库（如需创建 demo 库，提前创建好）
2. 根据实际环境修改 etc/config.yaml
3. 项目根目录执行：

```bash
cd cmd
go run main.go
# 或编译后运行
go build -o server main.go && ./server
```

4. API测试：
   - 健康检查：GET /ping
   - 图书列表：GET /api/books
   - 新建/更新/删除见 handler 路由定义，支持全部 RESTful

---

## 六、接口返回体样例

- 成功（有 data）：

```json
{
  "code": 0,
  "message": "success",
  "data": {...}
}
```

- 成功（无 data，仅 code/message）：

```json
{
  "code": 0,
  "message": "success"
}
```

- 失败：

```json
{
  "code": -1,
  "message": "失败信息"
}
```

---

## 七、后续扩展建议

- 用户/权限/登录模块可基于同样分层快速扩展
- 可引入Swagger自动化接口文档
- 增加自定义中间件：日志、鉴权、限流
- 集成单元测试、接口测试以保证迭代安全

---

如需自动生成 API 测试脚本、文档、mock 数据、持续集成、k8s部署等辅助服务，可随时沟通！  
如有格式、内容等需要微调请说明，我会帮你持续完善！
