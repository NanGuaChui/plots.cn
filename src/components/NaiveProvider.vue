<script setup lang="ts">
import { defineComponent, h } from 'vue';
import {
  NDialogProvider,
  NLoadingBarProvider,
  NMessageProvider,
  NModalProvider,
  NNotificationProvider,
  useDialog,
  useLoadingBar,
  useMessage,
  useNotification,
} from 'naive-ui';

const NaiveProviderContent = defineComponent({
  name: 'NaiveProviderContent',
  setup() {
    window.$loadingBar = useLoadingBar();
    window.$dialog = useDialog();
    window.$message = useMessage();
    window.$notification = useNotification();
  },
  render() {
    return h('div');
  },
});
</script>

<template>
  <NLoadingBarProvider>
    <NModalProvider>
      <NDialogProvider>
        <NNotificationProvider>
          <NMessageProvider>
            <slot></slot>
            <NaiveProviderContent />
          </NMessageProvider>
        </NNotificationProvider>
      </NDialogProvider>
    </NModalProvider>
  </NLoadingBarProvider>
</template>
