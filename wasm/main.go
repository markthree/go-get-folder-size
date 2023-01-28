// js-wasm
package main

import (
	"os"
	"syscall/js"

	"code.cloudfoundry.org/bytefmt"
	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	value, err := getFolderSize.Parallel(os.Getenv("base"))

	if err != nil {
		js.Global().Set("$folderSizeError", err)
		return
	}

	if os.Getenv("pretty") == "true" {
		js.Global().Set("$folderSize", bytefmt.ByteSize(uint64(value)))
		return
	}

	js.Global().Set("$folderSize", value)
}
