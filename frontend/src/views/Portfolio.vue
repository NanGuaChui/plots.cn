<template>
  <div class="portfolio-page">
    <div class="page-header">
      <h1>作品集</h1>
      <p>我的项目与作品展示</p>
    </div>

    <div class="container">
      <div class="portfolio-content">
        <div class="loading" v-if="loading"></div>

        <div class="portfolio-grid" v-else-if="portfolios.length">
          <div class="card portfolio-card" v-for="portfolio in portfolios" :key="portfolio.id">
            <div class="portfolio-image" v-if="portfolio.image_url">
              <img :src="portfolio.image_url" :alt="portfolio.title" />
            </div>
            <div class="portfolio-body">
              <h3>{{ portfolio.title }}</h3>
              <p class="description">{{ portfolio.description }}</p>
              <div class="tech-stack" v-if="portfolio.tech_stack">
                <span v-for="tech in parseTechStack(portfolio.tech_stack)" :key="tech" class="tech-tag">
                  {{ tech }}
                </span>
              </div>
              <div class="portfolio-links">
                <a v-if="portfolio.project_url" :href="portfolio.project_url" target="_blank" class="btn btn-primary">
                  🔗 查看项目
                </a>
                <a v-if="portfolio.github_url" :href="portfolio.github_url" target="_blank" class="btn btn-secondary">
                  📦 GitHub
                </a>
              </div>
            </div>
          </div>
        </div>

        <div v-else class="empty-state">
          <p>暂无作品</p>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/utils/api';

const portfolios = ref([]);
const loading = ref(true);

const parseTechStack = (techStack) => {
  try {
    return JSON.parse(techStack);
  } catch {
    return techStack.split(',').map((t) => t.trim());
  }
};

onMounted(async () => {
  try {
    const res = await api.get('/portfolios');
    portfolios.value = res.data?.data || [];
  } catch (error) {
    console.error('加载作品集失败:', error);
  } finally {
    loading.value = false;
  }
});
</script>

<style lang="scss" scoped>
.portfolio-page {
  .portfolio-content {
    padding: 3rem 0;
  }
}

.portfolio-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 2rem;

  @media (max-width: 768px) {
    grid-template-columns: 1fr;
  }
}

.portfolio-card {
  display: flex;
  flex-direction: column;

  .portfolio-image {
    height: 220px;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: transform 0.3s;
    }
  }

  &:hover .portfolio-image img {
    transform: scale(1.05);
  }

  .portfolio-body {
    padding: 1.5rem;
    flex: 1;
    display: flex;
    flex-direction: column;

    h3 {
      font-size: 1.25rem;
      margin-bottom: 0.75rem;
      color: #2d3748;
    }

    .description {
      color: #718096;
      margin-bottom: 1rem;
      flex: 1;
    }

    .tech-stack {
      display: flex;
      flex-wrap: wrap;
      gap: 0.5rem;
      margin-bottom: 1rem;

      .tech-tag {
        padding: 0.25rem 0.5rem;
        background: #eef2ff;
        color: #667eea;
        border-radius: 4px;
        font-size: 0.8rem;
      }
    }

    .portfolio-links {
      display: flex;
      gap: 0.75rem;

      .btn {
        padding: 0.5rem 1rem;
        font-size: 0.875rem;
      }
    }
  }
}

.empty-state {
  text-align: center;
  padding: 4rem;
  color: #a0aec0;
}
</style>
