package engine

import (
	"log"
)

type ConcurrentEngine struct {
	Scheduler  Scheduler
	WokerCount int
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	//in := make(chan Request)
	out := make(chan ParserResult)

	//e.Scheduler.ConfigureMasterWorkerChan(in)

	//run for tell e.Scheduler's requestChan chan workChan ready go
	e.Scheduler.Run()

	//创建Worker
	for i := 0; i < e.WokerCount; i++ {
		//createWorker(in, out)
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//任务分发给Worker
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	count := 0
	for {
		//打印out的items
		result := <-out
		for _, item := range result.Items {
			log.Printf("Get Count:%d; Items: %v\n", count, item)
			count++
		}

		//将out里的Request送给Scheduler
		for _, r := range result.Requests {
			e.Scheduler.Submit(r)
		}

	}
}

//workerConut goroutine to exec worker for Loop
/*func createWorker(in chan Request, out chan ParserResult) {
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
}*/

//workerConut goroutine to exec worker for Loop
func createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			// 这里会卡住,直到QueuedScheduler.Run()里第三个case被执行,即把requestQ[0]放进workerQ[0]
			// 后每个这里定义的in(chan of Request)就有值不阻塞了
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
