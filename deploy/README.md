# 🚀 plots 服务器部署指南

## 📋 部署前准备

### 服务器要求
- **系统**: Ubuntu 24.04 LTS Server
- **配置**: 2核CPU / 2GB内存 / 50GB硬盘
- **网络**: 开放 22(SSH), 80(HTTP), 443(HTTPS) 端口

### 域名准备
1. 购买域名
2. 添加 A 记录指向服务器 IP
3. 等待 DNS 生效（通常几分钟到几小时）

---

## 🔧 部署步骤

### 步骤 1: 上传部署脚本

将 `deploy/` 目录下的脚本上传到服务器：

```bash
# 方式1: 使用 scp
scp -r deploy/* root@139.155.244.16:/root/deploy/

# 方式2: 使用 git clone
ssh root@139.155.244.16
git clone https://github.com/yourusername/plots.git
cd plots/deploy
```

### 步骤 2: 执行初始化脚本

```bash
# 登录服务器
ssh root@139.155.244.16

# 进入脚本目录
cd ~/deploy  # 或 cd ~/plots/deploy

# 给脚本添加执行权限
chmod +x *.sh

# 执行初始化脚本
sudo ./01-init-server.sh
```

脚本将自动安装：
- Go 1.22
- Node.js 20 LTS
- Nginx
- SQLite
- Certbot (SSL)
- 配置防火墙

### 步骤 3: 重新登录

```bash
exit
ssh root@139.155.244.16
```

验证安装：
```bash
go version      # 应显示 go1.22.x
node -v         # 应显示 v20.x
nginx -v        # 应显示 nginx/1.24.x
```

### 步骤 4: 配置 Nginx

```bash
sudo ./02-setup-nginx.sh
# 输入你的域名（如 example.com）
```

### 步骤 5: 申请 SSL 证书

```bash
sudo certbot --nginx -d yourdomain.com -d www.yourdomain.com
# 按提示输入邮箱，同意条款
```

### 步骤 6: 配置 systemd 服务

```bash
sudo ./03-setup-service.sh
```

**重要**: 修改 JWT 密钥
```bash
# 生成随机密钥
openssl rand -base64 32

# 编辑配置文件
sudo vim /opt/plots/configs/config.yaml
# 将 jwt.secret 的值替换为生成的随机字符串
```

### 步骤 7: 配置 GitHub Actions 自动部署

```bash
sudo ./04-setup-deploy-user.sh
```

脚本执行完成后会显示：
1. 服务器 IP
2. 部署用户名
3. SSH 私钥

在 GitHub 仓库中配置 Secrets：
1. 进入仓库 → Settings → Secrets and variables → Actions
2. 添加以下 Secrets：

| Secret 名称 | 值 |
|-------------|-----|
| `SERVER_HOST` | 服务器 IP |
| `SERVER_USER` | `deploy` |
| `SSH_PRIVATE_KEY` | 脚本输出的私钥（包含 BEGIN 和 END 行） |

---

## ✅ 部署完成检查清单

- [ ] 系统更新完成
- [ ] Go 安装成功 (`go version`)
- [ ] Node.js 安装成功 (`node -v`)
- [ ] Nginx 运行正常 (`systemctl status nginx`)
- [ ] 防火墙已启用 (`ufw status`)
- [ ] 域名已解析到服务器
- [ ] SSL 证书已申请
- [ ] JWT 密钥已修改
- [ ] GitHub Secrets 已配置

---

## 📁 目录结构

```
/opt/plots/
├── bin/              # 可执行文件
│   └── plots-server
├── web/              # 前端静态文件
│   ├── index.html
│   └── assets/
├── data/             # 数据文件
│   └── plots.db      # SQLite 数据库
├── logs/             # 日志文件
│   ├── app.log
│   ├── app_error.log
│   ├── nginx_access.log
│   └── nginx_error.log
└── configs/          # 配置文件
    └── config.yaml
```

---

## 🔄 常用运维命令

### 服务管理
```bash
sudo systemctl start plots    # 启动服务
sudo systemctl stop plots     # 停止服务
sudo systemctl restart plots  # 重启服务
sudo systemctl status plots   # 查看状态
```

### 日志查看
```bash
tail -f /opt/plots/logs/app.log           # 实时查看应用日志
sudo journalctl -u plots -f               # 查看系统日志
tail -f /opt/plots/logs/nginx_access.log  # 查看 Nginx 访问日志
```

### SSL 证书
```bash
sudo certbot certificates     # 查看证书状态
sudo certbot renew            # 手动续期（通常自动）
sudo certbot renew --dry-run  # 测试续期
```

### 数据库备份
```bash
# 备份数据库
cp /opt/plots/data/plots.db /opt/plots/data/plots.db.backup.$(date +%Y%m%d)

# 恢复数据库
cp /opt/plots/data/plots.db.backup.20240224 /opt/plots/data/plots.db
```

---

## 🚨 故障排查

### SSH 连接问题

#### 问题1: "Permission denied" 密码被拒绝

**可能原因**: root 账户被锁定

检查账户状态：
```bash
# 通过云服务商 VNC 控制台登录后执行
passwd -S root
# 如果显示 "root L ..." (L表示Locked)，账户被锁定
```

解决方法：
```bash
# 解锁并重设密码
sudo passwd root
# 输入两次新密码

# 确保 SSH 配置允许密码登录
sudo vim /etc/ssh/sshd_config
# 确认以下配置为 yes:
# PasswordAuthentication yes
# PermitRootLogin yes

# 重启 SSH 服务
sudo systemctl restart sshd
```

#### 问题2: "REMOTE HOST IDENTIFICATION HAS CHANGED" 警告

**原因**: 服务器重装系统后 SSH 密钥变化

解决方法（在本地 Windows 执行）：
```bash
# 删除旧的密钥记录
ssh-keygen -R 服务器IP

# 例如
ssh-keygen -R 139.155.244.16

# 然后重新连接，输入 yes 信任新密钥
ssh root@服务器IP
```

#### 问题3: 只能通过密钥登录

如果服务器只允许密钥登录：
```bash
# 使用密钥文件连接
ssh -i C:\Users\你的用户名\.ssh\密钥文件.pem root@服务器IP
```

---

### 服务无法启动
```bash
sudo journalctl -u plots -n 50  # 查看详细错误
ls -la /opt/plots/bin/plots-server  # 检查权限
```

### 502 Bad Gateway
```bash
sudo systemctl status plots     # 检查后端服务是否运行
netstat -tlnp | grep 8080       # 检查端口是否监听
```

### SSL 证书问题
```bash
sudo certbot --nginx -d yourdomain.com --force-renewal  # 重新申请证书
```
