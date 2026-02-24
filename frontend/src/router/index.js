import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
  },
  {
    path: '/articles',
    name: 'Articles',
    component: () => import('@/views/Articles.vue'),
  },
  {
    path: '/articles/:id',
    name: 'ArticleDetail',
    component: () => import('@/views/ArticleDetail.vue'),
  },
  {
    path: '/portfolio',
    name: 'Portfolio',
    component: () => import('@/views/Portfolio.vue'),
  },
  {
    path: '/about',
    name: 'About',
    component: () => import('@/views/About.vue'),
  },
  {
    path: '/admin/login',
    name: 'AdminLogin',
    component: () => import('@/views/admin/Login.vue'),
  },
  {
    path: '/admin',
    name: 'Admin',
    component: () => import('@/views/admin/Layout.vue'),
    meta: { requiresAuth: true },
    children: [
      {
        path: '',
        name: 'AdminDashboard',
        component: () => import('@/views/admin/Dashboard.vue'),
      },
      {
        path: 'articles',
        name: 'AdminArticles',
        component: () => import('@/views/admin/Articles.vue'),
      },
      {
        path: 'portfolios',
        name: 'AdminPortfolios',
        component: () => import('@/views/admin/Portfolios.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    name: 'NotFound',
    component: () => import('@/views/NotFound.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition;
    } else {
      return { top: 0 };
    }
  },
});

// 路由守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token');

  if (to.meta.requiresAuth && !token) {
    next({ name: 'AdminLogin' });
  } else {
    next();
  }
});

export default router;
