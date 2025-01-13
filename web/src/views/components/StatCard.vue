<template>
  <div class="stateBox">
    <el-row :gutter="12">
      <el-col :span="6" style="padding-right: 16px;padding-left: 16px">
        <el-card shadow="always">
          <div class="tit">总资产</div>
          <div class="num">
            <AnimatedNumber :from="0" :to="325600" :duration=2000>
              <template #default="{ option, item }">
                <span>￥{{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
          <div class="info">
            <AnimatedNumber :from="0" :to="10000000" :duration=2000>
              <template #default="{ option, item }">
                <span>目标资产：￥{{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6" style="padding-right: 16px;padding-left: 16px">
        <el-card shadow="always">
          <div class="tit">MySql</div>
          <div class="num">
            <AnimatedNumber :from="0" :to="5564" :duration="2000">
              <template #default="{ option, item }">
                <span>总实例: {{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
          <div class="info">
            <AnimatedNumber :from="0" :to="1869" :duration="2000">
              <template #default="{ option, item }">
                <span>正常运行：{{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6" style="padding-right: 16px;padding-left: 16px">
        <el-card shadow="always">
          <div class="tit">MongoDB</div>
          <div class="num">
            <AnimatedNumber :from="0" :to="578" :duration="2000">
              <template #default="{ option, item }">
                <span>总实例: {{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
          <div class="info">
            <AnimatedNumber :from="0" :to="467" :duration="2000">
              <template #default="{ option, item }">
                <span>正常运行：{{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
        </el-card>
      </el-col>
      <el-col :span="6" style="padding-right: 16px;padding-left: 16px">
        <el-card shadow="always">
          <div class="tit">Redis</div>
          <div class="num">
            <AnimatedNumber :from="0" :to="8743" :duration="2000">
              <template #default="{ option, item }">
                <span>总实例: {{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
          <div class="info">
            <AnimatedNumber :from="0" :to="7869" :duration="2000">
              <template #default="{ option, item }">
                <span>正常运行：{{ item.number }}</span>
              </template>
            </AnimatedNumber>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
  <div class="gva-card-box" style="padding-top: 2px;padding-bottom: 10px">
    <el-row :gutter="12">
      <el-col :span="24" style="padding-right: 16px;padding-left: 16px">
        <el-card shadow="always">
          <div class="gva-card gva-top-card" style="padding: 20px">
            <div class="gva-top-card-left">
              <div class="gva-top-card-left-title" style="padding-top:10px">早安，管理员，请开始一天的工作吧</div>
              <div class="gva-top-card-left-dot" style="padding-top:10px">今日晴，0℃ - 10℃，天气寒冷，注意添加衣物。</div>
              <div >
                <div class="gva-top-card-left-item" style="padding-top:10px"> 使用教学： <a
                    target="view_window" href="https://www.baidu.com"
                    style="color: rgb(64, 158, 255);">https://www.baidu.com</a></div>
                <div class="gva-top-card-left-item" style="padding-top:10px"> 代码仓库： <a
                    target="view_window" href="https://github.com/weiquanpeng/xiaohongshu"
                    style="color: rgb(64, 158, 255);">https://github.com/weiquanpeng/xiaohongshu</a></div>
              </div>
            </div>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>

</script>

<style scoped>
.stateBox {
  .el-card {
    margin-bottom: 5px;
  }

  .tit {
    height: 22px;
    color: rgba(0, 0, 0, .45);
    font-size: 14px;
    line-height: 22px;
  }

  .num {
    height: 28px;
    margin-top: 5px;
    color: rgba(0, 0, 0, .75);
    font-size: 20px;
    line-height: 38px;
    white-space: nowrap;
    font-weight: bold;
    text-overflow: ellipsis;
    word-break: break-all;
  }

  .info {
    position: relative;
    width: 100%;
    margin-top: 20px;
  }
}

.gva-card{
  height: 190px;
  background: linear-gradient(rgba(255, 255, 255, 0.6), rgba(255, 255, 255, 0.6)), url("@/assets/images/8.png");
}

</style>