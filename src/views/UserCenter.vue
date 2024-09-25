<template>
  <div class="main-body">
    <n-row style="justify-content: left">
      <n-page-header subtitle="" @back="back">
        <template #title>
          用户中心
        </template>
      </n-page-header>
    </n-row>

    <n-spin :show="isLoading">
      <n-row>
        <n-card class="profile-title-card">
          <div class="title-main">
            <div class="title-nickname">
              <b>{{user.name}}</b>
              <n-button
                  text
                  type="primary"
                  size="small"
                  style="margin-left: auto"
                  icon="Edit"
                  @click="openEditForm"
              >
                修改密码
              </n-button>
            </div>
            <div class="small-text" style="margin-bottom: 10px">{{`用户ID：${user.id}`}}</div>
            <div class="small-text">{{`创建时间：${new Date(user.created_at).toLocaleString()}`}}</div>
          </div>
        </n-card>
      </n-row>

      <n-row>
        <n-card class="profile-body-card">
          <div class="title-nickname">
            <b>个人资料</b>
            <n-button
                text
                type="primary"
                size="small"
                style="margin-left: auto"
                icon="Edit"
                @click="openInfoForm"
            >
              编辑
            </n-button>
          </div>
          <n-divider></n-divider>
          <div
              v-for="item in infoDataCol"
          >
          <span class="body-data">
            <span class="data-label">{{item.label}}</span>
            <span v-if="user && user[item.property]" class="data-prop">{{`${mapping(item, user[item.property])}${item.suffix ? item.suffix : ''}`}}</span>
            <span v-else class="data-prop-null">未填写</span>
          </span>
          </div>
        </n-card>
      </n-row>

      <n-row>
        <n-card class="profile-body-card">
          <div class="title-nickname">
            <b>方向信息</b>
            <n-button
                text
                type="primary"
                size="small"
                style="margin-left: auto"
                icon="Edit"
                @click="openDirectionForm"
            >
              编辑
            </n-button>
          </div>
          <n-divider></n-divider>
          <div
              v-for="item in directionDataCol"
          >
          <span v-if="direction.includes(item.direction)" class="body-data">
            <span class="data-label">{{item.label}}</span>
            <span v-if="item.direction === 1 && pg_goal && pg_goal[item.property]" class="data-prop">{{ `${mapping(item, pg_goal[item.property])}${item.suffix ? item.suffix : ''}` }}</span>
            <span v-else-if="item.direction === 2 && empl_goal && empl_goal[item.property]" class="data-prop">{{ `${mapping(item, empl_goal[item.property])}${item.suffix ? item.suffix : ''}` }}</span>
            <span v-else class="data-prop-null">未填写</span>
          </span>
          </div>
        </n-card>
      </n-row>
    </n-spin>
  </div>

  <n-modal v-model:show="editFormVisible">
    <n-card
        style="width: 85%"
        title="修改密码"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="editFormRef" :model="editForm" :rules="rules">
        <n-form-item path="old_password" label="旧密码">
          <n-input
              v-model:value="editForm.old_password"
              type="password"
              placeholder="请输入旧密码"
              show-password-on="click"
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="new_password" label="新密码">
          <n-input
              v-model:value="editForm.new_password"
              type="password"
              placeholder="请输入新密码"
              show-password-on="click"
              @keydown.enter.prevent
          />
        </n-form-item>
        <n-form-item path="confirm" label="确认密码">
          <n-input
              v-model:value="editForm.confirm"
              type="password"
              placeholder="请确认密码"
              show-password-on="click"
              @keydown.enter.prevent
          />
        </n-form-item>

        <n-row style="justify-content: center">
          <n-button type="primary" @click="submitEditForm">
            修改密码
          </n-button>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
  <n-modal v-model:show="directionFormVisible">
    <n-card
        style="width: 85%"
        title="编辑方向信息"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="directionRef" :model="directionForm" :rules="rules">
        <n-form-item path="direction" label="方向">
          <n-select
              v-model:value="directionForm.direction"
              :options="[{label: '考研', value: 1},
                         {label: '就业', value: 2},]"
              placeholder="请选择方向"
          />
        </n-form-item>

        <n-form-item v-if="directionForm.direction === 1" path="target_university" label="目标院校">
          <n-input
              v-model:value="directionForm.target_university"
              placeholder="请输入目标院校"
          />
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 1" path="target_major" label="目标专业">
          <n-input
              v-model:value="directionForm.target_major"
              placeholder="请输入目标专业"
          />
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 1" path="target_score" label="目标分数">
          <n-input-number
              v-model:value="directionForm.target_score"
              :show-button="false"
              placeholder="请输入目标分数"
              style="width: 100%"
          />
        </n-form-item>

        <n-form-item v-if="directionForm.direction === 2" path="status" label="状态">
          <n-select
              v-model:value="directionForm.status"
              :options="[{label: '未拿到offer', value: 1},
                         {label: '已拿到offer', value: 2},]"
              placeholder="请选择就业状态"
          />
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 2" path="target_company" label="目标公司">
          <n-input
              v-model:value="directionForm.target_company"
              placeholder="请输入目标公司"
          />
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 2" path="target_job" label="目标职位">
          <n-input
              v-model:value="directionForm.target_job"
              placeholder="请输入目标职位"
          />
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 2" path="target_salary" label="理想薪资">
          <n-input-number
              v-model:value="directionForm.target_salary"
              :show-button="false"
              placeholder="请输入理想薪资"
              style="width: 100%"
          >
            <template #suffix>
              千元
            </template>
          </n-input-number>
        </n-form-item>
        <n-form-item v-if="directionForm.direction === 2" path="target_area" label="目标地区">
          <n-input
              v-model:value="directionForm.target_area"
              placeholder="请输入目标地区"
          />
        </n-form-item>

        <n-row style="justify-content: center">
          <n-button type="primary" @click="submitDirectionForm">
            确定方向
          </n-button>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
  <n-modal v-model:show="infoFormVisible">
    <n-card
        style="width: 85%"
        title="编辑个人资料"
        :bordered="false"
        role="dialog"
        aria-modal="true"
    >
      <n-form ref="infoRef" :model="infoForm" :rules="rules">
        <n-form-item path="class" label="班级">
          <n-input
              v-model:value="infoForm.class"
              placeholder="请输入班级"
          />
        </n-form-item>
        <n-form-item path="major" label="专业">
          <n-input
              v-model:value="infoForm.major"
              placeholder="请输入专业"
          />
        </n-form-item>

        <n-row style="justify-content: center">
          <n-button type="primary" @click="submitInfoForm">
            确定
          </n-button>
        </n-row>
      </n-form>
    </n-card>
  </n-modal>
