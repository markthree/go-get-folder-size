// js-wasm
package main

import (
	"os"
	"syscall/js"

	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	value, err := getFolderSize.Parallel(os.Getenv("base"))

	if err != nil {
		js.Global().Set("$folderSizeError", err)
		return
	}

	js.Global().Set("$folderSize", value)
}
