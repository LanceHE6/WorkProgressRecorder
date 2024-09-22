<script setup>
import {ref} from "vue";
import {useRoute} from "vue-router";
import { dateZhCN, zhCN } from 'naive-ui'

const transitionName = ref('')
const route = useRoute();

</script>

<!--test-->
<template>
  <n-config-provider :locale="zhCN" :date-locale="dateZhCN">
    <n-dialog-provider>
      <n-message-provider>
        <router-view v-slot="{ Component }">
          <!--        <keep-alive>-->
          <component :is="Component"/>
          <!--        </keep-alive>-->
        </router-view>
      </n-message-provider>
    </n-dialog-provider>
  </n-config-provider>
</template>

<style scoped>

</style>
