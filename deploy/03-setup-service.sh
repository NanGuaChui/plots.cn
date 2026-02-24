#!/bin/bash
# ============================================
# systemd 服务配置脚本
# ============================================

set -e

if [ "$EUID" -ne 0 ]; then
    echo "请使用 sudo 运行此脚本"
    exit 1
fi

GREEN='\033[0;32m'
NC='\033[0m'

print_status() { echo -e "${GREEN}[✓]${NC} $1"; }

echo "=========================================="
echo "  systemd 服务配置"
echo "=========================================="

# 创建 systemd 服务文件
cat > /etc/systemd/system/plots.service << 'EOF'
[Unit]
Description=plots Personal Website
After=network.target

[Service]
Type=simple
User=www-data
Group=www-data
WorkingDirectory=/opt/plots
ExecStart=/opt/plots/bin/plots-server
Restart=always
RestartSec=5

StandardOutput=append:/opt/plots/logs/app.log
StandardError=append:/opt/plots/logs/app_error.log

NoNewPrivileges=true
PrivateTmp=true
ProtectSystem=strict
ReadWritePaths=/opt/plots

Environment=GIN_MODE=release
Environment=PORT=8080
Environment=DB_PATH=/opt/plots/data/plots.db
Environment=CONFIG_PATH=/opt/plots/configs/config.yaml

MemoryMax=256M
CPUQuota=80%

[Install]
WantedBy=multi-user.target
EOF

print_status "systemd 服务文件创建完成"

# 创建日志轮转配置
cat > /etc/logrotate.d/plots << 'EOF'
/opt/plots/logs/*.log {
    daily
    missingok
    rotate 14
    compress
    delaycompress
    notifempty
    create 0640 www-data www-data
}
EOF

print_status "日志轮转配置完成"

# 创建默认配置文件
cat > /opt/plots/configs/config.yaml << 'EOF'
server:
  port: 8080
  mode: release

database:
  path: /opt/plots/data/plots.db

jwt:
  secret: "请修改为随机字符串"
  expire: 72h

upload:
  max_size: 10485760
  allowed_types:
    - image/jpeg
    - image/png
    - image/gif
    - image/webp

log:
  level: info
  path: /opt/plots/logs/app.log
EOF

chown www-data:www-data /opt/plots/configs/config.yaml
chmod 600 /opt/plots/configs/config.yaml

systemctl daemon-reload
systemctl enable plots

print_status "systemd 配置完成"

echo ""
echo "⚠️  重要: 请修改 /opt/plots/configs/config.yaml 中的 jwt.secret"
echo "   生成密钥: openssl rand -base64 32"
echo ""
echo "下一步: sudo ./04-setup-deploy-user.sh"
