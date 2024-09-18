/**
 * @Author: chentong
 * @Date: 2024/09/18 15:54
 */

package svc

import (
	task "github.com/ch3nnn/xxl-job-zero/internal/task"
	"github.com/ch3nnn/xxl-job-zero/pkg/xxlx"
)

func RegisterTasks(exec xxlx.Executor) {
	exec.RegTasks(
		[]xxlx.Task{
			{
				Pattern: "task.test",
				TaskFn:  task.Test,
			},
			{
				Pattern: "task.test2",
				TaskFn:  task.Test2,
			},
			{
				Pattern: "task.panic",
				TaskFn:  task.Panic,
			},
		},
	)
}
