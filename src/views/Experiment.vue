<template>
  <n-float-button
      @click="openAddPostModal"
      right="30px"
      bottom="30px"
      type="primary"
      width="48px"
      height="48px"
      style="z-index: 1"
  >
    <n-icon size="32px">
      <Add />
    </n-icon>
  </n-float-button>
  <n-float-button
      @click="refresh"
      right="30px"
      bottom="90px"
      width="48px"
      height="48px"
      style="z-index: 1; background-color: #2080f0; color: white"
  >
    <n-icon size="32px">
      <Refresh />
    </n-icon>
  </n-float-button>
  <div class="main-body">
      <n-row style="justify-content: left; height: 5vh">
        <n-page-header subtitle="" @back="back">
          <template #title>
            找工作心得
          </template>
        </n-page-header>
      </n-row>
      <n-spin :show="isLoading">
      <n-row>
        <n-tabs type="line" size="large" animated>
          <n-tab-pane name="square" tab="讨论广场">
            <n-infinite-scroll ref="postScroll" style="height: 82vh" :distance="10" @load="handlePostLoad">
              <n-card
                  v-for="item in postList"
                  :segmented="{
                  footer: 'soft',
                }"
                  embedded
                  style="text-align: start; padding: 0; margin-bottom: 10px; border-color: #d2d2d2"
              >
                <template #header>
                  {{ item.title }}
                </template>
                <template #header-extra>
                <span v-if="!item.anonymous" class="very-small-text">
                  {{ item.user.name }}
                </span>
                  <span style="margin-left: 20px">
                  {{ toLocaleString(item.created_at) }}
                </span>
                </template>
                <n-ellipsis expand-trigger="click" line-clamp="2" :tooltip="false" style="height: auto">
                  {{ item.content }}
                </n-ellipsis>
                <template #footer>
                  <n-card
                      v-for="(i, index) in item.comments"
                      embedded
                      :bordered="false"
                      style="text-align: start"
                      class="card-comment"
                  >
                    <div v-if="index < item.limitNum">
                    <span>
                      {{ !i.anonymous ? i.user.name : '匿名用户' }}：{{ i.content }}
                    </span>
                    </div>
                  </n-card>
                  <n-button
                      v-if="item.comments && item.comments.length > 3"
                      type="info"
                      text
                      @click="expand(item)"
                  >
                    {{ item.limitText }}
                  </n-button>
                  <n-row v-if="CURRENT_USER" style="margin-top: 10px; justify-content: space-between">
                    <n-input
                        v-model:value="item.inputText"
                        placeholder="发表你的看法"
                    >
                      <template #suffix>
                        <n-checkbox
                            v-model:checked="item.inputAnonymous"
                            size="small"
                            label="匿名"
                            class="checkbox-comment"
                        />
                        <n-button @click="sendComment(item)" color="#18a058" quaternary circle>
                          <template #icon>
                            <n-icon><SendSharp /></n-icon>
                          </template>
                        </n-button>
                      </template>
                    </n-input>
                  </n-row>
                </template>
              </n-card>
              <div v-if="postLoading" class="text">
                加载中...
              </div>
              <div v-if="postNoMore" class="text">
                已经到底啦
              </div>
            </n-infinite-scroll>
          </n-tab-pane>
          <n-tab-pane v-if="CURRENT_USER" name="myExperiment" tab="我的心得">
            <n-infinite-scroll ref="myPostScroll" style="height: 82vh" :distance="10" @load="handleMyPostLoad">
              <n-card
                  v-for="item in myPostList"
                  :segmented="{
                  footer: 'soft',
                }"
                  embedded
                  style="text-align: start; padding: 0; margin-bottom: 10px; border-color: #d2d2d2">
                <template #header>
                  {{ item.title }}
                </template>
                <template #header-extra>
                  <n-popconfirm :negative-text="null" @positive-click="deletePost(item)">
                    <template #trigger>
                      <n-button
                          text
                          type="error"
                      >
                        删除
                      </n-button>
                    </template>
                    你确定要删除这条心得吗？删除后将无法恢复！
                  </n-popconfirm>
                  <span style="margin-left: 20px">
                  {{ toLocaleString(item.created_at) }}
                </span>
                </template>
                <n-ellipsis expand-trigger="click" line-clamp="2" :tooltip="false" style="height: auto">
                  {{ item.content }}
                </n-ellipsis>
                <template #footer>
                  <n-card
                      v-for="(i, index) in item.comments"
                      embedded
                      :bordered="false"
                      style="text-align: start"
                      class="card-comment"
                  >
                    <div v-if="index < item.limitNum">
                    <span>
                      {{ !i.anonymous ? i.user.name : '匿名用户' }}：{{ i.content }}
                    </span>
                    </div>
                  </n-card>
                  <n-button
                      v-if="item.comments && item.comments.length > 3"
                      type="info"
                      text
                      @click="expand(item)"
                  >
                    {{ item.limitText }}
                  </n-button>
                  <n-row style="margin-top: 10px; justify-content: space-between">
                    <n-input
                        v-model:value="item.inputText"
                        placeholder="发表你的看法"
                    >
                      <template #suffix>
                        <n-checkbox
                            v-model:checked="item.inputAnonymous"
                            size="small"
                            label="匿名"
                            class="checkbox-comment"
                        />
                        <n-button @click="sendComment(item)" color="#18a058" quaternary circle>
                          <template #icon>
                            <n-icon><SendSharp /></n-icon>
                          </template>
                        </n-button>
                      </template>
                    </n-input>
                  </n-row>
                </template>
              </n-card>
              <div v-if="myPostLoading" class="text">
                加载中...
              </div>
              <div v-if="myPostNoMore" class="text">
                已经到底啦
              </div>
            </n-infinite-scroll>
          </n-tab-pane>
        </n-tabs>
      </n-row>
    </n-spin>
  </div>

  <n-modal
      v-model:show="showAddPostModal"
      preset="dialog"
      title="发表心得"
      positive-text="发表"
      @positive-click="submitAddPostForm"
      :mask-closable="false"
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
      <n-checkbox
          v-model:checked="addPostForm.anonymous"
          label="匿名发表"
      />
    </n-form>
  </n-modal>
