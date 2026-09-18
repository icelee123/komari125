# Komari 1.2.5-fix2 安全裁剪版

本项目基于 [komari-monitor/komari](https://github.com/komari-monitor/komari) 的 `1.2.5-fix2` 版本裁剪，目标是在保留 Komari 监控面板主要维护能力的同时，移除服务端被攻破后可能导致 Agent 主机远程执行命令的功能。

## 修改内容

- 移除服务端远程命令执行、Web 终端、任务执行与任务结果接口。
- 移除 Agent 自动升级和被控相关逻辑。
- 保留 Agent 与服务端的 WebSocket 通信，Token 仍放在 URL query，兼容现有 Nginx WebSocket 反代配置。
- v2 事件白名单只允许 `agent.ping`，远程执行类事件会被拒绝。
- 服务端与 Agent 双重限制 Ping 任务频率，最小间隔为 30 秒。
- 保留机器管理、离线通知、负载/流量通知、延迟检查和数据记录等原版维护功能。
- 内置 `default` 与 `/admin` 均使用 [icelee123/komari-web125](https://github.com/icelee123/komari-web125) 的修改版前端，并已移除远控相关页面和入口。
- 不再内置 `komari-next-main` 主题，默认主题为 Komari 官方前端裁剪版，主题 `short` 为 `default`。

## 构建

仓库已包含构建后的前端资源，可直接构建服务端：

```bash
go build -trimpath -o komari-server .
```

如需重新构建前端：

```bash
cd ../komari-web125
npm ci
npm run build
cp -a dist ../komari125/web/public/defaultTheme/dist
cp komari-theme.json ../komari125/web/public/defaultTheme/
cp -a dist ../komari125/web/public/adminTheme/dist
```

GitHub Actions 会从 `icelee123/komari-web125` 拉取修改版前端，并同时更新 `defaultTheme` 与 `adminTheme`。

## 安装

```bash
sudo bash install-komari.sh
```

或：

```bash
curl -fsSL https://raw.githubusercontent.com/icelee123/komari125/main/install-komari.sh | sudo bash
```

默认监听端口为 `25774`，管理端地址为 `/admin`。

## 测试

```bash
go test ./web/public ./web/api/admin ./web/agent
go test ./...
```

`utils/geoip` 依赖外部 MaxMind 数据下载，在网络受限环境中可能超时。

## 上游与许可

原项目版权归 Komari 作者所有，本项目保留原许可证并基于原版修改。
