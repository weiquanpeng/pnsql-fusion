import request from "@/request/index.js"

export const getCaptchaApi=()=>{
    return request.get("/v1/public/login/get")
}

