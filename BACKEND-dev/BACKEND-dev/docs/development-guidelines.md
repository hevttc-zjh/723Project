# 情指行风险洞察系统开发规范文档

## 1. 项目概述

### 1.1 项目目标
打造一个集成化、可视化、智能化的社会人员画像平台，帮助公安机关通过多种方式快速查询、分析社会人员的风险画像与社会关系网络。

### 1.2 技术架构
- **后端**: Gin + GORM + PostgreSQL
- **认证**: JWT
- **配置**: Viper
- **开发语言**: Go 1.21+

### 1.3 团队规模
- 后端开发: 3人
- 前端开发: 2人
- 项目经理: 1人
- 总计: 6人

## 2. 代码规范

### 2.1 命名规范

#### 2.1.1 包命名
- 使用小写字母，单词间用下划线分隔
- 包名应该简洁明了，表达包的功能
```go
// 正确
package user_service
package data_repository

// 错误
package UserService
package dataRepository
```

#### 2.1.2 文件命名
- 使用小写字母和下划线
- 文件名应该反映其内容
```go
// 正确
user_handler.go
person_service.go
database_config.go

// 错误
UserHandler.go
personService.go
```

#### 2.1.3 变量命名
- 使用驼峰命名法
- 私有变量以小写字母开头
- 公共变量以大写字母开头
```go
// 正确
var userName string
var UserID int
var isActive bool

// 错误
var user_name string
var userid int
```

#### 2.1.4 函数命名
- 使用驼峰命名法
- 函数名应该清晰表达其功能
```go
// 正确
func GetUserByID(id int) (*User, error)
func CreatePersonProfile(person *Person) error
func ValidateIDCard(idCard string) bool

// 错误
func getUser(id int) (*User, error)
func create_person(person *Person) error
```

### 2.2 目录结构规范

```
risk-insight-system/
├── config/                 # 配置管理
│   ├── config.go          # 配置初始化
│   └── config.yaml        # 配置文件
├── internal/              # 内部包
│   ├── handler/           # 处理器层 (HTTP请求处理)
│   │   ├── person.go      # 人员画像相关
│   │   ├── data.go        # 数据接入相关
│   │   ├── feedback.go    # 反馈信息相关
│   │   └── stats.go       # 统计分析相关
│   ├── middleware/        # 中间件
│   │   ├── cors.go        # CORS中间件
│   │   ├── auth.go        # 认证中间件
│   │   ├── logger.go      # 日志中间件
│   │   └── validator.go   # 参数验证中间件
│   ├── model/             # 数据模型层
│   │   ├── person.go      # 人员模型
│   │   ├── case.go        # 案件模型
│   │   ├── feedback.go    # 反馈模型
│   │   └── common.go      # 公共模型
│   ├── repository/        # 数据访问层
│   │   ├── person.go      # 人员数据访问
│   │   ├── case.go        # 案件数据访问
│   │   ├── feedback.go    # 反馈数据访问
│   │   └── base.go        # 基础数据访问
│   ├── service/           # 业务逻辑层
│   │   ├── person.go      # 人员业务逻辑
│   │   ├── data.go        # 数据处理逻辑
│   │   ├── feedback.go    # 反馈业务逻辑
│   │   └── stats.go       # 统计业务逻辑
│   ├── router/            # 路由层
│   │   ├── router.go      # 路由初始化
│   │   ├── person.go      # 人员路由
│   │   ├── data.go        # 数据路由
│   │   └── stats.go       # 统计路由
│   ├── server/            # 服务器层
│   │   ├── database.go    # 数据库连接
│   │   ├── redis.go       # Redis连接
│   │   └── server.go      # 服务器配置
│   └── utils/             # 工具包
│       ├── jwt.go         # JWT工具
│       ├── validator.go   # 验证工具
│       ├── logger.go      # 日志工具
│       └── response.go    # 响应工具
├── logs/                  # 日志文件
├── docs/                  # 文档
│   ├── api/               # API文档
│   ├── database/          # 数据库文档
│   └── deployment/        # 部署文档
├── scripts/               # 脚本文件
│   ├── build.sh           # 构建脚本
│   ├── deploy.sh          # 部署脚本
│   └── migrate.sh         # 数据库迁移脚本
├── tests/                 # 测试文件
│   ├── unit/              # 单元测试
│   ├── integration/       # 集成测试
│   └── e2e/               # 端到端测试
├── go.mod                 # Go模块文件
├── go.sum                 # Go依赖文件
├── main.go                # 主程序入口
└── README.md              # 项目说明
```

## 3. 开发流程规范

### 3.1 分支管理

#### 3.1.1 分支命名
- `main`: 主分支，用于生产环境
- `develop`: 开发分支，用于集成测试
- `feature/功能名称`: 功能开发分支
- `hotfix/问题描述`: 紧急修复分支
- `release/版本号`: 发布分支

#### 3.1.2 工作流程
1. 从 `develop` 分支创建 `feature` 分支
2. 在 `feature` 分支上开发功能
3. 完成开发后提交 Pull Request 到 `develop` 分支
4. 代码审查通过后合并到 `develop` 分支
5. 定期从 `develop` 分支创建 `release` 分支
6. 测试通过后合并到 `main` 分支

### 3.2 提交规范

#### 3.2.1 提交信息格式
```
<type>(<scope>): <subject>

<body>

<footer>
```

#### 3.2.2 类型说明
- `feat`: 新功能
- `fix`: 修复bug
- `docs`: 文档更新
- `style`: 代码格式调整
- `refactor`: 代码重构
- `test`: 测试相关
- `chore`: 构建过程或辅助工具的变动

