<template>
  <RouterView v-slot="{ Component }">
    <template v-if="Component">
      <Transition name="slide-fade">
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

</script>

<style lang="scss">
.slide-fade-enter-active {
  transition: all 0.3s ease-out;
}

.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-enter-from,
.slide-fade-leave-to {
  transform: translateX(20px);
  opacity: 0;
}
</style>