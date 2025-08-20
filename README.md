<div align="center">

# plots.cn

基于 Vue 3 + TypeScript + Vite 的前端模板工程，集成 Pinia（含持久化与防抖写入）、vue-router、vue-i18n、Naive UI、Tailwind CSS 与 FullCalendar，提供多语言与日历示例，适合作为中后台或工具类应用的起点。

</div>

## ✨ 特性概览
- 现代技术栈：Vite 7 + Vue 3 `<script setup>` + TypeScript。
- 状态管理：Pinia + 自定义防抖本地持久化（减少频繁写入 localStorage）。
- 多语言：vue-i18n（简体中文 / English），支持运行时切换。
- UI 框架：Naive UI 按需全局服务注入（message / dialog / notification / loading bar）。
- 样式方案：Tailwind CSS + SCSS 组合；统一 `@` 路径别名。
- 日历模块：FullCalendar 集成 + 自定义事件状态/颜色/右键处理示例。
- 扁平 ESLint + Prettier 配置，允许单词组件名与 `any`。

## 🧱 目录结构（节选）
```
src/
	main.ts                # 启动：store -> i18n -> router -> mount
	App.vue                # 全局 NConfigProvider + NaiveProvider + RouterView
	router/                # 路由定义与权限钩子 (hash 模式)
	stores/                # Pinia 根与 app-store（theme / language）
	i18n/                  # 多语言初始化与消息文件
	hooks/                 # 复用逻辑 (布局 / 语言)
	components/            # 全局组件 (SvgIcon / NaiveProvider)
	views/                 # 页面 (layout / home / calendar)
	assets/css/            # 全局样式与 Tailwind 入口
```

## ⚙️ 启动与脚本
```bash
yarn            # 安装依赖
yarn dev        # 开发环境（含 HMR）
yarn build      # 类型检查 + 生产构建
yarn preview    # 本地预览构建产物
yarn lint       # ESLint 自动修复
yarn format     # Prettier 格式化 src 下代码
```

> 类型检查通过 `vue-tsc --build` 执行（见 `package.json` 的 `build` 脚本）。

## 🧠 架构要点
| 模块      | 说明                                                                                  |
| --------- | ------------------------------------------------------------------------------------- |
| 启动流程  | `main.ts`：创建 `app` -> `setupStore` -> `setupI18n` -> `setupRouter` -> `mount`      |
| 路由      | `createWebHashHistory`；根布局 `views/layout/Index.vue` + 子路由；未匹配重定向 Home   |
| 状态      | `stores/index.ts` 注册 Pinia + 持久化插件；`debounceStorage` 1s 防抖写入 localStorage |
| app-store | 键名 `app-store`；字段：`theme`（不持久化）、`language`（持久化）                     |
| 多语言    | `i18n/index.ts` 设置 `legacy: false`；启动时根据 store 设定 locale                    |
| UI 服务   | `components/NaiveProvider.vue` 将 Naive UI service 注入 `window.$message` 等          |
| 样式      | Tailwind 实用类 + 自定义 SCSS；全局入口在 `main.ts` 引入                              |
| 日历      | `views/calendar/Calendar.vue`：自定义事件状态枚举 + 颜色映射 + 右键示例 + 动态 locale |

## 🗂 状态持久化说明
自定义 storage：
```ts
setItem: useDebounceFn((key, value) => localStorage.setItem(key, value), 1000)
```
含义：连续频繁修改（如快速切换语言）只会在最后一次 1 秒后写入，读取时需考虑可能的延迟。排除字段通过 store 内 `persist.omit` 配置。

## 🌍 i18n 使用
```ts
const { t } = useI18n();
t('calendar.event.timed');
```
切换语言：
```ts
useAppStore().setLanguage('en-US'); // 自动更新 i18n locale
```
新增词条：在 `en-US.json` 与 `zh-CN.json` 同步添加；保持 key 结构一致。

## 🗓 Calendar 模块速览
- 事件状态类型：`initial | ongoing | completed | settled`。
- 颜色与文案映射：`statusColorMap` / `statusTextMap`；渲染后在 `eventDidMount` 内直接写入 `style`。
- 添加事件：日期选择回调 `handleDateSelect` (`prompt`)；可改为 `window.$dialog` 自定义表单。
- 右键扩展：`dayCellDidMount` 添加 `contextmenu` 监听，可接入菜单组件。
- 若改为 CSS 类，需在 `tailwind.config.js` safelist 动态类名，或保留内联样式。

## 🔌 扩展指南
### 新增页面
1. 创建 `src/views/feature/Feature.vue`
2. 在 `router/routes.ts` 的 `children` 中添加：
```ts
{ path: 'feature', name: 'Feature', component: () => import('@/views/feature/Feature.vue') }
```

### 新增 Store
```ts
export const useXStore = defineStore('x-store', {
	persist: {},
	state: () => ({}),
	actions: {},
});
```
放置于 `src/stores/x/index.ts`；如需排除字段持久化添加 `persist.omit`。

### 新增多语言词条
1. `en-US.json` / `zh-CN.json` 各加 key。
2. 组件使用 `t('group.key')`。

## 🎨 主题（当前状态）
`app-store.theme` 仅存储值（`light | dark | auto`），尚未接入 DOM class 切换。扩展建议：
```ts
watch(() => appStore.theme, v => {
	document.documentElement.classList.toggle('dark', v === 'dark');
});
```
配合 Tailwind 的 `darkMode: 'class'` 可实现暗色主题（需在 `tailwind.config.js` 修改）。

## 🛠 代码风格与约定
- 导入统一使用 `@` 别名指向 `src`。
- 允许单词组件名（`vue/multi-word-component-names` 已关闭）。
- 允许 `any`（减少初期约束）。
- 避免直接改写 `i18n.global.locale.value`，使用封装方法或 store action。
- 修改持久化行为时，谨慎移除当前的防抖逻辑，避免性能退化。

## 🚧 常见注意点
| 场景                                | 提示                                           |
| ----------------------------------- | ---------------------------------------------- |
| 快速修改语言后立即读取 localStorage | 可能尚未写入（1s 防抖）                        |
| 新增动态状态类名                    | 使用内联 style 或在 Tailwind 配置 safelist     |
| 清理本地存储                        | 目前仅在启动时手动移除 `app-store`，保持选择性 |

## 📌 Roadmap (可选)
- [ ] 主题暗色模式实际样式
- [ ] 路由权限 / 动态标题
- [ ] 日历事件 CRUD 后端对接
- [ ] 国际化词条补充（更完整业务域）

## 🧪 调试建议
- 使用 Vue Devtools（`vite-plugin-vue-devtools` 已安装）。
- 全局 Naive UI 服务可在控制台测试：`window.$message.success('Hi')`。

## 🤝 贡献
欢迎提交 Issue / PR。统一使用 yarn，确保 `yarn lint` 与构建通过。

## 📄 许可证
仓库未声明 License，如需开源请补充 `LICENSE` 文件（例如 MIT）。

---
若需补充部署指引 / CI / 主题策略，请提出需求再迭代 README。
