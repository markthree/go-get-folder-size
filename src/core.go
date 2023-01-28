package core

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"
)

// Synchronous execution, slow
func Sync(folder string) (totalSize int64, e error) {
	dirEntrys, err := os.ReadDir(folder)

	if err != nil {
		return 0, err
	}

	if len(dirEntrys) == 0 {
		return 0, nil
	}

	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			size, err := Sync(path.Join(folder, dirEntry.Name()))
			if err != nil {
				return 0, err
			}
			totalSize += size
			continue
		}

		info, err := dirEntry.Info()
		if err != nil {
			return 0, err
		}
		totalSize += info.Size()
	}

	return totalSize, nil
}

// Parallel execution, fast enough
func Parallel(folder string) (totalSize int64, e error) {
	var wg sync.WaitGroup
	dirEntrys, err := os.ReadDir(folder)

	// catch panic
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

	if err != nil {
		return 0, err
	}

	dirEntrysLen := len(dirEntrys)

	if dirEntrysLen == 0 {
		return 0, nil
	}

	wg.Add(dirEntrysLen)

	for _, dirEntry := range dirEntrys {
		go func(dirEntry fs.DirEntry) {
			defer wg.Done()

			if dirEntry.IsDir() {
				size, err := Parallel(path.Join(folder, dirEntry.Name()))
				if err != nil {
					panic(err)
				}
				atomic.AddInt64(&totalSize, size)
				return
			}

			info, err := dirEntry.Info()
			if err != nil {
				panic(err)
			}
			atomic.AddInt64(&totalSize, info.Size())
		}(dirEntry)
	}

	wg.Wait()
	return totalSize, nil
}
