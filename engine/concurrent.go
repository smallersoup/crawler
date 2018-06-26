package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WokerCount int
}

type Scheduler interface {
	Submit(Request)
	ConfigureMasterWorkerChan(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	in := make(chan Request)
	out := make(chan ParserResult)

	e.Scheduler.ConfigureMasterWorkerChan(in)

	//创建Worker
	for i := 0; i < e.WokerCount; i++ {
		createWorker(in, out)
	}


	//任务分发给Worker
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}


	for  {

		//打印out的items
		result := <- out
		for _, item := range result.Items {
			log.Printf("Get Items: %v\n", item)
		}

		//将out里的Request送给Scheduler
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}

	}
}

//workerConut goroutine to exec worker for Loop
func createWorker(in chan Request, out chan ParserResult) {
	go func() {
		for {
			request := <-in

			parserResult, err := worker(request)

			//发生了错误继续下一个
			if err != nil {
				continue
			}

			//将parserResult送出
			out <- parserResult
		}
	}()
}