</template>

<script setup>
import {router} from "../router/index.js";
import {Add} from '@vicons/ionicons5'
import { SendSharp, Refresh } from '@vicons/ionicons5'


import {onMounted, ref} from "vue";
import {axiosDelete, axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {useMessage} from "naive-ui";
import {toLocaleString} from "../utils/other.js";
import {CURRENT_USER} from "../utils/appManager.js";

const message = useMessage()
const pageSize = 10
const isLoading = ref(true)



let postListPage = 1
let postListTotal = 1
const postScroll = ref(null)
const postLoading = ref(false)
const postNoMore = ref(false)
const postList = ref([])
const postListParams = ref({
  page_size: pageSize,
})

let myPostListPage = 1
let myPostListTotal = 1
const myPostScroll = ref(null)
const myPostLoading = ref(false)
const myPostNoMore = ref(false)
const myPostList = ref([])
const myPostListParams = ref({
  uid: '',
  page_size: pageSize,
})

const addPostFormRef = ref(null)
const showAddPostModal = ref(false)
const addPostForm = ref({
  title: '',
  content: '',
  anonymous: false,
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
  await refresh()
})

async function getPostList(page = 1){
  postListParams.value.page = page
  const result = await axiosGet({
    url: '/post/search',
    params: postListParams.value,
    name: 'postList-search'
  })
  if(result && result.data){
    for(let i in result.data.rows){
      result.data.rows[i].limitText = '展开'
      result.data.rows[i].limitNum = 3
      result.data.rows[i].inputAnonymous = false
      result.data.rows[i].inputText = ''
    }
    postListTotal = result.data.total
    return result.data.rows
  }
  else{
    return null
  }
}

async function initPostList(){
  postListPage = 1
  const initList = await getPostList(postListPage.value)
  if(!initList){
    message.warning("获取心得列表失败，请刷新重试")
  }
  postList.value = initList
  postLoading.value = false
  postNoMore.value = false
}

async function getMyPostList(page = 1){
  myPostListParams.value.page = page
  const result = await axiosGet({
    url: '/post/search',
    params: myPostListParams.value,
    name: 'myPostList-search'
  })
  if(result && result.data){
    for(let i in result.data.rows){
      result.data.rows[i].limitText = '展开'
      result.data.rows[i].limitNum = 3
      result.data.rows[i].inputAnonymous = false
      result.data.rows[i].inputText = ''
    }
    myPostListTotal = result.data.total
    return result.data.rows
  }
  else{
    return null
  }
}
async function initMyPostList(){
  myPostListPage = 1
  const initList = await getMyPostList(myPostListPage)
  if(!initList){
    message.warning("获取个人心得列表失败，请刷新重试")
  }
  myPostList.value = initList
  myPostLoading.value = false
  myPostNoMore.value = false
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
        }
        else{
          message.error("发表心得失败")
        }
        showAddPostModal.value = false
        await initPostList()
        await initMyPostList()
      }
      else{
        message.error("网络请求出错了")
      }
    }
  })
}

