# WatchVideo

基于 Hertz 框架的视频网站API

使用了 GORM, SQLite, jwt, Valkey

## 接口

### 用户模块
- [x] 注册
- [x] 登录
- [x] 用户信息
- [x] 上传头像

### 视频模块
- [x] 投稿
- [x] 发布列表
- [x] 搜索视频
- [x] 热门排行榜

### 互动模块
- [x] 点赞操作
- [x] 点赞列表
- [ ] 发表评论
- [ ] 删除评论
- [ ] 评论列表

### 社交模块
- [ ] 关注操作
- [ ] 关注列表
- [ ] 粉丝列表
- [ ] 好友列表

## 代办
- [ ] 完成接口
- [ ] 完成Docker部署

## 项目结构树
```
.
├── biz
│   ├── dao
│   │   ├── db
│   │   │   ├── db_init.go
│   │   │   ├── like_dao.go
│   │   │   ├── user_dao.go
│   │   │   └── video_dao.go
│   │   └── redis
│   ├── handler
│   │   ├── like_handler.go
│   │   ├── ping.go
│   │   ├── user_handler.go
│   │   └── video_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── model
│   │   ├── api
│   │   │   ├── like_api.go
│   │   │   ├── user_api.go
│   │   │   └── video_api.go
│   │   ├── domain
│   │   └── store
│   │       ├── like_model.go
│   │       ├── user_model.go
│   │       └── video_model.go
│   ├── router
│   │   └── register.go
│   └── service
│       ├── like_service.go
│       ├── user_service.go
│       └── video_service.go
├── build.sh
├── docs
│   └── WatchVideo.md
├── go.mod
├── go.sum
├── main.go
├── pkg
│   ├── cache
│   │   └── redis.go
│   ├── response
│   │   └── response.go
│   └── utils
│       └── jwt.go
├── README.md
├── router.go
├── router_gen.go
├── script
│   └── bootstrap.sh
├── storage
│   ├── avatars
│   └── videos
└── watchvideo.db
```