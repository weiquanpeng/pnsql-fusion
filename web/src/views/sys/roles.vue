<template>
  <div class="pwq-search2">
    <el-form :inline="true" :model="queryParams" size="small">
      <el-form-item>
        <el-button type="primary" @click="onAdd" icon="edit" size="small" style="margin-left: 5px">批量修改</el-button>
      </el-form-item>
      <el-form-item label="关键词">
        <el-input v-model="queryParams.keyword" placeholder="请输入用户" maxlength="12" @keydown.enter="onSearch" clearable />
      </el-form-item>
      <el-form-item label="关键词">
        <el-input v-model="queryParams.data" placeholder="请输入权限" maxlength="12" @keydown.enter="onSearch" clearable />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSearch" :icon="Search" size="small">查询</el-button>
      </el-form-item>
    </el-form>
  </div>
  <div class="pwq-tables-roles" style="padding-top: 5px">
    <el-table
        :data="tableData"
        style="width: 100%; height: calc(100vh - 194px)"
        :size="'small'"
        :row-style="{ height: `40px` }"
        empty-text=" "
        v-loading="loading"
        element-loading-text="加载中..."
        element-loading-background="rgba(255, 255, 255, 0.1)"
    >
      <el-table-column fixed prop="account" label="用户"  align="center" width="200"/>
      <el-table-column label="是否启用"  align="center" width="300px">
        <template #default="status">
          <el-switch v-model="status.row.enable" :active-value="1" :inactive-value="2" active-text="启用" size="small" @change="statusEnable(status.row, status.row.enable)"/>
        </template>
      </el-table-column>
      <el-table-column label="用户权限" align="center" width="280px">
        <template #default="{row,$index}">
        <el-cascader
            v-model="row.roles"
            :options="rolesData"
            :props="props"
            @change="handleChangeRole(row)"
            size="small"
            popper-class="ssaassd"
            style="width: 210px;"
            placeholder="选择权限"
        />
        </template>
      </el-table-column>

      <el-table-column fixed="right" label="操作" align="center" min-width="200px">
        <template #default="{ row }">
          <el-button text icon="edit" type="primary" size="small">授予权限</el-button>
          <el-button text icon="switch" type="success" size="small" @click="handleReset(row.id)" style="margin: 0">重置密码</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>

  <div class="pwq-block2">
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
  <UpdateDialog ref="diaLog"></UpdateDialog>
</template>

<script setup>
import { ref, reactive, onMounted } from 'vue';
import { Plus, Delete, Search } from "@element-plus/icons-vue";
import PVA from "@/utils/pwq_Message.js";
import {DelSysUser, LoadSysUserData, UptSysUser, UptSysUserFiled} from "@/api/sys_user.js";
import {GetSysRoles} from "@/api/sys_roles.js";
import {zhCn} from "element-plus/es/locale/index";
import UpdateDialog from "@/views/components/UpdatePassWord.vue";

const queryParams = reactive({
  page: 1,
  pageSize: 20,
  keyword: "",
});

const tableData = ref([]);
const rolesData = ref([]);
const total = ref(0);
const loading = ref(false);

//新增按钮
const onAdd = () => {
  PVA.confirm("提示", "确认删除", "warning", "删除成功");
};

//删除按钮
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
  loading.value = true;
  const resp = await LoadSysUserData(queryParams);
  tableData.value = resp.data.list;
  total.value = resp.data.total;
  queryParams.pageSize = resp.data.pageSize;
  loading.value = false;
};

//查询权限
const props = ref({multiple: true,emitPath: false,label:'roleName',value:'id'})
const handleSysRolesList = async () => {
  const resp = await GetSysRoles();
  rolesData.value = resp.data
};

//状态修改
const statusEnable = async (row, newValue) => {
  if (newValue === 2){
    row.enable = 2
    await UptSysUser(row)
  }else {
    row.enable = 1
    await UptSysUser(row)
  }
};

//编辑权限
const handleChangeRole = async (row) => {
  row.roles.sort((a, b) => a - b);
  const handleRow = reactive({
    id: '',
    field: { roles: "" },
  });
  const stringRoles = JSON.stringify(row.roles);
  handleRow.id=row.id
  handleRow.field.roles=stringRoles
  await UptSysUserFiled(handleRow)
}

//右边栏重置密码
const diaLog = ref({})    //获取对话框组件方法
const handleReset = async (id) => {
  diaLog.value.handleOpen(id)
}

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

onMounted(() => {
  handleLoadSysUserData();
  handleSysRolesList();
});
</script>

<style scoped>
.pwq-search2{
  height: 35px;
  border-bottom: 1px solid #e7e7e7;
}

.ssaassd{
  .el-cascader-panel {
    padding: 0;
    height: 100px;
    .el-cascader-menu {font-size: 12px;height: 100px;}
    .el-cascader-node{padding-left: 8px;height: 30px}
  }
}

.cell{
  width: 240px;
  .el-cascader.el-cascader--small.el-tooltip__trigger.el-tooltip__trigger{
    width: 200px;
  }
  .el-icon.el-input__icon.icon-arrow-down{
    margin: 0;
  }
}
.pwq-block2{
  margin-top: 5px;
  margin-right: 40px;
  display: flex;
  justify-content: right;
}

</style>