async function sendComment(item){
  if (!item.inputText) return
  const body = {
    id: item.id,
    content: item.inputText,
    anonymous: item.inputAnonymous
  }
  const result = await axiosPost({
    url: '/post/comment/add',
    data: body,
    name: 'add-comment'
  })
  if(result){
    if(result.data.code === 0){
      message.success("评论成功")
    }
    else{
      message.error("评论失败")
    }
    await initPostList()
    await initMyPostList()
  }
  else{
    message.error("网络请求出错了")
  }
}

async function deletePost(item){
  const body = {
    id: item.id,
  }
  const result = await axiosDelete({
    url: '/post/del',
    data: body,
    name: 'del-post'
  })
  if(result){
    if(result.data.code === 0){
      message.success("删除心得成功")
    }
    else{
      message.error("删除心得失败")
    }
    await initPostList()
    await initMyPostList()
  }
  else{
    message.error("网络请求出错了")
  }
}

function expand(item){
  if(item.limitNum === 3){
    item.limitNum = 65535
    item.limitText = '收回'
  }
  else{
    item.limitNum = 3
    item.limitText = '展开'
  }
}

function openAddPostModal(){
  if(CURRENT_USER.value){
    showAddPostModal.value = true
  }
  else{
    message.warning("请先登录")
  }
}

async function refresh(){
  isLoading.value = true
  await initPostList()
  await initMyPostList()
  postListToTop()
  myPostListToTop()
  isLoading.value = false
}

async function handlePostLoad(){
  if (postLoading.value || postNoMore.value){
    return
  }
  if(postList.value.length >= postListTotal){
    postNoMore.value = true
    return
  }
  postLoading.value = true
  postListPage += 1
  const newPage = await getPostList(postListPage)
  if(!newPage){
    message.error("加载心得列表失败，请稍后重试")
  }
  else{
    Array.prototype.push.apply(postList.value, newPage)
  }
  postLoading.value = false
}

async function handleMyPostLoad(){
  if (myPostLoading.value || myPostNoMore.value){
    return
  }
  if(myPostList.value.length >= myPostListTotal){
    myPostNoMore.value = true
    return
  }
  myPostLoading.value = true
  myPostListPage += 1
  const newPage = await getMyPostList(myPostListPage)
  if(!newPage){
    message.error("加载个人心得列表失败，请稍后重试")
  }
  else{
    Array.prototype.push.apply(myPostList.value, newPage)
  }
  myPostLoading.value = false
}

function postListToTop(){
  postScroll.value?.scrollbarInstRef?.scrollTo({ top: 0, behavior: 'smooth' })
}

function myPostListToTop(){
  myPostScroll.value?.scrollbarInstRef?.scrollTo({ top: 0, behavior: 'smooth' })
}

</script>

<style scoped>
.card-comment :deep(.n-card__content){
  padding: 0;
  color: #8c939d;
}
.checkbox-comment :deep(.n-checkbox__label){
  color: #8c939d;
}
</style>