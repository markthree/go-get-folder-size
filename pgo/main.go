package main

import (
	"fmt"
	"os"
	"runtime/pprof"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	f, _ := os.OpenFile("../default.pgo", os.O_CREATE|os.O_RDWR, 0644)
	defer f.Close()
	pprof.StartCPUProfile(f)
	defer pprof.StopCPUProfile()

	root, err := os.Getwd()
	if err != nil {
		fmt.Print(err)
		return
	}

	size, err := getFolderSize.Parallel(root)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(size)
}
