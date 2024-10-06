---
title: WorkProgressRecorder
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"


---

# WorkProgressRecorder

Base URLs:

* <a href="http://127.0.0.1:8070">开发环境: http://127.0.0.1:8070</a>

# Authentication

- HTTP Authentication, scheme: bearer

# 用户

## POST 登录

POST /api/v1/user/login

> Body 请求参数

```yaml
account: "8888888888"
password: admin

```

### 请求参数

| 名称         | 位置   | 类型   | 必选 | 说明 |
| ------------ | ------ | ------ | ---- | ---- |
| Content-Type | header | string | 否   | none |
| body         | body   | object | 否   | none |
| » account    | body   | string | 是   | 学号 |
| » password   | body   | string | 是   | 密码 |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "token": "enim non Excepteur veniam eu",
    "user": {
      "id": 123456879,
      "created_at": "1989-01-07 09:10:45",
      "updated_at": "2012-12-14 08:32:57",
      "account": "do deserunt ad quis adipisicing",
      "name": "务素文国起化南",
      "permission": 1
    }
  }
}
```

```json
{
  "code": 1,
  "msg": "账号或密码错误",
  "data": null
}
```

```json
{
  "code": -1,
  "msg": "生成token失败",
  "data": {
    "err": ""
  }
}
```

```json
{
  "code": -2,
  "msg": "更新session_id失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称           | 类型    | 必选  | 约束 | 中文名 | 说明 |
| -------------- | ------- | ----- | ---- | ------ | ---- |
| » code         | integer | true  | none |        | none |
| » msg          | string  | true  | none |        | none |
| » data         | object  | false | none |        | none |
| »» token       | string  | true  | none |        | none |
| »» user        | object  | true  | none |        | none |
| »»» id         | integer | true  | none |        | none |
| »»» created_at | string  | true  | none |        | none |
| »»» updated_at | string  | true  | none |        | none |
| »»» account    | string  | true  | none |        | none |
| »»» name       | string  | true  | none |        | none |
| »»» permission | integer | true  | none |        | none |

## POST 导入用户

POST /api/v1/user/import

单个用户数据结构
{
"account":"2021214444" ,
"name": "zzt",
"class":"13202107",
"major":"软工"
}

> Body 请求参数

```json
{
  "list": [
    {
      "account": "2021214286",
      "name": "hml",
      "class": "13202107",
      "major": "软工"
    },
    {
      "account": "2021214285",
      "name": "lj",
      "class": "13202107",
      "major": "软工"
    }
  ]
}
```

### 请求参数

| 名称   | 位置 | 类型   | 必选 | 中文名   | 说明 |
| ------ | ---- | ------ | ---- | -------- | ---- |
| body   | body | object | 否   |          | none |
| » list | body | object | 是   | 用户列表 | none |

> 返回示例

```json
{
  "code": 100,
  "msg": "导入存在错误",
  "data": {
    "errs": {
      "2021214285": "Error 1062: Duplicate entry '2021214285' for key 'users.account'",
      "2021214286": "Error 1062: Duplicate entry '2021214286' for key 'users.account'"
    }
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## PUT 更改密码

PUT /api/v1/user/update/psw

> Body 请求参数

```json
{
  "id": 2875660435460096,
  "old_password": "admin",
  "new_password": "123456"
}
```

### 请求参数

| 名称           | 位置 | 类型    | 必选 | 中文名 | 说明   |
| -------------- | ---- | ------- | ---- | ------ | ------ |
| body           | body | object  | 否   |        | none   |
| » id           | body | integer | 否   |        | 用户id |
| » old_password | body | string  | 否   |        | 原密码 |
| » new_password | body | string  | 是   |        | 新密码 |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": -2,
  "msg": "原密码错误",
  "data": null
}
```

```json
{
  "code": -1,
  "msg": "用户不存在",
  "data": null
}
```

```json
{
  "code": -3,
  "msg": "更新密码失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## PUT 更新用户信息

PUT /api/v1/user/update/info

> Body 请求参数

```json
{
  "class": "13002107",
  "major": "微电子"
}
```

### 请求参数

| 名称    | 位置 | 类型   | 必选 | 中文名 | 说明 |
| ------- | ---- | ------ | ---- | ------ | ---- |
| body    | body | object | 否   |        | none |
| » class | body | string | 是   |        | 班级 |
| » major | body | string | 是   |        | 专业 |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": -1,
  "msg": "用户不存在",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "更新失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## GET 获取用户目标信息

GET /api/v1/user/target

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "direction": [
      1,
      2
    ],
    "goal": {
      "empl_goal": {
        "id": 2867451477889024,
        "created_at": "2024-09-22T20:02:33+08:00",
        "updated_at": "2024-09-22T20:09:59+08:00",
        "uid": 2866331728744448,
        "status": 1,
        "target_company": "菜鸟",
        "target_job": "in dolore reprehenderit do",
        "target_salary": "67",
        "target_area": "officia deserunt est proident eiusmod"
      },
      "pg_goal": {
        "id": 2868208361017344,
        "created_at": "2024-09-22T20:05:33+08:00",
        "updated_at": "2024-09-22T20:09:53+08:00",
        "uid": 2866331728744448,
        "target_university": "右转",
        "target_major": "elit sunt commodo velit",
        "target_score": 87
      }
    }
  }
}
```

```json
{
  "code": -1,
  "msg": "用户不存在",
  "data": null
}
```

```json
{
  "code": -2,
  "msg": "查询目标信息失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称                   | 类型      | 必选  | 约束 | 中文名 | 说明                   |
| ---------------------- | --------- | ----- | ---- | ------ | ---------------------- |
| » code                 | integer   | true  | none |        | none                   |
| » msg                  | string    | true  | none |        | none                   |
| » data                 | object    | true  | none |        | none                   |
| »» direction           | [integer] | true  | none |        | 用户方向1为考研2为就业 |
| »» goal                | object    | false | none |        | 目标信息               |
| »»» id                 | integer   | true  | none |        | none                   |
| »»» created_at         | string    | true  | none |        | none                   |
| »»» updated_at         | string    | true  | none |        | none                   |
| »»» uid                | integer   | true  | none |        | none                   |
| »»» target_university  | string    | true  | none |        | none                   |
| »»» target_major       | string    | true  | none |        | none                   |
| »»» target_score       | integer   | true  | none |        | none                   |
| »»» empl_goal          | object    | true  | none |        | none                   |
| »»»» id                | integer   | true  | none |        | none                   |
| »»»» created_at        | string    | true  | none |        | none                   |
| »»»» updated_at        | string    | true  | none |        | none                   |
| »»»» uid               | integer   | true  | none |        | none                   |
| »»»» status            | integer   | true  | none |        | none                   |
| »»»» target_company    | string    | true  | none |        | none                   |
| »»»» target_job        | string    | true  | none |        | none                   |
| »»»» target_salary     | string    | true  | none |        | none                   |
| »»»» target_area       | string    | true  | none |        | none                   |
| »»» pg_goal            | object    | true  | none |        | none                   |
| »»»» id                | integer   | true  | none |        | none                   |
| »»»» created_at        | string    | true  | none |        | none                   |
| »»»» updated_at        | string    | true  | none |        | none                   |
| »»»» uid               | integer   | true  | none |        | none                   |
| »»»» target_university | string    | true  | none |        | none                   |
| »»»» target_major      | string    | true  | none |        | none                   |
| »»»» target_score      | integer   | true  | none |        | none                   |

## GET 获取用户信息

GET /api/v1/user/info

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "user": {
      "id": 358091589685248,
      "created_at": "2024-09-15T21:51:15+08:00",
      "updated_at": "2024-09-18T10:06:44+08:00",
      "account": "2021214285",
      "name": "lj",
      "permission": 1,
      "direction": 0
    }
  }
}
```

```json
{
  "code": -1,
  "msg": "用户不存在",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称                  | 类型    | 必选 | 约束 | 中文名 | 说明                   |
| --------------------- | ------- | ---- | ---- | ------ | ---------------------- |
| » code                | integer | true | none |        | none                   |
| » msg                 | string  | true | none |        | none                   |
| » data                | object  | true | none |        | none                   |
| »» direction          | integer | true | none |        | 用户方向1为考研2为就业 |
| »» goal               | object  | true | none |        | 目标信息               |
| »»» id                | integer | true | none |        | none                   |
| »»» created_at        | string  | true | none |        | none                   |
| »»» updated_at        | string  | true | none |        | none                   |
| »»» uid               | integer | true | none |        | none                   |
| »»» target_university | string  | true | none |        | none                   |
| »»» target_major      | string  | true | none |        | none                   |
| »»» target_score      | integer | true | none |        | none                   |
| »» user               | object  | true | none |        | none                   |
| »»» id                | integer | true | none |        | none                   |
| »»» created_at        | string  | true | none |        | none                   |
| »»» updated_at        | string  | true | none |        | none                   |
| »»» account           | string  | true | none |        | none                   |
| »»» name              | string  | true | none |        | none                   |
| »»» permission        | integer | true | none |        | none                   |
| »»» direction         | integer | true | none |        | none                   |

## GET 查询用户信息

GET /api/v1/user/search

### 请求参数

| 名称      | 位置  | 类型   | 必选 | 中文名 | 说明                                    |
| --------- | ----- | ------ | ---- | ------ | --------------------------------------- |
| page      | query | string | 否   |        | 页码                                    |
| page_size | query | string | 否   |        | 每页大小                                |
| account   | query | string | 否   |        | 依账号/学号查询                         |
| name      | query | string | 否   |        | 依姓名查询                              |
| direction | query | string | 否   |        | 依目标方向查询0为未填写,3未两个方向都有 |
| keyword   | query | string | 否   |        | none                                    |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "page": 1,
    "page_size": 10,
    "rows": [
      {
        "goal": {
          "id": 636899496562688,
          "created_at": "2024-09-16T16:19:08+08:00",
          "updated_at": "2024-09-20T20:31:32+08:00",
          "uid": 358091535159296,
          "status": 1,
          "target_company": "华为",
          "target_job": "保安",
          "target_salary": "2k",
          "target_area": "重庆"
        },
        "user": {
          "id": 358091535159296,
          "created_at": "2024-09-15T21:51:15+08:00",
          "updated_at": "2024-09-20T20:31:32+08:00",
          "account": "2021214286",
          "name": "hml",
          "permission": 1,
          "direction": 2
        }
      }
    ],
    "total": 1
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称                | 类型     | 必选  | 约束 | 中文名 | 说明 |
| ------------------- | -------- | ----- | ---- | ------ | ---- |
| » code              | integer  | true  | none |        | none |
| » msg               | string   | true  | none |        | none |
| » data              | object   | true  | none |        | none |
| »» page             | integer  | true  | none |        | none |
| »» page_size        | integer  | true  | none |        | none |
| »» rows             | [object] | true  | none |        | none |
| »»» goal            | object   | false | none |        | none |
| »»»» id             | integer  | true  | none |        | none |
| »»»» created_at     | string   | true  | none |        | none |
| »»»» updated_at     | string   | true  | none |        | none |
| »»»» uid            | integer  | true  | none |        | none |
| »»»» status         | integer  | true  | none |        | none |
| »»»» target_company | string   | true  | none |        | none |
| »»»» target_job     | string   | true  | none |        | none |
| »»»» target_salary  | string   | true  | none |        | none |
| »»»» target_area    | string   | true  | none |        | none |
| »»» user            | object   | false | none |        | none |
| »»»» id             | integer  | true  | none |        | none |
| »»»» created_at     | string   | true  | none |        | none |
| »»»» updated_at     | string   | true  | none |        | none |
| »»»» account        | string   | true  | none |        | none |
| »»»» name           | string   | true  | none |        | none |
| »»»» permission     | integer  | true  | none |        | none |
| »»»» direction      | integer  | true  | none |        | none |
| »» total            | integer  | true  | none |        | none |

## DELETE 删除用户方向

DELETE /api/v1/user/dire/del

> Body 请求参数

```json
{
  "direction": 4
}
```

### 请求参数

| 名称        | 位置 | 类型    | 必选 | 中文名 | 说明 |
| ----------- | ---- | ------- | ---- | ------ | ---- |
| body        | body | object  | 否   |        | none |
| » direction | body | integer | 是   |        | none |

> 返回示例

```json
{
  "code": 1,
  "msg": "删除失败",
  "data": {
    "err": "不能删除用户没有的方向"
  }
}
```

```json
{
  "code": 2,
  "msg": "未知的方向",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

# 目标/就业目标

## POST 添加或更新就业目标

POST /api/v1/emgl/add

> Body 请求参数

```json
{
  "status": 1,
  "target_company": "菜鸟",
  "target_job": "in dolore reprehenderit do",
  "target_salary": "67",
  "target_area": "officia deserunt est proident eiusmod"
}
```

### 请求参数

| 名称             | 位置 | 类型    | 必选 | 中文名 | 说明                                |
| ---------------- | ---- | ------- | ---- | ------ | ----------------------------------- |
| body             | body | object  | 否   |        | none                                |
| » status         | body | integer | 是   |        | 就业状态1:未拿到offer;2:已拿到offer |
| » target_company | body | string  | 是   |        | 目标公司                            |
| » target_job     | body | string  | 是   |        | 目标职位                            |
| » target_salary  | body | string  | 是   |        | 理想薪资                            |
| » target_area    | body | string  | 是   |        | 目标地区                            |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "添加失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

# 目标/考研目标

## POST 添加或更新考研目标 

POST /api/v1/pggl/add

> Body 请求参数

```json
{
  "target_university": "右转",
  "target_major": "elit sunt commodo velit",
  "target_score": 87
}
```

### 请求参数

| 名称                | 位置 | 类型   | 必选 | 中文名 | 说明     |
| ------------------- | ---- | ------ | ---- | ------ | -------- |
| body                | body | object | 否   |        | none     |
| » target_university | body | string | 是   |        | 目标院校 |
| » target_major      | body | string | 是   |        | 目标专业 |
| » target_score      | body | number | 是   |        | 目标分数 |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "添加失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

# 打卡

## POST 打卡

POST /api/v1/clock_in/add

> Body 请求参数

```json
{
  "location": "ea"
}
```

### 请求参数

| 名称       | 位置 | 类型   | 必选 | 中文名 | 说明     |
| ---------- | ---- | ------ | ---- | ------ | -------- |
| body       | body | object | 否   |        | none     |
| » location | body | string | 是   |        | 打卡地点 |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "您在当前时间段已经打过卡了哦",
  "data": null
}
```

```json
{
  "code": -1,
  "msg": "无法获取用户id",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## GET 判断用户在当前时间段是否已打卡

GET /api/v1/clock_in/is_clock_in

> 返回示例

```json
{
  "code": 1,
  "msg": "您在当前时间段已经打过卡了哦",
  "data": null
}
```

```json
{
  "code": 0,
  "msg": "您还未打卡哦",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## GET 查询打卡数据

GET /api/v1/clock_in/search

### 请求参数

| 名称      | 位置  | 类型    | 必选 | 中文名 | 说明                               |
| --------- | ----- | ------- | ---- | ------ | ---------------------------------- |
| page      | query | string  | 否   |        | 页码                               |
| page_size | query | string  | 否   |        | 每页数据个数                       |
| uid       | query | integer | 否   |        | 用户id                             |
| date      | query | string  | 否   |        | 打卡日期格式为 yyyy-mm-dd          |
| time_slot | query | string  | 否   |        | 时间段取值为["上午","下午","晚上"] |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "page": 1,
    "page_size": 10,
    "rows": [
      {
        "id": 1268080080588800,
        "created_at": "2024-09-18T10:07:13+08:00",
        "updated_at": "2024-09-18T10:07:13+08:00",
        "uid": 358091589685248,
        "clock_in_time": 1726625232,
        "clock_in_location": "ea"
      },
      {
        "id": 1268230337335296,
        "created_at": "2024-09-18T10:07:49+08:00",
        "updated_at": "2024-09-18T10:07:49+08:00",
        "uid": 358091589685248,
        "clock_in_time": 1726625268,
        "clock_in_location": "ea"
      },
      {
        "id": 2089919690313728,
        "created_at": "2024-09-20T16:32:55+08:00",
        "updated_at": "2024-09-20T16:32:55+08:00",
        "uid": 358091589685248,
        "clock_in_time": 1726821174,
        "clock_in_location": "ea"
      },
      {
        "id": 2136496911945728,
        "created_at": "2024-09-20T19:37:59+08:00",
        "updated_at": "2024-09-20T19:37:59+08:00",
        "uid": 358091535159296,
        "clock_in_time": 1726832279,
        "clock_in_location": "数字图书馆"
      },
      {
        "id": 2136589706727424,
        "created_at": "2024-09-20T19:38:22+08:00",
        "updated_at": "2024-09-20T19:38:22+08:00",
        "uid": 358091535159296,
        "clock_in_time": 1726832301,
        "clock_in_location": "教室"
      },
      {
        "id": 2184314150195200,
        "created_at": "2024-09-20T22:48:00+08:00",
        "updated_at": "2024-09-20T22:48:00+08:00",
        "uid": 358091589685248,
        "clock_in_time": 1726843679,
        "clock_in_location": "老图书馆"
      }
    ],
    "total": 6
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称                  | 类型     | 必选  | 约束 | 中文名 | 说明 |
| --------------------- | -------- | ----- | ---- | ------ | ---- |
| » code                | integer  | true  | none |        | none |
| » msg                 | string   | true  | none |        | none |
| » data                | object   | true  | none |        | none |
| »» page               | integer  | true  | none |        | none |
| »» page_size          | integer  | true  | none |        | none |
| »» rows               | [object] | false | none |        | none |
| »»» id                | integer  | true  | none |        | none |
| »»» created_at        | string   | true  | none |        | none |
| »»» updated_at        | string   | true  | none |        | none |
| »»» uid               | integer  | true  | none |        | none |
| »»» clock_in_time     | integer  | true  | none |        | none |
| »»» clock_in_location | string   | true  | none |        | none |
| »» total              | integer  | true  | none |        | none |

# 考研倒计时

## GET 获取考研倒计时

GET /api/v1/pgee/cd

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "days": 91
  }
}
```

```json
{
  "code": 1,
  "msg": "解析考研日期失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称    | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ------- | ------- | ----- | ---- | ------ | ---- |
| » code  | integer | true  | none |        | none |
| » msg   | string  | true  | none |        | none |
| » data  | object  | false | none |        | none |
| »» days | integer | true  | none |        | none |

# 工作日志

## POST 添加工作日志

POST /api/v1/wl/add

> Body 请求参数

```json
{
  "company_name": "严情方种布党两",
  "job": "adipisicing incididunt sit enim",
  "salary": "enim ad",
  "location": "aliqua ea irure et"
}
```

### 请求参数

| 名称           | 位置 | 类型   | 必选 | 中文名 | 说明     |
| -------------- | ---- | ------ | ---- | ------ | -------- |
| body           | body | object | 否   |        | none     |
| » company_name | body | string | 是   |        | 公司名称 |
| » job          | body | string | 是   |        | 工作岗位 |
| » salary       | body | string | 是   |        | 薪资     |
| » location     | body | string | 是   |        | 工作地区 |

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

## PUT 更新工作状态

PUT /api/v1/wl/status/add

> Body 请求参数

```json
{
  "id": 2072736864145408,
  "create_time": 228304523029,
  "status": "拿到offer"
}
```

### 请求参数

| 名称          | 位置 | 类型    | 必选 | 中文名 | 说明         |
| ------------- | ---- | ------- | ---- | ------ | ------------ |
| body          | body | object  | 否   |        | none         |
| » id          | body | integer | 是   |        | 工作日志id   |
| » create_time | body | integer | 是   |        | 创建的时间戳 |
| » status      | body | string  | 是   |        | 状态         |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "添加失败",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ------ | ------- | ----- | ---- | ------ | ---- |
| » code | integer | true  | none |        | none |
| » msg  | string  | true  | none |        | none |
| » data | null    | false | none |        | none |

## GET 获取用户工作日志

GET /api/v1/wl/list

### 请求参数

| 名称 | 位置  | 类型    | 必选 | 中文名 | 说明   |
| ---- | ----- | ------- | ---- | ------ | ------ |
| uid  | query | integer | 否   |        | 用户id |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "work_logs": [
      {
        "id": 2072736864145408,
        "created_at": "2024-09-20T15:24:38+08:00",
        "updated_at": "2024-09-20T16:12:17+08:00",
        "uid": 358091589685248,
        "company_name": "严情方种布党两",
        "job": "adipisicing incididunt sit enim",
        "salary": "enim ad",
        "location": "aliqua ea irure et",
        "status_time_line": [
          {
            "created_at": 228304523029,
            "status": "失败"
          },
          {
            "created_at": 228304523029,
            "status": "成功"
          },
          {
            "created_at": 228304523029,
            "status": "拿到offer"
          }
        ]
      }
    ]
  }
}
```

```json
{
  "code": 1,
  "msg": "获取失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称                 | 类型     | 必选  | 约束 | 中文名 | 说明 |
| -------------------- | -------- | ----- | ---- | ------ | ---- |
| » code               | integer  | true  | none |        | none |
| » msg                | string   | true  | none |        | none |
| » data               | object   | false | none |        | none |
| »» work_logs         | [object] | true  | none |        | none |
| »»» id               | integer  | false | none |        | none |
| »»» created_at       | string   | false | none |        | none |
| »»» updated_at       | string   | false | none |        | none |
| »»» uid              | integer  | false | none |        | none |
| »»» company_name     | string   | false | none |        | none |
| »»» job              | string   | false | none |        | none |
| »»» salary           | string   | false | none |        | none |
| »»» location         | string   | false | none |        | none |
| »»» status_time_line | [object] | false | none |        | none |
| »»»» created_at      | integer  | true  | none |        | none |
| »»»» status          | string   | true  | none |        | none |

# 心得/帖子

## POST 添加帖子

POST /api/v1/post/add

> Body 请求参数

```json
{
  "title": "non officia elit ut dolor",
  "content": "ipsum proident ad"
}
```

### 请求参数

| 名称      | 位置 | 类型   | 必选 | 中文名 | 说明 |
| --------- | ---- | ------ | ---- | ------ | ---- |
| body      | body | object | 否   |        | none |
| » title   | body | string | 是   |        | none |
| » content | body | string | 是   |        | none |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "添加失败",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ------ | ------- | ----- | ---- | ------ | ---- |
| » code | integer | true  | none |        | none |
| » msg  | string  | true  | none |        | none |
| » data | null    | false | none |        | none |

## POST 添加评论

POST /api/v1/post/comment/add

> Body 请求参数

```json
{
  "id": 2416718815170560,
  "content": "Ut esse in aliquip"
}
```

### 请求参数

| 名称      | 位置 | 类型    | 必选 | 中文名 | 说明 |
| --------- | ---- | ------- | ---- | ------ | ---- |
| body      | body | object  | 否   |        | none |
| » id      | body | integer | 是   |        | none |
| » content | body | string  | 是   |        | none |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

```json
{
  "code": 1,
  "msg": "添加失败",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选  | 约束 | 中文名 | 说明 |
| ------ | ------- | ----- | ---- | ------ | ---- |
| » code | integer | true  | none |        | none |
| » msg  | string  | true  | none |        | none |
| » data | null    | false | none |        | none |

## GET 查询帖子

GET /api/v1/post/search

### 请求参数

| 名称      | 位置  | 类型    | 必选 | 中文名 | 说明             |
| --------- | ----- | ------- | ---- | ------ | ---------------- |
| page      | query | string  | 否   |        | 页码             |
| page_size | query | string  | 否   |        | 每页数据个数     |
| uid       | query | integer | 否   |        | 发表帖子的用户id |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": {
    "page": 1,
    "page_size": 10,
    "rows": [
      {
        "id": 2416718815170560,
        "created_at": "2024-09-21T14:11:30+08:00",
        "updated_at": "2024-09-21T14:31:06+08:00",
        "uid": 2006707475582976,
        "tittle": "non officia elit ut dolor",
        "content": "ipsum proident ad",
        "comments": [
          {
            "uid": 2006707475582976,
            "created_at": 1726900265,
            "content": "Ut esse in aliquip"
          }
        ]
      }
    ],
    "total": 1
  }
}
```

```json
{
  "code": 1,
  "msg": "查询失败",
  "data": {
    "err": ""
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称            | 类型     | 必选  | 约束 | 中文名 | 说明 |
| --------------- | -------- | ----- | ---- | ------ | ---- |
| » code          | integer  | true  | none |        | none |
| » msg           | string   | true  | none |        | none |
| » data          | object   | false | none |        | none |
| »» page         | integer  | true  | none |        | none |
| »» page_size    | integer  | true  | none |        | none |
| »» rows         | [object] | true  | none |        | none |
| »»» id          | integer  | false | none |        | none |
| »»» created_at  | string   | false | none |        | none |
| »»» updated_at  | string   | false | none |        | none |
| »»» uid         | integer  | false | none |        | none |
| »»» tittle      | string   | false | none |        | none |
| »»» content     | string   | false | none |        | none |
| »»» comments    | [object] | false | none |        | none |
| »»»» uid        | integer  | false | none |        | none |
| »»»» created_at | integer  | false | none |        | none |
| »»»» content    | string   | false | none |        | none |
| »» total        | integer  | true  | none |        | none |

## DELETE 删除帖子

DELETE /api/v1/post/del

> Body 请求参数

```json
{
  "id": 2416718815170560
}
```

### 请求参数

| 名称 | 位置 | 类型    | 必选 | 中文名 | 说明   |
| ---- | ---- | ------- | ---- | ------ | ------ |
| body | body | object  | 否   |        | none   |
| » id | body | integer | 是   |        | 帖子id |

> 返回示例

```json
{
  "code": 0,
  "msg": "success",
  "data": null
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称   | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ------ | ------- | ---- | ---- | ------ | ---- |
| » code | integer | true | none |        | none |
| » msg  | string  | true | none |        | none |
| » data | null    | true | none |        | none |

# 版本

## GET 获取后端版本号

GET /api/v1/ver

> 返回示例

> 200 Response

```json
{
  "code": 0,
  "msg": "string",
  "data": {
    "version": "string"
  }
}
```

### 返回结果

| 状态码 | 状态码含义                                              | 说明 | 数据模型 |
| ------ | ------------------------------------------------------- | ---- | -------- |
| 200    | [OK](https://tools.ietf.org/html/rfc7231#section-6.3.1) | none | Inline   |

### 返回数据结构

状态码 **200**

| 名称       | 类型    | 必选 | 约束 | 中文名 | 说明 |
| ---------- | ------- | ---- | ---- | ------ | ---- |
| » code     | integer | true | none |        | none |
| » msg      | string  | true | none |        | none |
| » data     | object  | true | none |        | none |
| »» version | string  | true | none |        | none |

# 数据模型

