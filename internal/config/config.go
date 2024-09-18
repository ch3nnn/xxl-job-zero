package config

import (
	"github.com/ch3nnn/xxl-job-zero/pkg/xxlx"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf

	XxlJobConf xxlx.XxlJobConf
}
