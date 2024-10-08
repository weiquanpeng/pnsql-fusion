<template>
  <div class="el-pwq-header">
    <a-page-header
        :style="{ background: 'var(--color-bg-2)'}"
        title="PnSql"
        subtitle="Devops"
        :show-back="false"
        class="arco-layout-sider-children"
        style="padding: 0;height: 60px; "
    >
      <template #breadcrumb>
        <a-breadcrumb style="display: flex;justify-content: flex-start; margin-left: 30px;font-style: italic;">
          <a-breadcrumb-item>
            <icon-home/>
          </a-breadcrumb-item>
          <a-breadcrumb-item v-if="route.path.substring(0,route.path.indexOf('/',2))">
            <a>{{useStore.getMenuTree().find(obj=>obj.path===route.path.substring(0,route.path.indexOf('/',2)))?.name || '没有主路由'}}</a>
          </a-breadcrumb-item>
          <a-breadcrumb-item v-if="route.path.substring(route.path.indexOf('/', 1))">
            <a>{{useStore.getMenuTree().find(obj => obj.path === route.path.substring(0, route.path.indexOf('/', 2)))?.children?.find(child => child.path === route.path.substring(route.path.indexOf('/', 0)))?.name || '没有子路由'}}</a>
          </a-breadcrumb-item>
        </a-breadcrumb>
      </template>
    </a-page-header>
  </div>

  <div style="position: absolute; top: 20px; right: 30px; ">
    <el-dropdown @command="handleCommand">
    <span class="pwq-el-dropdown-link">
      {{ useStore.user }}
      <el-icon class="el-icon--right">
        <arrow-down />
      </el-icon>
    </span>
      <template #dropdown>
        <el-dropdown-menu>
          <el-dropdown-item command="toPersonal">个人中心</el-dropdown-item>
          <el-dropdown-item command="toLogout">注销</el-dropdown-item>
        </el-dropdown-menu>
      </template>
    </el-dropdown>
    </div>
</template>

<script setup>
import {
  IconHome
} from '@arco-design/web-vue/es/icon';
import {useCounterStore} from "@/stores/user.js";
import {useRoute} from "vue-router";
import {ArrowDown} from "@element-plus/icons-vue";
import PVA from "@/utils/pwq_Message.js";
const useStore = useCounterStore()
const route = useRoute()

const commands = ({
  toPersonal:()=>{
    ElMessage(`click on item 个人中心`)
  },
  toLogout:()=>{
    PVA.confirm("注销","确认退出登录","warning","注销成功").then(res=>{
      useStore.toLogout()
    })
  }
});


const handleCommand = (command) => {
  commands[command] && commands[command]()
}

</script>

<style>
.arco-layout-sider-children{
  font-style: italic;
}
.arco-dropdown-open .arco-icon-down {
  transform: rotate(180deg);
}

.el-pwq-header{
  position: absolute;
  width: 500px;
  left: 10px;
}

.pwq-el-dropdown-link:focus {
  outline: none;
}


</style>