</template>

<script setup>
import {router} from "../router/index.js";
import {onMounted, ref} from "vue";
import {axiosGet, axiosPost, axiosPut} from "../utils/axiosUtil.js";
import {CURRENT_USER, setUser} from "../utils/appManager.js";
import {useMessage} from "naive-ui";
import {isSame, pIntValidatorRequire, pNumValidatorRequire} from "../utils/validator.js";
import {mapping} from "../utils/other.js";

const message = useMessage()
const isLoading = ref(true)

const editFormRef = ref(null)
const editFormVisible = ref(false)
const editForm = ref({
  old_password: '',
  new_password: '',
  confirm: '',
})

const directionRef = ref(null)
const directionFormVisible = ref(false)
const directionForm = ref({
  direction: 1,
  status: 1,
  target_university: '',
  target_major: '',
  target_score: null,
  target_company: '',
  target_job: '',
  target_salary: null,
  target_area: '',
})

const infoRef = ref(null)
const infoFormVisible = ref(false)
const infoForm = ref({
  class: '',
  major: '',
})

const rules = {
  old_password: [
    { required: true, message: '请输入旧密码', trigger: ['blur', 'input']},
  ],
  new_password: [
    { required: true, message: '请输入新密码', trigger: ['blur', 'input']},
  ],
  confirm: [
    { required: true, message: '请确认新密码', trigger: ['blur', 'input']},
    { validator: (rule, value, callback) => isSame(rule, value, callback, editForm.value.new_password, '两次输入的密码不一致'), trigger: 'blur' },
  ],
  target_university: [
    { required: true, message: '请输入目标院校', trigger: ['blur', 'input']},
  ],
  target_major: [
    { required: true, message: '请输入目标专业', trigger: ['blur', 'input']},
  ],
  target_score: [
    { validator: pIntValidatorRequire, trigger: 'blur' },
  ],
  target_company: [
    { required: true, message: '请输入目标公司', trigger: ['blur', 'input']},
  ],
  target_job: [
    { required: true, message: '请输入目标职位', trigger: ['blur', 'input']},
  ],
  target_salary: [
    { validator: pNumValidatorRequire, trigger: 'blur' },
  ],
  target_area: [
    { required: true, message: '请输入目标地区', trigger: ['blur', 'input']},
  ],
  class: [
    { required: true, message: '请输入班级', trigger: ['blur', 'input']},
  ],
  major: [
    { required: true, message: '请输入专业', trigger: ['blur', 'input']},
  ],
}

