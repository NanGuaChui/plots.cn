<template>
  <div class="dashboard-container">
    <div class="welcome-section">
      <h2>欢迎回来，{{ characterStore.currentCharacter?.name }}</h2>
      <el-tag :type="characterStore.currentCharacter?.type === 'main' ? 'primary' : 'warning'">
        {{ characterStore.currentCharacter?.type === 'main' ? '主号' : '铁人号' }}
      </el-tag>
    </div>

    <el-row :gutter="24">
      <!-- 角色状态卡片 -->
      <el-col :span="12">
        <el-card class="stat-card">
          <template #header>
            <div class="card-header">
              <span>角色状态</span>
            </div>
          </template>
          <div class="stat-grid">
            <div class="stat-item">
              <span class="label">战斗等级</span>
              <span class="value">{{ characterStore.currentCharacter?.combat_level || 1 }}</span>
            </div>
            <div class="stat-item">
              <span class="label">力量</span>
              <span class="value">{{ characterStore.currentCharacter?.strength || 10 }}</span>
            </div>
            <div class="stat-item">
              <span class="label">敏捷</span>
              <span class="value">{{ characterStore.currentCharacter?.agility || 10 }}</span>
            </div>
            <div class="stat-item">
              <span class="label">耐力</span>
              <span class="value">{{ characterStore.currentCharacter?.endurance || 10 }}</span>
            </div>
            <div class="stat-item">
              <span class="label">智力</span>
              <span class="value">{{ characterStore.currentCharacter?.intelligence || 10 }}</span>
            </div>
            <div class="stat-item">
              <span class="label">幸运</span>
              <span class="value">{{ characterStore.currentCharacter?.luck || 10 }}</span>
            </div>
          </div>
        </el-card>
      </el-col>

      <!-- 任务队列卡片 -->
      <el-col :span="12">
        <el-card class="task-card">
          <template #header>
            <div class="card-header">
              <span>任务队列</span>
              <span class="queue-count">(0/10)</span>
            </div>
          </template>
          <div class="empty-queue">
            <el-icon :size="48" color="#ccc">
              <List />
            </el-icon>
            <p>暂无任务</p>
            <p class="hint">从左侧菜单选择技能开始工作</p>
          </div>
        </el-card>
      </el-col>
    </el-row>

    <!-- 快捷操作 -->
    <el-card class="quick-actions">
      <template #header>
        <div class="card-header">
          <span>快捷操作</span>
        </div>
      </template>
      <div class="action-grid">
        <div class="action-item" @click="goTo('/game/skills/woodcutting')">
          <el-icon :size="32">
            <SetUp />
          </el-icon>
          <span>伐木</span>
        </div>
        <div class="action-item" @click="goTo('/game/skills/mining')">
          <el-icon :size="32">
            <Coin />
          </el-icon>
          <span>采矿</span>
        </div>
        <div class="action-item" @click="goTo('/game/skills/fishing')">
          <el-icon :size="32">
            <Ship />
          </el-icon>
          <span>钓鱼</span>
        </div>
        <div class="action-item" @click="goTo('/game/battle')">
          <el-icon :size="32">
            <Aim />
          </el-icon>
          <span>战斗</span>
        </div>
        <div class="action-item" @click="goTo('/game/inventory')">
          <el-icon :size="32">
            <Box />
          </el-icon>
          <span>仓库</span>
        </div>
        <div class="action-item" @click="goTo('/game/character')">
          <el-icon :size="32">
            <User />
          </el-icon>
          <span>角色</span>
        </div>
      </div>
    </el-card>
  </div>
</template>

<script setup>
import { useRouter } from 'vue-router';
import { useCharacterStore } from '@/stores/character';
import { List, SetUp, Coin, Ship, Aim, Box, User } from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

const router = useRouter();
const characterStore = useCharacterStore();

function goTo(path) {
  // 暂时提示功能未实现
  ElMessage.info('该功能将在后续任务中实现');
  // router.push(path);
}
</script>

<style scoped>
.dashboard-container {
  padding: 24px;
}

.welcome-section {
  display: flex;
  align-items: center;
  gap: 12px;
  margin-bottom: 24px;
}

.welcome-section h2 {
  margin: 0;
  color: #333;
}

.stat-card,
.task-card,
.quick-actions {
  margin-bottom: 24px;
}

.card-header {
  display: flex;
  align-items: center;
  gap: 8px;
  font-weight: 600;
}

.queue-count {
  color: #999;
  font-weight: normal;
}

.stat-grid {
  display: grid;
  grid-template-columns: repeat(3, 1fr);
  gap: 16px;
}

.stat-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 12px;
  background: #f5f7fa;
  border-radius: 8px;
}

.stat-item .label {
  font-size: 12px;
  color: #999;
  margin-bottom: 4px;
}

.stat-item .value {
  font-size: 24px;
  font-weight: 600;
  color: #667eea;
}

.empty-queue {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 40px 0;
  color: #999;
}

.empty-queue p {
  margin: 8px 0 0 0;
}

.empty-queue .hint {
  font-size: 12px;
  color: #ccc;
}

.action-grid {
  display: grid;
  grid-template-columns: repeat(6, 1fr);
  gap: 16px;
}

.action-item {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px;
  background: #f5f7fa;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
}

.action-item:hover {
  background: #667eea;
  color: #fff;
}

.action-item span {
  margin-top: 8px;
  font-size: 14px;
}

@media (max-width: 1200px) {
  .action-grid {
    grid-template-columns: repeat(3, 1fr);
  }
}

@media (max-width: 768px) {
  .stat-grid {
    grid-template-columns: repeat(2, 1fr);
  }

  .action-grid {
    grid-template-columns: repeat(2, 1fr);
  }
}
</style>
