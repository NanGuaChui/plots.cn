#!/bin/bash
# ============================================
# Ubuntu 24.04 服务器初始化脚本
# 用途: Go + SQLite + Nginx 个人网站部署
# ============================================

set -e

if [ "$EUID" -ne 0 ]; then
    echo "请使用 sudo 运行此脚本"
    exit 1
fi

echo "=========================================="
echo "  Ubuntu 24.04 服务器初始化脚本"
echo "=========================================="

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_status() { echo -e "${GREEN}[✓]${NC} $1"; }
print_warning() { echo -e "${YELLOW}[!]${NC} $1"; }

# 步骤 1: 系统更新
echo ""
echo ">>> 步骤 1/8: 系统更新..."
apt update && apt upgrade -y
print_status "系统更新完成"

# 步骤 2: 安装基础工具
echo ""
echo ">>> 步骤 2/8: 安装基础工具..."
apt install -y curl wget git vim htop unzip build-essential sqlite3 tree net-tools
print_status "基础工具安装完成"

# 步骤 3: 安装 Nginx
echo ""
echo ">>> 步骤 3/8: 安装 Nginx..."
apt install -y nginx
systemctl enable nginx
systemctl start nginx
print_status "Nginx 安装完成"

# 步骤 4: 安装 Go
echo ""
echo ">>> 步骤 4/8: 安装 Go 1.22..."
GO_VERSION="1.22.0"

if [ -d "/usr/local/go" ]; then
    print_warning "Go 已安装，跳过..."
else
    wget -q --show-progress https://golang.google.cn/dl/go${GO_VERSION}.linux-amd64.tar.gz -O /tmp/go.tar.gz || \
    wget -q --show-progress https://mirrors.aliyun.com/golang/go${GO_VERSION}.linux-amd64.tar.gz -O /tmp/go.tar.gz
    
    tar -C /usr/local -xzf /tmp/go.tar.gz
    rm /tmp/go.tar.gz
    
    cat > /etc/profile.d/go.sh << 'EOF'
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
export GOPROXY=https://goproxy.cn,direct
EOF
    print_status "Go ${GO_VERSION} 安装完成"
fi

source /etc/profile.d/go.sh

# 步骤 5: 安装 Node.js
echo ""
echo ">>> 步骤 5/8: 安装 Node.js 20 LTS..."

if command -v node &> /dev/null; then
    print_warning "Node.js 已安装，跳过..."
else
    curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
    apt install -y nodejs
    npm config set registry https://registry.npmmirror.com
    print_status "Node.js $(node -v) 安装完成"
fi

# 步骤 6: 安装 Certbot
echo ""
echo ">>> 步骤 6/8: 安装 Certbot..."
apt install -y certbot python3-certbot-nginx
print_status "Certbot 安装完成"

# 步骤 7: 配置防火墙
echo ""
echo ">>> 步骤 7/8: 配置防火墙..."
ufw default deny incoming
ufw default allow outgoing
ufw allow 22/tcp
ufw allow 80/tcp
ufw allow 443/tcp
ufw --force enable
print_status "防火墙配置完成"

# 步骤 8: 创建应用目录
echo ""
echo ">>> 步骤 8/8: 创建应用目录..."
mkdir -p /opt/plots/{bin,web,data,logs,configs}
chown -R www-data:www-data /opt/plots
chmod -R 755 /opt/plots
print_status "应用目录创建完成"

echo ""
echo "=========================================="
echo "  初始化完成！"
echo "=========================================="
echo ""
echo "下一步: 重新登录后执行 sudo ./02-setup-nginx.sh"
