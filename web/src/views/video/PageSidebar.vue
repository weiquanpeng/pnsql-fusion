<template>
  <div class="menu-demo" >
    <a-menu
        :style="{ width: '150px', height: '648px' }"
        :default-open-keys=[selectKeys()]
        :selected-keys=[openKeys()]
        breakpoint="xl"
        :accordion="true"
        :collapsed="false"
    >
      <a-sub-menu :index="menu.path" v-for="menu in menuTreeData()" :key="menu.id">
        <template #icon>
          <el-icon><component :is="menu.icon"/></el-icon>
        </template>
        <template #title>{{menu.name}}</template>
        <a-menu-item v-for="(citem,cindex) in menu.children" :key="citem.id" @click="navigateToMenu(citem.path,citem.name,menu.id,citem.id)">{{ citem.name }}</a-menu-item>
      </a-sub-menu>
    </a-menu>
  </div>
</template>
<script>
import router from "@/router/index.js";
import {useCounterStore} from "@/stores/user.js";
const useMenuTabs = useMenuTabStore()
const useStore = useCounterStore()


import {
  IconMenuFold,
  IconMenuUnfold,
  IconApps,
  IconBug,
  IconBulb,
} from '@arco-design/web-vue/es/icon';
import {useMenuTabStore} from "@/stores/menuTab.js";

export default {
  methods: {
    menuTreeData() {
      return useStore.getMenuTree()
    },
    selectKeys(){
      return useMenuTabs.getSelectKeys()
    },
    openKeys(){
      return useMenuTabs.getOpenKeys()
    },
  },
  components: {
    IconMenuFold,
    IconMenuUnfold,
    IconApps,
    IconBug,
    IconBulb,
  },
  setup() {
    const r = router; // 确保正确导入 useRouter
    const navigateToMenu = async (path, name, selID, openID) => {
      useMenuTabs.addMenuTab(name, path, selID, openID);
      try {
        await r.push(path);
      } catch (error) {
        console.error('Navigation failed', error);
      }
    };
    return {
      navigateToMenu,
    };
  }
};

</script>
<style>
.menu-demo {
  box-sizing: border-box;
  height: 648px;
  background-color: var(--color-neutral-2);
}


</style>