const direction = ref([1])
const pg_goal = ref(null)
const empl_goal = ref(null)
const user = CURRENT_USER

//个人信息显示格式列表
const infoDataCol = [
  {property: "account", label: "账号"},
  {property: "class", label: "班级"},
  {property: "major", label: "专业"},
  {property: "direction", label: "方向", isMappingList: true, mappingList:[
      {label: '未填写', value: 0},
      {label: '考研', value: 1},
      {label: '就业', value: 2},
    ]},
]

//方向信息显示格式列表
const directionDataCol = [
  {property: "target_university", label: "目标院校", direction: 1},
  {property: "target_major", label: "目标专业", direction: 1},
  {property: "target_score", label: "目标分数", direction: 1},
  {property: "target_company", label: "目标公司", direction: 2},
  {property: "target_job", label: "目标岗位", direction: 2},
  {property: "target_salary", label: "期望薪资", direction: 2, suffix: 'k'},
  {property: "target_area", label: "目标地区", direction: 2}
]

function back(){
  router.push("/")
}

onMounted(async () => {
  await refresh()
})

const openEditForm = () => {
  editFormVisible.value = true
}
const submitEditForm = async () => {
  if (!editFormRef) return
  await editFormRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await axiosPut({
        url:'/user/update/psw',
        data: editForm.value,
        name: 'userCenter-change-password'
      })
      if(result){
        if(result.data.code === 0){
          message.success("密码修改成功")
          editFormVisible.value = false
          for(const item in editForm.value){
            editForm.value[item] = ''
          }
          await refresh()
        }
        else if(result.data.code === -2){
          message.error("原密码错误")
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

const openDirectionForm = () => {
  directionForm.value.direction = direction.value[0]
  for(const i in directionForm.value){
    if(pg_goal.value && pg_goal.value[i]){
      directionForm.value[i] = pg_goal.value[i]
    }
    else if(empl_goal.value && empl_goal.value[i]){
      directionForm.value[i] = empl_goal.value[i]
    }
  }
  directionFormVisible.value = true
}
const submitDirectionForm = async () => {
  if (!directionRef) return
  console.log(directionForm.value)
  await directionRef.value?.validate(async (errors) => {
    if (!errors) {
      let url = ''
      if(directionForm.value.direction === 1){
        url = '/pggl/add'
      }
      else if(directionForm.value.direction === 2){
        url = '/emgl/add'
        directionForm.value.target_salary = directionForm.value.target_salary.toString()
      }
      else{
        message.error("方向类型错误：", directionForm.value.direction)
        return
      }
      const result = await axiosPost({
        url: url,
        data: directionForm.value,
        name: 'userCenter-edit-direction'
      })
      if(result){
        if(result.data.code === 0){
          message.success("方向确定成功")
          directionFormVisible.value = false
          await refresh()
        }
        else {
          message.error("后端处理出错了")
        }
      }
      else{
        message.error("网络请求出错了")
      }
    }
  })
}

const openInfoForm = () => {
  for(const i in infoForm.value){
    if(user.value && user.value[i]){
      infoForm.value[i] = user.value[i]
    }
  }
  infoFormVisible.value = true
}
const submitInfoForm = async () => {
  if (!infoRef) return
  await infoRef.value?.validate(async (errors) => {
    if (!errors) {
      const result = await axiosPut({
        url: '/user/update/info',
        data: infoForm.value,
        name: 'userCenter-edit-info'
      })
      if(result){
        if(result.data.code === 0){
          message.success("编辑个人信息成功")
          infoFormVisible.value = false
          await refresh()
        }
        else {
          message.error("后端处理出错了")
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
    name: 'userCenter-refresh'
  })
  if(result && result.data && result.data.user){
    setUser(result.data.user)
  }

  const result2 = await axiosGet({
    url: '/user/target',
    name: 'userCenter-get-empl_goal'
  })
  if(result2.data && result2.data.direction && result2.data.direction.length > 0){
    direction.value = result2.data.direction
    empl_goal.value = result2.data.goal.empl_goal
    pg_goal.value = result2.data.goal.pg_goal
  }
  else{
    message.warning("您还未确定方向！")
    directionFormVisible.value = true
  }
  isLoading.value = false

}
</script>

<style scoped>

</style>