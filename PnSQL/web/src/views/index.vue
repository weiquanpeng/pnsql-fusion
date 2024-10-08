<template>
  <div class="common-layout">
    <el-container>
      <el-header class="pwq-header"><page-header/></el-header>
      <el-container >
        <el-aside class="pwq-sidebar" style="width:150px"><PageSidebar/></el-aside>
        <el-main class="pwq-main">
          <page-menu-tab></page-menu-tab>
          <page-main></page-main>
        </el-main>
      </el-container>
    </el-container>
  </div>
</template>

<script setup>
import PageHeader from '@/views/video/PageHeader.vue'
import PageMenuTab from '@/views/video/PageMenuTab.vue'
import PageMain from '@/views/video/PageMain.vue'
import PageSidebar from '@/views/video/PageSidebar.vue';

</script>

<style scoped>
.pwq-header{
  display: flex;
}

.pwq-sidebar{
  height: calc(100vh - 70px);
  overflow: hidden;
  border-top: 3px solid var(--color-neutral-2);
  border-right: 3px solid var(--color-neutral-2)
}

.pwq-main{
  padding: 0 0;
  height: calc(100vh - 70px);
  border-top: 3px solid var(--color-neutral-2)
}

</style>