package job

import (
	"lizard/internal/controller/job/middleware"
	"lizard/internal/utils/cronjob"
)

func ProvideJobController(trendJobCtrl ITrendJobCtrl) *Controller {
	return &Controller{
		TrendJobCtrl: trendJobCtrl,
	}
}

type Controller struct {
	TrendJobCtrl ITrendJobCtrl
}

func SetJobError(ctx *cronjob.Context, data error) {
	middleware.SetJobError(ctx, data)
}

func SetJobFunc(ctx *cronjob.Context) {
	middleware.SetJobFunc(ctx)
}
