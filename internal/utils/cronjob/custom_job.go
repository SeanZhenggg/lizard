package cronjob

func NewCustomJobFunc(c *cronJob, cmd FuncJob) *customJob {
	return &customJob{
		handlers: append(c.handlers, cmd),
	}
}

type customJob struct {
	ctx      *Context
	handlers handlerChain
}

func (job *customJob) Run() {
	job.ctx = &Context{}
	job.ctx.reset(job.handlers)
	job.ctx.Next()
}
