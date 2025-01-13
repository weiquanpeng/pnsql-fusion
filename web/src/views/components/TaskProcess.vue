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
    <a-descriptions :data="dynamicData" :column="2" bordered />
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
import {IconLoading} from '@arco-design/web-vue/es/icon'; // 仅导入 IconLoading
import {ref, computed, watch, nextTick} from 'vue';

const visible = ref(false);
const taskTitle = ref();
const scrollContainer = ref(null);
const confirmLoading = ref(false);
import formatDateTime from "@/utils/formatDateTime.js";

// 步骤条数据
const steps = ref([]);

// 动态生成 data
const dynamicData = computed(() => {
  return [
    { label: '标题', value: taskTitle.value || '' }, // 使用 taskTitle 的值
    { label: '子工单号', value: steps.value[0]?.id || '' }, // 使用 steps 的数据
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
  return enableScroll.value ? {minWidth: '200%'} : {minWidth: '100%'};
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

const handleClick = (list, title) => {
  visible.value = true;
  steps.value = list;
  taskTitle.value = title;
};

const handleOk = async () => {
  confirmLoading.value = true; // 点击后显示 loading
  try {
    // 模拟异步操作（例如 API 请求）
    await new Promise((resolve) => setTimeout(resolve, 2000)); // 2 秒延迟
    // 处理确认逻辑
    visible.value = false;
    emit('refreshTable'); // 触发父组件的刷新逻辑
  } catch (error) {
    console.error('操作失败', error);
  } finally {
    confirmLoading.value = false; // 完成后关闭 loading
  }
};

const handleCancel = () => {
  visible.value = false;
};

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
}

/* 自定义底部样式 */
.footer-wrapper {
  display: flex;
  justify-content: flex-start; /* 按钮靠左 */
  padding: 16px; /* 调整内边距 */
}
</style>