import { defineStore } from 'pinia';
import router from '@/router/index.js';
import { useCounterStore } from '@/stores/user.js'; // 引入 useCounterStore

export const useMenuTabStore = defineStore('menuTab', {
    state: () => ({
        addTabsValue: '/index/dashboard', // 当前选中的标签页路径
        menuTableTabs: [{ name: '首页', path: '/index/dashboard' }], // 标签页列表
        selectKeys: "", // 当前选中的菜单项
        openKeys: ""    // 当前展开的菜单项
    }),
    actions: {
        // 根据路由地址查找父菜单和子菜单 ID
        findMenuIdsByPath(path) {
            const useStore = useCounterStore(); // 获取 useCounterStore 实例
            const menuTree = useStore.getMenuTree(); // 获取菜单树数据

            // 遍历菜单树，查找匹配的路由地址
            for (const parentMenu of menuTree) {
                for (const childMenu of parentMenu.children) {
                    if (childMenu.path === path) {
                        return {
                            parentId: parentMenu.id, // 父菜单 ID
                            childId: childMenu.id    // 子菜单 ID
                        };
                    }
                }
            }
            return null; // 如果未找到匹配的路由，返回 null
        },
        // 获取当前选中的标签页路径
        getAddTabsValue() {
            return this.addTabsValue;
        },
        // 获取当前选中的菜单项
        getSelectKeys() {
            return this.selectKeys;
        },
        // 获取当前展开的菜单项
        getOpenKeys() {
            return this.openKeys;
        },
        // 添加标签页
        addMenuTab(name, path, selID, openID) {
            this.selectKeys = selID;
            this.openKeys = openID;
            const index = this.menuTableTabs.findIndex(m => m.path === path);
            if (index === -1) {
                // 如果已经有5个标签页，则替换掉最后一个
                if (this.menuTableTabs.length >= 6) {
                    this.menuTableTabs[5] = { name, path }; // 直接替换第五个标签
                } else {
                    this.menuTableTabs.push({ name, path }); // 添加新标签
                }
            }
            this.addTabsValue = path;
        },
        // 处理标签点击事件
        handleTabClick(path) {
            this.addTabsValue = path;
            router.push(path);
        },
        // 移除标签页
        removeMenuTab(path) {
            // 检查是否只剩最后一个标签
            if (this.menuTableTabs.length === 1) {
                return; // 直接返回，不进行删除
            }
            const index = this.menuTableTabs.findIndex(m => m.path === path);
            if (index > -1) {
                this.menuTableTabs.splice(index, 1); // 删除标签

                // 跳转到删除标签的前一个标签
                this.addTabsValue = index === this.menuTableTabs.length ?
                    this.menuTableTabs[index - 1].path :
                    this.menuTableTabs[index].path;
                router.push(this.addTabsValue);
            }
        }
    },
    persist: {
        key: 'pwq-pinia-menuTabs',
        storage: localStorage,
    }
});