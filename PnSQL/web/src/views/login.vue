<template>
  <div class="login-box">
    <div class="img_box">
      <div class="bg_blue"></div>
      <ul class="circles">
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
        <li></li>
      </ul>
      <div class="w-full max-w-md">
        <div class="fz48 cof fw" style="margin-top:10px; line-height:64px; color: white; font-size: 50px; font-weight: bold;text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);">欢迎使用</div>
        <div class="fz14 cof" style="margin-top:10px; line-height:24px; color: white; font-size: 20px; font-weight: bold;text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5);">欢迎来到好玩俱乐部，在这里和志同道合的朋友一起分享有趣的故事，一起组织有趣的活动...</div>
      </div>
    </div>
    <div class="login_box">
      <div class="login-wrap">
        <h1 class="header_fz32">{{ cTitle }}</h1>
        <form action="#">
          <div class="ksd-el-items"><input type="text" v-model="loginUser.username" class="ksd-login-input"  placeholder="请输入账号"></div>
          <div class="ksd-el-items"><input type="password" v-model="loginUser.password" class="ksd-login-input" placeholder="请输入密码" @keydown.enter="onSubmit"></div>
          <div class="ksd-el-items pr">
            <input type="text" class="ksd-login-input" maxlength="6" v-model="loginUser.code" placeholder="请输入验证码"  @keydown.enter="onSubmit">
            <img v-if="codeURL" class="code_url"  :src="codeURL" @click="getCaptcha">
          </div>
          <div class="ksd-el-items"><input type="button" @click.prevent="onSubmit" class="ksd-login-btn" value="登录"></div>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup>
import {useRouter, useRoute} from "vue-router";
import {useCounterStore} from '@/stores/user.js'
import request from "@/request/index.js";
import {getCaptchaApi} from "@/api/code.js";
const cTitle = ref(import.meta.env.VITE_APP_TITLE)

const userStores = useCounterStore()

var router = useRouter()
var route = useRoute()

const codeURL = ref("")
const loginUser = reactive({
  username: "",
  password: "",
  id: "",
  code: ""
})

const getCaptcha = async () => {
  const resp = await getCaptchaApi()
  const {baseURL, id} = resp.data
  codeURL.value = baseURL
  loginUser.id = id
}

const onSubmit = async () => {
  try {
    await userStores.toLogin(loginUser)
    // router.push("/index?user=" + resp.data.user + "&token=" + resp.data.token)    //replace不能返回登录页面
    await router.push('/index')
  } catch (e) {
    await getCaptcha()
  }
}

onMounted(() => {
  getCaptcha()
})
</script>



<style scoped lang="scss">
.pr{position: relative;}
.code_url{position: absolute;top:5px;right:5px;width: 140px;}
.ksd-el-items{margin: 15px 0;}
.ksd-login-input{border:1px solid #eee;padding:16px 8px;width: 100%;box-sizing: border-box;outline: none;border-radius: 4px;}
.ksd-login-btn{border:1px solid #eee;padding:16px 8px;width: 100%;box-sizing: border-box;
  background:#2196F3;color:#fff;border-radius:6px;cursor: pointer;}
.ksd-login-btn:hover{background:#1789e7;}
.login-box{
  display: flex;
  flex-wrap: wrap;
  background:#fff;
  .login_box{
    width: 35%;height: 100vh;
    display: flex;
    align-items: center;
    justify-content: center;
    .header{margin-bottom: 30px;}
    .login-wrap{
      width: 560px;
      height: 444px;
      padding:20px 100px;
      box-sizing: border-box;
      border-radius: 8px;
      box-shadow: 0 0 10px #fafafa;
      background: rgba(255,255,255,0.6);
      text-align: center;
      display: flex;
      flex-direction: column;
      justify-content: center;
    }
  }
  .img_box{
    width: 65%;
    height: 100vh;
    display: flex;
    align-items: center;
    justify-content:center;
    position:relative;
    background: url("../assets/images/login_background.jpg");
    background-size:cover;
    background-repeat:no-repeat;

    .bg_blue{
      background-image:linear-gradient(to bottom,#4f46e5,#3b82f6);
      position:absolute;
      top:0;
      left:0;
      bottom:0;
      right:0;
      opacity:0.75;
    }
  }
}

.circles {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
}
.circles li {
  position: absolute;
  display: block;
  list-style: none;
  width: 20px;
  height: 20px;
  background: rgba(255, 255, 255, 0.2);
  animation: animate 15s linear infinite;
  bottom: -150px;
}
.circles li:nth-child(1) {
  left: 25%;
  width: 80px;
  height: 80px;
  animation-delay: 0s;
  animation-duration: 10s;
}
.circles li:nth-child(2) {
  left: 10%;
  width: 120px;
  height: 120px;
  animation-delay: 3s;
  animation-duration: 12s;
}
.circles li:nth-child(3) {
  left: 70%;
  width: 20px;
  height: 20px;
  animation-delay: 4s;
}
.circles li:nth-child(4) {
  left: 40%;
  width: 60px;
  height: 60px;
  animation-delay: 0s;
  animation-duration: 18s;
}
.circles li:nth-child(5) {
  left: 65%;
  width: 20px;
  height: 20px;
  animation-delay: 0s;
}
.circles li:nth-child(6) {
  left: 75%;
  width: 110px;
  height: 110px;
  animation-delay: 3s;
}
.circles li:nth-child(7) {
  left: 35%;
  width: 150px;
  height: 150px;
  animation-delay: 7s;
}
.circles li:nth-child(8) {
  left: 50%;
  width: 25px;
  height: 25px;
  animation-delay: 15s;
  animation-duration: 45s;
}
.circles li:nth-child(9) {
  left: 20%;
  width: 15px;
  height: 15px;
  animation-delay: 2s;
  animation-duration: 35s;
}
.circles li:nth-child(10) {
  left: 85%;
  width: 150px;
  height: 150px;
  animation-delay: 0s;
  animation-duration: 11s;
}
@keyframes animate {
  0% {
    transform: translateY(0) rotate(0deg);
    opacity: 1;
    border-radius: 0;
  }

  100% {
    transform: translateY(-1000px) rotate(720deg);
    opacity: 0;
    border-radius: 50%;
  }
}
.max-w-md {
  max-width: 28rem;
  position:relative;
  z-index:10;
}

.header_fz32 {
  font-size: 36px;
  font-weight: bold;
  color: black;            /* 字体颜色 */
  text-shadow: 2px 2px 4px rgba(0, 0, 0, 0.5); /* 添加阴影 */
  line-height: 1.5;
  text-align: center;
  margin: 20px 0;
  font-family: 'Georgia', serif; /* 可选字体 */
}
</style>
