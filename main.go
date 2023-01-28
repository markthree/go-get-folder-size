// go-cli
package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func sliceIncludes(s []string, o string) bool {
	for _, v := range s {
		if v == o {
			return true
		}
	}
	return false
}

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	var _size int64

	if sliceIncludes(os.Args, "-s") {
		size, err := getFolderSize.Sync(root)
		if err != nil {
			fmt.Println(err)
			return
		}
		_size = size
	} else {
		size, err := getFolderSize.Parallel(root)
		if err != nil {
			fmt.Println(err)
			return
		}
		_size = size
	}

	if sliceIncludes(os.Args, "-p") {
		print(bytefmt.ByteSize(uint64(_size)))
	} else {
		print(uint64(_size))
	}
}
