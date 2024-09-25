<template>
  <div class="main-body">
    <n-row style="padding-bottom: 10px">
      <n-gradient-text type="info" font-size="medium">
        您好，{{state.user ? state.user.name : ''}}
      </n-gradient-text>
      <n-button
          v-if="!state.user"
          type="warning"
          text
          @click="state.showModal = true"
      >
        点击此处登录
      </n-button>
      <n-button
          v-if="state.user"
          type="error"
          text
          @click="logout"
          style="margin-left: auto"
      >
        退出登录
      </n-button>
    </n-row>

    <n-row>
      <n-icon size="30" color="#2080f0">
        <CalendarOutline />
      </n-icon>
      <n-gradient-text type="info" font-size="large" style="margin-left: 10px">
        {{getCurrentDateFormatted()}}
      </n-gradient-text>
    </n-row>

    <n-row>
      <n-carousel autoplay>
        <img
            class="carousel-img"
            src="/1.png"
            alt=""
        >
        <img
            class="carousel-img"
            src="/2.png"
            alt=""
        >
        <img
            class="carousel-img"
            src="/3.png"
            alt=""
        >
        <img
            class="carousel-img"
            src="/4.png"
            alt=""
        >
      </n-carousel>
    </n-row>

    <n-row style="padding-top: 5px">
      <n-icon size="30" color="#2080f0">
        <AppsOutline />
      </n-icon>
      <n-gradient-text type="info" font-size="large" style="margin-left: 10px">
        功能列表
      </n-gradient-text>
    </n-row>

    <n-row style="margin-top: 10px">
      <n-grid :cols="4">
        <n-gi
          v-for="item in toolsList"
        >
          <div v-if="item.permission === 0 || (item.permission <= state.user.permission)">
            <n-row style="justify-content: center;">
              <n-button class="main-btn" :color="item.color" @click="jumpTo(item.path)">
                <template #icon>
                  <n-icon :component="item.icon" size="24px"></n-icon>
                </template>
              </n-button>
            </n-row>
            <n-row style="justify-content: center;">
              <n-text style="font-size: 15px">
                {{item.name}}
              </n-text>
            </n-row>
          </div>
        </n-gi>
      </n-grid>
    </n-row>
    <n-text class="footer">
      {{`${FRONTEND_VERSION}_${BACKEND_VERSION}`}}
    </n-text>
  </div>

  <n-modal
      v-model:show="state.showModal"
      preset="dialog"
      title="用户登录"
      positive-text="登录"
      @positive-click="login"
      :mask-closable="false"
  >
      <n-form ref="formRef" :model="state.loginForm" :rules="state.rules">
        <n-form-item path="account" label="学号">
          <n-input v-model:value="state.loginForm.account" @keydown.enter.prevent placeholder="请输入10位学号"/>
        </n-form-item>
        <n-form-item path="password" label="密码">
          <n-input
              v-model:value="state.loginForm.password"
              type="password"
              placeholder="请输入密码"
              show-password-on="click"
              @keydown.enter.prevent
          />
        </n-form-item>

        <n-row style="justify-content: center; padding-bottom: 10px">
          <n-gradient-text type="warning" font-size="small">
            初始密码为学号后6位
          </n-gradient-text>
        </n-row>
      </n-form>
  </n-modal>
</template>

<script setup>
import { CalendarOutline, AppsOutline, GolfOutline, ReceiptOutline, PersonOutline, PeopleOutline, SettingsOutline } from '@vicons/ionicons5'
import {router} from "../router/index.js";
import {onMounted, reactive, ref} from "vue";
import { useMessage } from 'naive-ui'
import {axiosGet, axiosPost} from "../utils/axiosUtil.js";
import {BACKEND_VERSION, CURRENT_USER, FRONTEND_VERSION, getUser, setUser} from "../utils/appManager.js";

const message = useMessage();
const formRef = ref(null)

const state = reactive({
  loginForm: {
    account: '',
    password: ''
  },
  rules: {
    account: [
      { required: true, message: '请输入学号', trigger: ['blur', 'input']},
      { max: 10, min: 10, message: '请输入10位学号', trigger: ['blur', 'input']}
    ],
    password: [
      { required: true, message: '请输入密码', trigger: ['blur', 'input']},
    ],
  },
  user: CURRENT_USER,
  showModal: false,
})

onMounted( async () =>{
  if(localStorage.getItem("token")){
    const result = await axiosGet({
      url: '/user/info',
      name: 'auto_login'
    })
    if(result && result.code === 0){
      setUser(result.data.user)
    }
    else{
      message.warning("登录已过期，请重新登录")
      setUser('')
      localStorage.setItem("token", '')
    }
  }
  await getBackendVersion()
})

function getCurrentDateFormatted() {
  // 获取当前日期
  const now = new Date();

  // 获取年、月、日
  // 注意：getMonth() 方法返回的月份是从 0（一月）到 11（十二月）的整数
  const year = now.getFullYear();
  const month = String(now.getMonth() + 1).padStart(2, '0'); // 月份加1，并且保证是两位数
  const day = String(now.getDate()).padStart(2, '0'); // 保证是两位数

  // 格式化日期

  // 返回格式化后的日期
  return `${year}年${month}月${day}日`;
}

function jumpTo(path){
  if(path === '/experiment'){
    message.warning('下次更新绝对做出来QAQ，马上搞完了！！！')
    return
  }
  if(getUser() || path === '/experiment'){
    router.push(path)
  }
  else{
    message.warning('请先登录！')
    state.showModal = true
  }
}

async function login(){
  if (!formRef) return
  await formRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await axiosPost(
          {
            url: '/user/login',
            data: state.loginForm,
            name: 'login',
          }
      )
      if (result) {
        if (result.data.code === 0) {
          setUser(result.data.data.user)
          localStorage.setItem("token", result.data.data.token)
          message.success('登录成功！')
          state.showModal = false
        } else {
          message.error('账号或密码错误！')
        }
      } else {
        message.error('网络请求错误！')
      }
    }
  })
}

async function getBackendVersion(){
  const result = await axiosGet({
    url: '/ping',
    name: 'get-backend-version',
  })
  if(result && result.data){
    BACKEND_VERSION.value = result.data.version
  }
  else{
    console.error("获取后端版本失败")
  }
}

function logout(){
  setUser('')
  localStorage.setItem("token", '')
  message.success('已退出登录')
}

const toolsList = [
  {name: '考研打卡', path: '/postgraduate', color: '#2080f0', icon: GolfOutline, permission: 0},
  {name: '找工作日志', path: '/occupation', color: '#8a2be2', icon: ReceiptOutline, permission: 0},
  {name: '个人中心', path: '/userCenter', color: '#ff69b4', icon: PersonOutline, permission: 0},
  {name: '找工作心得', path: '/experiment', color: '#36ad6a', icon: PeopleOutline, permission: 0},
  {name: '用户管理', path: '/userManagement', color: '#f0a020', icon: SettingsOutline, permission: 2},
]
</script>

<style scoped>

</style>