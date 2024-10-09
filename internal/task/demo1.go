/**
 * @Author: chentong
 * @Date: 2024/09/18 15:38
 */

package task

import (
	"context"
	"fmt"

	xxl "github.com/xxl-job/xxl-job-executor-go"
)

func Test1(cxt context.Context, param *xxl.RunReq) (msg string) {
	fmt.Println("test one task" + param.ExecutorHandler + " param：" + param.ExecutorParams + " log_id:" + xxl.Int64ToStr(param.LogID))
	return "test done"
}
