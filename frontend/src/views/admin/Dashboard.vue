<template>
  <div class="dashboard">
    <h1>仪表盘</h1>

    <div class="stats-grid">
      <div class="stat-card">
        <div class="stat-icon">📝</div>
        <div class="stat-info">
          <h3>{{ stats.articles }}</h3>
          <p>文章数量</p>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">🎨</div>
        <div class="stat-info">
          <h3>{{ stats.portfolios }}</h3>
          <p>作品数量</p>
        </div>
      </div>
      <div class="stat-card">
        <div class="stat-icon">👁️</div>
        <div class="stat-info">
          <h3>{{ stats.views }}</h3>
          <p>总浏览量</p>
        </div>
      </div>
    </div>

    <div class="quick-actions">
      <h2>快捷操作</h2>
      <div class="actions-grid">
        <router-link to="/admin/articles" class="action-card">
          <span>📝</span>
          <p>管理文章</p>
        </router-link>
        <router-link to="/admin/portfolios" class="action-card">
          <span>🎨</span>
          <p>管理作品</p>
        </router-link>
        <a href="/" target="_blank" class="action-card">
          <span>🌐</span>
          <p>访问网站</p>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';

const stats = ref({
  articles: 0,
  portfolios: 0,
  views: 0,
});

onMounted(() => {
  stats.value = {
    articles: 0,
    portfolios: 0,
    views: 0,
  };
});
</script>

<style lang="scss" scoped>
.dashboard {
  h1 {
    margin-bottom: 2rem;
    color: #2d3748;
  }

  h2 {
    margin-bottom: 1rem;
    color: #4a5568;
  }
}

.stats-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1.5rem;
  margin-bottom: 3rem;
}

.stat-card {
  background: white;
  padding: 1.5rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  display: flex;
  align-items: center;
  gap: 1rem;

  .stat-icon {
    font-size: 2.5rem;
  }

  .stat-info {
    h3 {
      font-size: 2rem;
      color: #2d3748;
    }

    p {
      color: #718096;
      font-size: 0.9rem;
    }
  }
}

.actions-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 1rem;
}

.action-card {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  text-align: center;
  text-decoration: none;
  transition: all 0.2s;

  &:hover {
    transform: translateY(-4px);
    box-shadow: 0 8px 16px rgba(0, 0, 0, 0.1);
  }

  span {
    font-size: 2.5rem;
    display: block;
    margin-bottom: 0.5rem;
  }

  p {
    color: #4a5568;
    font-weight: 500;
  }
}

@media (max-width: 768px) {
  .stats-grid,
  .actions-grid {
    grid-template-columns: 1fr;
  }
}
</style>
