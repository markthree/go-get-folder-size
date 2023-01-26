package goGetFolderSize

import (
	"os"
	"path"
	"sync"
)

// Synchronous execution, slow
func GetFolderSize(folder string) (int64, error) {
	var totalSize int64
	dirEntrys, err := os.ReadDir(folder)

	if len(dirEntrys) == 0 {
		return 0, nil
	}

	if err != nil {
		return 0, err
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
func GetFolderSizeParallel(folder string) (int64, error) {
	var totalSize int64
	var wg sync.WaitGroup

	// TODO Parallel
	wg.Wait()

	return totalSize, nil
}
