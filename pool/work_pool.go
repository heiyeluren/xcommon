package pool

type Pool struct {
	work chan func()   // 任务
	sem  chan struct{} // 数量
}

func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		sem:  make(chan struct{}, size),
	}
}

func (p *Pool) worker(f func()) {
	defer func() { <-p.sem }()
	for {
		f()
		f = <-p.work
	}
}

func (p *Pool) NewTask(f func()) {
	select {
	case p.work <- f:
	case p.sem <- struct{}{}:
		go p.worker(f)
	}
}

