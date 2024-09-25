<template>
  <n-back-top :right="100" />
  <div class="main-body">
    <n-float-button
        @click="showAddPostModal = true"
        right="30px"
        bottom="30px"
        type="primary"
        width="48px"
        height="48px"
    >
      <n-icon size="32px">
        <Add />
      </n-icon>
    </n-float-button>
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          找工作心得
        </template>
      </n-page-header>
    </n-row>
    <n-row>
      <n-tabs type="line" size="large" animated>
        <n-tab-pane name="square" tab="讨论广场">
          讨论广场
        </n-tab-pane>
        <n-tab-pane name="myExperiment" tab="我的心得">
          我的心得
        </n-tab-pane>
      </n-tabs>
    </n-row>
  </div>
  <n-modal v-model:show="showAddPostModal">
    <n-card
        style="width: 85%"
        title="发表心得"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="addPostFormRef" :model="addPostForm" :rules="rules">
        <n-form-item path="title" label="标题">
          <n-input
              v-model:value="addPostForm.title"
              placeholder="请输入心得标题"
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="content" label="内容">
          <n-input
              v-model:value="addPostForm.content"
              type="textarea"
              maxlength="300"
              placeholder="请输入心得内容"
              show-count
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-row style="justify-content: center; margin-top: 10px">
          <n-button
              type="primary"
              @click="submitAddPostForm"
          >
            发表
          </n-button>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
</template>

<script setup>
import {router} from "../router/index.js";
import {Add} from '@vicons/ionicons5'

import {onMounted, ref} from "vue";
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {NButton, useMessage} from "naive-ui";

const message = useMessage()

const postList = ref(null)

const addPostFormRef = ref(null)
const showAddPostModal = ref(false)
const addPostForm = ref({
  title: '',
  content: ''
})

const rules = {
  title: [
    {required: true, message: '请输入标题', trigger: ['blur', 'input']},
  ],
  content: [
    {required: true, message: '请输入内容', trigger: ['blur', 'input']},
  ]
}

function back(){
  router.push("/")
}

onMounted(async () => {
  await getPostList()
})

async function getPostList(){
  const result = await axiosGet({
    url: '/post/search',
    name: 'postList-search'
  })
  if(result && result.data){
    postList.value = result.data.rows
  }
  else{
    message.warning("获取找工作心得列表失败，请刷新重试")
  }
}

async function submitAddPostForm(){
  if (!addPostFormRef) return
  await addPostFormRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await axiosPost({
        url: '/post/add',
        data: addPostForm.value,
        name: 'add-post'
      })
      if(result){
        if(result.data.code === 0){
          message.success("发表心得成功")
          await getPostList()
          showAddPostModal.value = false
        }
        else{
          message.error("发表心得失败")
        }
      }
      else{
        message.error("网络请求出错了")
      }
    }
  })
}

</script>

<style scoped>

</style>