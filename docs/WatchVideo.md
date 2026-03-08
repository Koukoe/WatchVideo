---
title: 默认模块
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
generator: "@tarslib/widdershins v4.0.30"

---

# 默认模块

Base URLs:

* <a href="http://localhost:8888">开发环境: http://localhost:8888</a>

# Authentication

# 用户

## POST 注册

POST /user/register

> Body 请求参数

```yaml
username: Amy
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "code": 10000,
    "msg": "success"
  },
  "data": {
    "user_id": "10002"
  }
}
```

> 400 Response

```json
{
  "base": {
    "code": 10001,
    "msg": "user already exists"
  },
  "data": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|object|true|none||none|
|»» user_id|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|null|true|none||none|

## POST 登录

POST /user/login

> Body 请求参数

```yaml
username: Amy
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 是 |none|
|» username|body|string| 是 |none|
|» password|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "code": 10000,
    "msg": "success"
  },
  "data": {
    "user_id": "10002",
    "username": "Amy",
    "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAwMDIiLCJ0b2tlbl90eXBlIjoiYWNjZXNzIiwiZXhwIjoxNzcyOTg0MjAyLCJpYXQiOjE3NzI5NzcwMDJ9.xjZj2brB3K9WFIXqBe4mnxQmUG6eglEBWIoqNg9bvu0",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMTAwMDIiLCJ0b2tlbl90eXBlIjoicmVmcmVzaCIsImV4cCI6MTc3MzU4MTgwMiwiaWF0IjoxNzcyOTc3MDAyfQ.Ep1-5bJlo--_e_hg3GRm8BWzH4c-ZV0vEpA-qQFhoRc"
  }
}
```

> 400 Response

```json
{
  "base": {
    "code": 10001,
    "msg": "invalid username or password"
  },
  "data": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|object|true|none||none|
|»» user_id|string|true|none||none|
|»» username|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|null|true|none||none|

## GET 用户信息

GET /user/info

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|user_id|query|string| 是 |none|
|Access-Token|header|string| 是 |none|
|Refresh-Token|header|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "code": 10000,
    "msg": "success"
  },
  "data": {
    "user_id": "10002",
    "username": "Amy",
    "avatar_url": "",
    "created_at": "2026-03-08 21:31:19",
    "updated_at": "2026-03-08 21:31:19",
    "deleted_at": ""
  }
}
```

> 403 Response

```json
{
  "base": {
    "code": 10003,
    "msg": "forbidden"
  },
  "data": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|403|[Forbidden](https://tools.ietf.org/html/rfc7231#section-6.5.3)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|object|true|none||none|
|»» user_id|string|true|none||none|
|»» username|string|true|none||none|
|»» avatar_url|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» deleted_at|string|true|none||none|

状态码 **403**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|null|true|none||none|

## PUT 上传头像

PUT /user/avatar/upload

> Body 请求参数

```yaml
data: /home/koukoe/Pictures/ArchTan.png

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Access-Token|header|string| 是 |none|
|Refresh-Token|header|string| 是 |none|
|body|body|object| 是 |none|
|» data|body|string(binary)| 是 |none|

> 返回示例

> 200 Response

```json
{
  "base": {
    "code": 10000,
    "msg": "success"
  },
  "data": {
    "user_id": "10002",
    "username": "Amy",
    "avatar_url": "/storage/avatars/1772977254646115217.png",
    "created_at": "2026-03-08 21:31:19",
    "updated_at": "2026-03-08 21:40:54",
    "deleted_at": ""
  }
}
```

> 400 Response

```json
{
  "base": {
    "code": 10001,
    "msg": "missing file"
  },
  "data": null
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|object|true|none||none|
|»» user_id|string|true|none||none|
|»» username|string|true|none||none|
|»» avatar_url|string|true|none||none|
|»» created_at|string|true|none||none|
|»» updated_at|string|true|none||none|
|»» deleted_at|string|true|none||none|

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» base|object|true|none||none|
|»» code|integer|true|none||none|
|»» msg|string|true|none||none|
|» data|null|true|none||none|

# 数据模型

