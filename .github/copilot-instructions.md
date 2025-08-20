# Copilot Instructions for This Repo

Concise, repo-specific guidance for AI coding agents. Focus on patterns actually present here.

## Project Overview

- Tech stack: Vue 3 + TypeScript + Vite 7 + Pinia (with persisted state) + Vue Router + vue-i18n + Naive UI + Tailwind CSS + FullCalendar.
- Entry bootstrap: `src/main.ts` creates app, sets up store -> i18n -> router, then mounts `#app`.
- Routing: Hash history (`createWebHashHistory`) with a layout route wrapping child views (`src/router/routes.ts`). 404 is caught and redirected to `Home`.
- State: Pinia root in `src/stores`, app-level store `useAppStore` in `src/stores/app` holds `theme` + `language`, persisted (except `theme`). Persistence uses a debounced custom storage adapter.
- i18n: Initialized in `src/i18n/index.ts` with `legacy: false`; locale taken from store on setup; messages in `en-US.json` & `zh-CN.json`.
- UI Providers: `App.vue` wraps the app in `NConfigProvider` and a custom `NaiveProvider` that attaches Naive UI global APIs to `window`.
- Calendar feature: `src/views/calendar/Calendar.vue` demonstrates FullCalendar integration with dynamic event coloring by custom status.

## Conventions & Patterns

- Path alias: `@` => `src`. Use for all intra-project imports (already consistently used).
- Lazy-loaded views: Route components defined with dynamic imports inside route records.
- Layout pattern: Root `/` route uses `views/layout/Index.vue` which renders a persistent header + `<RouterView/>` for children.
- Store persistence: Use `persist` option inside `defineStore` plus plugin in `stores/index.ts`. To exclude properties from persistence, add them to `omit` (see `theme`).
- Debounced localStorage writes: Added via `useDebounceFn` (1s) to reduce rapid writes; consider this when expecting immediate reflection of state in storage.
- i18n locale changes: Call `useAppStore().setLanguage(...)` or `setLocale(locale)` (the former updates store + i18n). Avoid direct `i18n.global.locale.value =` outside helpers.
- Global UI services: Access Naive UI instances via `window.$message`, `window.$dialog`, etc., after `NaiveProvider` renders. Do not recreate providers per view.
- Styling: Tailwind utility classes for layout; scoped SCSS for component-specific styles; global base styles in `assets/css/main.scss` + generated Tailwind file.
- Component naming: Single-word or simple names allowed (lint rule disabling `vue/multi-word-component-names`).
- Explicit `any` allowed (rule disabled). Avoid unnecessary casting but you won't get lint failures for `any`.

## Calendar Module Details

- Status-driven styling: Color mapping via `statusColorMap`; DOM styles applied in `eventDidMount`. If expanding statuses, update: type union, `statusTextMap`, `statusColorMap`, any template badges.
- Event creation: `handleDateSelect` prompts user; for production replace `prompt` with custom modal using `window.$dialog`.
- Reactive configuration: `calendarOptions` is a `ref<CalendarOptions>`; modify `calendarOptions.value` to adjust runtime behavior.

## Adding Features

- New view: Create `src/views/<Feature>/<Feature>.vue` and add a child route under the layout in `routes.ts` with lazy import. Keep route names unique.
- New store: Create in `src/stores/<domain>/index.ts`, export a `useXStore`, add persistence config if needed (omit sensitive/volatile fields).
- New locale strings: Add keys to both `en-US.json` & `zh-CN.json`; reference in components via `useI18n().t('path.key')`.
- UI feedback: Use `window.$message.success('...')` etc., only after app mounted (e.g., inside `onMounted`).

## Build & Tooling

- Scripts: `yarn dev`, `yarn build` (runs type-check + `vite build`), `yarn preview`, `yarn lint`, `yarn format`.
- Type checking uses `vue-tsc --build` with project refs (`tsconfig.json` referencing `tsconfig.app.json`).
- ESLint flat config in `eslint.config.ts`; do not add a legacy `.eslintrc`.
- Tailwind scan paths defined in `tailwind.config.js`; include new file patterns if placing Vue files outside `src`.

## Guard & Extensibility

- Route guard placeholder in `router/permission.ts`. Implement auth / dynamic title logic there (currently only logs router instance). Use `router.beforeEach`.
- Theme switching: `setTheme` currently only sets state; extend to toggle `document.documentElement.classList` for `dark` if adding dark mode CSS.

## Global Types & Env

- Global typings can be added under `src/typings` or `env.d.ts`. Keep alias updates synchronized in `tsconfig.app.json`.

## Common Pitfalls

- Remember debounced persistence: test race conditions when reading immediately after writing.
- Avoid direct localStorage clear that might remove other stores—current `setupStore` clears only `app-store` explicitly (comment shows example of another store). Keep this selective.
- FullCalendar event styles: Using inline styles; if switching to CSS classes ensure Tailwind safelist for dynamic class names (e.g., `status-*`).

## Quick Examples

- Access translation: `const { t } = useI18n(); t('calendar.event.timed')`.
- Add route: `{ path: 'reports', name: 'Reports', component: () => import('@/views/reports/Reports.vue') }`.
- Update locale: `useAppStore().setLanguage('en-US')`.
- Use message: `window.$message.success('Saved')`.

## When Unsure

Prefer following existing file patterns; mirror folder structure; keep imports using `@` alias. Ask for clarification if a needed cross-cutting concern (auth, error handling) isn't yet implemented.

---

Feedback welcome: identify unclear areas (e.g., desired auth/permission model, theming approach) so this doc can be refined.
