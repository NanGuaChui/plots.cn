<template>
  <div>
    <FullCalendar :options="calendarOptions">
      <template #eventContent="arg">
        <div
          class="fc-event-inner flex justify-between"
          :class="'status-' + (arg.event.extendedProps.status || 'initial')"
        >
          <div class="flex items-center gap-1">
            <b v-if="arg.timeText">{{ arg.timeText }}</b>
            <span class="title">{{ arg.event.title }}</span>
            <span class="status-badge" :class="'status-' + (arg.event.extendedProps.status || 'initial')">
              {{ getStatusText(arg.event.extendedProps.status) }}
            </span>
          </div>
          <div>
            <NButton quaternary circle size="tiny" color="#fff">
              <template #icon>
                <SvgIcon icon="ri:delete-bin-line" />
              </template>
            </NButton>
            <NDropdown trigger="hover" size="small" :options="menuOptions" @select="handleMenuSelect">
              <NButton quaternary circle size="tiny" color="#fff">
                <template #icon>
                  <SvgIcon icon="ri:dashboard-horizontal-fill" />
                </template>
              </NButton>
            </NDropdown>
          </div>
        </div>
      </template>
    </FullCalendar>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted, ref, type Ref } from 'vue';
import FullCalendar from '@fullcalendar/vue3';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import type { CalendarOptions } from '@fullcalendar/core';
import { useAppStore } from '@/stores/app';
import { NButton, NDropdown } from 'naive-ui';
import SvgIcon from '@/components/SvgIcon.vue';

// Event status type & color / text mapping
type EventStatus = 'initial' | 'ongoing' | 'completed' | 'settled';

const statusTextMap: Record<EventStatus, string> = {
  initial: '初始',
  ongoing: '进行中',
  completed: '已完成',
  settled: '已结算',
};

const menuOptions: Ref<Array<{ label: string; key: EventStatus }>> = computed(() =>
  Object.entries(statusTextMap).map(([value, label]) => ({ label, key: value as EventStatus })),
);

function getStatusText(rawStatus: unknown): string {
  const status = (rawStatus || 'initial') as EventStatus;
  return statusTextMap[status] ?? statusTextMap.initial;
}

// calendar options (reactive)
const calendarOptions = ref<CalendarOptions>({
  plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
  locale: useAppStore().language,
  height: 'calc(max(100vh - 200px, 700px))',
  headerToolbar: {
    left: 'prev,next today',
    center: 'title',
    right: 'dayGridMonth,timeGridWeek,timeGridDay',
  },
  initialView: 'dayGridMonth',
  editable: true,
  selectable: true,
});
const handleMenuSelect = () => {};

onMounted(async () => {
  const res = await fetch('/data.json');
  if (res.ok) {
    const events = await res.json();
    calendarOptions.value.events = events;
  }
});
</script>

<style lang="scss" scoped>
:deep(.fc-event) {
  transition:
    background-color 0.2s ease,
    border-color 0.2s ease;
  display: flex;
  padding: 0;
  align-items: center;
  font-size: 12px;
  margin-bottom: 5px;
  &:hover {
    filter: brightness(1.1);
  }
}

.fc-event-inner {
  display: flex;
  gap: 4px;
  align-items: center;
  flex-wrap: wrap;
  flex: 1;

  padding: 4px;
  border-radius: 3px;
  line-height: 1.2;
  font-size: 10px;
  font-weight: 500;
  background: rgba(255, 255, 255, 0.2);
  color: #fff;
  backdrop-filter: blur(2px);
}

.status-initial {
  background: #64748b;
}
.status-ongoing {
  background: #3b82f6;
}
.status-completed {
  background: #10b981;
}
.status-settled {
  background: #f59e0b;
}
</style>
