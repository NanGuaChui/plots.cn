import { createRouter, createWebHistory } from 'vue-router';
import { useUserStore } from '@/stores/user';
import { useCharacterStore } from '@/stores/character';

// 视图组件
import Login from '@/views/Login.vue';
import Register from '@/views/Register.vue';
import CharacterSelect from '@/views/CharacterSelect.vue';
import GameLayout from '@/components/GameLayout.vue';
import Dashboard from '@/views/Dashboard.vue';

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: Login,
    meta: { requiresGuest: true },
  },
  {
    path: '/register',
    name: 'Register',
    component: Register,
    meta: { requiresGuest: true },
  },
  {
    path: '/characters',
    name: 'CharacterSelect',
    component: CharacterSelect,
    meta: { requiresAuth: true },
  },
  {
    path: '/game',
    component: GameLayout,
    meta: { requiresAuth: true, requiresCharacter: true },
    children: [
      {
        path: '',
        name: 'Dashboard',
        component: Dashboard,
      },
      // 后续添加其他路由
      // { path: 'skills/woodcutting', name: 'Woodcutting', component: () => import('@/views/skills/Woodcutting.vue') },
      // { path: 'skills/mining', name: 'Mining', component: () => import('@/views/skills/Mining.vue') },
      // { path: 'inventory', name: 'Inventory', component: () => import('@/views/Inventory.vue') },
      // { path: 'character', name: 'Character', component: () => import('@/views/Character.vue') },
    ],
  },
  {
    path: '/',
    redirect: '/login',
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/login',
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach(async (to, from, next) => {
  const userStore = useUserStore();
  const characterStore = useCharacterStore();

  // 检查登录状态
  const isLoggedIn = userStore.isLoggedIn;

  // 如果页面需要登录但用户未登录
  if (to.meta.requiresAuth && !isLoggedIn) {
    next('/login');
    return;
  }

  // 如果页面需要未登录状态但用户已登录
  if (to.meta.requiresGuest && isLoggedIn) {
    // 检查是否有选中的角色
    if (characterStore.hasCharacter) {
      next('/game');
    } else {
      next('/characters');
    }
    return;
  }

  // 如果页面需要选中角色但未选中
  if (to.meta.requiresCharacter && !characterStore.hasCharacter) {
    // 尝试恢复角色
    const restored = await characterStore.restoreCharacter();
    if (!restored) {
      next('/characters');
      return;
    }
  }

  next();
});

export default router;
