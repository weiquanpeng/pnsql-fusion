<template>
  <el-tabs
      v-model="currentTabsValue"
      @tab-click="handleTabClick"
      type="border-card"
      class="demo-tabs"
      closable
      @tab-remove="removeTab"
      style="--el-tabs-header-height: 30px ;"
  >
    <el-tab-pane
        v-for="item in menuList"
        :key="item.path"
        :label="item.name"
        :name="item.path"
    >
    </el-tab-pane>
  </el-tabs>
</template>

<script setup>
import {useRoute} from "vue-router";
import {useMenuTabStore} from "@/stores/menuTab.js";
const route = useRoute()
const menuTabs = useMenuTabStore()
// const addTabsValue = computed(()=>menuTabs.addTabsValue);
const currentTabsValue = ref(menuTabs.addTabsValue);
const menuList = computed(()=>menuTabs.menuTableTabs);
const maxTabs = 5

const handleTabClick = (tab) => {
  menuTabs.handleTabClick(tab.props.name)
}

watch(
    () => menuTabs.addTabsValue,  // 监视 menuTabs.addTabsValue
    (newValue) => {
      currentTabsValue.value = newValue;  // 更新 currentTabsValue
    }
);

const removeTab = (path) => {
  menuTabs.removeMenuTab(path)
}
</script>

<style>
.demo-tabs{
  margin-bottom: 10px;
  .el-tabs__content{
    padding: 0;
  }
}
</style>
