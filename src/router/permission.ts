import type { Router } from 'vue-router';

export const setupPageGuard = (router: Router) => {
  console.log('🚀 ~ setupPageGuard ~ router:', router);
  // router.beforeEach(async (to, from, next) => {});
};
