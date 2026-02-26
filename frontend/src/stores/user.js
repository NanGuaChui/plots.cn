import { defineStore } from 'pinia';
import { ref, computed } from 'vue';
import api from '@/utils/api';

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '');
  const user = ref(null);

  const isLoggedIn = computed(() => !!token.value);

  async function login(username, password) {
    const res = await api.post('/auth/login', { username, password });
    token.value = res.token;
    user.value = res.user;
    localStorage.setItem('token', res.token);
    return res;
  }

  async function register(username, password) {
    const res = await api.post('/auth/register', { username, password });
    return res;
  }

  function logout() {
    token.value = '';
    user.value = null;
    localStorage.removeItem('token');
    localStorage.removeItem('currentCharacterId');
  }

  // 初始化时尝试获取用户信息
  async function fetchUser() {
    if (!token.value) return null;
    try {
      const res = await api.get('/auth/me');
      user.value = res;
      return res;
    } catch (error) {
      logout();
      return null;
    }
  }

  return { token, user, isLoggedIn, login, register, logout, fetchUser };
});
