# WatchVideo

基于 Hertz 框架的视频网站后端服务（GORM + SQLite）

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
- [ ] 热门排行榜

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
│   │   │   ├── user_dao.go
│   │   │   └── video_dao.go
│   │   └── redis
│   ├── handler
│   │   ├── ping.go
│   │   ├── user_handler.go
│   │   └── video_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── model
│   │   ├── api
│   │   │   ├── user_api.go
│   │   │   └── video_api.go
│   │   ├── domain
│   │   └── store
│   │       ├── user_model.go
│   │       └── video_model.go
│   ├── router
│   │   └── register.go
│   └── service
│       ├── user_service.go
│       └── video_service.go
├── build.sh
├── docs
│   └── WatchVideo.md
├── go.mod
├── go.sum
├── main.go
├── pkg
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