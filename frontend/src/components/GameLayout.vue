<template>
  <el-container class="game-layout">
    <!-- 侧边栏 -->
    <el-aside :width="isCollapsed ? '64px' : '200px'" class="sidebar">
      <div class="logo" @click="toggleCollapse">
        <el-icon :size="24">
          <Compass />
        </el-icon>
        <span v-if="!isCollapsed" class="logo-text">Poetize</span>
      </div>
      <SideMenu :is-collapsed="isCollapsed" />
    </el-aside>

    <!-- 主内容区 -->
    <el-container class="main-container">
      <!-- 顶部栏 -->
      <el-header class="header">
        <div class="header-left">
          <el-button :icon="isCollapsed ? Expand : Fold" text @click="toggleCollapse" />
        </div>
        <div class="header-right">
          <el-dropdown @command="handleCommand">
            <span class="user-info">
              <el-avatar :size="32" :icon="User" />
              <span class="username">{{ characterStore.currentCharacter?.name }}</span>
              <el-icon>
                <ArrowDown />
              </el-icon>
            </span>
            <template #dropdown>
              <el-dropdown-menu>
                <el-dropdown-item command="switchCharacter">切换角色</el-dropdown-item>
                <el-dropdown-item command="logout" divided>退出登录</el-dropdown-item>
              </el-dropdown-menu>
            </template>
          </el-dropdown>
        </div>
      </el-header>

      <!-- 内容区 -->
      <el-main class="main-content">
        <router-view />
      </el-main>
    </el-container>
  </el-container>
</template>

<script setup>
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { Compass, User, ArrowDown, Expand, Fold } from '@element-plus/icons-vue';
import { useUserStore } from '@/stores/user';
import { useCharacterStore } from '@/stores/character';
import SideMenu from './SideMenu.vue';

const router = useRouter();
const userStore = useUserStore();
const characterStore = useCharacterStore();

const isCollapsed = ref(false);

function toggleCollapse() {
  isCollapsed.value = !isCollapsed.value;
}

function handleCommand(command) {
  switch (command) {
    case 'switchCharacter':
      characterStore.selectCharacter(null);
      router.push('/characters');
      break;
    case 'logout':
      userStore.logout();
      characterStore.clearState();
      router.push('/login');
      break;
  }
}
</script>

<style scoped>
.game-layout {
  height: 100vh;
}

.sidebar {
  background: #304156;
  transition: width 0.3s;
  overflow: hidden;
}

.logo {
  height: 60px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  color: #fff;
  font-size: 18px;
  font-weight: bold;
  cursor: pointer;
  border-bottom: 1px solid rgba(255, 255, 255, 0.1);
}

.logo-text {
  white-space: nowrap;
}

.main-container {
  flex-direction: column;
}

.header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 0 16px;
  background: #fff;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
}

.header-left,
.header-right {
  display: flex;
  align-items: center;
}

.user-info {
  display: flex;
  align-items: center;
  gap: 8px;
  cursor: pointer;
  padding: 4px 8px;
  border-radius: 4px;
}

.user-info:hover {
  background: #f5f5f5;
}

.username {
  color: #333;
  font-size: 14px;
}

.main-content {
  background: #f0f2f5;
  overflow-y: auto;
}

/* Element Plus Menu 深色主题 */
:deep(.el-menu) {
  background-color: #304156;
  border-right: none;
}

:deep(.el-menu-item),
:deep(.el-sub-menu__title) {
  color: #bfcbd9;
}

:deep(.el-menu-item:hover),
:deep(.el-sub-menu__title:hover) {
  background-color: #263445;
}

:deep(.el-menu-item.is-active) {
  color: #409eff;
  background-color: #263445;
}

:deep(.el-sub-menu .el-menu) {
  background-color: #1f2d3d;
}

:deep(.el-menu-item-group__title) {
  color: #909399;
  padding-left: 20px;
}

:deep(.el-menu--collapse .el-sub-menu__title span),
:deep(.el-menu--collapse .el-menu-item span) {
  display: none;
}
</style>
