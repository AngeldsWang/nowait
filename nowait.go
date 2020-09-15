package nowait

import "sync"

type NowaitGroup struct {
	wg *sync.WaitGroup
}

func NewNowaitGroup() *NowaitGroup {
	return &NowaitGroup{&sync.WaitGroup{}}
}

func (no *NowaitGroup) Add(delta int) {
	no.wg.Add(delta)
}

func (no *NowaitGroup) Done() {
	no.wg.Done()
}

func (no *NowaitGroup) Wait() {
	// do wait only in tests
	if runInTest() {
		no.wg.Wait()
	}
}
