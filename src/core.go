package core

import (
	"io/fs"
	"os"
	"path"
)

type EntrySizeChan chan int64

func Invoke(folder string) (size int64, e error) {
	gracefulExit := func(err error) {
		e = err
		size = 0
	}

	entrys, err := os.ReadDir(folder)
	if err != nil {
		gracefulExit(err)
		return
	}
	entrysLen := len(entrys)
	if entrysLen == 0 {
		gracefulExit(nil)
		return
	}
	errChan := make(chan error)
	sizeChan := make(EntrySizeChan, entrysLen)

	for i := 0; i < entrysLen; i++ {
		go func(entry fs.DirEntry) {
			if entry.IsDir() {
				subFolderSize, err := Invoke(path.Join(folder, entry.Name()))
				if err != nil {
					errChan <- err
					return
				}
				sizeChan <- subFolderSize
				return
			}

			info, err := entry.Info()
			if err != nil {
				errChan <- err
				return
			}
			sizeChan <- info.Size()
		}(entrys[i])
	}

	for i := 0; i < entrysLen; i++ {
		select {
		case value := <-sizeChan:
			size += value
		case newErr := <-errChan:
			if newErr != nil {
				gracefulExit(newErr)
				return
			}
		}
	}
	return size, nil
}

func LooseInvoke(folder string) int64 {
	size := int64(0)

	entrys, err := os.ReadDir(folder)
	if err != nil {
		return 0
	}
	entrysLen := len(entrys)
	if entrysLen == 0 {
		return 0
	}
	sizeChan := make(EntrySizeChan, entrysLen)

	for i := 0; i < entrysLen; i++ {
		go func(entry fs.DirEntry) {
			if entry.IsDir() {
				subFolderSize := LooseInvoke(path.Join(folder, entry.Name()))
				sizeChan <- subFolderSize
				return
			}
			info, err := entry.Info()
			if err != nil {
				sizeChan <- 0
				return
			}
			sizeChan <- info.Size()
		}(entrys[i])
	}
	for i := 0; i < entrysLen; i++ {
		size += <-sizeChan
	}
	return size
}
