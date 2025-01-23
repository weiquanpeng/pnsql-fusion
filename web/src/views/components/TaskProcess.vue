<template>
  <a-drawer
      :width="'66%'"
      :visible="visible"
      placement="right"
      @cancel="handleCancel"
      unmountOnClose
  >
    <!-- 抽屉内容 -->
    <div v-if="visible" ref="scrollContainer" class="scroll-container" :data-simplebar="enableScroll">
      <!-- 显示步骤条内容 -->
      <div v-show="displayMode === 1">
        <a-steps class="steps-wrapper" small :style="stepsWrapperStyle">
          <a-step
              v-for="(step, index) in steps"
              :key="index"
              :description="step.approve"
              :status="getStepStatus(step.status)"
          >
            {{ step.title }}
            <template #icon v-if="step.status === 'doing'">
              <component :is="IconLoading" />
            </template>
          </a-step>
        </a-steps>
      </div>

      <!-- 显示骨架屏 -->
      <div v-show="displayMode === 2">
        <a-skeleton :animation="true">
          <a-space direction="vertical" :style="{ width: 'calc(100% - 32px)', margin: '0 16px', height: '60px' }" size="large">
            <a-skeleton-line :rows="2" />
          </a-space>
        </a-skeleton>
      </div>
    </div>

    <!-- 动态数据 -->
    <a-descriptions :column="2" bordered style="margin-top: 16px;">
      <a-descriptions-item v-for="item in dynamicData" :key="item.label" :label="item.label">
        <template v-if="isLoading">
          <a-skeleton-line :widths="['80%']" :rows="1" />
        </template>
        <template v-else>
          {{ item.value }}
        </template>
      </a-descriptions-item>
    </a-descriptions>

    <!-- 工单参数卡片 -->
    <a-card class="parameter-card" title="工单参数" :bordered="false" style="margin-top: 16px;">
      <div style="display: flex; align-items: center; justify-content: space-between;">
        <!-- 查看参数 -->
        <a-trigger :popup-translate="[-100, 20]">
          <a-button type="outline" style="border-radius: 4px;">
            <template #icon>
              <icon-info-circle />
            </template>
            查看参数
          </a-button>
          <template #content>
            <div class="trigger-demo-translate">
              <a-scrollbar style="height: 200px; overflow: auto;">
                <a-descriptions :column="1" bordered>
                  <a-descriptions-item
                      v-for="(value, key) in steps[0]?.parameter || {}"
                      :key="key"
                      :label="key"
                  >
                    <a-tag>{{ value }}</a-tag>
                  </a-descriptions-item>
                </a-descriptions>
              </a-scrollbar>
            </div>
          </template>
        </a-trigger>

        <!-- 执行进度 -->
        <div style="display: flex; align-items: center; gap: 8px;">
          <span style="font-size: 14px; color: var(--color-text-1);">任务执行进度</span>
          <a-progress :percent="0.2" :stroke-width="6" :style="{ width: '500px' }" />
        </div>
      </div>
    </a-card>

    <!-- 日志组件 -->
    <a-card class="log-viewer" title="工单日志" :bordered="false" style="margin-top: 16px;">
      <!-- 加载状态 -->
      <a-spin v-if="logDisplayMode === 2" tip="加载日志中..." />

      <!-- 日志内容 -->
      <a-textarea
          v-if="logDisplayMode === 1"
          :model-value="logs"
          readonly
          class="log-container"
      />
    </a-card>

    <!-- 抽屉底部 -->
    <template #footer>
      <div class="footer-wrapper">
        <a-button type="primary" @click="handleOk" :loading="confirmLoading">确认</a-button>
        <a-button type="dashed" status="danger" style="margin-left: 18px;" @click="handleCancel">关闭</a-button>
      </div>
    </template>
  </a-drawer>
</template>

<script setup>
import 'simplebar/dist/simplebar.min.css';
import SimpleBar from 'simplebar';
import { IconLoading, IconInfoCircle } from '@arco-design/web-vue/es/icon';
import { ref, computed, watch, nextTick, getCurrentInstance } from 'vue';
import formatDateTime from '@/utils/formatDateTime.js';
import {uptSubTaskConfigStatus} from "@/api/task_config.js";

// 获取当前组件实例
const { proxy } = getCurrentInstance();

// 抽屉是否可见
const visible = ref(false);
// 任务标题
const taskTitle = ref('');
// 滚动容器
const scrollContainer = ref(null);
// 确认按钮的加载状态
const confirmLoading = ref(false);
// 数据加载状态
const isLoading = ref(false);
// 显示模式：1-显示数据，2-显示骨架屏
const displayMode = ref(2);
// 步骤条数据
const steps = ref([]);
// 日志内容
const logs = ref('');
// 日志显示模式：1-显示日志，2-显示加载动画
const logDisplayMode = ref(2);

