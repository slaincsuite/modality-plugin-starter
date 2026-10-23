# modality-plugin-starter

modality Go 插件开发快速开始模板。

## 开发环境

所需环境：

1. Go 1.27.3
2. Node 22
3. npm 11
4. modality 宿主（可选）

克隆项目：

```bash
git clone git@github.com:slaincsuite/modality-plugin-starter.git
cd modality-plugin-starter
```

## 本地调试

打包为 `.mgpn` 单文件插件：

```bash
make package
# 输出: build/plugin/modality-plugin-starter-v1.0.0.mgpn
```

将 `.mgpn` 文件放入宿主 `plugin.dir` 目录，宿主会自动扫描 `.mgpn` 格式的插件包。
在宿主插件管理接口启用 `modality-plugin-starter` 后，
访问 `/plugins/modality-plugin-starter/#/example`。

Go 单测与 UI 单测：

```bash
make verify
```

## 插件声明

`plugin/plugin.json` 声明插件名、服务、宿主能力申请与 UI 挂载点。
插件如需访问宿主，只申请必要的 `grants`；宿主会重新鉴权并默认拒绝未授权访问。

### 宿主访问示例

`src/starter.go` 展示了宿主能力调用的标准模式：

```go
client := NewHostClient(caller) // caller 由 go-plugin broker 注入
data, err := QueryModuleRegistry(ctx, client)
```

`NewHostClient` 内部按 `plugin.json` grants 声明能力集，宿主逐项重新鉴权。
未声明的能力一律拒绝；宿主不可用时返回 `ErrHostUnavailable`。

## 目录结构

- `plugin.json`：插件清单（项目根目录）
- `src/apps/main.go`：插件入口
- `src/starter.go`：插件业务实现、宿主访问示例与测试
- `ui/`：插件前端工程（路由声明在 `ui/src/routes.json`）

## 打包

```bash
make package
```

输出单文件插件包：`build/plugin/modality-plugin-starter-v1.0.0.mgpn`。

`.mgpn` 文件是包含插件二进制、`plugin.json` 和 UI 静态资源的压缩包。
宿主只读取 `.mgpn` 后缀的文件作为插件载入，其他文件会被忽略。

## 平台说明

`make package` 默认按当前操作系统构建。
Go 只为当前操作系统构建插件二进制，且所有平台均输出无扩展名的
`modality-plugin-starter`，名称与 `plugin.json` 的 `binary` 一致。
