<script setup lang="ts">
import SvgIcon from '@/components/SvgIcon.vue';
import { NButton } from 'naive-ui';
</script>

<template>
  <div class="flex justify-center items-center h-screen gap-4">
    <n-button icon-placement="right">
      <template #icon>
        <SvgIcon icon="ri:anticlockwise-2-line" />
      </template>
      +100 元
    </n-button>
  </div>
</template>
