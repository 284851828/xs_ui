<template>
  <div class="page">
    <DashboardSimple/>
  </div>
</template>

<script setup>
import DashboardSimple from './simple/index.vue';
</script>

<style lang="scss" scoped>
.page {
  height: calc(100vh - 101px);
}
.admin-box {
  padding: 0;
}
</style>