package scheduler

import (
	"crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (q *QueuedScheduler) Submit(r engine.Request) {
	q.requestChan <- r
}

func (q *QueuedScheduler) WorkReady(wChan chan engine.Request) {
	//这里不会卡住
	q.workerChan <- wChan
}

func (q *QueuedScheduler) ConfigureMasterWorkerChan(rChan chan engine.Request) {
	panic("implement me")
}

func (q *QueuedScheduler) Run() {
	//chan 需要make
	q.workerChan = make(chan chan engine.Request)
	q.requestChan = make(chan engine.Request)

	go func() {

		//所有worker在队列workerQ里取request进行fetch
		var requestQ []engine.Request
		var workerQ []chan engine.Request

		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request

			//if requestQ and workerQ is not empty, send first r of requestQ to first work chan of workerQ
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeRequest = requestQ[0]
				activeWorker = workerQ[0]
			}

			select {
			case r := <-q.requestChan:
				//firstly, r send to request_queue, then both request and worker is not empty, will send r to w
				requestQ = append(requestQ, r)
			case w := <-q.workerChan:
				//firstly, w send to worker_queue, then both request and worker is not empty, will send r to w
				workerQ = append(workerQ, w)
			case activeWorker <- activeRequest:
				requestQ = requestQ[1:]
				workerQ = workerQ[1:]
			}
		}
	}()
}
