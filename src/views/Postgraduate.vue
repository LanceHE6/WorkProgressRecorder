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
    <n-row style="margin-top: 20px">
      <n-date-picker
          v-model:formatted-value="checkinListParams.date"
          value-format="yyyy-MM-dd"
          type="date"
          placeholder="按日期查询"
          clearable
          style="width: 35%; margin-right: 15px; text-align: start"
      >

      </n-date-picker>
      <n-select
          v-model:value="checkinListParams.time_slot"
          placeholder="按时间段查询"
          :options="[
            {label: '上午', value: '上午'},
            {label: '下午', value: '下午'},
            {label: '晚上', value: '晚上'},
          ]"
          clearable
          style="width: 35%; margin-right: 30px; text-align: start"
      >

      </n-select>
      <n-button
        type="primary"
        @click="getCheckinList"
        secondary
      >
         搜索
      </n-button>
    </n-row>
    <n-row style="justify-content: center; margin-top: 10px">
      <n-data-table
          :columns="checkinHead"
          :data="checkinList"
          :bordered="false"
          max-height="40vh"
          style="width: 100%; height: 40vh; background-color: #f9f9f9c0;"
      />
    </n-row>
    <n-pagination
        :display-order="['pages', 'quick-jumper']"
        :item-count="total"
        :page-sizes="[pageSize]"
        :page-slot="5"
        show-quick-jumper
        style="margin-top: 10px"
        :on-update-page="checkinPageChange"
    />
  </div>

  <n-modal
      v-model:show="showCheckInModal"
      preset="dialog"
      title="学习打卡"
      positive-text="打卡"
      negative-text="算了"
      @positive-click="checkIn">
          <div>
            <n-button
                v-for="item in options"
                @click="place = item.value"
                type="success"
                secondary
                style="width: 45%; margin: 5px">
              {{ item.label }}
            </n-button>
          </div>

    <n-form-item path="place" label="打卡地点(选择后可以手动补充详细地点)">
      <n-input
          v-model:value="place"
          placeholder="请选择或手动输入打卡地点"
      />
    </n-form-item>

  </n-modal>
</template>

<script setup>
import {GolfOutline} from '@vicons/ionicons5'
import {router} from "../router/index.js";
import {onMounted, ref} from "vue";
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {NButton, useMessage} from "naive-ui";
import {getUser} from "../utils/appManager.js";

const message = useMessage()
const pageSize = 5

const examDay = ref(null)
const goal = ref(null)
const showCheckInModal = ref(false)

const isCheck = ref(true)

const place = ref(null)
const total = ref(0)
const checkinList = ref([])
const checkinListParams = ref({
  uid: '',
  date: null,
  time_slot: null,
  page: 1,
})

function back(){
  router.push("/")
}

const options = [
  {label: '数字图书馆', value: '数字图书馆'},
  {label: '老图书馆', value: '老图书馆'},
  {label: '教学楼', value: '教学楼'},
  {label: '实验楼', value: '实验楼'},
  {label: '信科楼', value: '信科楼'},
]

const checkinHead = [
  {title: "序号", key: "number",
    render: (_, index) => {
      return `${index + 1}`;
    }
  },
  { title: "打卡地点", key: "clock_in_location"},
  { title: "打卡时间", key: "clock_in_time", render: (row) => {
      return new Date(row.clock_in_time * 1000).toLocaleString()
    }},
]

onMounted(async () =>{
  const result = await axiosGet({
    url: '/user/target',
    name: 'postList-get-empl_goal'
  })
  if(result && result.data && result.data.direction.includes(1)){
    goal.value = result.data.goal.pg_goal
  }
  else{
    message.warning("您没有确定考研目标")
  }

  const result2 = await axiosGet({
    url: '/clock_in/is_clock_in',
    name: 'postList-get-clock'
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
    name: 'postList-get-exam-day'
  })
  if(result3 && result3.data){
    examDay.value = result3.data.days
  }
  else{
    message.error("考研时间获取失败")
  }
  await getCheckinList()
})

async function getCheckinList(){
  checkinListParams.value.uid = getUser().id
  if(!checkinListParams.value.uid){
    message.error("身份验证失败，请重新登录")
    return
  }
  const result = await axiosGet({
    url: '/clock_in/search',
    params: checkinListParams.value,
    name: 'get-checkin-list'
  })
  if(result && result.data){
    checkinList.value = result.data.rows
    total.value = result.data.total
  }
  else{
    message.error("打卡记录获取失败")
  }
}

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
      await getCheckinList()
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

async function checkinPageChange(page){
  checkinListParams.value.page = page
  await getCheckinList()
}

</script>

<style scoped>

</style>