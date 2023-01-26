package main

import (
	"fmt"
	"io/fs"
	"os"
	"path"
	"sync"
	"sync/atomic"

	"code.cloudfoundry.org/bytefmt"
)

// Synchronous execution, slow
func GetFolderSize(folder string) (int64, error) {
	var totalSize int64
	dirEntrys, err := os.ReadDir(folder)

	if err != nil {
		return 0, err
	}

	if len(dirEntrys) == 0 {
		return 0, nil
	}

	for _, dirEntry := range dirEntrys {
		if dirEntry.IsDir() {
			size, err := GetFolderSize(path.Join(folder, dirEntry.Name()))
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
func GetFolderSizeParallel(folder string) (totalSize int64, e error) {
	var wg sync.WaitGroup
	dirEntrys, err := os.ReadDir(folder)

	// try panic
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
		go func(_dirEntry fs.DirEntry) {
			if _dirEntry.IsDir() {
				size, err := GetFolderSizeParallel(path.Join(folder, _dirEntry.Name()))
				if err != nil {
					panic(err)
				}
				atomic.AddInt64(&totalSize, size)
			} else {
				info, err := _dirEntry.Info()
				if err != nil {
					panic(err)
				}
				atomic.AddInt64(&totalSize, info.Size())
			}
			wg.Done()
		}(dirEntry)
	}

	wg.Wait()
	return totalSize, nil
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	size, err := GetFolderSizeParallel(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	print(bytefmt.ByteSize(uint64(size)))
}
