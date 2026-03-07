```
.
├── biz
│   ├── dao
│   │   ├── db
│   │   │   ├── db_init.go
│   │   │   └── user_dao.go
│   │   └── redis
│   ├── handler
│   │   ├── ping.go
│   │   └── user_handler.go
│   ├── middleware
│   │   └── auth.go
│   ├── model
│   │   ├── api
│   │   │   └── user_api.go
│   │   ├── domain
│   │   └── store
│   │       └── user_model.go
│   ├── router
│   │   └── register.go
│   └── service
│       └── user_service.go
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
│   └── avatars
└── watchvideo.db
```