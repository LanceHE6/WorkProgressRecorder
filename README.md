<p align="center">
<img alt="logo" src="./imgs/logo.ico">
</p>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">WorkProgressRecorder</h1>
<h1 align="center" style="margin: 30px 0 30px; font-weight: bold;">WPR考研就业信息管理平台</h1>
<div align="center"> 



</div>

<h4 align="center">基于gin+Vue3前后端分离式考研就业信息管理平台</h4>
<div align="center">


![Static Badge](https://img.shields.io/badge/Licence-MIT-blue)
![Static Badge](https://img.shields.io/badge/前端-vue-orange)
![Static Badge](https://img.shields.io/badge/后端-gin-green)

</div>


## 平台简介

* 前端技术栈 [Vue3](https://v3.cn.vuejs.org) + [Naive UI](https://www.naiveui.com/zh-CN/os-theme) + [Vite](https://cn.vitejs.dev) 。
* 后端[gin](https://gin-gonic.com/zh-cn/)+[gorm](https://gorm.io/zh_CN/docs/index.html)。
* 数据库[mysql]([MySQL](https://www.mysql.com/cn/))。

## 内置功能

### 权限组

#### 普通用户

* [x] 学习打卡
* [x] 本人打卡记录查询
* [x] 管理考研或就业目标信息
* [x] 管理就业日志
* [x] 管理找工作心得

#### 管理员

* [x] 继承普通用户权限
* [x] 导入用户账号
* [x] 导出用户数据
* [x] 查看用户信息(考研,就业目标信息,就业日志)
* [x] 修改用户密码
* [x] 查看用户打卡记录

## 部署

#### 直接部署

##### 前端

```bash
# 克隆项目
git clone https://github.com/LanceHE6/WorkProgressRecorder.git

# 进入项目目录
cd vue-web

# 安装依赖
npm install

# 启动服务
npm run dev

# 前端默认访问地址 http://localhost:5173
```

##### 后端

```bash
# 进入项目目录
cd go-server

# 下载依赖
go mod download

# 构建项目
go build -o main .
```

----

#### Docker部署

##### 前端

```bash
# 克隆项目
git clone https://github.com/LanceHE6/WorkProgressRecorder.git

# 进入项目目录
cd vue-web

# 构建镜像
docker build -t wpr-web:latest .

# 运行容器
docker run --name wpr-web -p 5173:5173 -d wpr-web:latest
```

##### 后端

```bash
# 进入项目目录
cd go-server

# 构建镜像
docker build -t wpr-server:latest .

# 运行容器
docker run --name wpr-server -p 8080:8080 -d wpr-server:latest
```



## 文档

[后端接口文档](./api.md)

[后端更新日志](./change_log.md)