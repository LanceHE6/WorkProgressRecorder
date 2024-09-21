<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          用户管理
        </template>
      </n-page-header>
    </n-row>

    <n-row>
      <n-tabs type="line" size="large" animated>
        <n-tab-pane name="user" tab="用户列表">
          <n-row>
            <n-input
                v-model:value="userState.params.account"
                placeholder="按学号查询"
                style="width: 15%; margin-right: 10px; text-align: start;"
            />
            <n-input
                v-model:value="userState.params.name"
                placeholder="按姓名查询"
                style="width: 15%; margin-right: 10px; text-align: start;"
            />
            <n-select
                v-model:value="userState.params.direction"
                placeholder="按方向查询"
                :options="[
                    {label: '未选择', value: null},
                    {label: '未确认', value: 0},
                    {label: '考研', value: 1},
                    {label: '就业', value: 2},
                ]"
                style="width: 15%; margin-right: 10px; text-align: start"
                class="my-select"
            />

            <n-button
                type="info"
                size="large"
                @click="getUserList"
                style="margin-right: 10px"
            >
              搜索
            </n-button>
            <n-button
                type="primary"
                size="large"
                @click="uploadUser"
            >
              导入
            </n-button>
          </n-row>
          <n-data-table
              :columns="userTableHeader"
              :data="userList"
              :bordered="false"
              striped
              style="font-size: 18px"
          />
        </n-tab-pane>
        <n-tab-pane name="checkin" tab="打卡记录">
          <n-row>
            <n-date-picker
                placeholder="按日期查询"
                v-model:formatted-value="checkinState.params.date"
                value-format="yyyy-MM-dd"
                type="date"
                clearable
                style="width: 15%; margin-right: 10px; text-align: start"
                class="my-select"
            />
            <n-select
                v-model:value="checkinState.params.time_slot"
                placeholder="按时间段查询"
                :options="[
                    {label: '未选择', value: null},
                    {label: '上午', value: '上午'},
                    {label: '下午', value: '下午'},
                    {label: '晚上', value: '晚上'},
                ]"
                style="width: 15%; margin-right: 10px"
                class="my-select"
            />

            <n-button
                type="info"
                size="large"
                @click="getCheckinList"
                style="margin-right: 10px"
            >
              搜索
            </n-button>
          </n-row>
          <n-data-table
              :columns="checkinTableHeader"
              :data="checkinList"
              :bordered="false"
              striped
              style="font-size: 18px"
          />
        </n-tab-pane>
      </n-tabs>
    </n-row>
  </div>

  <n-modal v-model:show="showGoalModal">
    <n-card
        style="width: 85%"
        title="方向信息"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <div
        v-for="item in dataCol"
      >
          <span v-if="user.direction === item.direction" class="body-data">
            <span class="data-label">{{item.label}}</span>
            <span v-if="goal && goal[item.property]" class="data-prop">{{goal[item.property]}}</span>
            <span v-else class="data-prop-null">未填写</span>
          </span>
      </div>
      <div
        v-if="user.direction === 2"
      >
        <span class="body-data">找工作日志</span>
        <n-data-table
            :columns="logTableHeader"
            :data="logs"
            :bordered="false"
            style="font-size: 16px"
        />
      </div>
    </n-card>
  </n-modal>
  <n-modal v-model:show="showUploadModal">
    <n-card
        style="width: 85%"
        title="导入用户"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-upload
          ref="myUpload"
          :default-upload="false"
          directory-dnd
          :custom-request="excelToJson"
      >
        <n-upload-dragger>
          <div style="margin-bottom: 12px">
            <n-icon size="48" :depth="3">
              <ArchiveOutline />
            </n-icon>
          </div>
          <n-text style="font-size: 16px">
            点击或者拖动文件到该区域来上传，重复的数据将会被忽略！
          </n-text>
        </n-upload-dragger>
      </n-upload>
      <n-row style="justify-content: center">
        <n-button
            type="warning"
            size="large"
            @click="downloadTemplate"
            text
        >
          点击此处下载样例模版
        </n-button>
      </n-row>
      <n-row style="justify-content: center; margin-top: 10px">
        <n-button
            type="primary"
            @click="submitUpload"
        >
          导入用户
        </n-button>
      </n-row>
    </n-card>
  </n-modal>