#### 3.2.3 示例
```
feat(person): 添加人员画像查询功能

- 实现按身份证号查询人员信息
- 添加人员基础信息展示
- 集成警情记录查询

Closes #123
```

### 3.3 代码审查规范

#### 3.3.1 审查要点
- 代码是否符合项目规范
- 功能实现是否正确
- 是否有潜在的安全问题
- 性能是否合理
- 测试覆盖率是否足够

#### 3.3.2 审查流程
1. 开发者提交 Pull Request
2. 至少一名团队成员进行代码审查
3. 审查通过后合并代码
4. 如有问题，开发者修改后重新提交

## 4. API设计规范

### 4.1 RESTful API设计

#### 4.1.1 URL设计
```
GET    /api/v1/persons          # 获取人员列表
GET    /api/v1/persons/:id      # 获取特定人员
POST   /api/v1/persons          # 创建人员
PUT    /api/v1/persons/:id      # 更新人员
DELETE /api/v1/persons/:id      # 删除人员
```

#### 4.1.2 响应格式
```json
{
  "code": 200,
  "message": "操作成功",
  "data": {
    // 具体数据
  },
  "timestamp": "2024-01-01T00:00:00Z"
}
```

#### 4.1.3 错误处理
```json
{
  "code": 400,
  "message": "参数错误",
  "errors": [
    {
      "field": "id_card",
      "message": "身份证号格式不正确"
    }
  ],
  "timestamp": "2024-01-01T00:00:00Z"
}
```

### 4.2 状态码规范
- `200`: 成功
- `201`: 创建成功
- `400`: 请求参数错误
- `401`: 未授权
- `403`: 禁止访问
- `404`: 资源不存在
- `500`: 服务器内部错误

## 5. 数据库设计规范

### 5.1 表命名规范
- 使用小写字母和下划线
- 表名应该清晰表达其用途
```sql
-- 正确
persons
police_cases
social_relations
person_feedbacks

-- 错误
Persons
policeCases
socialRelations
```

### 5.2 字段命名规范
- 使用小写字母和下划线
- 主键统一使用 `id`
- 外键使用 `表名_id` 格式
```sql
-- 正确
CREATE TABLE persons (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    id_card VARCHAR(18),
    phone VARCHAR(20),
    created_at TIMESTAMP,
    updated_at TIMESTAMP
);

-- 错误
CREATE TABLE Persons (
    ID SERIAL PRIMARY KEY,
    Name VARCHAR(100),
    idCard VARCHAR(18)
);
```

### 5.3 索引规范
- 主键自动创建索引
- 外键字段创建索引
- 经常查询的字段创建索引
- 复合查询创建复合索引

## 6. 安全规范

### 6.1 认证授权
- 使用JWT进行身份认证
- 实现基于角色的访问控制(RBAC)
- 敏感操作需要二次验证

### 6.2 数据安全
- 敏感数据加密存储
- 数据传输使用HTTPS
- 实现SQL注入防护
- 实现XSS防护

### 6.3 日志安全
- 记录用户操作日志
- 敏感信息脱敏处理
- 定期清理日志文件

## 7. 性能规范

### 7.1 数据库性能
- 合理使用索引
- 避免N+1查询问题
- 使用连接池
- 定期优化慢查询

### 7.2 缓存策略
- 使用Redis缓存热点数据
- 实现合理的缓存失效策略
- 避免缓存穿透和雪崩

### 7.3 并发处理
- 使用goroutine处理并发请求
- 实现合理的限流策略
- 避免死锁和竞态条件

## 8. 测试规范

### 8.1 单元测试
- 测试覆盖率不低于80%
- 每个函数都要有对应的测试
- 使用mock对象隔离依赖

### 8.2 集成测试
- 测试API接口功能
- 测试数据库操作
- 测试第三方服务集成

### 8.3 性能测试
- 测试接口响应时间
- 测试并发处理能力
- 测试内存使用情况

## 9. 部署规范

### 9.1 环境配置
- 开发环境: 本地开发使用
- 测试环境: 功能测试使用
- 预生产环境: 生产前验证使用
- 生产环境: 正式运行使用

### 9.2 部署流程
1. 代码审查通过
2. 自动化测试通过
3. 构建Docker镜像
4. 部署到测试环境
5. 测试验证通过
6. 部署到生产环境

### 9.3 监控告警
- 应用性能监控
- 错误日志监控
- 数据库性能监控
- 服务器资源监控

## 10. 文档规范

### 10.1 代码注释
- 公共函数必须有注释
- 复杂逻辑必须有注释
- 注释要清晰明了

### 10.2 API文档
- 使用Swagger生成API文档
- 详细描述接口参数和返回值
- 提供接口调用示例

### 10.3 技术文档
- 架构设计文档
- 数据库设计文档
- 部署运维文档

## 11. 团队协作规范

### 11.1 沟通机制
- 每日站会同步进度
- 周会讨论技术方案
- 及时沟通问题和风险

### 11.2 任务分配
- 明确任务目标和时间
- 合理分配工作量
- 及时跟踪任务进度

### 11.3 知识分享
- 定期技术分享
- 代码审查学习
- 文档知识沉淀

## 12. 质量保证

### 12.1 代码质量
- 使用静态代码分析工具
- 遵循代码规范
- 定期代码重构

### 12.2 功能质量
- 充分的功能测试
- 用户体验优化
- 性能指标达标

### 12.3 安全质量
- 安全漏洞扫描
- 权限控制验证
- 数据安全保护

---

**注意**: 本规范文档会随着项目发展持续更新，请团队成员及时关注最新版本。 