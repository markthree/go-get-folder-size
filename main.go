// go-cli
package main

import (
	"fmt"
	"os"

	"code.cloudfoundry.org/bytefmt"
	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	root, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "-s" {
		size, err := getFolderSize.Sync(root)
		if err != nil {
			fmt.Println(err)
			return
		}
		print(bytefmt.ByteSize(uint64(size)))
	} else {
		size, err := getFolderSize.Parallel(root)
		if err != nil {
			fmt.Println(err)
			return
		}
		print(bytefmt.ByteSize(uint64(size)))
	}
}