</template>

<script setup>
import {router} from "../router/index.js";
import {h, onMounted, reactive, ref} from "vue";
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {getUserPermission, setUser} from "../utils/appManager.js";
import {NButton, useMessage} from "naive-ui";
import { ArchiveOutline } from "@vicons/ionicons5";
import * as XLSX from "xlsx"

const message = useMessage()

const myUpload = ref(null)

const showGoalModal = ref(false)
const showUploadModal = ref(false)
const logs = ref([])
const goal = ref({})
const user = ref({})

//用户信息显示格式列表
const dataCol = [
  {property: "target_university", label: "目标院校", direction: 1},
  {property: "target_major", label: "目标专业", direction: 1},
  {property: "target_score", label: "目标分数", direction: 1},
  {property: "target_company", label: "目标公司", direction: 2},
  {property: "target_job", label: "目标岗位", direction: 2},
  {property: "target_salary", label: "目标薪资", direction: 2},
  {property: "target_area", label: "目标地区", direction: 2}
]

const userState = reactive({
  params: {
    page: 1,
    account: null,
    name: null,
    direction: null
  },
  total: 1
})

const checkinState = reactive({
  params: {
    page: 1,
    date: null,
    time_slot: null,
  },
  total: 1
})

const userTableHeader = [
  {title: "学号", key: "account"},
  {title: "姓名", key: "name"},
  {title: "方向", key: "direction",
    sorter: (row1, row2) => row1.direction - row2.direction,
    render(row) {
      const direction = row.direction
      switch (direction){
        case 1: {
          return '考研'
        }
        case 2: {
          return '就业'
        }
        default: {
          return '未确认'
        }
      }
    }
  },
  {title: "目标信息", key: "goal",
    render(_, index) {
      return h(
          NButton,
          {
            disabled: !hasGoal(index),
            strong: true,
            tertiary: true,
            type: hasGoal(index) ? "info" : "warning",
            size: "small",
            onClick: () => showGoal(index)
          },
          { default: () => hasGoal(index) ? "查看目标" : "未填写" }
      );
    }
  },
]

const logTableHeader = [
  { title: "公司名称", key: "company_name"},
  { title: "工作岗位", key: "job"},
  { title: "薪资", key: "salary"},
  { title: "工作地区", key: "location"},
  // { title: "操作", key: "actions",
  //   render(row) {
  //     return h(
  //         NButton,
  //         {
  //           strong: true,
  //           tertiary: true,
  //           type: 'info',
  //           size: "small",
  //           onClick: () => showLogDetail(row)
  //         },
  //         { default: () => "详细信息" }
  //     );
  //   }
  // }
]

const checkinTableHeader = [
  { title: "学号", key: "user",
    render(row) {
      return row.user.account
    }
  },
  { title: "姓名", key: "user",
    render(row) {
      return row.user.name
    }
  },
  { title: "打卡地点", key: "clock_in_location"},
  { title: "打卡时间", key: "clock_in_time",
    render(row) {
      return new Date(row.clock_in_time * 1000).toLocaleString()
    }
  },
]

function showLogDetail(row){
  localStorage.setItem("logDetail", JSON.stringify(row))
  router.push("/logDetail")
}

const userList = ref([])
const goalList = ref([])
const checkinList = ref([])

function back(){
  router.push("/")
}

onMounted( async () => {
  const result = await axiosGet({
    url: '/user/info',
    name: 'authorization'
  })
  if(result && result.data && result.data.user){
    setUser(result.data.user)
  }
  else{
    await router.push("/")
    message.warning("登录已过期，请重新登录")
    setUser('')
    localStorage.setItem("token", '')
    return
  }

  if(getUserPermission() < 2){
    await router.push("/")
    message.warning("权限不足")
    return
  }
  await getUserList()
  await getCheckinList()
})

