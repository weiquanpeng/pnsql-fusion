<template>
  <div style="display: flex; justify-content: space-between; align-items: center; padding-bottom: 10px;">
    <!-- 靠左固定的 Radio 按钮组 -->
    <div style="padding-left: 40px;">
      <a-radio-group v-model="position" type="button" @change="handleRadioChange">
        <a-radio value="mine">我的工单</a-radio>
        <a-radio value="approve">我的审批</a-radio>
        <a-radio value="all">所有工单</a-radio>
      </a-radio-group>
    </div>
    <div style="padding-right: 40px;">
      <a-input-group>
        <a-select
            v-model="selectedColumn"
            :options="columnOptions"
            :style="{ width: '160px' }"
            placeholder="选择列"
        />
        <a-input
            v-model:model-value="searchValue"
            :style="{ width: '320px' }"
            placeholder="输入搜索关键词"
            search-button
            :loading="false"
            :table-layout="'fixed'"
        />
      </a-input-group>
    </div>
  </div>

  <div> <!-- 确保外层 div 高度固定 -->
    <a-table
        :stripe="true"
        :data="filteredTaskConfigs"
        :loading="loadStatus"
        :pagination="pagination"
    >
      <template #columns>
        <a-table-column title="ID" data-index="id" :width=100 align="center">
          <template #cell="{ record }">
            <div
                @click="handleIdOpen(record.id, record.title)"
                style="color: blue; cursor: pointer; font-size: 16px; padding: 8px; display: inline-block;">
              {{ record.id }}
            </div>
          </template>
        </a-table-column>
        <a-table-column title="标题" data-index="title" align="center"></a-table-column>
        <a-table-column title="提交人" data-index="owner" align="center"></a-table-column>
        <a-table-column title="状态" data-index="status" :width=60 align="center">
          <template #cell="{ record }">
      <span v-if="record.status === 'todo'">
        <icon-play-arrow-fill :style="{ fontSize: '25px', color: 'green' }" />
      </span><span v-else-if="record.status === 'doing'">
        <icon-refresh :style="{ fontSize: '25px', color: 'blue' }" spin />
      </span><span v-else-if="record.status === 'error'">
        <icon-exclamation-circle-fill :style="{ fontSize: '25px', color: 'red' }" />
      </span><span v-else-if="record.status === 'done'">
        <icon-check-circle-fill :style="{ fontSize: '25px', color: 'black' }" />
      </span><span v-else-if="record.status === 'close'">
        <icon-close-circle-fill :style="{ fontSize: '25px', color: 'grey' }" />
      </span>
            <span v-else>
        {{ record.status }}
      </span>
          </template>
        </a-table-column>
        <a-table-column title="描述" data-index="describe" align=center :ellipsis="true" :tooltip="true"></a-table-column>
        <a-table-column title="创建时间" data-index="createdAt" :minWidth=300 align="center" :sortable="{sortDirections: ['ascend', 'descend']}"></a-table-column>
        <a-table-column title="更新时间" data-index="updatedAt" :minWidth=300 align="center" :sortable="{sortDirections: ['ascend', 'descend']}"></a-table-column>
        <a-table-column title="操作" align="center">
          <template #cell="{ record }">
            <a-button
                type="primary"
                shape="round"
                size="small"
                hoverable
                @click="handleIdOpen(record.id, record.title)"
                style="margin: 0 4px;"
            >
              <template #icon>
                <icon-search />
              </template>
              详情
            </a-button>
          </template>
        </a-table-column>
      </template>
    </a-table>
  </div>
  <task-process ref="taskProcess" @callParentMethod="handleRadioChange" @refreshTable="refreshTableData"></task-process>
</template>

<script setup>
import {loadtaskconfig, loadOwnerTaskConfig, loadApproveTaskConfig, getSubTaskData} from "@/api/task_config.js";
import {useCounterStore} from "@/stores/user.js";
import TaskProcess from "@/views/components/TaskProcess.vue";
import formatDateTime from "@/utils/formatDateTime.js";

