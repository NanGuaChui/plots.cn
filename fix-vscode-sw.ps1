# VSCode Service Worker 修复脚本

Write-Host "正在关闭 VSCode..." -ForegroundColor Yellow
Get-Process "Code" -ErrorAction SilentlyContinue | Stop-Process -Force

Write-Host "等待进程完全关闭..." -ForegroundColor Yellow
Start-Sleep -Seconds 2

$vscodeDataPaths = @(
    "$env:APPDATA\Code\Cache",
    "$env:APPDATA\Code\CachedData",
    "$env:APPDATA\Code\CachedExtensions",
    "$env:APPDATA\Code\CachedExtensionVSIXs",
    "$env:APPDATA\Code\Code Cache",
    "$env:APPDATA\Code\GPUCache",
    "$env:APPDATA\Code\Service Worker",
    "$env:APPDATA\Code\storage.json",
    "$env:APPDATA\Code\logs"
)

Write-Host "正在清理缓存..." -ForegroundColor Yellow
foreach ($path in $vscodeDataPaths) {
    if (Test-Path $path) {
        Write-Host "删除: $path" -ForegroundColor Cyan
        Remove-Item $path -Recurse -Force -ErrorAction SilentlyContinue
    }
}

Write-Host "`n清理完成！请重新启动 VSCode" -ForegroundColor Green
Write-Host "按任意键退出..." -ForegroundColor Yellow
$null = $Host.UI.RawUI.ReadKey("NoEcho,IncludeKeyDown")