async function getUserList(){
  const result2 = await axiosGet({
    url: '/user/search',
    params: userState.params,
    name: 'get-user-list'
  })
  if(result2 && result2.data){
    //提取user
    userList.value = result2.data.rows.map(item => item.user);
    //提取goal
    goalList.value = result2.data.rows.map(item => item.goal);
  }
  else{
    message.error("网络请求出错了")
  }
}

async function getCheckinList(){
  const result3 = await axiosGet({
    url: '/clock_in/search',
    params: checkinState.params,
    name: 'get-checkin-list'
  })
  if(result3 && result3.data){
    checkinList.value = result3.data.rows
  }
  else{
    message.error("网络请求出错了")
  }
}

async function showGoal(index){
  goal.value = goalList.value[index]
  user.value = userList.value[index]

  if(user.value.direction === 2){
    const result = await axiosGet({
      url: '/wl/list',
      params: {
        uid: user.value.id
      },
      name: 'user-get-log'
    })
    if(result && result.data){
      logs.value = result.data.work_logs
    }
    else{
      message.error("获取用户找工作日志失败，请稍后再试")
      return
    }
  }
  showGoalModal.value = true
}

function hasGoal(index){
  return goalList.value[index] !== null
}

function uploadUser(){
  showUploadModal.value = true
}

function submitUpload(){
  myUpload.value.submit()
  showUploadModal.value = false
}

function excelToJson(e){
  let file = e.file.file // 文件信息
  if (!file) {
    // 没有文件
    return false
  } else if (!/\.(xls|xlsx)$/.test(file.name.toLowerCase())) {
    // 格式根据自己需求定义
    message.error('上传格式不正确，请上传xls或者xlsx格式')
    return false
  }

  const fileReader = new FileReader()
  fileReader.onload = async (ev) => {
    try {
      const data = ev.target.result
      const workbook = XLSX.read(data, {
        type: 'binary' // 以字符编码的方式解析
      })
      const excelName = workbook.SheetNames[0] // 取第一张表
      const exl = XLSX.utils.sheet_to_json(workbook.Sheets[excelName]) // 生成json表格内容
      //console.log("excelToJson:", exl)
      // 将 JSON 数据上传给服务器
      let dataList = []
      for (const item in exl) {
        const data = exl[item]
        const user = {
          account: data.account.toString(),
          name: data.name
        }
        //console.log("data:", data)
        dataList.push(user)
      }
      dataList = Array.from(new Set(dataList))
      console.log("result", dataList)
      await uploadUserList(dataList)
    } catch (error) {
      console.error("excelToJson:", error)
      return false
    }
  }
  fileReader.readAsBinaryString(file)
}

function downloadTemplate(){
  const workbook = XLSX.utils.book_new();
  const worksheet = XLSX.utils.json_to_sheet([{
    account: '',
    name: ''
  }])
  XLSX.utils.book_append_sheet(workbook, worksheet, "sheet1")
  worksheet["!cols"] = new Array(2).fill({ wch: 15 });
  XLSX.writeFileXLSX(workbook, "提交模版.xlsx")
}

async function uploadUserList(userList){
  const result = await axiosPost({
    url: '/user/import',
    data: {
      list: userList
    },
    name: 'user-import'
  })

  if(result && result.data){
    if(result.data.code === 0){
      message.success("用户导入成功")
    }
    else{
      message.warning("导入存在错误")
    }
  }
  else{
    message.error("网络请求出错了")
  }
  await getUserList()
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
.body-data{
  display: flex;
  margin: 20px 0 20px 10px;
  justify-content: flex-start;
}
.body-data-span{
  display: flex;
  justify-content: flex-start;
}
.data-label{
  display: flex;
  align-items: center; /* 垂直居中 */
  width: 100px;
  font-size: 15px;
}
.data-prop{
  display: flex;
  align-items: center; /* 垂直居中 */
  font-size: 15px;
}
.data-prop-null{
  display: flex;
  align-items: center; /* 垂直居中 */
  color: #8c939d;
}
.my-select :deep(.n-base-selection-label) {
  height: 40px;
}
.my-select :deep(.n-input__input-el){
  height: 40px;
}
</style>