// 监听 steps 的变化，动态切换 displayMode
watch(steps, (newSteps) => {
  if (newSteps && newSteps.length > 0) {
    displayMode.value = 1; // 有数据时显示步骤条
    isLoading.value = false; // 数据加载完成
  } else {
    displayMode.value = 2; // 无数据时显示骨架屏
    isLoading.value = true; // 数据加载中
  }
});

// 动态生成 data
const dynamicData = computed(() => {
  return [
    { label: '标题', value: taskTitle.value || '' },
    { label: '子工单号', value: steps.value[0]?.id || '' },
    { label: '申请人', value: steps.value[0]?.owner || '' },
    { label: '创建时间', value: formatDateTime(steps.value[0]?.createdAt) || '' },
    { label: '审批人', value: steps.value[0]?.approve || '' },
    { label: '更新时间', value: formatDateTime(steps.value[0]?.updatedAt) || '' },
    { label: '描述', value: steps.value[0]?.describe || '' },
  ];
});

// 判断是否需要启用滚动条
const enableScroll = computed(() => steps.value.length > 6);

// 动态设置步骤条宽度
const stepsWrapperStyle = computed(() => {
  return enableScroll.value ? { minWidth: '200%' } : { minWidth: '100%' };
});

// 根据 step.status 返回对应的状态
const getStepStatus = (status) => {
  switch (status) {
    case 'tosplit':
      return 'wait';
    case 'todo':
      return 'process';
    case 'doing':
      return 'process';
    case 'error':
      return 'error';
    case 'done':
      return 'finish';
    default:
      return 'error';
  }
};

// 处理抽屉打开
const handleClick = (list, logsData, title) => {
  if (logsData === '') {
    // 第一次调用：打开抽屉并显示骨架屏
    visible.value = true; // 打开抽屉
    taskTitle.value = title; // 设置任务标题
    steps.value = []; // 清空步骤条数据
    logs.value = ''; // 清空日志内容
    logDisplayMode.value = 2; // 显示加载动画
  } else {
    // 第二次调用：更新数据
    steps.value = list; // 更新步骤条数据
    if (logsData) {
      logs.value = Array.isArray(logsData) ? logsData.join('\n') : logsData; // 处理日志数据
      logDisplayMode.value = 1; // 显示日志内容
    } else {
      logs.value = ''; // 清空日志内容
      logDisplayMode.value = 2; // 显示加载动画
    }
  }
};

// 处理确认按钮点击
const handleOk = async () => {
  confirmLoading.value = true; // 点击后显示 loading
  try {
    // 找到当前正在处理的节点
    const currentStep = steps.value.find(step => step.status === 'todo');
    if (currentStep) {
      await uptSubTaskConfigStatus(currentStep.id);
    } else {
      console.log('没有找到正在处理的子工单');
    }
    visible.value = false;
    emit('refreshTable'); // 触发父组件的刷新逻辑
  } catch (error) {
    console.error('操作失败', error);
  } finally {
    confirmLoading.value = false; // 完成后关闭 loading
  }
};

// 处理抽屉关闭
const handleCancel = () => {
  visible.value = false;
  steps.value = []; // 清空步骤条数据
  taskTitle.value = ''; // 清空任务标题
  logs.value = ''; // 清空日志内容
  logDisplayMode.value = 2; // 重置日志显示模式为加载动画
};

// 监听抽屉的 visible 状态
watch(visible, (newVal) => {
  if (newVal) {
    nextTick(() => {
      if (scrollContainer.value && enableScroll.value) {
        new SimpleBar(scrollContainer.value);
      }
    });
  }
});

// 暴露 handleClick 方法给父组件
defineExpose({
  handleClick,
});

// 定义 emit
const emit = defineEmits(['refreshTable']);
</script>

<style scoped>
.scroll-container {
  width: 100%;
  overflow-x: auto;
  height: 60px;
  padding-bottom: 20px;
}

.steps-wrapper {
  display: inline-flex;
  white-space: nowrap;
  margin-top: 16px;
  padding: 8px;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.footer-wrapper {
  display: flex;
  justify-content: flex-start;
  padding: 16px;
}

.log-viewer {
  background-color: #ffffff;
  border-radius: 12px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
  border: 1px solid #e8e8e8;
}

.log-container {
  width: 100%;
  height: 420px;
  font-family: monospace;
  font-size: 14px;
  line-height: 1.5;
  color: #333333;
  background-color: #f9f9f9;
  border: 1px solid #d9d9d9;
  border-radius: 8px;
  padding: 12px;
  resize: none;
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1);
  overflow-y: auto;
}

.parameter-card {
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.trigger-demo-translate {
  padding: 16px;
  background-color: #ffffff;
  border-radius: 8px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}
</style>