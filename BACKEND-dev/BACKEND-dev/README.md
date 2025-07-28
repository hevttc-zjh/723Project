# 情指行风险洞察系统

## 项目简介

情指行风险洞察系统是一个集成化、可视化、智能化的社会人员画像平台，帮助公安机关通过警情关联、身份证号、姓名、手机号等多种方式快速查询、分析社会人员的风险画像与社会关系网络，实现"多源数据整合 + 重点人员研判 + 风险预警响应"的一体化支撑。

## 技术栈

- **后端框架**: Gin
- **数据库**: PostgreSQL
- **ORM**: GORM
- **认证**: JWT
- **配置管理**: Viper
- **开发语言**: Go 1.21+

## 项目结构

```
risk-insight-system/
├── config/                 # 配置管理
│   ├── config.go          # 配置初始化
│   └── config.yaml        # 配置文件
├── internal/              # 内部包
│   ├── api/           # 处理器
│   │   ├── person.go      # 人员画像相关
│   │   ├── data.go        # 数据接入相关
│   │   ├── feedback.go    # 反馈信息相关
│   │   └── stats.go       # 统计分析相关
│   ├── middleware/        # 中间件
│   │   └── cors.go        # CORS中间件
│   ├── model/             # 数据模型
│   ├── repository/        # 数据访问层
│   ├── service/           # 业务逻辑层
│   ├── router/            # 路由
│   │   └── router.go      # 路由初始化
│   └── server/            # 服务器
│       └── database.go    # 数据库连接
├── logs/                  # 日志文件
├── docs/                  # 文档
├── scripts/               # 脚本文件
├── go.mod                 # Go模块文件
├── go.sum                 # Go依赖文件
├── main.go                # 主程序入口
└── README.md              # 项目说明
```

## 快速开始

### 环境要求

- Go 1.21+
- PostgreSQL 12+
- Git

### 安装步骤

1. 克隆项目
```bash
git clone <repository-url>
cd risk-insight-system
```

2. 安装依赖
```bash
go mod tidy
```

3. 配置数据库
```bash
# 创建数据库
createdb risk_insight

# 修改配置文件 config/config.yaml 中的数据库连接信息
```

4. 运行项目
```bash
go run main.go
```

5. 访问服务
```
健康检查: http://localhost:8080/health
API文档: http://localhost:8080/api/v1
```

## API接口

### 人员画像相关
- `GET /api/v1/person/search` - 搜索人员
- `GET /api/v1/person/:id` - 获取人员画像
- `GET /api/v1/person/:id/relations` - 获取人员社会关系
- `GET /api/v1/person/:id/cases` - 获取人员案件记录

### 数据接入相关
- `POST /api/v1/data/police` - 接收警情数据
- `POST /api/v1/data/social` - 接收社会数据
- `POST /api/v1/data/case` - 接收案件数据
- `POST /api/v1/data/internal` - 接收公安内部数据

### 反馈信息相关
- `POST /api/v1/feedback/` - 创建反馈信息
- `GET /api/v1/feedback/:person_id` - 获取人员反馈信息
- `PUT /api/v1/feedback/:id` - 更新反馈信息

### 统计分析相关
- `GET /api/v1/stats/dashboard` - 获取首页大屏统计数据
- `GET /api/v1/stats/distribution` - 获取人员分布统计
- `GET /api/v1/stats/risk-levels` - 获取风险等级统计

## 开发规范

请参考 `docs/development-guidelines.md` 文件了解详细的开发规范。

## 贡献指南

1. Fork 项目
2. 创建功能分支 (`git checkout -b feature/AmazingFeature`)
3. 提交更改 (`git commit -m 'Add some AmazingFeature'`)
4. 推送到分支 (`git push origin feature/AmazingFeature`)
5. 打开 Pull Request

## 许可证

本项目采用 MIT 许可证 - 查看 [LICENSE](LICENSE) 文件了解详情。 