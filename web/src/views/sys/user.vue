<template>
  <div class="pwq-search">
    <el-form :inline="true" :model="queryParams" size="small">
      <el-form-item>
        <el-button type="primary" @click="onAdd" :icon="Plus" size="small" style="margin-left: 5px;">添加</el-button>
        <el-button type="danger" @click="onDelete" :icon="Delete" size="small" v-if="useStore.roles.includes(1)">删除</el-button>
      </el-form-item>
      <el-form-item label="关键词">
        <el-input v-model="queryParams.keyword" placeholder="请输入用户" maxlength="12" @input="onSearch" clearable />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch" :icon="Search" size="small">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="pwq-tables">
    <el-table :data="tableData"
              style="width: 100%; height: calc(100vh - 194px)"
              :row-style="{ height: `40px` }"
              size="small"
              border
              stripe
              empty-text=" "
              v-loading="loading"
              element-loading-text="加载中..."
              element-loading-background="rgba(255, 255, 255, 0.1)"
    >
      <el-table-column type="selection" fixed="left" width="50" />
      <el-table-column fixed prop="account" label="用户" width="100"  align="center" />
      <el-table-column prop="phone" label="电话" width="160px"  align="center" />
      <el-table-column prop="email" label="邮箱" width="180px"  align="center" />
      <el-table-column prop="createdAt" label="创建时间" align="center">
        <template #default="{ row }">
          {{ formatDate(row.createdAt) }}
        </template>
      </el-table-column>
      <el-table-column prop="updatedAt" label="更新时间" align="center">
        <template #default="{ row }">
          {{ formatDate(row.updatedAt) }}
        </template>
      </el-table-column>
      <el-table-column fixed="right" label="操作" align="center">
        <template #default="{ row }">
          <el-button text icon="edit" type="primary" size="small">编辑</el-button>
          <el-button text icon="remove" type="danger" size="small" @click="handleDel(row)" style="margin: 0">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
  <div class="pwq-block">
    <el-config-provider :locale="zhCn">
      <el-pagination
          v-model:page-size="queryParams.pageSize"
          v-model:current-change="queryParams.page"
          class="demonstration"
          size="small"
          :page-sizes="[20, 40, 80, 200]"
          layout="total,prev, pager, next, sizes"
          :total="total"
          @current-change="handlePageData"
          @size-change="handleSizeData"
      />
    </el-config-provider>
  </div>
  <add-users ref="addUser" @callParentMethod="handleLoadSysUserData"></add-users>
</template>

<script setup>
import { Plus, Delete, Search } from "@element-plus/icons-vue";
import PVA from "@/utils/pwq_Message.js";
import {DelSysUser, LoadSysUserData, UptSysUser, UptSysUserFiled} from "@/api/sys_user.js";
import {zhCn} from "element-plus/es/locale/index";
import AddUsers from "@/views/components/AddUsers.vue";
import {useCounterStore} from "@/stores/user.js";
const queryParams = reactive({
  page: 1,
  pageSize: 20,
  keyword: "",
});

const tableData = ref([]);
const total = ref(0);
const loading = ref(false);
const useStore = useCounterStore()

//添加用户
const  addUser = ref({})
const onAdd = async () => {
  addUser.value.AddUsersOnSubmit()
};

//批量删除用户
const onDelete = () => {
  PVA.confirm("提示", "确认删除", "warning", "删除成功");
};


const onSearch = () => {
  queryParams.page = 1
  queryParams.total = 0
  queryParams.pageSize = 20
  handleLoadSysUserData()
};

//查询列表
const handleLoadSysUserData = async () => {
  tableData.value=[]
  loading.value = true; // 请求开始前设置加载状态为 true
  try {
    const resp = await LoadSysUserData(queryParams);
    tableData.value = resp.data.list;
    total.value = resp.data.total;
    queryParams.pageSize = resp.data.pageSize;
  } catch (error) {
    // 处理错误，例如显示消息
    console.error(error);
  } finally {
    loading.value = false; // 请求完成后设置加载状态为 false
  }
};

//右边栏删除
const handleDel = async (row) => {
  const user = `确认删除用户： ${row.account}`
  PVA.confirm("提示", user, "warning", "删除成功")
      .then(async () => {
        await DelSysUser(row.id);
        handleLoadSysUserData()
      })
      .catch(() => {
        ElMessage('取消操作');
      });
};

//分页选择
const handlePageData = (pSize) => {
  queryParams.page=pSize
  handleLoadSysUserData()
};
const handleSizeData = (pageSize) => {
  queryParams.pageSize=pageSize
  handleLoadSysUserData()
};

//时间字段显示
const formatDate = (dateString) => {
  const date = new Date(dateString);
  const year = date.getFullYear();
  const month = String(date.getMonth() + 1).padStart(2, '0');
  const day = String(date.getDate()).padStart(2, '0');
  const hours = String(date.getHours()).padStart(2, '0');
  const minutes = String(date.getMinutes()).padStart(2, '0');
  const seconds = String(date.getSeconds()).padStart(2, '0');
  return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
};

//监听变化
watch(() => queryParams.keyword, (newKeyword, oldKeyword) => {
  onSearch();
}, { immediate: true });

onMounted(() => {
  handleLoadSysUserData();
});
</script>

<style scoped>
.pwq-search{
  height: 30px;
}
</style>