<template>
  <div class="menu-demo">
    <a-menu
        :style="{ width: '150px', height: '648px' }"
        :default-open-keys="defaultOpenKeys"
        :selected-keys="[selectedKeys]"
        v-model:open-keys="openKeys"
        breakpoint="xl"
        :accordion="true"
        :collapsed="false"
    >
      <a-sub-menu :index="menu.path" v-for="menu in menuTreeData" :key="menu.id">
        <template #icon>
          <el-icon><component :is="menu.icon" /></el-icon>
        </template>
        <template #title>{{ menu.name }}</template>
        <a-menu-item
            v-for="(citem, cindex) in menu.children"
            :key="citem.id"
            :index="citem.path"
            @click="handleMenuItemClick(citem.path, menu.id)"
        >
          {{ citem.name }}
        </a-menu-item>
      </a-sub-menu>
    </a-menu>
  </div>
</template>

<script>
import { computed, ref, watch } from 'vue';
import { useRoute } from 'vue-router';
import router from "@/router/index.js";
import { useCounterStore } from "@/stores/user.js";
import { useMenuTabStore } from "@/stores/menuTab.js";
import {
  IconMenuFold,
  IconMenuUnfold,
  IconApps,
  IconBug,
  IconBulb,
} from '@arco-design/web-vue/es/icon';

export default {
  setup() {
    const route = useRoute();
    const useStore = useCounterStore();
    const useMenuTabs = useMenuTabStore();

    // 获取菜单树数据
    const menuTreeData = computed(() => useStore.getMenuTree());

    // 根据路由地址查找父子菜单 ID
    const findMenuIds = () => {
      const currentPath = route.path; // 获取当前路由地址
      return useMenuTabs.findMenuIdsByPath(currentPath);
    };

    // 默认展开的菜单项（响应式变量）
    const defaultOpenKeys = ref([]);

    // 当前选中的菜单项（响应式变量）
    const selectedKeys = ref("");

    // 动态展开的菜单项（响应式变量）
    const openKeys = ref([]);

    // 初始化菜单状态
    const initMenuState = () => {
      const menuIds = findMenuIds();
      if (menuIds) {
        defaultOpenKeys.value = [menuIds.parentId]; // 设置默认展开的父菜单
        selectedKeys.value = menuIds.childId; // 设置当前选中的子菜单
      }
    };

    // 初始化
    initMenuState();

    // 监听路由变化，动态更新菜单状态和标签页状态
    watch(
        () => route.path,
        (newPath) => {
          const menuIds = useMenuTabs.findMenuIdsByPath(newPath);
          if (menuIds) {
            defaultOpenKeys.value = [menuIds.parentId]; // 更新默认展开的父菜单
            selectedKeys.value = menuIds.childId; // 更新当前选中的子菜单

            // 更新标签页状态
            const menuName = menuTreeData.value
                .find(menu => menu.id === menuIds.parentId)
                .children.find(child => child.id === menuIds.childId).name;
            useMenuTabs.addMenuTab(menuName, newPath, menuIds.childId, menuIds.parentId);

            // 动态设置 openKeys，确保菜单栏展开
            openKeys.value = [menuIds.parentId];
          }
        },
        { immediate: true } // 立即执行一次，确保初始化时同步状态
    );

    // 处理菜单项点击事件
    const handleMenuItemClick = (path, parentId) => {
      // 设置 openKeys 为父菜单 ID，展开对应的父菜单
      openKeys.value = [parentId];

      // 跳转到对应的路由
      router.push(path);
    };

    return {
      menuTreeData,
      defaultOpenKeys,
      selectedKeys,
      openKeys,
      handleMenuItemClick,
    };
  },
  components: {
    IconMenuFold,
    IconMenuUnfold,
    IconApps,
    IconBug,
    IconBulb,
  },
};
</script>

<style>
.menu-demo {
  box-sizing: border-box;
  height: 648px;
  background-color: var(--color-neutral-2);
}
</style>