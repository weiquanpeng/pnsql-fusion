<template>
  <a-card class="log-viewer" title="工单日志" :bordered="false" style="margin-top: 30px">
    <!-- 加载状态 -->
    <a-spin v-if="loading" tip="Loading logs..." />

    <!-- 日志内容 -->
    <a-textarea
        v-else
        :model-value="logs"
        readonly
        class="log-container"
    />
  </a-card>
</template>

<script setup>
import { ref } from 'vue';

// 日志相关状态
const logs = ref('');
const loading = ref(true);

// 模拟日志数据
const mockLogs = `
[2023-10-25 10:00:00] INFO: Server started on port 8080
[2023-10-25 10:01:23] DEBUG: Received request from 192.168.1.1
[2023-10-25 10:02:45] WARNING: Disk usage is above 80%
[2023-10-25 10:03:12] ERROR: Failed to connect to database
[2023-10-25 10:05:00] INFO: Database connection restored
[2023-10-25 10:10:00] INFO: User "admin" logged in
[2023-10-25 10:15:00] DEBUG: Processing background task
[2023-10-25 10:20:00] INFO: Background task completed
[2023-10-25 10:25:00] WARNING: High memory usage detected
[2023-10-25 10:30:00] INFO: Memory usage normalized
`.trim();

// 模拟异步获取日志数据
const fetchLogs = () => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(mockLogs);
    }, 1000); // 模拟 1 秒延迟
  });
};

// 初始化加载日志
fetchLogs().then((data) => {
  logs.value = data;
  loading.value = false;
});
</script>

<style scoped>
.log-viewer {
  background-color: #ffffff; /* 白色背景 */
  border-radius: 12px; /* 更大的圆角 */
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15); /* 更柔和的阴影 */
  border: 1px solid #e8e8e8; /* 浅灰色边框 */
  margin-top: 20px;
}

.log-container {
  width: 100%;
  height: 300px; /* 固定高度 */
  font-family: monospace;
  font-size: 14px;
  line-height: 1.5;
  color: #333333; /* 深灰色文字 */
  background-color: #f9f9f9; /* 浅灰色背景 */
  border: 1px solid #d9d9d9; /* 浅灰色边框 */
  border-radius: 8px; /* 圆角 */
  padding: 12px; /* 更大的内边距 */
  resize: none; /* 禁止用户调整大小 */
  box-shadow: inset 0 1px 3px rgba(0, 0, 0, 0.1); /* 内阴影 */
  overflow-y: auto; /* 启用垂直滚动条 */
}

/* 自定义滚动条样式 */
.log-container::-webkit-scrollbar {
  width: 8px;
}

.log-container::-webkit-scrollbar-track {
  background: #f1f1f1; /* 滚动条轨道颜色 */
  border-radius: 4px;
}

.log-container::-webkit-scrollbar-thumb {
  background: #c1c1c1; /* 滚动条滑块颜色 */
  border-radius: 4px;
}

.log-container::-webkit-scrollbar-thumb:hover {
  background: #a8a8a8; /* 滚动条滑块悬停颜色 */
}
</style>