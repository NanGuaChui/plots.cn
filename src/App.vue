<script setup lang="ts">
import { useLanguage } from '@/hooks/useLanguage';
import NaiveProvider from '@/components/NaiveProvider.vue';
import { NConfigProvider } from 'naive-ui';

const { language } = useLanguage();
</script>

<template>
  <NConfigProvider class="h-full" :locale="language">
    <NaiveProvider>
      <RouterView />
    </NaiveProvider>
  </NConfigProvider>
</template>
