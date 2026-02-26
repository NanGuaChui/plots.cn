<template>
  <div class="character-select-container">
    <div class="header">
      <h1>选择角色</h1>
      <el-button type="danger" @click="handleLogout">退出登录</el-button>
    </div>

    <div class="character-slots">
      <div v-for="slotIndex in 4" :key="slotIndex - 1" class="character-slot">
        <template v-if="getCharacterBySlot(slotIndex - 1)">
          <div class="character-card" @click="selectAndEnter(getCharacterBySlot(slotIndex - 1))">
            <div class="character-avatar">
              <el-icon :size="48">
                <User />
              </el-icon>
            </div>
            <div class="character-info">
              <h3>{{ getCharacterBySlot(slotIndex - 1).name }}</h3>
              <el-tag :type="slotIndex === 1 ? 'primary' : 'warning'" size="small">
                {{ slotIndex === 1 ? '主号' : '铁人号' }}
              </el-tag>
              <p class="level">战斗等级: {{ getCharacterBySlot(slotIndex - 1).combat_level }}</p>
            </div>
            <el-button
              type="danger"
              size="small"
              class="delete-btn"
              @click.stop="confirmDelete(getCharacterBySlot(slotIndex - 1))"
            >
              删除
            </el-button>
          </div>
        </template>
        <template v-else>
          <div class="empty-slot" @click="openCreateDialog(slotIndex - 1)">
            <el-icon :size="48" color="#ccc">
              <Plus />
            </el-icon>
            <p>创建{{ slotIndex === 1 ? '主号' : '铁人号' }}</p>
          </div>
        </template>
      </div>
    </div>

    <!-- 创建角色对话框 -->
    <el-dialog v-model="createDialogVisible" title="创建角色" width="400px">
      <el-form ref="createFormRef" :model="createForm" :rules="createRules" label-width="80px">
        <el-form-item label="角色名" prop="name">
          <el-input v-model="createForm.name" placeholder="请输入角色名" maxlength="50" />
        </el-form-item>
        <el-form-item label="角色类型">
          <el-tag :type="createForm.slotIndex === 0 ? 'primary' : 'warning'">
            {{ createForm.slotIndex === 0 ? '主号' : '铁人号' }}
          </el-tag>
          <p class="type-desc">
            {{ createForm.slotIndex === 0 ? '主号可参与市场交易' : '铁人号完全自给自足，无法交易' }}
          </p>
        </el-form-item>
      </el-form>
      <template #footer>
        <el-button @click="createDialogVisible = false">取消</el-button>
        <el-button type="primary" :loading="createLoading" @click="handleCreateCharacter"> 创建 </el-button>
      </template>
    </el-dialog>
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { User, Plus } from '@element-plus/icons-vue';
import { ElMessage, ElMessageBox } from 'element-plus';
import { useUserStore } from '@/stores/user';
import { useCharacterStore } from '@/stores/character';

const router = useRouter();
const userStore = useUserStore();
const characterStore = useCharacterStore();

const createDialogVisible = ref(false);
const createLoading = ref(false);
const createFormRef = ref(null);

const createForm = reactive({
  name: '',
  slotIndex: 0,
});

const createRules = {
  name: [
    { required: true, message: '请输入角色名', trigger: 'blur' },
    { min: 2, max: 50, message: '角色名长度为 2-50 个字符', trigger: 'blur' },
  ],
};

onMounted(async () => {
  try {
    await characterStore.fetchCharacters();
  } catch (error) {
    ElMessage.error('获取角色列表失败');
  }
});

function getCharacterBySlot(slotIndex) {
  return characterStore.characters.find((c) => c.slot_index === slotIndex);
}

function openCreateDialog(slotIndex) {
  createForm.name = '';
  createForm.slotIndex = slotIndex;
  createDialogVisible.value = true;
}

async function handleCreateCharacter() {
  const valid = await createFormRef.value.validate().catch(() => false);
  if (!valid) return;

  createLoading.value = true;
  try {
    await characterStore.createCharacter(createForm.name, createForm.slotIndex);
    ElMessage.success('创建成功');
    createDialogVisible.value = false;
  } catch (error) {
    ElMessage.error(error.response?.data?.error || '创建失败');
  } finally {
    createLoading.value = false;
  }
}

async function confirmDelete(character) {
  try {
    await ElMessageBox.confirm(`确定要删除角色 "${character.name}" 吗？此操作不可恢复！`, '删除角色', {
      confirmButtonText: '确定删除',
      cancelButtonText: '取消',
      type: 'warning',
    });
    await characterStore.deleteCharacter(character.id);
    ElMessage.success('删除成功');
  } catch (error) {
    if (error !== 'cancel') {
      ElMessage.error(error.response?.data?.error || '删除失败');
    }
  }
}

function selectAndEnter(character) {
  characterStore.selectCharacter(character);
  router.push('/game');
}

function handleLogout() {
  userStore.logout();
  characterStore.clearState();
  router.push('/login');
}
</script>

<style scoped>
.character-select-container {
  min-height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 40px;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 40px;
}

.header h1 {
  color: #fff;
  margin: 0;
}

.character-slots {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(280px, 1fr));
  gap: 24px;
  max-width: 1200px;
  margin: 0 auto;
}

.character-slot {
  aspect-ratio: 1;
  min-height: 200px;
}

.character-card {
  height: 100%;
  background: #fff;
  border-radius: 12px;
  padding: 24px;
  display: flex;
  flex-direction: column;
  align-items: center;
  cursor: pointer;
  transition: all 0.3s ease;
  position: relative;
}

.character-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.2);
}

.character-avatar {
  width: 80px;
  height: 80px;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #fff;
  margin-bottom: 16px;
}

.character-info {
  text-align: center;
  flex: 1;
}

.character-info h3 {
  margin: 0 0 8px 0;
  color: #333;
}

.character-info .level {
  margin: 8px 0 0 0;
  color: #666;
  font-size: 14px;
}

.delete-btn {
  position: absolute;
  top: 12px;
  right: 12px;
}

.empty-slot {
  height: 100%;
  background: rgba(255, 255, 255, 0.9);
  border-radius: 12px;
  border: 2px dashed #ccc;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: all 0.3s ease;
}

.empty-slot:hover {
  border-color: #667eea;
  background: #fff;
}

.empty-slot p {
  margin: 12px 0 0 0;
  color: #999;
}

.type-desc {
  margin: 8px 0 0 0;
  font-size: 12px;
  color: #999;
}
</style>
