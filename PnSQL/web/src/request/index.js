import config from './config.js'
import {useCounterStore} from "@/stores/user.js";
import router from '@/router'
import axios from "axios";
const instance = axios.create({...config})


// 添加请求拦截器
instance.interceptors.request.use(function (config) {
    const useStore = useCounterStore()
    const token=useStore.getToken()
    if(token){
        config.headers['Authorization']=token
    }else {
        router.push('/login')
    }
    return config;
}, function (error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
instance.interceptors.response.use(function (response) {
    const useStore = useCounterStore();
    // 检查并设置 token
    if (response.headers['PanGu-token']) {
        useStore.setToken(response.headers['PanGu-token']);
    }

    // 检查返回的状态码
    if (response.data.code === 200) {
        return response.data;
    } else if (response.data.code === 801) {
        useStore.toLogout()
        ElMessage({message: response.data.msg, type: "error", showClose: true});
        return Promise.reject(response.data);
    } else {
        if (response.data.msg) {
            ElMessage({message: response.data.msg, type: "error", showClose: true});
        }
        return Promise.reject(response.data);
    }
}, function (err) {
    if (err && err.response) {
        switch (err.response.status) {
            case 400:
                err.message = "请求错误";
                break;
            case 401:
                err.message = "未授权，请登录";
                break;
            case 403:
                err.message = "拒绝访问";
                break;
            case 404:
                err.message = `请求地址出错: ${err.response.config.url}`;
                break;
            case 408:
                err.message = "请求超时";
                break;
            case 500:
                err.message = "服务器内部错误";
                break;
            case 501:
                err.message = "服务未实现";
                break;
            case 502:
                err.message = "网关错误";
                break;
            case 503:
                err.message = "服务不可用";
                break;
            case 504:
                err.message = "网关超时";
                break;
            case 505:
                err.message = "HTTP版本不受支持";
                break;
            default:
        }

        if (err.message) {
            ElMessage({message: err.message, type: "error", showClose: true});
        }
    }
    return Promise.reject(err);
});


export default instance


