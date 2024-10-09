package handler

import (
	"net/http"

	"github.com/ch3nnn/xxl-job-zero/internal/svc"
	"github.com/zeromicro/go-zero/rest"
)

// RegisterHandlers 注册路由
func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/run",
				Handler: serverCtx.XxlExec.RunTask,
			},
			{
				Method:  http.MethodPost,
				Path:    "/kill",
				Handler: serverCtx.XxlExec.KillTask,
			},
			{
				Method:  http.MethodPost,
				Path:    "/log",
				Handler: serverCtx.XxlExec.TaskLog,
			},
			{
				Method:  http.MethodPost,
				Path:    "/beat",
				Handler: serverCtx.XxlExec.Beat,
			},
			{
				Method:  http.MethodPost,
				Path:    "/idleBeat",
				Handler: serverCtx.XxlExec.IdleBeat,
			},
		},
	)
}
