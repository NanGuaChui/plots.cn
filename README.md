# plots - 个人网站

一个使用 **Go + Vue 3** 构建的轻量级个人网站，适合低配服务器部署。

## ✨ 特性

- 🚀 **轻量高效** - Go 后端内存占用 ~30MB，适合 2GB 内存服务器
- 📝 **博客系统** - 支持文章发布、分类、标签
- 🎨 **作品展示** - 展示项目作品集
- 🔐 **管理后台** - 简洁的后台管理界面
- 🔄 **自动部署** - GitHub Actions 自动构建部署
- 📱 **响应式设计** - 支持移动端访问

## 🛠️ 技术栈

### 后端
- **Go 1.22** + Gin 框架
- **SQLite** 数据库
- **JWT** 认证

### 前端
- **Vue 3** + Vite
- **Vue Router** 路由
- **Pinia** 状态管理
- **SCSS** 样式

### 部署
- **Nginx** 反向代理
- **systemd** 服务管理
- **GitHub Actions** CI/CD
- **Let's Encrypt** SSL 证书

## 📁 项目结构

```
plots/
├── backend/                 # Go 后端
│   ├── cmd/server/         # 程序入口
│   └── internal/           # 内部模块
│       ├── config/         # 配置管理
│       ├── handler/        # HTTP 处理器
│       ├── model/          # 数据模型
│       └── repository/     # 数据访问
├── frontend/               # Vue 前端
│   ├── src/
│   │   ├── views/         # 页面组件
│   │   ├── router/        # 路由配置
│   │   ├── utils/         # 工具函数
│   │   └── styles/        # 样式文件
│   └── public/
├── deploy/                 # 部署脚本
│   ├── 01-init-server.sh  # 服务器初始化
│   ├── 02-setup-nginx.sh  # Nginx 配置
│   ├── 03-setup-service.sh # systemd 配置
│   └── 04-setup-deploy-user.sh # 部署用户
└── Makefile               # 构建命令
```

## 🚀 快速开始

### 本地开发

```bash
# 克隆项目
git clone https://github.com/yourusername/plots.git
cd plots

# 安装依赖
make install

# 启动后端开发服务器 (终端1)
make dev-backend

# 启动前端开发服务器 (终端2)
make dev-frontend
```

访问 http://localhost:3000

### 构建

```bash
# 构建前后端
make build

# 仅构建后端
make backend

# 仅构建前端
make frontend
```

## 📦 服务器部署

### 服务器要求

- **系统**: Ubuntu 24.04 LTS
- **配置**: 2核 CPU / 2GB 内存 / 50GB 硬盘
- **端口**: 22, 80, 443

### 部署步骤

详细说明请查看 [deploy/README.md](deploy/README.md)

## ⚙️ 配置

配置文件位于 `/opt/plots/configs/config.yaml`:

```yaml
server:
  port: 8080
  mode: release

database:
  path: /opt/plots/data/plots.db

jwt:
  secret: "your-secret-key"  # 请修改！
  expire: 72h
```

## 🔑 API 接口

### 公开接口

| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/articles | 获取文章列表 |
| GET | /api/articles/:id | 获取文章详情 |
| GET | /api/portfolios | 获取作品列表 |
| POST | /api/auth/login | 登录 |

### 认证接口

| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/articles | 创建文章 |
| PUT | /api/articles/:id | 更新文章 |
| DELETE | /api/articles/:id | 删除文章 |
| POST | /api/portfolios | 创建作品 |
| PUT | /api/portfolios/:id | 更新作品 |
| DELETE | /api/portfolios/:id | 删除作品 |

## 📝 常用命令

```bash
# 服务管理
sudo systemctl start plots    # 启动
sudo systemctl stop plots     # 停止
sudo systemctl restart plots  # 重启
sudo systemctl status plots   # 状态

# 查看日志
tail -f /opt/plots/logs/app.log
sudo journalctl -u plots -f

# 数据库备份
cp /opt/plots/data/plots.db ~/plots.db.backup

# 触发CI
# 创建空提交
git commit --allow-empty -m "触发 CI 构建"
```

## 📄 License

MIT License

## 🤝 贡献

欢迎提交 Issue 和 Pull Request！
