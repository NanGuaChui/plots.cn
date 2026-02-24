<template>
  <div class="admin-login">
    <div class="login-card">
      <h1>{{ isInitMode ? '初始化管理员' : '管理后台' }}</h1>
      <p v-if="isInitMode" class="init-hint">首次使用，请创建管理员账号</p>

      <div class="loading" v-if="checking"></div>

      <form v-else @submit.prevent="handleSubmit">
        <div class="form-group">
          <label>用户名</label>
          <input
            type="text"
            v-model="form.username"
            :placeholder="isInitMode ? '设置管理员用户名 (至少3位)' : '请输入用户名'"
            required
          />
        </div>
        <div class="form-group">
          <label>密码</label>
          <input
            type="password"
            v-model="form.password"
            :placeholder="isInitMode ? '设置密码 (至少6位)' : '请输入密码'"
            required
          />
        </div>
        <div class="form-group" v-if="isInitMode">
          <label>确认密码</label>
          <input type="password" v-model="confirmPassword" placeholder="再次输入密码" required />
        </div>
        <button type="submit" class="btn btn-primary btn-block" :disabled="loading">
          {{ loading ? '处理中...' : isInitMode ? '创建管理员' : '登录' }}
        </button>
        <p v-if="error" class="error-msg">{{ error }}</p>
      </form>

      <div class="back-link">
        <router-link to="/">← 返回首页</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { api } from '@/utils/api';

const router = useRouter();
const form = reactive({
  username: '',
  password: '',
});
const confirmPassword = ref('');
const loading = ref(false);
const checking = ref(true);
const error = ref('');
const isInitMode = ref(false);

// 检查是否需要初始化
const checkInit = async () => {
  try {
    const res = await api.get('/auth/check');
    isInitMode.value = !res.data?.data?.initialized;
  } catch (err) {
    console.error('检查初始化状态失败:', err);
  } finally {
    checking.value = false;
  }
};

const handleSubmit = async () => {
  error.value = '';

  // 初始化模式下检查密码确认
  if (isInitMode.value) {
    if (form.password !== confirmPassword.value) {
      error.value = '两次输入的密码不一致';
      return;
    }
    if (form.username.length < 3) {
      error.value = '用户名至少3位';
      return;
    }
    if (form.password.length < 6) {
      error.value = '密码至少6位';
      return;
    }
  }

  loading.value = true;

  try {
    const endpoint = isInitMode.value ? '/auth/init' : '/auth/login';
    const res = await api.post(endpoint, form);
    const { token } = res.data.data;
    localStorage.setItem('token', token);
    router.push('/admin');
  } catch (err) {
    error.value = err.response?.data?.message || (isInitMode.value ? '创建失败' : '登录失败，请检查用户名和密码');
  } finally {
    loading.value = false;
  }
};

onMounted(checkInit);
</script>

<style lang="scss" scoped>
.admin-login {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 1rem;
}

.login-card {
  background: white;
  padding: 3rem;
  border-radius: 12px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  width: 100%;
  max-width: 400px;

  h1 {
    text-align: center;
    margin-bottom: 0.5rem;
    color: #2d3748;
  }

  .init-hint {
    text-align: center;
    color: #667eea;
    margin-bottom: 1.5rem;
    font-size: 0.9rem;
  }

  .btn-block {
    width: 100%;
    margin-top: 1rem;
  }

  .error-msg {
    color: #e53e3e;
    text-align: center;
    margin-top: 1rem;
    font-size: 0.9rem;
  }

  .back-link {
    text-align: center;
    margin-top: 1.5rem;

    a {
      color: #718096;
      text-decoration: none;

      &:hover {
        color: #667eea;
      }
    }
  }

  .loading {
    display: flex;
    justify-content: center;
    padding: 2rem;

    &::after {
      content: '';
      width: 30px;
      height: 30px;
      border: 3px solid #e2e8f0;
      border-top-color: #667eea;
      border-radius: 50%;
      animation: spin 1s linear infinite;
    }
  }
}

@keyframes spin {
  to {
    transform: rotate(360deg);
  }
}
</style>
