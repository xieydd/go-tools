package random

import (
	"time"
	"os"
	"sync"
	"strconv"
)

var (
	rand uint32
	randmu  sync.Mutex
)

func reseed() uint32 {
	return uint32(time.Now().UnixNano() + int64(os.Getpid()))
}

func Randomint32() string {
	randmu.Lock()
	defer randmu.Unlock()
	r := rand
	if (r == 0) {
		r = reseed()
	}
	r = 2637485*r + 1013904223
	rand = r
	return strconv.Itoa(int(1e9 + r%1e9))[1:]
}


