export default {
    method:"get",
    baseURL: import.meta.env.VITE_BASE_PATH + import.meta.env.VITE_BASE_API,
    timeout: 10000,
    headers:{
        'Content-Type':'application/json;charset=UTF-8',
    },
    responseType:"json"
}