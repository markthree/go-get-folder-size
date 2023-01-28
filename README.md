# go-get-folder-size

Get the size of a folder by recursively iterating through all its sub(files && folders). Use go, so high-speed

<br />

## motivation

To quickly know the folder size，but [get-folder-size](https://github.com/alessioalex/get-folder-size) is implemented by nodejs, which is too slow。

<br />

## Usage

### npm

#### install

```shell
npm install go-get-folder-size
```

#### cli

```shell
# Binary go, fastest
npx go-get-folder-size
```

#### program

```ts
import {
	getFolderSize,
	getFolderSizeBin,
	getFolderSizeWasm
} from 'go-get-folder-size'

const base = './' // The directory path you want to get

await getFolderSizeBin(base) // Binary go, fastest

await getFolderSize(base) // native node

await getFolderSizeWasm(base) // Wasm go，slowest
```

### go

#### cli

```shell
go install github.com/markthree/go-get-folder-size
```

```shell
go-get-folder-size
```

#### program

```shell
# Super invincible fast
go get github.com/markthree/go-get-folder-size
```

```go
package main

import (
	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	size, err := getFolderSize.Parallel("./") // Concurrent running, invincible fast
}
```

<br />

## License

Made with [markthree](https://github.com/markthree)

Published under [MIT License](./LICENSE).
