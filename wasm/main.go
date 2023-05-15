// js-wasm
package main

import (
	"os"
	"syscall/js"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	base := os.Getenv("base")
	isLoose := os.Getenv("loose") == "true"

	if isLoose {
		value := getFolderSize.LooseParallel(base)
		js.Global().Set("$folderSize", value)
	} else {
		value, err := getFolderSize.Parallel(base)

		if err != nil {
			js.Global().Set("$folderSizeError", err)
			return
		}

		js.Global().Set("$folderSize", value)
	}
}
