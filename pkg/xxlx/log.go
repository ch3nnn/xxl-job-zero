/**
 * @Author: chentong
 * @Date: 2024/09/18 17:44
 */

package xxlx

import (
	xxl "github.com/xxl-job/xxl-job-executor-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type logger struct {
	xxl.Logger
}

func NewLogger() xxl.Logger {
	return &logger{}
}

func (l *logger) Info(format string, a ...interface{}) {
	logx.Infof(format, a...)
}

func (l *logger) Error(format string, a ...interface{}) {
	logx.Errorf(format, a...)
}
