package core

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"
)

func calc(entry fs.DirEntry, wg *sync.WaitGroup, folder string, total *int64) {
	defer wg.Done()

	if entry.IsDir() {
		size, err := Parallel(path.Join(folder, entry.Name()))
		if err != nil {
			panic(err)
		}
		atomic.AddInt64(total, size)
		return
	}

	info, err := entry.Info()
	if err != nil {
		panic(err)
	}
	atomic.AddInt64(total, info.Size())
}

// Parallel execution, fast enough
func Parallel(folder string) (total int64, e error) {
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
		go calc(entrys[i], &wg, folder, &total)
	}

	wg.Wait()
	return total, nil
}
