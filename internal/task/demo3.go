/**
 * @Author: chentong
 * @Date: 2024/09/18 15:38
 */

package task

import (
	"context"

	xxl "github.com/xxl-job/xxl-job-executor-go"
)

func Panic(cxt context.Context, param *xxl.RunReq) (msg string) {
	panic("test panic")
	return
}
