# xxl-job-zero
> 与 go-zero 框架集成

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
1	测试panic	BEAN：task.panic	* 0 * * * ?	admin	STOP	
2	测试耗时任务	BEAN：task.test2	* * * * * ?	admin	STOP	
3	测试golang	BEAN：task.test		* * * * * ?	admin	STOP
```
