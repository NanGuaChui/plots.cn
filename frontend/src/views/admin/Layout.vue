<template>
  <div class="admin-layout">
    <aside class="admin-sidebar">
      <div class="sidebar-header">
        <h2>plots</h2>
        <span>管理后台</span>
      </div>
      <nav class="sidebar-nav">
        <router-link to="/admin" exact>📊 仪表盘</router-link>
        <router-link to="/admin/articles">📝 文章管理</router-link>
        <router-link to="/admin/portfolios">🎨 作品集</router-link>
      </nav>
      <div class="sidebar-footer">
        <button @click="handleLogout" class="logout-btn">🚪 退出登录</button>
        <router-link to="/" class="back-link">← 返回前台</router-link>
      </div>
    </aside>

    <main class="admin-main">
      <router-view />
    </main>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';

const router = useRouter();

const handleLogout = () => {
  localStorage.removeItem('token');
  router.push('/admin/login');
};
</script>

<style lang="scss" scoped>
.admin-layout {
  display: flex;
  min-height: 100vh;
}

.admin-sidebar {
  width: 250px;
  background: #1a202c;
  color: white;
  display: flex;
  flex-direction: column;

  .sidebar-header {
    padding: 1.5rem;
    border-bottom: 1px solid #2d3748;

    h2 {
      font-size: 1.5rem;
      margin-bottom: 0.25rem;
    }

    span {
      font-size: 0.875rem;
      color: #a0aec0;
    }
  }

  .sidebar-nav {
    flex: 1;
    padding: 1rem 0;

    a {
      display: block;
      padding: 0.75rem 1.5rem;
      color: #a0aec0;
      text-decoration: none;
      transition: all 0.2s;

      &:hover {
        background: #2d3748;
        color: white;
      }

      &.router-link-active {
        background: #667eea;
        color: white;
      }
    }
  }

  .sidebar-footer {
    padding: 1rem 1.5rem;
    border-top: 1px solid #2d3748;

    .logout-btn {
      width: 100%;
      padding: 0.75rem;
      background: transparent;
      border: 1px solid #4a5568;
      color: #a0aec0;
      border-radius: 6px;
      cursor: pointer;
      margin-bottom: 0.75rem;
      transition: all 0.2s;

      &:hover {
        background: #e53e3e;
        border-color: #e53e3e;
        color: white;
      }
    }

    .back-link {
      display: block;
      text-align: center;
      color: #718096;
      font-size: 0.875rem;
      text-decoration: none;

      &:hover {
        color: white;
      }
    }
  }
}

.admin-main {
  flex: 1;
  background: #f7fafc;
  padding: 2rem;
  overflow-y: auto;
}
</style>
