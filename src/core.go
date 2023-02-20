package core

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"
)

// Parallel execution, fast enough
func Parallel(folder string) (totalSize int64, e error) {
	var wg sync.WaitGroup
	entrys, err := os.ReadDir(folder)

	// catch panic
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	if err != nil {
		return 0, err
	}

	entrysLen := len(entrys)

	if entrysLen == 0 {
		return 0, nil
	}

	wg.Add(entrysLen)

	for i := 0; i < entrysLen; i++ {
		go func(entry fs.DirEntry) {
			defer wg.Done()

			if entry.IsDir() {
				size, err := Parallel(path.Join(folder, entry.Name()))
				if err != nil {
					panic(err)
				}
				atomic.AddInt64(&totalSize, size)
				return
			}

			info, err := entry.Info()
			if err != nil {
				panic(err)
			}
			atomic.AddInt64(&totalSize, info.Size())
		}(entrys[i])
	}

	wg.Wait()
	return totalSize, nil
}
