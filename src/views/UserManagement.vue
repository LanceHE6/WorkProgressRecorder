<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          数据管理
        </template>
      </n-page-header>
    </n-row>

    <n-row>
      <n-tabs type="line" size="large" animated>
        <n-tab-pane name="user" tab="用户管理">
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
                clearable
                :options="[
                    {label: '未填写', value: 0},
                    {label: '考研', value: 1},
                    {label: '就业', value: 2},
                    {label: '考研&就业', value: 3},
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
                style="margin-right: 10px"
            >
              导入
            </n-button>
            <n-button
                type="primary"
                size="large"
                @click="downloadUserList"
            >
              导出
            </n-button>
          </n-row>
          <n-data-table
              :columns="userTableHeader"
              :data="userList"
              :bordered="false"
              :scroll-x="1200"
              max-height="65vh"
              striped
              style="font-size: 18px"
          />
          <n-pagination
              :display-order="['pages', 'quick-jumper']"
              :item-count="userState.total"
              :page-sizes="[pageSize]"
              :page-slot="5"
              show-quick-jumper
              style="margin-top: 10px"
              :on-update-page="userPageChange"
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
                clearable
                :options="[
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
              :scroll-x="1200"
              max-height="65vh"
              striped
              style="font-size: 18px"
          />
          <n-pagination
              :display-order="['pages', 'quick-jumper']"
              :item-count="checkinState.total"
              :page-sizes="[pageSize]"
              :page-slot="5"
              show-quick-jumper
              style="margin-top: 10px"
              :on-update-page="checkinPageChange"
          />
        </n-tab-pane>
      </n-tabs>
    </n-row>
  </div>

  <n-modal v-model:show="showGoalModal">
    <n-card
        style="width: 85%"
        title="目标信息"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <div
        v-for="item in goalCol"
      >
          <span v-if="user.direction.includes(item.direction)" class="body-data">
            <span class="data-label">{{item.label}}</span>
            <span v-if="goal && goal[item.property]" class="data-prop">{{`${goal[item.property]}${item.suffix ? item.suffix : ''}`}}</span>
            <span v-else class="data-prop-null">未填写</span>
          </span>
      </div>
      <div
        v-if="user.direction.includes(2)"
      >
        <div class="title-nickname">
          <b>找工作日志</b>
        </div>
        <n-data-table
            :columns="logTableHeader"
            :data="logList"
            :bordered="false"
            :scroll-x="1000"
            max-height="40%"
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
        <n-text>
          请使用样例模版导入，
        </n-text>
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
  <n-modal v-model:show="showLogDetailModal">
    <n-card
        style="width: 80%"
        title="进度详情"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <div class="title-nickname">
        <b>进度</b>
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
  </n-modal>
  <n-modal :show="isExporting">
    <n-card
        style="width: auto; justify-content: center"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-spin>
        <template #description>
          表格导出中，请稍后...
        </template>
      </n-spin>
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
import {mapping} from "../utils/other.js";

const message = useMessage()

const isExporting = ref(false)

const myUpload = ref(null)
const pageSize = 30

const showGoalModal = ref(false)
const showUploadModal = ref(false)
const showLogDetailModal = ref(false)
const logList = ref([])  //日志列表
const log = ref({})  // 选中的日志
const timeline = ref([])  // 选中日志的进度
const goal = ref({})
const user = ref({})

