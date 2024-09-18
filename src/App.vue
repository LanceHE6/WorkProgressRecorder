<script setup>

</script>

<template>
  <n-config-provider :theme-overrides="{ common: { fontWeightStrong: '600' } }">
    <router-view/>
  </n-config-provider>
</template>

<style scoped>
</style>
