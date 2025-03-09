# xxl-job-zero
> 与 go-zero 框架集成

## 目录结构
```shell
./
├── LICENSE
├── README.md
├── deploy  // 部署相关
│   ├── docker-compose.yml
│   └── tables_xxl_job.sql
├── etc // 配置文件
│   └── job.yaml
├── go.mod
├── go.sum
├── internal 
│   ├── config  
│   │   └── config.go
│   ├── handler 
│   │   └── routes.go  // xxl-job-admin http交互 
│   ├── svc
│   │   ├── regtask.go  // 注册任务
│   │   └── servicecontext.go 
│   └── task  // 定时任务
│       ├── demo1.go
│       ├── demo2.go
│       └── demo3.go
├── job.go  // 主函数
├── makefile
└── pkg
    └── xxlx
        ├── log.go
        └── xxl.go


```

## xxl-job-admin 配置

### 1. 添加执行器

执行器管理->新增执行器,执行器列表如下：

```
AppName		名称		注册方式	OnLine 		机器地址 		操作
golang-jobs	golang执行器	自动注册 		查看 ( 1 ）   
```

查看->注册节点

```
http://127.0.0.1:8888
```

### 2. 添加任务

任务管理->新增(注意，使用BEAN模式，`JobHandler`与`RegTask`名称一致)

```
1	测试panic	BEAN：task.demo3	* 0 * * * ?	admin	STOP	
2	测试耗时任务	BEAN：task.demo2	* * * * * ?	admin	STOP	
3	测试golang	BEAN：task.demo1	* * * * * ?	admin	STOP
```
