package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
	getFolderSize "github.com/markthree/go-get-folder-size/core"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}
	size, err := getFolderSize.Parallel(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	print(bytefmt.ByteSize(uint64(size)))
}
