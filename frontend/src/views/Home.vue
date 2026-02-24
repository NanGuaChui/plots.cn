<template>
  <div class="home">
    <!-- Hero Section -->
    <section class="hero">
      <div class="container">
        <h1>欢迎来到我的个人网站</h1>
        <p>分享技术、记录生活、展示作品</p>
        <div class="hero-buttons">
          <router-link to="/articles" class="btn btn-primary">浏览博客</router-link>
          <router-link to="/portfolio" class="btn btn-secondary">查看作品</router-link>
        </div>
      </div>
    </section>

    <!-- Latest Articles -->
    <section class="section">
      <div class="container">
        <h2 class="section-title">最新文章</h2>
        <div class="grid grid-3" v-if="articles.length">
          <article class="card article-card" v-for="article in articles" :key="article.id">
            <div class="article-cover" v-if="article.cover">
              <img :src="article.cover" :alt="article.title" />
            </div>
            <div class="article-body">
              <span class="article-category">{{ article.category || '未分类' }}</span>
              <h3>
                <router-link :to="`/articles/${article.id}`">{{ article.title }}</router-link>
              </h3>
              <p>{{ article.summary }}</p>
              <div class="article-meta">
                <span>{{ formatDate(article.created_at) }}</span>
                <span>{{ article.views }} 阅读</span>
              </div>
            </div>
          </article>
        </div>
        <div v-else class="empty-state">
          <p>暂无文章</p>
        </div>
        <div class="section-footer" v-if="articles.length">
          <router-link to="/articles" class="btn btn-secondary">查看更多文章</router-link>
        </div>
      </div>
    </section>

    <!-- Portfolio Preview -->
    <section class="section section-alt">
      <div class="container">
        <h2 class="section-title">精选作品</h2>
        <div class="grid grid-2" v-if="portfolios.length">
          <div class="card portfolio-card" v-for="portfolio in portfolios" :key="portfolio.id">
            <div class="portfolio-image" v-if="portfolio.image_url">
              <img :src="portfolio.image_url" :alt="portfolio.title" />
            </div>
            <div class="portfolio-body">
              <h3>{{ portfolio.title }}</h3>
              <p>{{ portfolio.description }}</p>
              <div class="portfolio-links">
                <a v-if="portfolio.project_url" :href="portfolio.project_url" target="_blank" class="btn btn-primary">
                  查看项目
                </a>
                <a v-if="portfolio.github_url" :href="portfolio.github_url" target="_blank" class="btn btn-secondary">
                  GitHub
                </a>
              </div>
            </div>
          </div>
        </div>
        <div v-else class="empty-state">
          <p>暂无作品</p>
        </div>
      </div>
    </section>
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue';
import { api } from '@/utils/api';

const articles = ref([]);
const portfolios = ref([]);

const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString('zh-CN');
};

onMounted(async () => {
  try {
    const [articlesRes, portfoliosRes] = await Promise.all([
      api.get('/articles?page=1&page_size=6'),
      api.get('/portfolios'),
    ]);
    articles.value = articlesRes.data?.data || [];
    portfolios.value = (portfoliosRes.data?.data || []).slice(0, 4);
  } catch (error) {
    console.error('加载数据失败:', error);
  }
});
</script>

<style lang="scss" scoped>
.hero {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  text-align: center;
  padding: 6rem 1rem;

  h1 {
    font-size: 3rem;
    margin-bottom: 1rem;
  }

  p {
    font-size: 1.25rem;
    opacity: 0.9;
    margin-bottom: 2rem;
  }

  .hero-buttons {
    display: flex;
    gap: 1rem;
    justify-content: center;

    .btn-secondary {
      background: rgba(255, 255, 255, 0.2);
      color: white;
      border: 1px solid rgba(255, 255, 255, 0.3);

      &:hover {
        background: rgba(255, 255, 255, 0.3);
      }
    }
  }
}

.section {
  padding: 4rem 0;

  &-alt {
    background: #f7fafc;
  }

  &-title {
    text-align: center;
    font-size: 2rem;
    margin-bottom: 3rem;
  }

  &-footer {
    text-align: center;
    margin-top: 3rem;
  }
}

.article-card {
  .article-cover {
    height: 200px;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }

  .article-body {
    padding: 1.5rem;
  }

  .article-category {
    display: inline-block;
    padding: 0.25rem 0.75rem;
    background: #eef2ff;
    color: #667eea;
    border-radius: 4px;
    font-size: 0.875rem;
    margin-bottom: 0.75rem;
  }

  h3 {
    margin-bottom: 0.75rem;

    a {
      color: #2d3748;
      text-decoration: none;

      &:hover {
        color: #667eea;
      }
    }
  }

  p {
    color: #718096;
    font-size: 0.95rem;
    margin-bottom: 1rem;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .article-meta {
    display: flex;
    gap: 1rem;
    font-size: 0.875rem;
    color: #a0aec0;
  }
}

.portfolio-card {
  display: flex;
  flex-direction: column;

  .portfolio-image {
    height: 200px;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }

  .portfolio-body {
    padding: 1.5rem;
    flex: 1;
    display: flex;
    flex-direction: column;

    h3 {
      margin-bottom: 0.75rem;
    }

    p {
      color: #718096;
      flex: 1;
      margin-bottom: 1rem;
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

.empty-state {
  text-align: center;
  padding: 3rem;
  color: #a0aec0;
}

@media (max-width: 768px) {
  .hero {
    padding: 4rem 1rem;

    h1 {
      font-size: 2rem;
    }
  }

  .hero-buttons {
    flex-direction: column;
  }
}
</style>
