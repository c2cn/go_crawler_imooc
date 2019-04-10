package scheduler

import "distributedCrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChannel(c chan engine.Request) {
	s.workerChan = c
}

func (s *SimpleScheduler) Submit( rq engine.Request){
	//send resquest to worker channel
	//不用go 会导致循环等待
	go func() {
		s.workerChan <- rq
	}()

}

