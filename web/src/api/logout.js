import request from "@/request/index.js"

export const handleLogout = ()=>{
    return request.post("/v1/public/logout/post")
}

