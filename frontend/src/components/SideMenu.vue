<template>
  <el-menu :default-active="activeMenu" class="side-menu" :collapse="isCollapsed" router @select="handleSelect">
    <!-- 主页 -->
    <el-menu-item index="/game">
      <el-icon>
        <House />
      </el-icon>
      <template #title>主页</template>
    </el-menu-item>

    <!-- 生活技能 -->
    <el-sub-menu index="skills">
      <template #title>
        <el-icon>
          <Tools />
        </el-icon>
        <span>生活技能</span>
      </template>

      <!-- 采集类 -->
      <el-menu-item-group title="采集">
        <el-menu-item index="/game/skills/woodcutting" disabled>
          <el-icon>
            <SetUp />
          </el-icon>
          伐木
        </el-menu-item>
        <el-menu-item index="/game/skills/mining" disabled>
          <el-icon>
            <Coin />
          </el-icon>
          采矿
        </el-menu-item>
        <el-menu-item index="/game/skills/fishing" disabled>
          <el-icon>
            <Ship />
          </el-icon>
          钓鱼
        </el-menu-item>
        <el-menu-item index="/game/skills/gathering" disabled>
          <el-icon>
            <Cherry />
          </el-icon>
          采集
        </el-menu-item>
        <el-menu-item index="/game/skills/farming" disabled>
          <el-icon>
            <Chicken />
          </el-icon>
          养殖
        </el-menu-item>
      </el-menu-item-group>

      <!-- 合成类 -->
      <el-menu-item-group title="合成">
        <el-menu-item index="/game/skills/cooking" disabled>
          <el-icon>
            <Bowl />
          </el-icon>
          烹饪
        </el-menu-item>
        <el-menu-item index="/game/skills/smithing" disabled>
          <el-icon>
            <MagicStick />
          </el-icon>
          锻造
        </el-menu-item>
        <el-menu-item index="/game/skills/tailoring" disabled>
          <el-icon>
            <ScaleToOriginal />
          </el-icon>
          缝制
        </el-menu-item>
      </el-menu-item-group>
    </el-sub-menu>

    <!-- 战斗 -->
    <el-menu-item index="/game/battle" disabled>
      <el-icon>
        <Aim />
      </el-icon>
      <template #title>战斗</template>
    </el-menu-item>

    <!-- 强化 -->
    <el-menu-item index="/game/enhancement" disabled>
      <el-icon>
        <MagicStick />
      </el-icon>
      <template #title>强化</template>
    </el-menu-item>

    <!-- 仓库 -->
    <el-menu-item index="/game/inventory" disabled>
      <el-icon>
        <Box />
      </el-icon>
      <template #title>仓库</template>
    </el-menu-item>

    <!-- 角色 -->
    <el-menu-item index="/game/character" disabled>
      <el-icon>
        <User />
      </el-icon>
      <template #title>角色</template>
    </el-menu-item>
  </el-menu>
</template>

<script setup>
import { computed } from 'vue';
import { useRoute } from 'vue-router';
import {
  House,
  Tools,
  SetUp,
  Coin,
  Ship,
  Cherry,
  Chicken,
  Bowl,
  MagicStick,
  ScaleToOriginal,
  Aim,
  Box,
  User,
} from '@element-plus/icons-vue';
import { ElMessage } from 'element-plus';

defineProps({
  isCollapsed: {
    type: Boolean,
    default: false,
  },
});

const route = useRoute();

const activeMenu = computed(() => {
  return route.path;
});

function handleSelect(index) {
  // 如果是禁用的菜单项，显示提示
  if (index !== '/game') {
    // 路由会自动处理，这里不需要额外操作
  }
}
</script>

<style scoped>
.side-menu {
  height: 100%;
  border-right: none;
}

.side-menu:not(.el-menu--collapse) {
  width: 200px;
}

.el-menu-item.is-disabled {
  opacity: 0.6;
}
</style>
