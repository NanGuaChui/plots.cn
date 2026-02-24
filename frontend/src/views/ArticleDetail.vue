<template>
  <div class="article-detail">
    <div class="loading" v-if="loading"></div>

    <template v-else-if="article">
      <div class="article-header" :style="headerStyle">
        <div class="container">
          <span class="article-category">{{ article.category || '未分类' }}</span>
          <h1>{{ article.title }}</h1>
          <div class="article-meta">
            <span>{{ formatDate(article.created_at) }}</span>
            <span>{{ article.views }} 阅读</span>
          </div>
        </div>
      </div>

      <div class="container">
        <article class="article-body">
          <div class="article-content" v-html="article.content"></div>
        </article>

        <div class="article-footer">
          <router-link to="/articles" class="btn btn-secondary">← 返回文章列表</router-link>
        </div>
      </div>
    </template>

    <div v-else class="not-found">
      <h2>文章不存在</h2>
      <router-link to="/articles" class="btn btn-primary">返回文章列表</router-link>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import { useRoute } from 'vue-router';
import { api } from '@/utils/api';

const route = useRoute();
const article = ref(null);
const loading = ref(true);

const headerStyle = computed(() => {
  if (article.value?.cover) {
    return {
      backgroundImage: `linear-gradient(rgba(0,0,0,0.6), rgba(0,0,0,0.6)), url(${article.value.cover})`,
      backgroundSize: 'cover',
      backgroundPosition: 'center',
    };
  }
  return {};
});

const formatDate = (dateStr) => {
  const date = new Date(dateStr);
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
  });
};

onMounted(async () => {
  try {
    const res = await api.get(`/articles/${route.params.id}`);
    article.value = res.data?.data;
  } catch (error) {
    console.error('加载文章失败:', error);
  } finally {
    loading.value = false;
  }
});
</script>

<style lang="scss" scoped>
.article-detail {
  min-height: 60vh;
}

.article-header {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  padding: 4rem 0;
  text-align: center;

  .article-category {
    display: inline-block;
    padding: 0.25rem 0.75rem;
    background: rgba(255, 255, 255, 0.2);
    border-radius: 4px;
    font-size: 0.875rem;
    margin-bottom: 1rem;
  }

  h1 {
    font-size: 2.5rem;
    margin-bottom: 1rem;
    max-width: 800px;
    margin-left: auto;
    margin-right: auto;
  }

  .article-meta {
    display: flex;
    justify-content: center;
    gap: 2rem;
    opacity: 0.9;
  }
}

.article-body {
  max-width: 800px;
  margin: 0 auto;
  padding: 3rem 1rem;
  background: white;
  margin-top: -2rem;
  border-radius: 12px;
  box-shadow: 0 4px 20px rgba(0, 0, 0, 0.1);
}

.article-footer {
  max-width: 800px;
  margin: 0 auto;
  padding: 2rem 1rem;
  text-align: center;
}

.not-found {
  text-align: center;
  padding: 6rem 1rem;

  h2 {
    margin-bottom: 2rem;
    color: #718096;
  }
}

@media (max-width: 768px) {
  .article-header {
    padding: 3rem 1rem;

    h1 {
      font-size: 1.75rem;
    }
  }
}
</style>
