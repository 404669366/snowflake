package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

func init() {
	NewSnowflake(1)
}

type IDS struct {
	rw   *sync.RWMutex
	data map[int64]struct{}
}

func NewIDS() *IDS {
	return &IDS{
		rw:   &sync.RWMutex{},
		data: make(map[int64]struct{}),
	}
}

func (i *IDS) Add(id int64) {
	i.rw.Lock()
	defer i.rw.Unlock()
	i.data[id] = struct{}{}
}

func (i *IDS) Len() int {
	i.rw.RLock()
	defer i.rw.RUnlock()
	return len(i.data)
}

func TestCreateId(t *testing.T) {
	wait := &sync.WaitGroup{}
	ids := NewIDS()
	for i := 0; i < 1000; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			ids.Add(Generate())
		}()
	}
	wait.Wait()
	fmt.Printf("ids.Len(): %v\n", ids.Len())
}
