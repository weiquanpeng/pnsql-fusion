<template>
  <el-dialog v-model="dialogFormVisible" title="修改密码" width="500">
    <el-form :model="form" style="padding-right: 40px">
      <el-form-item label="新密码" label-width="100">
        <el-input v-model="form.Password" autocomplete="off" placeholder="password" type="password" />
      </el-form-item>
      <el-form-item label="确认密码" label-width="100" required>
        <el-input v-model="form.newPassword" autocomplete="off" placeholder="password" type="password" @input="checkPasswords" />
      </el-form-item>
    </el-form>
    <template #footer>
      <div class="dialog-footer" style="display: flex;">
        <span class="password-message" v-show="showMismatchMessage">
          密码不一致
        </span>
        <div style="margin-left: auto;">
          <el-button @click="dialogFormVisible = false">取消</el-button>
          <el-button type="primary" :disabled="!passwordsMatch" @click="handleForm">
            确认
          </el-button>
        </div>
      </div>
    </template>
  </el-dialog>
</template>

<script setup>
import { ref, reactive, computed } from 'vue';
import {  UptSysUserFiled } from "@/api/sys_user.js";

const dialogFormVisible = ref(false);
const form = reactive({
  Password: '',
  newPassword: '',
});
const handleRow = reactive({
  id: '',
  field: { password: '' },
});

// 计算属性，判断新密码和确认密码是否一致
const passwordsMatch = computed(() => {
  return form.Password === form.newPassword;
});

// 计算属性，用于控制提示信息是否显示
const showMismatchMessage = computed(() => {
  return form.newPassword.length === form.Password.length && form.newPassword.length > 0 && !passwordsMatch.value;
});

const handleOpen = (id) => {
  handleRow.id = id;
  dialogFormVisible.value = true;
}

const handleForm = async () => {
  handleRow.field['password'] = form.newPassword;
  await UptSysUserFiled(handleRow);
  dialogFormVisible.value = false;
  ElMessage({message:"重置成功",type:"success",showClose:true})
}

// 校验密码是否一致的方法
const checkPasswords = () => {
  if (form.newPassword.length === form.Password.length) {
    return passwordsMatch.value;
  }
}

//将方法传给父组件
defineExpose({
  handleOpen
})

</script>


<style>
.el-dialog__footer {
  padding-top: 5px;
}

.password-message {
  font-style: italic;
  color: #d9534f;
  padding: 5px 6px; /* 减少内边距 */
  margin-left: 230px;
  background-color: #f9f2f2;
  border: 1px solid #d9534f;
  border-radius: 3px; /* 减小边角圆润度 */
  display: inline-block;
}

 .dialog-footer {
   display: flex;
   align-items: center; /* 使内容垂直居中 */
 }

.button-container {
  margin-left: auto; /* 将这个 div 右侧内容推到右边 */
}
</style>