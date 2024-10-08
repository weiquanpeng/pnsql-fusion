import request from "@/request/index.js"

export const GetSysUser = (id)=>{
    return request.get('/v1/sys/user/post/$(id)')
}

export const AddSysUser = (user)=>{
    return request.post('/v1/sys/user/addUser',user)
}

export const UptSysUser = (row)=>{
    return request.post('/v1/sys/user/uptUser',row)
}
export const UptSysUserFiled = (params)=>{
    return request.post('/v1/sys/user/uptFiled',params)
}

export const DelSysUser = (id)=>{
    return request.post(`/v1/sys/user/delUser?id=${id}`)
}

export const LoadSysUserData = (params)=>{
    return request.post("/v1/sys/user/load",params)
}

