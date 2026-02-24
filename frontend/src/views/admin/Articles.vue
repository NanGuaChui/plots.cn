<template>
  <div class="admin-articles">
    <div class="page-title">
      <h1>文章管理</h1>
      <button class="btn btn-primary" @click="showCreateModal = true">+ 新建文章</button>
    </div>

    <div class="articles-list">
      <div class="loading" v-if="loading"></div>
      <table v-else-if="articles.length" class="data-table">
        <thead>
          <tr>
            <th>标题</th>
            <th>分类</th>
            <th>状态</th>
            <th>浏览</th>
            <th>创建时间</th>
            <th>操作</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="article in articles" :key="article.id">
            <td>{{ article.title }}</td>
            <td>{{ article.category || '-' }}</td>
            <td>
              <span :class="['status', article.published ? 'published' : 'draft']">
                {{ article.published ? '已发布' : '草稿' }}
              </span>
            </td>
            <td>{{ article.views }}</td>
            <td>{{ formatDate(article.created_at) }}</td>
            <td class="actions">
              <button @click="editArticle(article)" class="btn-icon">✏️</button>
              <button @click="deleteArticle(article.id)" class="btn-icon danger">🗑️</button>
            </td>
          </tr>
        </tbody>
      </table>
      <div v-else class="empty-state">
        <p>暂无文章，点击上方按钮创建</p>
      </div>
    </div>

    <!-- 创建/编辑弹窗 -->
    <div class="modal-overlay" v-if="showCreateModal" @click.self="closeModal">
      <div class="modal">
        <h2>{{ editingArticle ? '编辑文章' : '新建文章' }}</h2>
        <form @submit.prevent="saveArticle">
          <div class="form-group">
            <label>标题</label>
            <input type="text" v-model="form.title" required />
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>分类</label>
              <input type="text" v-model="form.category" />
            </div>
            <div class="form-group">
              <label>标签 (逗号分隔)</label>
              <input type="text" v-model="form.tags" />
            </div>
          </div>
          <div class="form-group">
            <label>摘要</label>
            <textarea v-model="form.summary" rows="2"></textarea>
          </div>
          <div class="form-group">
            <label>内容</label>
            <textarea v-model="form.content" rows="10" required></textarea>
          </div>
          <div class="form-group">
            <label>封面图 URL</label>
            <input type="text" v-model="form.cover" />
          </div>
          <div class="form-group checkbox">
            <label>
              <input type="checkbox" v-model="form.published" />
              立即发布
            </label>
          </div>
          <div class="modal-actions">
            <button type="button" class="btn btn-secondary" @click="closeModal">取消</button>
            <button type="submit" class="btn btn-primary">保存</button>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { api } from '@/utils/api';

const articles = ref([]);
const loading = ref(true);
const showCreateModal = ref(false);
const editingArticle = ref(null);

const form = reactive({
  title: '',
  content: '',
  summary: '',
  cover: '',
  category: '',
  tags: '',
  published: false,
});

const formatDate = (dateStr) => {
  return new Date(dateStr).toLocaleDateString('zh-CN');
};

const fetchArticles = async () => {
  loading.value = true;
  try {
    const res = await api.get('/articles?page=1&page_size=100');
    articles.value = res.data?.data || [];
  } catch (error) {
    console.error('加载失败:', error);
  } finally {
    loading.value = false;
  }
};

const editArticle = (article) => {
  editingArticle.value = article;
  Object.assign(form, article);
  showCreateModal.value = true;
};

const closeModal = () => {
  showCreateModal.value = false;
  editingArticle.value = null;
  Object.assign(form, {
    title: '',
    content: '',
    summary: '',
    cover: '',
    category: '',
    tags: '',
    published: false,
  });
};

const saveArticle = async () => {
  try {
    if (editingArticle.value) {
      await api.put(`/articles/${editingArticle.value.id}`, form);
    } else {
      await api.post('/articles', form);
    }
    closeModal();
    fetchArticles();
  } catch (error) {
    alert('保存失败');
  }
};

const deleteArticle = async (id) => {
  if (!confirm('确定要删除这篇文章吗？')) return;
  try {
    await api.delete(`/articles/${id}`);
    fetchArticles();
  } catch (error) {
    alert('删除失败');
  }
};

onMounted(fetchArticles);
</script>

<style lang="scss" scoped>
.admin-articles {
  .page-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }
}

.data-table {
  width: 100%;
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);
  border-collapse: collapse;

  th,
  td {
    padding: 1rem;
    text-align: left;
    border-bottom: 1px solid #e2e8f0;
  }

  th {
    background: #f7fafc;
    font-weight: 600;
    color: #4a5568;
  }

  .status {
    padding: 0.25rem 0.5rem;
    border-radius: 4px;
    font-size: 0.8rem;

    &.published {
      background: #c6f6d5;
      color: #276749;
    }

    &.draft {
      background: #feebc8;
      color: #c05621;
    }
  }

  .actions {
    display: flex;
    gap: 0.5rem;
  }

  .btn-icon {
    background: none;
    border: none;
    cursor: pointer;
    font-size: 1rem;
    padding: 0.25rem;

    &.danger:hover {
      color: #e53e3e;
    }
  }
}

.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
}

.modal {
  background: white;
  padding: 2rem;
  border-radius: 12px;
  width: 100%;
  max-width: 700px;
  max-height: 90vh;
  overflow-y: auto;

  h2 {
    margin-bottom: 1.5rem;
  }

  .form-row {
    display: grid;
    grid-template-columns: 1fr 1fr;
    gap: 1rem;
  }

  .checkbox label {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    cursor: pointer;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 1.5rem;
  }
}

.empty-state {
  text-align: center;
  padding: 4rem;
  background: white;
  border-radius: 12px;
  color: #718096;
}
</style>
