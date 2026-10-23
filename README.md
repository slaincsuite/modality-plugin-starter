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

方式一：直接构建：

```bash
make package
```

方式二：源码运行宿主：

1. 将 `build/plugin/modality-plugin-starter` 复制或链接到宿主 `plugin.dir`；
2. 在宿主插件管理接口启用 `modality-plugin-starter`；
3. 访问 `/plugins/modality-plugin-starter/#/example`。

Go 单测与 UI 单测：

```bash
make verify
```

## 插件声明

`plugin/plugin.json` 声明插件名、服务、宿主能力申请与 UI 挂载点。
插件如需访问宿主，只申请必要的 `grants`；宿主会重新鉴权并默认拒绝未授权访问。

## 目录结构

- `src/modality-plugin-starter/main.go`：插件入口
- `src/modality-plugin-starter/`：插件业务实现与测试
- `plugin/`：插件清单与静态资源
- `ui/`：插件前端工程

## 打包

```bash
make package
```

输出目录：`build/plugin/modality-plugin-starter/`。
