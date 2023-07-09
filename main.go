// go-cli
package main

import (
	"bufio"
	"fmt"
	"os"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func success(base string, size int64) {
	fmt.Fprintf(os.Stdout, "%v,%v\n", base, size)
}

func fail(base string, err error) {
	fmt.Fprintf(os.Stderr, "%v,%v", base, err)
}

func handle(base string) {
	size, err := getFolderSize.Invoke(base)

	if err != nil {
		fail(base, err)
		return
	}

	success(base, size)
}

func looseHandle(base string) {
	size := getFolderSize.LooseInvoke(base)
	success(base, size)
}

func main() {
	isIpc := os.Getenv("ipc") == "true"
	isLoose := os.Getenv("loose") == "true"

	if isIpc {
		reader := bufio.NewReader(os.Stdin)
		for {
			base, err := reader.ReadString(',')
			base = base[:len(base)-1]
			if err != nil {
				fail(base, err)
				continue
			}

			if isLoose {
				go looseHandle(base)
			} else {
				go handle(base)
			}
		}
	} else {
		root, err := os.Getwd()
		if err != nil {
			fmt.Fprint(os.Stderr, err)
			return
		}
		if isLoose {
			size := getFolderSize.LooseInvoke(root)
			fmt.Fprint(os.Stdout, size)
		} else {
			size, err := getFolderSize.Invoke(root)
			if err != nil {
				fmt.Fprint(os.Stderr, err)
				return
			}
			fmt.Fprint(os.Stdout, size)
		}
	}
}
