package job

func ProvideJobController(trendJobCtrl ITrendJobCtrl) *Controller {
	return &Controller{
		TrendJobCtrl: trendJobCtrl,
	}
}

type Controller struct {
	TrendJobCtrl ITrendJobCtrl
}
