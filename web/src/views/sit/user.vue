<template>
  <a-row :gutter="16">
    <!-- 第一个输入框：显示确定的数据且不可编辑 -->
    <a-col :span="12">
      <a-input :value="fixedData" disabled />
    </a-col>

    <!-- 第二个输入框：用于输入新数据 -->
    <a-col :span="12">
      <a-input v-model:value="newData" placeholder="请输入新数据" />
    </a-col>

    <!-- 提交按钮 -->
    <a-col :span="24" style="margin-top: 16px;">
      <a-button type="primary" @click="handleSubmit">提交</a-button>
    </a-col>
  </a-row>
</template>

<script>
import { defineComponent, ref } from 'vue';

export default defineComponent({
  setup() {
    // 固定的数据（不可修改）
    const fixedData = '确定的数据';

    // 新数据（用户输入）
    const newData = ref('');

    // 提交逻辑
    const handleSubmit = () => {
      if (newData.value) {
        console.log('提交的新数据:', newData.value);
        // 这里可以执行数据修改逻辑，比如调用 API
        alert(`数据已修改为: ${newData.value}`);
        newData.value = ''; // 清空输入框
      } else {
        alert('请输入新数据');
      }
    };

    return {
      fixedData,
      newData,
      handleSubmit,
    };
  },
});
</script>

<style scoped>
/* 可以根据需要添加自定义样式 */
</style>