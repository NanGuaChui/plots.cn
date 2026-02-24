#!/bin/bash
# ============================================
# GitHub Actions 部署用户配置脚本
# ============================================

set -e

if [ "$EUID" -ne 0 ]; then
    echo "请使用 sudo 运行此脚本"
    exit 1
fi

GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

print_status() { echo -e "${GREEN}[✓]${NC} $1"; }
print_warning() { echo -e "${YELLOW}[!]${NC} $1"; }

echo "=========================================="
echo "  GitHub Actions 部署用户配置"
echo "=========================================="

DEPLOY_USER="deploy"

# 创建部署用户
if id "${DEPLOY_USER}" &>/dev/null; then
    print_warning "用户 ${DEPLOY_USER} 已存在"
else
    adduser --disabled-password --gecos "" ${DEPLOY_USER}
    print_status "用户 ${DEPLOY_USER} 创建完成"
fi

# 配置 sudo 权限（与 .github/workflows/deploy.yml 中的命令保持一致）
cat > /etc/sudoers.d/deploy << 'EOF'
# systemctl 服务管理
deploy ALL=(ALL) NOPASSWD: /bin/systemctl start plots
deploy ALL=(ALL) NOPASSWD: /bin/systemctl stop plots
deploy ALL=(ALL) NOPASSWD: /bin/systemctl restart plots
deploy ALL=(ALL) NOPASSWD: /bin/systemctl status plots

# 文件操作 - 移动部署文件
deploy ALL=(ALL) NOPASSWD: /bin/mv /tmp/plots-deploy/plots-server /opt/plots/bin/
deploy ALL=(ALL) NOPASSWD: /bin/mv /tmp/plots-deploy/web/* /opt/plots/web/

# 文件操作 - 清理和权限
deploy ALL=(ALL) NOPASSWD: /bin/rm -rf /opt/plots/web/*
deploy ALL=(ALL) NOPASSWD: /bin/rm -rf /tmp/plots-deploy
deploy ALL=(ALL) NOPASSWD: /bin/chmod +x /opt/plots/bin/plots-server
deploy ALL=(ALL) NOPASSWD: /bin/chown -R www-data\:www-data /opt/plots

# 备份操作
deploy ALL=(ALL) NOPASSWD: /bin/cp /opt/plots/bin/plots-server /opt/plots/bin/plots-server.bak
EOF

chmod 440 /etc/sudoers.d/deploy
print_status "sudo 权限配置完成"

# 配置 SSH
mkdir -p /home/${DEPLOY_USER}/.ssh
chmod 700 /home/${DEPLOY_USER}/.ssh
touch /home/${DEPLOY_USER}/.ssh/authorized_keys
chmod 600 /home/${DEPLOY_USER}/.ssh/authorized_keys
chown -R ${DEPLOY_USER}:${DEPLOY_USER} /home/${DEPLOY_USER}/.ssh

# 生成 SSH 密钥对
SSH_KEY_PATH="/home/${DEPLOY_USER}/.ssh/github_deploy_key"

if [ -f "${SSH_KEY_PATH}" ]; then
    print_warning "SSH 密钥已存在"
else
    ssh-keygen -t ed25519 -C "github-actions-deploy" -f ${SSH_KEY_PATH} -N ""
    chown ${DEPLOY_USER}:${DEPLOY_USER} ${SSH_KEY_PATH}*
    cat ${SSH_KEY_PATH}.pub >> /home/${DEPLOY_USER}/.ssh/authorized_keys
    print_status "SSH 密钥对生成完成"
fi

# 获取公网 IP
PUBLIC_IP=$(curl -s --connect-timeout 5 ip.sb 2>/dev/null || curl -s --connect-timeout 5 ifconfig.me 2>/dev/null || hostname -I | awk '{print $1}')

echo ""
echo "=========================================="
echo "  配置完成！"
echo "=========================================="
echo ""
echo "📋 GitHub Secrets 配置信息:"
echo ""
echo "SERVER_HOST: ${PUBLIC_IP}"
echo "SERVER_USER: deploy"
echo "SSH_PRIVATE_KEY: (见下方)"
echo ""
echo "=========================================="
echo "🔑 SSH 私钥:"
echo "=========================================="
cat ${SSH_KEY_PATH}
echo "=========================================="
