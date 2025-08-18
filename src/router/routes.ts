import type { RouteRecordRaw } from 'vue-router';

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Layout',
    component: () => import('@/views/layout/Index.vue'),
    children: [
      {
        path: '',
        name: 'Home',
        component: () => import('@/views/HomeView.vue'),
      },
      {
        path: 'calendar',
        name: 'calendar',
        component: () => import('@/views/calendar/Calendar.vue'),
      },
    ],
  },
  // { path: '/404', name: '404', component: () => import('@/views/NotFound.vue'), meta: { title: '404' } },
  { path: '/:catchAll(.*)*', name: 'CatchAll', redirect: { name: 'Home' } },
];

export default routes;
