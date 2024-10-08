<template>
  <el-drawer v-model="drawer" title="用户注册" :with-header="true" :size="500">
    <el-form :model="addUser" :rules="rules" ref="userForm" label-width="auto" style="display: flex; flex-direction: column; justify-content: center; align-items: center;">
      <div style="display: flex; justify-content: center; align-items: center; margin-bottom: 10px;">
        <el-form-item label="用户" prop="account">
          <el-input v-model="addUser.account" placeholder="username" style="width: 150px; padding-right: 20px" />
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input v-model="addUser.password" placeholder="password" style="width: 150px" />
        </el-form-item>
      </div>
      <div style="display: flex; justify-content: center; align-items: center;">
        <el-form-item label="电话">
          <el-input v-model="addUser.phone" placeholder="......" style="width: 150px; padding-right: 20px" />
        </el-form-item>
        <el-form-item label="邮箱">
          <el-input v-model="addUser.email" placeholder="...@xxx.com" style="width: 150px" />
        </el-form-item>
      </div>
      <div style="display: flex; flex-direction: column;">
        <el-form-item label="权限" prop="roles" style="padding-top: 14px;padding-right: 50px">
          <template #default="row">
          <el-cascader
              v-model="addUser.roles"
              :options="rolesData"
              :props="props"
              prop="roles"
              placeholder="请选择账号"
              style="width: 300px;"
          />
          </template>
        </el-form-item>
        <el-form-item label="激活" style="padding-top: 10px">
          <el-switch v-model="addUser.enable" :active-value="1" :inactive-value="2"></el-switch>
        </el-form-item>
      </div>
    </el-form>
    <div class="button-container">
      <el-button type="primary" :loading="suLoading" @click="userCommit">保存</el-button>
      <el-button @click="drawer = false">取消</el-button>
    </div>
  </el-drawer>
</template>

<script setup>
import {GetSysRoles} from "@/api/sys_roles.js";
import {AddSysUser} from "@/api/sys_user.js";
import throttle from "@/utils/throttle.js";

const drawer = ref(false);
const suLoading = ref(false)
const rolesData = ref([]);
const addUser = reactive({
  account: "",
  password: "",
  phone: "",
  email: "",
  enable: 1,
  roles: ""
});

// 定义验证规则
const rules = {
  account: [
    { required: true, message: '请输入用户名', trigger: 'blur' }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: 'blur' }
  ],
  roles: [
    { type: 'array', required: true, message: '请选择权限', trigger: 'change' }
  ]
};

const userForm = ref(null);

const AddUsersOnSubmit = () => {
  addUser.account = '';
  addUser.password = '';
  addUser.phone = '';
  addUser.email = '';
  addUser.enable = 1;
  addUser.roles = "";
  drawer.value = true
}

const emit = defineEmits(['callParentMethod']);   //父组件方法
const userCommit = () => {
  throttle(()=>{
    suLoading.value = true
    userForm.value.validate(async (valid) => {
      if (valid) {
        addUser.roles = addUser.roles.map(Number);
        addUser.roles.sort((a, b) => a - b);
        try {
          await AddSysUser(addUser);
          suLoading.value = false;
          drawer.value = false; // 关闭抽屉
          emit('callParentMethod');
          ElMessage({ message: "添加成功", type: "success", showClose: true });
        } catch (error) {
          console.error("添加用户失败:", error);
          suLoading.value = false;
          drawer.value = false;
        }
      } else {
        suLoading.value = false;
        return false; // 表单验证失败
      }
    });
  },2000)
};

const props = {
  multiple: true,
  emitPath: false,
  label: 'roleName', // 用于显示的文本
  value: 'roleCode' // 用于存储的值
};

const formSysRolesList = async () => {
  const resp = await GetSysRoles();
  rolesData.value = resp.data
};

onMounted(() => {
  formSysRolesList();
});

defineExpose({
  AddUsersOnSubmit
})

</script>

<style scoped>
.button-container {
  display: flex;
  padding-top: 300px;
  padding-left: 40px;
}
</style>