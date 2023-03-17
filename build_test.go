package snowflake

import (
	"fmt"
	"sync"
	"testing"
)

func init() {
	Init(1)
}

func TestCreateId(t *testing.T) {
	wait := &sync.WaitGroup{}
	for i := 0; i < 1000; i++ {
		wait.Add(1)
		go func() {
			defer wait.Done()
			fmt.Printf("CreateId(): %v\n", CreateId())
		}()
	}
	wait.Wait()
}
