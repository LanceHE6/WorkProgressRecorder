<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          找工作日志
        </template>
      </n-page-header>
    </n-row>


    <n-row>
      <n-card class="profile-title-card">
        <div class="title-main">
          <div class="small-text">{{`工作状态：${goal ? (goal.status === 1 ? '未拿到offer' : '已拿到offer') : ''}`}}</div>
        </div>
      </n-card>
    </n-row>

    <n-row>
      <n-card class="profile-body-card">
        <div class="title-main">
          <div class="small-text">{{`目标公司：${goal ? goal.target_company : ''}`}}</div>
          <div class="small-text">{{`理想薪资：${goal ? goal.target_salary : ''}k`}}</div>
          <div class="small-text">{{`目标地区：${goal ? goal.target_area : ''}`}}</div>
        </div>
      </n-card>
    </n-row>
    <n-row style="margin-top: 20px; justify-content: space-around">
      <n-button
          strong secondary
          type="warning"
          @click="showLogModal = true"
          style="width: 40%; height: 80px; font-size: 18px"
          :disabled="!goal"
      >
        新增工作日志
      </n-button>

      <n-button
          strong secondary
          type="warning"
          @click="showExperienceModal = true"
          style="width: 40%; height: 80px; font-size: 16px"
          :disabled="!goal"
      >
        找工作心得/困难
      </n-button>
    </n-row>

    <n-row style="justify-content: center; margin-top: 10px">
      <n-data-table
          :columns="logHead"
          :data="logList"
          :bordered="false"
          style="width: 95%; height: 50vh; background-color: #f9f9f9c0;"
      />
    </n-row>
  </div>

  <n-modal
      v-model:show="showLogModal"
      preset="dialog"
      title="新增工作日志"
      positive-text="确认"
      negative-text="算了"
      @positive-click="addLog"
  >
    <n-form ref="logFormRef" :model="logForm" :rules="rules">
      <n-form-item path="company_name" label="公司名称">
        <n-input v-model:value="logForm.company_name" @keydown.enter.prevent placeholder="请输入公司名称"/>
      </n-form-item>
      <n-form-item path="job" label="工作岗位">
        <n-input v-model:value="logForm.job" @keydown.enter.prevent placeholder="请输入工作岗位"/>
      </n-form-item>
      <n-form-item path="salary" label="薪资">
        <n-input v-model:value="logForm.salary" @keydown.enter.prevent placeholder="请输入薪资"/>
      </n-form-item>
      <n-form-item path="location" label="工作地区">
        <n-input v-model:value="logForm.location" @keydown.enter.prevent placeholder="请输入工作地区"/>
      </n-form-item>
    </n-form>
  </n-modal>

  <n-modal
      v-model:show="showExperienceModal"
      preset="dialog"
      title="找工作心得/困难"
      positive-text="确认"
  >
    <n-text>
      功能还在完善中，敬请期待...
    </n-text>
  </n-modal>
</template>

<script setup>
import {router} from "../router/index.js";
import {h, onMounted, ref} from "vue";
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {NButton, useMessage} from "naive-ui";
import {setUser} from "../utils/appManager.js";

const message = useMessage()

const isLoading = ref(true)

const showLogModal = ref(false)
const showExperienceModal = ref(false)

const goal = ref(null)
const logList = ref([])

const rules = {
  company_name:[
    { required: true, message: '请输入公司名称', trigger: ['blur', 'input']},
  ],
  job:[
    { required: true, message: '请输入工作岗位', trigger: ['blur', 'input']},
  ],
  salary:[
    { required: true, message: '请输入薪资', trigger: ['blur', 'input']},
  ],
  location:[
    { required: true, message: '请输入工作地区', trigger: ['blur', 'input']},
  ]
}

const logFormRef = ref(null)
const logForm = ref({
  company_name: '',
  job: '',
  salary: '',
  location: ''
})

const logHead = [
  { title: "公司名称", key: "company_name"},
  { title: "工作岗位", key: "job"},
  { title: "操作", key: "actions",
    render(row) {
      return h(
          NButton,
          {
            strong: true,
            tertiary: true,
            type: 'info',
            size: "small",
            onClick: () => showLogDetail(row)
          },
          { default: () => "详细信息" }
      );
    }
  }
]

function showLogDetail(row){
  localStorage.setItem("logDetail", JSON.stringify(row))
  router.push("/logDetail")
}

function back(){
  router.push("/")
}

onMounted(async () =>{
  await refresh()
  if(!goal.value){
    message.warning("您没有确定就业目标")
  }
})

const addLog = async () => {
  if (!logFormRef) return
  await logFormRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await axiosPost({
        url:'/wl/add',
        data: logForm.value,
        name: 'add-log'
      })
      if(result){
        if(result.data.code === 0){
          message.success("工作日志添加成功")
          for(const item in logForm.value){
            logForm.value[item] = ''
          }
          await refresh()
        }
        else {
          message.error("后端请求出错了")
        }
      }
      else{
        message.error("网络请求出错了")
      }
    }
  })
}

const refresh = async () => {
  isLoading.value = true
  const result = await axiosGet({
    url: '/user/info',
    name: 'occupation-refresh'
  })
  if(result && result.data && result.data.user){
    setUser(result.data.user)
  }

  const result2 = await axiosGet({
    url: '/user/target',
    name: 'occupation-get-empl_goal'
  })
  if(result2.data && result2.data.direction.includes(2)){
    goal.value = result2.data.goal.empl_goal

    const result3 = await axiosGet({
      url: '/wl/list',
      name: 'occupation-log-list'
    })
    if(result3 && result3.data){
      logList.value = result3.data.work_logs
    }
    else{
      message.error("获取工作日志失败")
    }
  }
  isLoading.value = false
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
.profile-title-card{
  margin: 30px 0 0 0;
  width: 95%;
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
</style>