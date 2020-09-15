package nowait

import "sync"

type NowaitGroup struct {
	*sync.WaitGroup
}

func NewNowaitGroup() *NowaitGroup {
	return &NowaitGroup{&sync.WaitGroup{}}
}

func (no *NowaitGroup) Add(delta int) {
	no.Add(delta)
}

func (no *NowaitGroup) Done() {
	no.Done()
}

func (no *NowaitGroup) Wait() {
	// do wait only in tests
	if runInTest() {
		no.Wait()
	}
}