//用户信息显示格式列表
const goalCol = [
  {property: "target_university", label: "目标院校", direction: 1},
  {property: "target_major", label: "目标专业", direction: 1},
  {property: "target_score", label: "目标分数", direction: 1},
  {property: "target_company", label: "目标公司", direction: 2},
  {property: "target_job", label: "目标岗位", direction: 2},
  {property: "target_salary", label: "目标薪资", direction: 2, suffix: 'k'},
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

//用户导出excel表格式
const userListDownloadTem = [
  {label: '学号', prop: 'account', listIndex: 1},
  {label: '班级', prop: 'class', listIndex: 1},
  {label: '姓名', prop: 'name', listIndex: 1},
  {label: '专业', prop: 'major', listIndex: 1},
  {label: '方向', prop: 'direction', listIndex: 1, isMapping: true, mappingList: [
      {label: '未填写', value: 0},
      {label: '考研', value: 1},
      {label: '就业', value: 2},
    ]},
  {label: '目标院校', prop: 'target_university', listIndex: 2},
  {label: '目标专业', prop: 'target_major', listIndex: 2},
  {label: '目标分数', prop: 'target_score', listIndex: 2},

  {label: '找工作状态', prop: 'status', listIndex: 2, isMapping: true, mappingList: [
      {label: '未填写', value: 0},
      {label: '未拿到offer', value: 1},
      {label: '已拿到offer', value: 2},
    ]},
  {label: '目标公司', prop: 'target_company', listIndex: 2},
  {label: '目标职位', prop: 'target_job', listIndex: 2},
  {label: '理想薪资', prop: 'target_salary', listIndex: 2},
  {label: '目标地区', prop: 'target_area', listIndex: 2},
]

const userTableHeader = [
  {title: "序号", key: "number",
    render: (_, index) => {
      return `${index + 1}`;
    },
    width: '80',
    fixed: "left"
  },
  {title: "姓名", key: "name", fixed: "left", width: '80'},
  {title: "学号", key: "account", width: '140'},
  {title: '班级', key: 'class', width: '120'},
  {title: '专业', key: 'major', width: '200'},
  {title: "方向", key: "direction", width: '120',
    render(row) {
      const direction = row.direction
      if(!direction || direction.length === 0){
        return '未填写'
      }
      else{
        let result = ''
        for(const item of direction){
          if(item === 1){
            result += '考研&'
          }
          else if(item === 2){
            result += '就业&'
          }
        }
        return result.slice(0, -1)
      }
    }
  },
  {title: "目标", key: "goal", width: '80',
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
          { default: () => hasGoal(index) ? "点击查看" : "未填写" }
      );
    }
  },
]

const logTableHeader = [
  { title: "公司名称", key: "company_name"},
  { title: "工作岗位", key: "job"},
  { title: "薪资", key: "salary"},
  { title: "工作地区", key: "location"},
  { title: "最新状态", key: "status",
    render(row) {
      let stage
      if(row.status_time_line.length > 0){
        const item = row.status_time_line[row.status_time_line.length - 1]
        switch (item.stage){
          case 1:{
            stage = '进行中'
            break
          }
          case 2:{
            stage = '成功'
            break
          }
          default:{
            stage = '失败'
            break
          }
        }
        return `${item.status}（${stage}）`
      }
      else{
        return '无'
      }
    }
  },
  { title: "最新时间", key: "time",
    render(row) {
      return row.status_time_line.length > 0 ? new Date(row.status_time_line[row.status_time_line.length - 1].created_at).toLocaleString() : '无'
    }
  },
  { title: "进度详情", key: "actions",
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
          { default: () => "点击查看" }
      );
    }
  }
]

const checkinTableHeader = [
  {title: "序号", key: "number",
    render: (_, index) => {
      return `${index + 1}`;
    }
  },
  { title: "学号", key: "user",
    render(row) {
      return row.user ? row.user.account : '无'
    }
  },
  { title: "姓名", key: "user",
    render(row) {
      return row.user ? row.user.name : '无'
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
  log.value = row
  timeline.value = [...row.status_time_line].reverse()
  showLogDetailModal.value = true
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
    params: {...userState.params, ...{page_size: pageSize}},
    name: 'get-user-list'
  })
  if(result2 && result2.data){
    //提取user
    userList.value = result2.data.rows.map(item => item.user);
    //提取goal
    goalList.value = result2.data.rows.map(item => item.goal);
    userState.total = result2.data.total
  }
  else{
    message.error("网络请求出错了")
  }
}

async function getCheckinList(){
  const result3 = await axiosGet({
    url: '/clock_in/search',
    params: {...checkinState.params, ...{page_size: pageSize}},
    name: 'get-checkin-list'
  })
  if(result3 && result3.data){
    checkinList.value = result3.data.rows

    checkinState.total = result3.data.total
  }
  else{
    message.error("网络请求出错了")
  }
}

