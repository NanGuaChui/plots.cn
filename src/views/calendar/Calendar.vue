<template>
  <div>
    <FullCalendar class="demo-app-calendar" :options="calendarOptions">
      <template #eventContent="arg">
        <div class="fc-event-inner">
          <b v-if="arg.timeText">{{ arg.timeText }}</b>
          <span class="title">{{ arg.event.title }}</span>
          <span class="status-badge" :class="'status-' + (arg.event.extendedProps.status || 'initial')">{{
            getStatusText(arg.event.extendedProps.status)
          }}</span>
        </div>
      </template>
    </FullCalendar>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { useI18n } from 'vue-i18n';
import FullCalendar from '@fullcalendar/vue3';
import dayGridPlugin from '@fullcalendar/daygrid';
import timeGridPlugin from '@fullcalendar/timegrid';
import interactionPlugin from '@fullcalendar/interaction';
import type { CalendarOptions, DateSelectArg, EventClickArg, EventApi, DayCellMountArg } from '@fullcalendar/core';
import { useAppStore } from '@/stores/app';

const { t } = useI18n();

// simple incremental id generator for demo events
let eventGuid = 0;
function createEventId() {
  return String(eventGuid++);
}

// today's date string (YYYY-MM-DD)
const todayStr = new Date().toISOString().replace(/T.*$/, '');

// Event status type & color / text mapping
type EventStatus = 'initial' | 'ongoing' | 'completed' | 'settled';

interface ExtendedEventInput {
  id: string;
  title: string;
  start: string;
  end?: string;
  allDay?: boolean;
  extendedProps?: {
    status: EventStatus;
  };
}

const statusTextMap: Record<EventStatus, string> = {
  initial: '初始',
  ongoing: '进行中',
  completed: '已完成',
  settled: '已结算',
};
function getStatusText(rawStatus: unknown): string {
  const status = (rawStatus || 'initial') as EventStatus;
  return statusTextMap[status] ?? statusTextMap.initial;
}

// 暴露方法与映射用于模板
defineExpose({ getStatusText, statusTextMap });

// 用于 FullCalendar 事件的颜色 (可同时使用 eventClassNames + 自定义 CSS)
const statusColorMap: Record<EventStatus, { bg: string; border: string; text: string }> = {
  initial: { bg: '#64748b', border: '#475569', text: '#ffffff' }, // slate
  ongoing: { bg: '#3b82f6', border: '#2563eb', text: '#ffffff' }, // blue
  completed: { bg: '#10b981', border: '#059669', text: '#ffffff' }, // green
  settled: { bg: '#f59e0b', border: '#d97706', text: '#ffffff' }, // amber
};

// initial demo events 带不同状态示例
const INITIAL_EVENTS: ExtendedEventInput[] = [
  {
    id: createEventId(),
    title: t('calendar.event.allDay'),
    start: todayStr,
    end: new Date().toISOString().replace(/T.*$/, ''),
    extendedProps: { status: 'initial' },
  },
  {
    id: createEventId(),
    title: t('calendar.event.timed'),
    start: `${todayStr}T12:00:00`,
    end: `${todayStr}T14:00:00`,
    extendedProps: { status: 'ongoing' },
  },
  {
    id: createEventId(),
    title: '会议总结',
    start: `${todayStr}T09:00:00`,
    end: `${todayStr}T10:00:00`,
    extendedProps: { status: 'completed' },
  },
  {
    id: createEventId(),
    title: '结算审核',
    start: `${todayStr}T16:00:00`,
    end: `${todayStr}T17:00:00`,
    extendedProps: { status: 'settled' },
  },
];

// reactive list of current events (kept in sync via eventsSet callback)
const currentEvents = ref<EventApi[]>([]);

function handleDateSelect(selectInfo: DateSelectArg) {
  const title = prompt(t('calendar.prompt.newTitle'));
  const calendarApi = selectInfo.view.calendar;
  calendarApi.unselect(); // clear date selection
  if (title) {
    calendarApi.addEvent({
      id: createEventId(),
      title,
      start: selectInfo.startStr,
      end: selectInfo.endStr,
      allDay: selectInfo.allDay,
      extendedProps: { status: 'initial' },
    });
  }
}

function handleEventClick(clickInfo: EventClickArg) {
  if (confirm(t('calendar.confirm.delete', { title: clickInfo.event.title }))) {
    clickInfo.event.remove();
  }
}

function handleEvents(events: EventApi[]) {
  currentEvents.value = events;
}

function handleDayCellDidMount(info: DayCellMountArg) {
  // 右键菜单示例
  info.el.addEventListener('contextmenu', (e) => {
    e.preventDefault();
    console.log('[dayCell contextmenu]', info);
    // TODO: 在此触发你的自定义菜单组件 / store 动作
  });
}

// calendar options (reactive)
const calendarOptions = ref<CalendarOptions>({
  plugins: [dayGridPlugin, timeGridPlugin, interactionPlugin],
  locale: useAppStore().language,
  headerToolbar: {
    left: 'prev,next today',
    center: 'title',
    right: 'dayGridMonth,timeGridWeek,timeGridDay',
  },
  initialView: 'dayGridMonth',
  initialEvents: INITIAL_EVENTS,
  editable: true,
  selectable: true,
  selectMirror: true,
  dayMaxEvents: true,
  dayCellDidMount: handleDayCellDidMount,
  select: handleDateSelect,
  eventClick: handleEventClick,
  eventsSet: handleEvents,
  height: 'calc(100vh - 60px)',
  // 根据状态应用颜色
  eventClassNames(arg) {
    const status: EventStatus = (arg.event.extendedProps.status || 'initial') as EventStatus;
    return [`status-${status}`];
  },
  eventDidMount(info) {
    const status: EventStatus = (info.event.extendedProps.status || 'initial') as EventStatus;
    const colors = statusColorMap[status];
    // 直接赋予背景色，避免与默认主题冲突
    info.el.style.backgroundColor = colors.bg;
    info.el.style.borderColor = colors.border;
    info.el.style.color = colors.text;
  },
  /* you can update a remote database when these fire:
  eventAdd:
  eventChange:
  eventRemove:
  */
});

// 已在上方统一暴露
</script>

<style lang="scss" scoped>
:deep(.fc-event) {
  transition:
    background-color 0.2s ease,
    border-color 0.2s ease;
  display: flex;
  align-items: center;
  padding: 2px 4px;
  font-size: 12px;
}

.fc-event-inner {
  display: flex;
  gap: 4px;
  align-items: center;
  flex-wrap: wrap;
}

.status-badge {
  padding: 0 4px;
  border-radius: 3px;
  line-height: 1.2;
  font-size: 10px;
  font-weight: 500;
  background: rgba(255, 255, 255, 0.2);
  color: #fff;
  backdrop-filter: blur(2px);
}

/* 可按需细化不同状态样式（目前颜色在 eventDidMount 已设置，仅微调 badge） */
.status-initial.status-badge {
  background: rgba(255, 255, 255, 0.25);
}
.status-ongoing.status-badge {
  background: rgba(255, 255, 255, 0.25);
}
.status-completed.status-badge {
  background: rgba(255, 255, 255, 0.25);
}
.status-settled.status-badge {
  background: rgba(255, 255, 255, 0.25);
}
</style>
