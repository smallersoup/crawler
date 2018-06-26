package scheduler

import "crawler/engine"

//SimpleScheduler one workChan to multi worker
type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan(r chan engine.Request) {
	s.workChan = r
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}
