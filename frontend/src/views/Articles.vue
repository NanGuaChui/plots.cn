<template>
  <div class="articles-page">
    <div class="page-header">
      <h1>博客文章</h1>
      <p>分享技术心得与生活感悟</p>
    </div>

    <div class="container">
      <div class="articles-content">
        <div class="loading" v-if="loading"></div>

        <div class="grid grid-3" v-else-if="articles.length">
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

        <!-- 分页 -->
        <div class="pagination" v-if="totalPages > 1">
          <button class="btn btn-secondary" :disabled="page === 1" @click="changePage(page - 1)">上一页</button>
          <span class="page-info">{{ page }} / {{ totalPages }}</span>
          <button class="btn btn-secondary" :disabled="page === totalPages" @click="changePage(page + 1)">
            下一页
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, watch } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import { api } from '@/utils/api';

const route = useRoute();
const router = useRouter();

const articles = ref([]);
const loading = ref(true);
const page = ref(1);
const pageSize = 9;
const totalPages = ref(1);

const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString('zh-CN');
};

const fetchArticles = async () => {
  loading.value = true;
  try {
    const res = await api.get(`/articles?page=${page.value}&page_size=${pageSize}`);
    articles.value = res.data?.data || [];
    totalPages.value = res.data?.total_pages || 1;
  } catch (error) {
    console.error('加载文章失败:', error);
  } finally {
    loading.value = false;
  }
};

const changePage = (newPage) => {
  page.value = newPage;
  router.push({ query: { page: newPage } });
  window.scrollTo({ top: 0, behavior: 'smooth' });
};

onMounted(() => {
  page.value = parseInt(route.query.page) || 1;
  fetchArticles();
});

watch(
  () => route.query.page,
  (newPage) => {
    page.value = parseInt(newPage) || 1;
    fetchArticles();
  },
);
</script>

<style lang="scss" scoped>
.articles-page {
  .articles-content {
    padding: 3rem 0;
  }
}

.article-card {
  .article-cover {
    height: 180px;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
      transition: transform 0.3s;
    }
  }

  &:hover .article-cover img {
    transform: scale(1.05);
  }

  .article-body {
    padding: 1.25rem;
  }

  .article-category {
    display: inline-block;
    padding: 0.2rem 0.6rem;
    background: #eef2ff;
    color: #667eea;
    border-radius: 4px;
    font-size: 0.8rem;
    margin-bottom: 0.5rem;
  }

  h3 {
    margin-bottom: 0.5rem;
    font-size: 1.1rem;

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
    font-size: 0.9rem;
    margin-bottom: 0.75rem;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    overflow: hidden;
  }

  .article-meta {
    display: flex;
    gap: 1rem;
    font-size: 0.8rem;
    color: #a0aec0;
  }
}

.pagination {
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 1rem;
  margin-top: 3rem;

  .page-info {
    color: #718096;
  }

  button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
  }
}

.empty-state {
  text-align: center;
  padding: 4rem;
  color: #a0aec0;
}
</style>
