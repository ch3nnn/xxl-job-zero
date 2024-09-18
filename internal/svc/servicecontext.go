package svc

import (
	"github.com/ch3nnn/xxl-job-zero/internal/config"
	"github.com/ch3nnn/xxl-job-zero/pkg/xxlx"
	xxl "github.com/xxl-job/xxl-job-executor-go"
)

type ServiceContext struct {
	Config  config.Config
	XxlExec xxl.Executor
}

func NewServiceContext(c config.Config) *ServiceContext {
	exec := xxlx.NewXxlExecutor(c.XxlJobConf)
	RegisterTasks(exec)

	return &ServiceContext{
		Config:  c,
		XxlExec: exec,
	}
}
