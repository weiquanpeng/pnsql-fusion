import request from "@/request/index.js"

export const GetSysRoles = ()=>{
    return request.get('/v1/sys/roles/list')
}

