package scheduler

import "crawler/engine"

//SimpleScheduler one workChan to multi worker
//所有worker公用一个chan
type SimpleScheduler struct {
	workChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	//要workerChan就return,即所有worker公用一个chan
	return s.workChan
}

func (s *SimpleScheduler) WorkReady(chan engine.Request) {
	//implement me
}

func (s *SimpleScheduler) Run() {
	s.workChan = make(chan engine.Request, 1024)
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.workChan <- r }()
}
