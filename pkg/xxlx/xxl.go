/**
 * @Author: chentong
 * @Date: 2024/09/18 16:09
 */

package xxlx

import (
	"fmt"

	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type (
	XxlJobConf struct {
		ServerAddr  string // xxl-job-admin 调度中心地址
		AccessToken string // xxl-job-admin 请求令牌
		RegistryKey string // 执行器名称
		Port        string // 本地(执行器)端口
	}

	Task struct {
		Pattern string       // 任务执行函数名称
		TaskFn  xxl.TaskFunc // 任务执行函数
	}

	Executor struct {
		xxl.Executor
	}
)

func NewXxlExecutor(c XxlJobConf) Executor {
	exec := xxl.NewExecutor(
		xxl.ServerAddr(c.ServerAddr),
		xxl.ExecutorPort(c.Port),
		xxl.AccessToken(c.AccessToken),
		xxl.RegistryKey(c.RegistryKey),
		xxl.SetLogger(NewLogger()),
	)
	exec.Init()

	return Executor{Executor: exec}
}

func (e *Executor) RegTasks(tasks []Task) {
	for _, task := range tasks {
		e.RegTask(task.Pattern, task.TaskFn)
		fmt.Printf("register task: %s \n", task.Pattern)
	}
}
