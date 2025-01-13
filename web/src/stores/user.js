import request from "@/request";
import router from "@/router/index.js";
import {handleLogout} from "@/api/logout.js";
import {GetSysMenus} from "@/api/sys_menu.js";

export const useCounterStore = defineStore('user', {
    state: () => ({
        user:'',
        token:'',
        roles:[],
        menuTree:[],
    }),
    actions: {
        setToken(newToken){
            this.token=newToken
        },
        getToken(){
            return this.token
        },
        getMenuTree(){
            return this.menuTree
        },
        async toLogin(loginUser) {
            const resp = await request.post("/v1/public/login/post",loginUser)
            this.user = resp.data.user
            this.setToken(resp.data.token)
            this.roles = resp.data.roles
            const menu = await GetSysMenus()
             // 过滤不具备权限菜单
            this.menuTree = menu.data.map(item => {
                const filteredChildren = item.children ? item.children.filter(child => resp.data.menus.includes(child.id)) : [];
                return {
                    ...item,
                    children: filteredChildren
                };
            }).filter(item => item.children.length > 0)
        },
        async toLogout() {
            await handleLogout()
            // localStorage.clear()     //清空全部缓存
            this.$reset()
            localStorage.removeItem("pwq-pinia-restore")
            localStorage.removeItem("pwq-pinia-menuTabs")
            router.push({name: 'login', replace: true})
            // window.location.reload()     //重新刷新页面
        },
    },
    //pinia 持久化缓存
    persist:{
        key:'pwq-pinia-restore',
        storage:localStorage,
    }
})

