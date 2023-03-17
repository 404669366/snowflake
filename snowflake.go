package snowflake

import (
	"sync"
	"time"
)

const (
	epoch       = 1577836800000 // 2020-01-01 00:00:00 UTC in milliseconds
	workerBits  = 10
	numberBits  = 12
	maxWorker   = -1 ^ (-1 << workerBits)
	maxNumber   = -1 ^ (-1 << numberBits)
	timeShift   = workerBits + numberBits
	workerShift = numberBits
)

type snowflake struct {
	lastTimestamp int64
	workerId      int
	number        int
	mutex         sync.Mutex
}

var snow *snowflake

func NewSnowflake(workerId int) {
	if workerId < 0 || workerId > maxWorker {
		panic("snowflake error : worker id out of range")
	}
	snow = &snowflake{workerId: workerId}
}

func Generate() int64 {
	snow.mutex.Lock()
	defer snow.mutex.Unlock()
	now := time.Now().UnixNano() / 1000000 // convert to milliseconds
	if snow.lastTimestamp == now {
		snow.number++
		if snow.number > maxNumber {
			for now <= snow.lastTimestamp {
				now = time.Now().UnixNano() / 1000000
			}
		}
	} else {
		snow.number = 0
	}
	if now < snow.lastTimestamp {
		panic("snowflake error : invalid system clock")
	}
	snow.lastTimestamp = now
	return (now-epoch)<<timeShift | int64(snow.workerId)<<workerShift | int64(snow.number)
}
