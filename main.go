// go-cli
package main

import (
	"fmt"
	"os"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
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