const useStore = useCounterStore();

// 原始数据存储
const taskConfigs = reactive([]);
const loadStatus = ref(false);
const pagination = {pageSize: 100};
const tableHeight = ref("60vh"); // 初始为50vh
const position = ref(localStorage.getItem('selectedPosition') || 'mine'); // 从 localStorage 中读取选中的工单类型
const subTaskList = ref([]);
const columnOptions = [
  {label: 'ID', value: 'id'},
  {label: '标题', value: 'title'},
  {label: '提交人', value: 'owner'},
  {label: '状态', value: 'status'},
  {label: '描述', value: 'describe'}
];

const selectedColumn = ref('id'); // 将默认值设置为 'id'
const searchValue = ref(''); // 当前输入的搜索词
const taskProcess = ref(null); // 初始化为 null

// 处理 ID 点击事件
const handleIdOpen = (id, taskTitle) => {
  // 第一次调用：打开抽屉并显示骨架屏
  taskProcess.value.handleClick([], '', taskTitle);

  // 获取子任务数据
  getSubTaskData(id)
      .then(response => {
        // 第二次调用：更新数据
        taskProcess.value.handleClick(response.data.data, response.data.logs, taskTitle);
      })
      .catch(error => {
        console.error('加载子任务数据失败:', error);
      });
};

// 处理 Radio 按钮组变化
const handleRadioChange = (newValue) => {
  localStorage.setItem('selectedPosition', newValue); // 将选中的工单类型存储到 localStorage
  taskConfigs.length = 0;
  switch (newValue) {
    case 'mine':
      fetchOwnerTaskConfig();
      break;
    case 'approve':
      fetchIdTaskConfig();
      break;
    case 'all':
      fetchTaskConfig();
      break;
  }
};

// 刷新表格数据
const refreshTableData = () => {
  taskConfigs.length = 0; // 清空当前数据
  handleRadioChange(position.value); // 重新加载数据
};

// 过滤任务配置
const filteredTaskConfigs = computed(() => {
  if (selectedColumn.value && searchValue.value) {
    return taskConfigs.filter(item => {
      const value = item[selectedColumn.value];
      return value?.toString().includes(searchValue.value);
    });
  }
  return taskConfigs;
});

// 获取任务配置
const fetchTaskConfigs = (apiCall) => {
  loadStatus.value = true;
  apiCall()
      .then(response => {
        response.data.list.forEach((item) => {
          taskConfigs.push({
            id: item.id,
            title: item.title,
            owner: item.owner,
            status: item.status,
            parameter: item.parameter,
            describe: item.describe,
            createdAt: formatDateTime(item.createdAt),
            updatedAt: formatDateTime(item.updatedAt)
          });
        });
      })
      .catch(error => {
        console.error('Error fetching data:', error);
      })
      .finally(() => {
        loadStatus.value = false;
      });
};

const fetchTaskConfig = () => {
  fetchTaskConfigs(loadtaskconfig);
};
const fetchIdTaskConfig = () => {
  fetchTaskConfigs(() => loadApproveTaskConfig(useStore.user));
};
const fetchOwnerTaskConfig = () => {
  fetchTaskConfigs(() => loadOwnerTaskConfig(useStore.user));
};

// 更新表格高度
const updateTableHeight = () => {
  const zoomLevel = (window.outerWidth / window.innerWidth) * 100;

  if (zoomLevel > 89 && zoomLevel < 91) {
    tableHeight.value = '67vh';
  } else if (zoomLevel > 80) {
    tableHeight.value = '64vh';
  } else {
    tableHeight.value = '70vh';
  }
};

// 组件挂载时初始化
onMounted(() => {
  const savedPosition = localStorage.getItem('selectedPosition') || 'mine';
  position.value = savedPosition;
  handleRadioChange(savedPosition); // 根据存储的值初始化数据
  updateTableHeight();
  window.addEventListener('resize', updateTableHeight);
});

// 组件卸载时清理
onBeforeUnmount(() => {
  window.removeEventListener('resize', updateTableHeight);
});
</script>