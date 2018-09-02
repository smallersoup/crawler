package engine

var visitedUrls = make(map[string]bool)

type ConcurrentEngine struct {
	Scheduler  Scheduler
	WokerCount int
	ItemChan   chan Item
	RequestProcessor Processor
}

type Processor func(Request) (ParserResult, error)

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
		//e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
		e.createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	//任务分发给Worker
	for _, r := range seeds {

		if isDuplicate(r.Url) {
			continue
		}

		e.Scheduler.Submit(r)
	}

	for {
		//打印out的items
		result := <-out
		for _, item := range result.Items {
			go func() { e.ItemChan <- item }()
		}

		//将out里的Request送给Scheduler
		for _, r := range result.Requests {

			if isDuplicate(r.Url) {
				continue
			}

			e.Scheduler.Submit(r)
		}
	}
}

func isDuplicate(url string) bool {

	//if url exist in map, return true
	if visitedUrls[url] {
		return true
	}

	//else not exist, set true, return false
	visitedUrls[url] = true

	return false
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
func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParserResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkReady(in)
			// 这里会卡住,直到QueuedScheduler.Run()里第三个case被执行,即把requestQ[0]放进workerQ[0]
			// 后每个这里定义的in(chan of Request)就有值不阻塞了
			request := <-in
			//parserResult, err := Worker(request)
			//Call WorkerRpc worker
			parserResult, err := e.RequestProcessor(request)

			//发生了错误继续下一个
			if err != nil {
				continue
			}
			//将parserResult送出
			out <- parserResult
		}
	}()
}
