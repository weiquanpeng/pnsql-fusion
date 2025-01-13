<template>
  <RouterView v-slot="{ Component }">
    <template v-if="Component">
      <Transition name="slide-fade" mode="out-in">
        <KeepAlive>
          <div :key="$route.path">
            <Suspense>
              <template #default>
                <component :is="Component"></component>
              </template>
              <template #fallback>
                <div>加载中...</div>
              </template>
            </Suspense>
          </div>
        </KeepAlive>
      </Transition>
    </template>
  </RouterView>
</template>

<script setup>
// No additional script needed in this setup unless specific logics or data are managed.
</script>

<style lang="scss">
.slide-fade-enter-active, .slide-fade-leave-active {
  transition: all 0.3s ease-out;
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}

.slide-fade-leave-from,
.slide-fade-enter-to {
  transform: translateX(0);
  opacity: 1;
}
</style>
