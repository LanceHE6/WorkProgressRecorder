<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          考研打卡
        </template>
      </n-page-header>
    </n-row>

    <n-row>
      <n-card class="profile-title-card">
        <div class="title-main">
          <div class="small-text">{{`25考研倒计时：${examDay ? examDay : 'xxx'}天`}}</div>
        </div>
      </n-card>
    </n-row>

    <n-row>
      <n-card class="profile-body-card">
        <div class="title-main">
          <div class="small-text">{{`目标院校：${goal ? goal.target_university : ''}`}}</div>
          <div class="small-text">{{`目标专业：${goal ? goal.target_major : ''}`}}</div>
          <div class="small-text">{{`目标分数：${goal ? goal.target_score : ''}`}}</div>
        </div>
      </n-card>
    </n-row>
    <n-row style="margin-top: 20px; justify-content: center">
      <n-button
          strong secondary
          type="info"
          @click="showCheckInModal = true"
          :disabled="isCheck"
          style="width: 80%; height: 80px; font-size: 24px">
        <template #icon>
          <n-icon size="24"><GolfOutline /></n-icon>
        </template>
        {{ isCheck ? '已打卡' : '学习打卡' }}
      </n-button>
    </n-row>
  </div>

  <n-modal
      v-model:show="showCheckInModal"
      preset="dialog"
      title="学习打卡"
      positive-text="确认"
      negative-text="算了"
      @positive-click="checkIn">
        <n-select
            v-model:value="place"
            :options="options"
            placeholder="请选择打卡地点"
            size="large"
            style="width: 100%"
            scrollable
        >
        </n-select>
  </n-modal>
</template>

<script setup>
import {GolfOutline} from '@vicons/ionicons5'
import {router} from "../router/index.js";
import {onMounted, ref} from "vue";
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {useMessage} from "naive-ui";

const message = useMessage()

const examDay = ref(null)
const goal = ref(null)
const showCheckInModal = ref(false)

const isCheck = ref(true)

const place = ref(null)

const options = [
  {label: '数字图书馆', value: '数字图书馆'},
  {label: '老图书馆', value: '老图书馆'},
  {label: '教学楼', value: '教学楼'},
  {label: '实验楼', value: '实验楼'},
  {label: '信科楼', value: '信科楼'},
]

function back(){
  router.push("/")
}

onMounted(async () =>{
  const result = await axiosGet({
    url: '/user/target',
    name: 'post-get-goal'
  })
  if(result && result.data && result.data.direction === 1){
    goal.value = result.data.goal
  }
  else{
    message.warning("您没有确定考研目标")
  }

  const result2 = await axiosGet({
    url: '/clock_in/is_clock_in',
    name: 'post-get-clock'
  })
  if(result2){
    if(result2.code === 0){
      isCheck.value = false
    }
    else if(result2.code === 1){
      isCheck.value = true
    }
  }
  else{
    message.error("打卡记录获取失败")
  }

  const result3 = await axiosGet({
    url: '/pgee/cd',
    name: 'post-get-exam-day'
  })
  if(result3 && result3.data){
    examDay.value = result3.data.days
  }
  else{
    message.error("考研时间获取失败")
  }
})

async function checkIn(){
  if(!place.value){
    message.warning("请选择打卡地点")
    return
  }
  const result = await axiosPost({
    url: '/clock_in/add',
    data: {
      location: place.value
    },
    name: 'check-in'
  })
  if(result){
    if(result.data.code === 0){
      message.success("打卡成功")
      isCheck.value = true
    }
    else{
      message.warning("你近期已经打过卡了")
    }
  }
  else{
    message.error("网络请求出错了")
  }
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