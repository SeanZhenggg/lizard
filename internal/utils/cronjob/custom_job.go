package cronjob

func NewCustomJobFunc(c *cronJob, cmd FuncJob) *CustomJob {
	return &CustomJob{
		handlers: append(c.handlers, cmd),
	}
}

type CustomJob struct {
	ctx      *Context
	handlers handlerChain
}

func (job *CustomJob) Run() {
	job.ctx = &Context{}
	job.ctx.reset(job.handlers)
	job.ctx.Next()
}
