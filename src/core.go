package core

import (
	"fmt"
	"os"
	"path"
	"sync"
	"sync/atomic"

	"github.com/panjf2000/ants/v2"
)

var pool, _ = ants.NewPool(1000000)

func calc(folder string) (total int64, e error) {
	entrys, err := os.ReadDir(folder)
	if err != nil {
		return 0, err
	}
	entrysLen := len(entrys)
	if entrysLen == 0 {
		return 0, nil
	}
	var wg sync.WaitGroup
	wg.Add(entrysLen)

	for i := 0; i < entrysLen; i++ {
		entry := entrys[i]
		ants.Submit(func() {
			defer wg.Done()
			if entry.IsDir() {
				size, err := calc(path.Join(folder, entry.Name()))
				if err != nil {
					panic(err)
				}
				atomic.AddInt64(&total, size)
				return
			}
			info, err := entry.Info()
			if err != nil {
				panic(err)
			}
			atomic.AddInt64(&total, info.Size())
		})
	}
	wg.Wait()
	return total, nil
}

// Parallel execution, fast enough
func Parallel(folder string) (total int64, e error) {
	// catch panic
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	total, err := calc(folder)
	if err != nil {
		return 0, err
	}
	return total, nil
}
