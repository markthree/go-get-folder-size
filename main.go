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

	size, err := getFolderSize.Parallel(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	if sliceIncludes(os.Args, "-p") {
		fmt.Print(bytefmt.ByteSize(uint64(size)))
	} else {
		fmt.Print(uint64(size))
	}
}
