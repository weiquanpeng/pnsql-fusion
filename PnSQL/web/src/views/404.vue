<template>
  <div class="not-found-container">
    <div class="not-found-content">
      <h1>4<span class="zero">0</span>4</h1>
      <p>看起来你迷路了，但至少这里的数字很有趣！</p>
      <router-link to="/" class="home-link">回到首页</router-link>
    </div>
  </div>
</template>

<script setup>
// 你的脚本逻辑
</script>

<style scoped>
.not-found-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #9ea8b1;
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
}

.not-found-content {
  text-align: center;
  max-width: 600px;
  min-height: 80vh;
}

h1 {
  font-size: 120px;
  color: #333;
  margin-bottom: 20px;
}

.zero {
  display: inline-block;
  position: relative;
  width: 60px;
  height: 60px;
  font-size: 60px;
  margin-left: 10px;
  margin-right: 10px;
  background-color: #409eff;
  color: white;
  border-radius: 50%;
  line-height: 60px;
}

p {
  font-size: 24px;
  color: #666;
  margin-bottom: 30px;
}

.home-link {
  display: inline-block;
  padding: 10px 20px;
  background-color: #409eff;
  color: white;
  text-decoration: none;
  border-radius: 5px;
  transition: background-color 0.3s;
}

.home-link:hover {
  background-color: #3e8ef7;
}
</style>