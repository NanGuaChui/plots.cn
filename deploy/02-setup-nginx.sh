#!/bin/bash
# ============================================
# Nginx 站点配置脚本
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
echo "  Nginx 站点配置"
echo "=========================================="

read -p "请输入你的域名 (例如: example.com): " DOMAIN

if [ -z "$DOMAIN" ]; then
    echo "域名不能为空！"
    exit 1
fi

echo "配置域名: ${DOMAIN}"

# 创建 Nginx 配置
cat > /etc/nginx/sites-available/plots << EOF
server {
    listen 80;
    listen [::]:80;
    server_name ${DOMAIN} www.${DOMAIN};
    
    root /opt/plots/web;
    index index.html;
    
    gzip on;
    gzip_vary on;
    gzip_min_length 1024;
    gzip_types text/plain text/css text/xml text/javascript application/javascript application/json application/xml;
    
    location / {
        try_files \$uri \$uri/ /index.html;
    }
    
    location /api/ {
        proxy_pass http://127.0.0.1:8080;
        proxy_http_version 1.1;
        proxy_set_header Host \$host;
        proxy_set_header X-Real-IP \$remote_addr;
        proxy_set_header X-Forwarded-For \$proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto \$scheme;
    }
    
    client_max_body_size 10M;
    
    location ~* \.(js|css|png|jpg|jpeg|gif|ico|svg|woff|woff2)$ {
        expires 30d;
        add_header Cache-Control "public, immutable";
    }
    
    access_log /opt/plots/logs/nginx_access.log;
    error_log /opt/plots/logs/nginx_error.log;
}
EOF

print_status "Nginx 配置文件创建完成"

rm -f /etc/nginx/sites-enabled/default
ln -sf /etc/nginx/sites-available/plots /etc/nginx/sites-enabled/

# 创建临时首页
cat > /opt/plots/web/index.html << 'EOF'
<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>plots - 部署成功</title>
    <style>
        body { font-family: sans-serif; background: linear-gradient(135deg, #667eea, #764ba2); min-height: 100vh; display: flex; align-items: center; justify-content: center; color: white; }
        .container { text-align: center; }
        h1 { font-size: 3em; }
    </style>
</head>
<body>
    <div class="container">
        <h1>🎉 plots</h1>
        <p>服务器配置成功！等待前端部署...</p>
    </div>
</body>
</html>
EOF

chown www-data:www-data /opt/plots/web/index.html

nginx -t && systemctl reload nginx
print_status "Nginx 配置完成"

echo ""
echo "=========================================="
echo "  Nginx 配置完成！"
echo "=========================================="
echo ""
echo "当前状态:"
echo "  - 域名: ${DOMAIN}"
echo "  - 站点目录: /opt/plots/web/"
echo "  - 日志目录: /opt/plots/logs/"
echo ""
echo "=========================================="
echo "  下一步操作"
echo "=========================================="
echo ""
echo "1. 确保域名已解析到此服务器 IP"
echo ""
echo "2. 申请 SSL 证书 (HTTPS):"
echo "   sudo certbot --nginx -d ${DOMAIN} -d www.${DOMAIN}"
echo ""
echo "3. 验证网站访问:"
echo "   curl http://${DOMAIN}"
echo "   或在浏览器访问: http://${DOMAIN}"
echo ""
echo "4. 继续运行下一个脚本:"
echo "   sudo ./03-setup-service.sh"
echo ""
