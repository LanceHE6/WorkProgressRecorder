<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          进度信息
        </template>
      </n-page-header>
    </n-row>
    <n-row>
      <n-card class="profile-body-card">
        <div class="title-nickname">
          <b>日志信息</b>
        </div>
        <div class="title-main">
          <div class="small-text">{{`公司名称：${log ? log.company_name : '-'}`}}</div>
          <div class="small-text">{{`岗位名称：${log ? log.job : '-'}`}}</div>
          <div class="small-text">{{`薪资：${log ? log.salary : '-'}k`}}</div>
          <div class="small-text">{{`工作地区：${log ? log.location : '-'}`}}</div>
        </div>
      </n-card>
    </n-row>
    <n-row>
      <n-card class="profile-body-card">
        <div class="title-nickname">
          <b>进度</b>
          <n-button
              text
              type="primary"
              size="small"
              style="margin-left: auto"
              icon="Edit"
              @click="openTimelineModal"
          >
            更新进度
          </n-button>
        </div>
        <div class="title-time">
          <n-empty v-if="timeline.length <= 0" description="无进度"/>
          <n-timeline>
            <n-timeline-item
                v-for="item in timeline"
                :type="item.stage === 1 ? 'info' : (item.stage === 2 ? 'success' : 'error')"
                :title="item.status"
                :time="new Date(item.created_at).toLocaleString()"
            />
          </n-timeline>
        </div>
      </n-card>
    </n-row>
  </div>

  <n-modal v-model:show="timelineModalVisible">
    <n-card
        style="width: 85%"
        title="更新进度"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="timelineFormRef" :model="timelineForm" :rules="rules">
        <n-form-item path="stage" label="状态">
          <n-select
              v-model:value="timelineForm.stage"
              placeholder="请选择状态"
              :options="[
                  {label: '进行中', value: 1},
                  {label: '成功', value: 2},
                  {label: '失败', value: 3},
              ]"
          />
        </n-form-item>
        <n-form-item path="status" label="状态名">
          <n-input
              v-model:value="timelineForm.status"
              placeholder="请输入状态名"
          />
        </n-form-item>

        <n-form-item path="create_time" label="时间">

          <n-date-picker
              v-model:value="timelineForm.create_time"
              type="datetime"
              placeholder="请选择时间"
              style="width: 100%"
          />
        </n-form-item>

        <n-row style="justify-content: center">
          <n-button type="primary" @click="submitTimelineForm">
            更新状态
          </n-button>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
</template>

<script setup>
import {onMounted, ref} from "vue";
import {router} from "../router/index.js";
import {axiosPut} from "../utils/axiosUtil.js";
import {useMessage} from "naive-ui";

const message = useMessage()

const timelineFormRef = ref(null)

const log = ref({})
const timeline = ref([])

const timelineModalVisible = ref(false)

const timelineForm = ref({
  id: '',
  stage: 1,
  status: '',
  create_time: '',
})

const rules = {
  status: [
    { required: true, message: '请输入状态名', trigger: ['blur', 'input']},
  ]
}

function back(){
  router.push("/occupation")
}

onMounted(() => {
  refresh()
})

function refresh(){
  log.value = JSON.parse(localStorage.getItem("logDetail") || '')
  timeline.value = [...log.value.status_time_line].reverse()
}


function openTimelineModal(){
  timelineForm.value.create_time = Date.now()
  timelineModalVisible.value = true
}
const submitTimelineForm = async () => {
  if (!timelineFormRef) return
  await timelineFormRef.value?.validate(async (errors) => {
    if (!errors) {
      timelineForm.value.id = log.value.id
      const result = await axiosPut({
        url:'/wl/status/add',
        data: timelineForm.value,
        name: 'log-status-add'
      })
      if(result){
        if(result.data.code === 0){
          message.success("进度更新成功")
          timelineModalVisible.value = false
          log.value.status_time_line.push({
            stage: timelineForm.value.stage,
            status: timelineForm.value.status,
            created_at: timelineForm.value.create_time
          })
          localStorage.setItem("logDetail", JSON.stringify(log.value))
          refresh()
        }
        else {
          message.error("网络请求出错了")
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
.main-body{
  display: flex;
  flex-direction: column;
  justify-content: flex-start; /* 子元素在父容器中垂直分布 */
  width: 100%;
  height: 100%;
  padding: 15px 10px 10px 10px;
  box-sizing: border-box;
}
.profile-body-card{
  margin: 10px 0 0 0;
  width: 95%;
}
.title-main{
  justify-content: center;
  padding: 0 0 0 15px;
}
.small-text{
  text-align: center;
  font-size: 18px;
}
.title-nickname{
  display: flex;
  justify-content: space-between;
  align-items: center; /* 垂直居中 */
  font-size: 16px;
  margin-bottom: 10px;
}
.title-time {
  justify-content: flex-start;
  text-align: start;
}
</style>