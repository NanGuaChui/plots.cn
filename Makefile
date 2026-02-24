# plots Makefile
# 用于本地开发和构建

.PHONY: all build backend frontend clean dev dev-backend dev-frontend install

# 默认目标
all: build

# 安装依赖
install:
	@echo ">>> 安装后端依赖..."
	cd backend && go mod download
	@echo ">>> 安装前端依赖..."
	cd frontend && npm install
	@echo ">>> 依赖安装完成"

# 构建所有
build: backend frontend
	@echo ">>> 构建完成"

# 构建后端
backend:
	@echo ">>> 构建后端..."
	cd backend && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o ../dist/plots-server ./cmd/server
	@echo ">>> 后端构建完成: dist/plots-server"

# 构建后端 (Windows)
backend-win:
	@echo ">>> 构建后端 (Windows)..."
	cd backend && go build -o ../dist/plots-server.exe ./cmd/server
	@echo ">>> 后端构建完成: dist/plots-server.exe"

# 构建前端
frontend:
	@echo ">>> 构建前端..."
	cd frontend && npm run build
	@echo ">>> 前端构建完成: frontend/dist/"

# 开发模式 - 同时运行前后端
dev:
	@echo "请在两个终端分别运行:"
	@echo "  make dev-backend"
	@echo "  make dev-frontend"

# 开发模式 - 后端
dev-backend:
	@echo ">>> 启动后端开发服务器..."
	cd backend && go run ./cmd/server

# 开发模式 - 前端
dev-frontend:
	@echo ">>> 启动前端开发服务器..."
	cd frontend && npm run dev

# 清理构建产物
clean:
	@echo ">>> 清理构建产物..."
	rm -rf dist/
	rm -rf frontend/dist/
	@echo ">>> 清理完成"

# 运行测试
test:
	@echo ">>> 运行后端测试..."
	cd backend && go test ./...

# 代码格式化
fmt:
	@echo ">>> 格式化后端代码..."
	cd backend && go fmt ./...
	@echo ">>> 格式化前端代码..."
	cd frontend && npm run format 2>/dev/null || true

# 帮助
help:
	@echo "可用命令:"
	@echo "  make install      - 安装依赖"
	@echo "  make build        - 构建前后端 (Linux)"
	@echo "  make backend      - 仅构建后端 (Linux)"
	@echo "  make backend-win  - 仅构建后端 (Windows)"
	@echo "  make frontend     - 仅构建前端"
	@echo "  make dev-backend  - 启动后端开发服务器"
	@echo "  make dev-frontend - 启动前端开发服务器"
	@echo "  make clean        - 清理构建产物"
	@echo "  make test         - 运行测试"
	@echo "  make fmt          - 代码格式化"
	@echo "  make help         - 显示帮助"
