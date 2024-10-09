/**
 * @Author: chentong
 * @Date: 2024/09/18 15:54
 */

package svc

import (
	task "github.com/ch3nnn/xxl-job-zero/internal/task"
	"github.com/ch3nnn/xxl-job-zero/pkg/xxlx"
)

// RegisterTasks 注册任务
func RegisterTasks(exec xxlx.Executor) {
	exec.RegTasks(
		[]xxlx.Task{
			{
				Pattern: "task.demo1",
				TaskFn:  task.Test1,
			},
			{
				Pattern: "task.demo2",
				TaskFn:  task.Test2,
			},
			{
				Pattern: "task.demo3",
				TaskFn:  task.Panic,
			},
		},
	)
}