async function showGoal(index){
  goal.value = {...goalList.value[index].empl_goal, ...goalList.value[index].pg_goal}
  user.value = userList.value[index]
  console.log("test", goal.value)

  if(user.value.direction.includes(2)){
    const result = await axiosGet({
      url: '/wl/list',
      params: {
        uid: user.value.id
      },
      name: 'user-get-log'
    })
    if(result && result.data){
      logList.value = result.data.work_logs
    }
    else{
      message.error("获取用户找工作日志失败，请稍后再试")
      return
    }
  }
  showGoalModal.value = true
}

function hasGoal(index){
  return goalList.value[index].empl_goal !== null || goalList.value[index].pg_goal !== null
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
          name: data.name,
          class: data.class,
          major: data.major
        }
        //console.log("data:", data)
        dataList.push(user)
      }
      dataList = Array.from(new Set(dataList))
      console.log("result", dataList)
      await uploadUserList(dataList)
    } catch (error) {
      console.error("excelToJson:", error)
      message.error("导入表格格式错误，建议使用样例模版进行导入")
      return false
    }
  }
  fileReader.readAsBinaryString(file)
}

function downloadTemplate(){
  const workbook = XLSX.utils.book_new();
  const worksheet = XLSX.utils.json_to_sheet([{
    account: '',
    name: '',
    class: '',
    major: ''
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

async function downloadUserList(){
  isExporting.value = true
  const result = await axiosGet({
    url: '/user/export',
    config: {
      responseType: 'blob'
    },
    name: 'export-userList'
  })

  if(!result){
    message.error("网络请求出错了")
    isExporting.value = false
    return
  }

  let blob = new Blob([result], { type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet' })

  let file = new File([blob], '导出表格.xlsx', {
    lastModified: new Date(),
    type: 'application/vnd.openxmlformats-officedocument.spreadsheetml.sheet'
  })

  const a = document.createElement('a')
  a.style.display = 'none'
  document.body.appendChild(a)
  const url = URL.createObjectURL(file)
  a.href = url
  a.download = file.name
  a.click()
  // 清理
  document.body.removeChild(a)
  URL.revokeObjectURL(url)

  isExporting.value = false
  // let userList
  // let goalList
  // const result = await axiosGet({
  //   url: '/user/search',
  //   params: {...userState.params , ...{page_size: 10086 * 10086}},
  //   name: 'get-all-user-list'
  // })
  // if(result && result.data){
  //   //提取user
  //   userList = result.data.rows.map(item => item.user);
  //   //提取goal
  //   goalList = result.data.rows.map(item => {
  //     return {...item.goal.empl_goal, ...item.goal.pg_goal}
  //   });
  // }
  // else{
  //   message.error("网络请求出错了")
  //   return
  // }
  //
  //
  // const myList = getListTem(userListDownloadTem, userList, goalList)
  //
  // const workbook = XLSX.utils.book_new();
  // const worksheet = XLSX.utils.json_to_sheet(myList.listResult)
  // XLSX.utils.book_append_sheet(workbook, worksheet, "sheet1")
  // XLSX.utils.sheet_add_aoa(worksheet, [myList.listHeader], { origin: "A1" });
  // worksheet["!cols"] = new Array(myList.listHeader.length).fill({ wch: 15 });
  // XLSX.writeFileXLSX(workbook, "导出表格.xlsx")
}

function getListTem(tem, list1, list2){
  if(list1.length !== list2.length){
    console.error("getListTem: the length of two lists are not equal")
    return
  }
  const listHeader = []
  const listResult = []
  for(const item of tem){
    listHeader.push(item.label)
  }
  for(let i = 0; i < list1.length; i++){
    const listItem = {}
    for(const item of tem){
      if(item.listIndex === 1){
        listItem[item.prop] = list1[i] && item.prop in list1[i] ? mapping(item, list1[i][item.prop]) : ''
      }
      else if(item.listIndex === 2){
        listItem[item.prop] = list2[i] && item.prop in list2[i] ? mapping(item, list2[i][item.prop]) : ''
      }
    }
    listResult.push(listItem)
  }
  return {
    listHeader: listHeader,
    listResult: listResult
  }
}

async function userPageChange(page){
  userState.params.page = page
  await getUserList()
}

async function checkinPageChange(page){
  checkinState.params.page = page
  await getCheckinList()
}
</script>

<style scoped>

</style>