// go-cli
package main

import (
	"bufio"
	"fmt"
	"os"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func success(base string, size int64) {
	fmt.Fprintf(os.Stdout, "%v,%v", base, size)
}

func fail(base string, err error) {
	fmt.Fprintf(os.Stderr, "%v,%v", base, err)
}

func handle(base string) {
	size, err := getFolderSize.Parallel(base)

	if err != nil {
		fail(base, err)
		return
	}

	success(base, size)
}

func main() {
	isIpc := os.Getenv("ipc")

	if isIpc == "true" {
		reader := bufio.NewReader(os.Stdin)
		for {
			base, err := reader.ReadString(',')
			base = base[:len(base)-1]
			if err != nil {
				fail(base, err)
				continue
			}

			go handle(base)
		}
	} else {
		root, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		size, err := getFolderSize.Parallel(root)
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}

		fmt.Fprint(os.Stdout, size)
	}
}
