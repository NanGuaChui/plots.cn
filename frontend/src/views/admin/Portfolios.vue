<template>
  <div class="admin-portfolios">
    <div class="page-title">
      <h1>作品集管理</h1>
      <button class="btn btn-primary" @click="showCreateModal = true">+ 新建作品</button>
    </div>

    <div class="portfolios-grid">
      <div class="loading" v-if="loading"></div>
      <template v-else-if="portfolios.length">
        <div class="portfolio-card" v-for="portfolio in portfolios" :key="portfolio.id">
          <div class="portfolio-image" v-if="portfolio.image_url">
            <img :src="portfolio.image_url" :alt="portfolio.title" />
          </div>
          <div class="portfolio-body">
            <h3>{{ portfolio.title }}</h3>
            <p>{{ portfolio.description }}</p>
            <div class="portfolio-actions">
              <button @click="editPortfolio(portfolio)" class="btn btn-secondary btn-sm">编辑</button>
              <button @click="deletePortfolio(portfolio.id)" class="btn btn-danger btn-sm">删除</button>
            </div>
          </div>
        </div>
      </template>
      <div v-else class="empty-state">
        <p>暂无作品，点击上方按钮创建</p>
      </div>
    </div>

    <!-- 创建/编辑弹窗 -->
    <div class="modal-overlay" v-if="showCreateModal" @click.self="closeModal">
      <div class="modal">
        <h2>{{ editingPortfolio ? '编辑作品' : '新建作品' }}</h2>
        <form @submit.prevent="savePortfolio">
          <div class="form-group">
            <label>标题</label>
            <input type="text" v-model="form.title" required />
          </div>
          <div class="form-group">
            <label>简介</label>
            <textarea v-model="form.description" rows="3"></textarea>
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>项目链接</label>
              <input type="text" v-model="form.project_url" />
            </div>
            <div class="form-group">
              <label>GitHub 链接</label>
              <input type="text" v-model="form.github_url" />
            </div>
          </div>
          <div class="form-group">
            <label>封面图 URL</label>
            <input type="text" v-model="form.image_url" />
          </div>
          <div class="form-group">
            <label>技术栈 (逗号分隔)</label>
            <input type="text" v-model="form.tech_stack" />
          </div>
          <div class="form-row">
            <div class="form-group">
              <label>排序</label>
              <input type="number" v-model.number="form.sort_order" />
            </div>
            <div class="form-group checkbox">
              <label>
                <input type="checkbox" v-model="form.published" />
                发布
              </label>
            </div>
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

const portfolios = ref([]);
const loading = ref(true);
const showCreateModal = ref(false);
const editingPortfolio = ref(null);

const form = reactive({
  title: '',
  description: '',
  image_url: '',
  project_url: '',
  github_url: '',
  tech_stack: '',
  sort_order: 0,
  published: true,
});

const fetchPortfolios = async () => {
  loading.value = true;
  try {
    const res = await api.get('/portfolios');
    portfolios.value = res.data?.data || [];
  } catch (error) {
    console.error('加载失败:', error);
  } finally {
    loading.value = false;
  }
};

const editPortfolio = (portfolio) => {
  editingPortfolio.value = portfolio;
  Object.assign(form, portfolio);
  showCreateModal.value = true;
};

const closeModal = () => {
  showCreateModal.value = false;
  editingPortfolio.value = null;
  Object.assign(form, {
    title: '',
    description: '',
    image_url: '',
    project_url: '',
    github_url: '',
    tech_stack: '',
    sort_order: 0,
    published: true,
  });
};

const savePortfolio = async () => {
  try {
    if (editingPortfolio.value) {
      await api.put(`/portfolios/${editingPortfolio.value.id}`, form);
    } else {
      await api.post('/portfolios', form);
    }
    closeModal();
    fetchPortfolios();
  } catch (error) {
    alert('保存失败');
  }
};

const deletePortfolio = async (id) => {
  if (!confirm('确定要删除这个作品吗？')) return;
  try {
    await api.delete(`/portfolios/${id}`);
    fetchPortfolios();
  } catch (error) {
    alert('删除失败');
  }
};

onMounted(fetchPortfolios);
</script>

<style lang="scss" scoped>
.admin-portfolios {
  .page-title {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 2rem;
  }
}

.portfolios-grid {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 1.5rem;
}

.portfolio-card {
  background: white;
  border-radius: 12px;
  overflow: hidden;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.05);

  .portfolio-image {
    height: 150px;
    overflow: hidden;

    img {
      width: 100%;
      height: 100%;
      object-fit: cover;
    }
  }

  .portfolio-body {
    padding: 1rem;

    h3 {
      margin-bottom: 0.5rem;
    }

    p {
      color: #718096;
      font-size: 0.9rem;
      margin-bottom: 1rem;
    }
  }

  .portfolio-actions {
    display: flex;
    gap: 0.5rem;
  }
}

.btn-sm {
  padding: 0.4rem 0.8rem;
  font-size: 0.85rem;
}

.btn-danger {
  background: #e53e3e;
  color: white;
  border: none;
  border-radius: 6px;
  cursor: pointer;

  &:hover {
    background: #c53030;
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
  max-width: 600px;
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
    padding-top: 1.5rem;
  }

  .modal-actions {
    display: flex;
    justify-content: flex-end;
    gap: 1rem;
    margin-top: 1.5rem;
  }
}

.empty-state {
  grid-column: 1 / -1;
  text-align: center;
  padding: 4rem;
  background: white;
  border-radius: 12px;
  color: #718096;
}

@media (max-width: 768px) {
  .portfolios-grid {
    grid-template-columns: 1fr;
  }
}
</style>
