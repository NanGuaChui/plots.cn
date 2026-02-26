# Plots 游戏后端

基于 Go + Gin + GORM + SQLite 的游戏后端服务。

## 技术栈

- **Web 框架**: Gin
- **ORM**: GORM
- **数据库**: SQLite
- **认证**: JWT (golang-jwt)

## 项目结构

```
backend/
├── cmd/
│   └── main.go           # 程序入口
├── internal/
│   ├── config/           # 配置
│   │   └── config.go
│   ├── handler/          # HTTP 处理器
│   │   ├── auth.go       # 认证接口
│   │   └── character.go  # 角色接口
│   ├── middleware/       # 中间件
│   │   └── auth.go       # JWT 认证中间件
│   ├── model/            # 数据模型
│   │   ├── user.go       # 用户模型
│   │   ├── character.go  # 角色模型
│   │   └── migrate.go    # 数据库迁移
│   ├── router/           # 路由配置
│   │   └── router.go
│   └── utils/            # 工具函数
│       └── jwt.go
├── go.mod
└── README.md
```

## 快速开始

### 1. 安装依赖

```bash
cd backend
go mod download
go mod tidy
```

### 2. 运行开发服务器

```bash
go run ./cmd
```

或使用 Makefile:

```bash
make dev-backend
```

服务将在 `http://localhost:8080` 启动。

### 3. 环境变量配置

| 变量名 | 默认值 | 说明 |
|--------|--------|------|
| PORT | 8080 | 服务端口 |
| DB_PATH | game.db | SQLite 数据库文件路径 |
| JWT_SECRET | your-secret-key-change-in-production | JWT 签名密钥 |

## API 接口

### 认证接口

| 方法 | 路径 | 说明 | 需要登录 |
|------|------|------|----------|
| POST | /api/auth/register | 用户注册 | 否 |
| POST | /api/auth/login | 用户登录 | 否 |
| GET | /api/auth/me | 获取当前用户信息 | 是 |

### 角色接口

| 方法 | 路径 | 说明 | 需要登录 |
|------|------|------|----------|
| GET | /api/characters | 获取当前用户所有角色 | 是 |
| POST | /api/characters | 创建新角色 | 是 |
| GET | /api/characters/:id | 获取角色详情 | 是 |
| DELETE | /api/characters/:id | 删除角色 | 是 |
| GET | /api/characters/:id/stats | 获取角色属性 | 是 |

### 示例请求

#### 注册用户

```bash
curl -X POST http://localhost:8080/api/auth/register \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "123456"}'
```

#### 用户登录

```bash
curl -X POST http://localhost:8080/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{"username": "testuser", "password": "123456"}'
```

响应将包含 JWT Token:
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
  "user": {
    "id": 1,
    "username": "testuser"
  }
}
```

#### 创建角色

```bash
curl -X POST http://localhost:8080/api/characters \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -d '{"name": "勇者", "slot_index": 0}'
```

#### 获取角色列表

```bash
curl http://localhost:8080/api/characters \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 数据模型

### 用户 (User)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| username | string | 用户名（唯一） |
| password | string | 加密后的密码 |
| created_at | time | 创建时间 |
| updated_at | time | 更新时间 |

### 角色 (Character)

| 字段 | 类型 | 说明 |
|------|------|------|
| id | uint | 主键 |
| user_id | uint | 所属用户 |
| name | string | 角色名 |
| slot_index | int | 槽位索引 (0-3) |
| type | string | 类型: main(主号) / ironman(铁人号) |
| strength | int | 力量 (默认10) |
| agility | int | 敏捷 (默认10) |
| endurance | int | 耐力 (默认10) |
| intelligence | int | 智力 (默认10) |
| luck | int | 幸运 (默认10) |
| combat_level | int | 战斗等级 (默认1) |
| combat_exp | int | 战斗经验 (默认0) |

### 衍生属性公式

- **物理攻击** = 力量 × 2 + 装备加成
- **魔法攻击** = 智力 × 2 + 装备加成
- **物理防御** = 耐力 × 1.5 + 装备加成
- **魔法防御** = 智力 × 1 + 装备加成
- **最大生命** = 100 + 耐力 × 10 + 装备加成
- **最大魔法** = 50 + 智力 × 5 + 装备加成
- **速度** = 敏捷 × 1 + 装备加成
- **闪避率** = 敏捷 × 0.05% + 装备加成
- **暴击率** = 幸运 × 0.1% + 装备加成
- **暴击伤害** = 150% + 幸运 × 0.5% + 装备加成
- **战利品掉落率** = 幸运 × 0.1% + 装备加成

## 业务规则

1. **角色限制**: 每个用户最多创建 4 个角色
2. **槽位规则**: 
   - 槽位 0: 主号 (main) - 可参与市场交易
   - 槽位 1-3: 铁人号 (ironman) - 完全自给自足，无法交易
3. **属性初始化**: 所有基础属性默认值为 10
