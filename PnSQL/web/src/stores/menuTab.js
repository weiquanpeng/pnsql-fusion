import router from '@/router/index.js'
export const useMenuTabStore = defineStore('menuTab', {
    state: () => ({
        addTabsValue:'/index/dashboard',
        menuTableTabs : [{name:'首页',path:'/index/dashboard'}],
        selectKeys:"",
        openKeys:""
    }),
    getters:{
    },
    actions: {
        getAddTabsValue(){
            return this.addTabsValue
        },
        getSelectKeys(){
            return this.selectKeys
        },
        getOpenKeys(){
            return this.openKeys
        },
        addMenuTab(name, path,selID,openID) {
            this.selectKeys=selID
            this.openKeys=openID
            var index = this.menuTableTabs.findIndex(m => m.path === path);
            if (index === -1) {
                // 如果已经有5个标签页，则替换掉最后一个
                if (this.menuTableTabs.length >= 6) {
                    // 替换第五个标签
                    this.menuTableTabs[5] = { name, path }; // 直接替换第五个标签
                } else {
                    this.menuTableTabs.push({ name, path }); // 添加新标签
                }
            }
            this.addTabsValue = path;
        },


        handleTabClick(path){
            this.addTabsValue=path
            router.push(path)
        },
        removeMenuTab(path) {
            // 检查是否只剩最后一个标签
            if (this.menuTableTabs.length === 1) {
                return; // 直接返回，不进行删除
            }
            var index = this.menuTableTabs.findIndex(m => m.path === path);
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
    persist:{
        key:'pwq-pinia-menuTabs',
        storage:localStorage,
    }
})

