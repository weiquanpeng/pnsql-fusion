import request from "@/request/index.js"

export const GetSysMenus = ()=>{
    return request.get('/v1/public/menu/list')
}