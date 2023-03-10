package core

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"

	"github.com/bytedance/gopkg/util/gopool"
)

var pool = gopool.NewPool("gopool.GetFolderSizePool", 10000, gopool.NewConfig())

func calc(entry fs.DirEntry, folder string, total *int64) {
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
	// catch panic
	defer func() {
		if err := recover(); err != nil {
			e = fmt.Errorf("%v", err)
		}
	}()

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
		func(entry fs.DirEntry) {
			pool.Go(func() {
				defer wg.Done()
				calc(entry, folder, &total)
			})
		}(entrys[i])
	}

	wg.Wait()
	return total, nil
}
