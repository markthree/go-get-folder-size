<div align="center">
    <img width="100%" height="100%" src="https://raw.githubusercontent.com/markthree/go-get-folder-size/main/snapshot.gif" />
    <h1><a href="https://github.com/markthree/go-get-folder-size">go-get-folder-size</a></h1>
    <p>递归获取文件夹大小，使用 go，足够快，可以跑在 node 中</p>
</div>

<br />

## 特性

- 🐉 [ipc go](./src/bin.ts)
- 🦕 [二进制 go](./src/bin.ts)
- 🦖 [原生 node](./src/node.ts)
- 🐊 [wasm go](./src/wasm.ts)

<br />

## 动机

想要快速知道文件夹大小，但 nodejs 实现的
[get-folder-size](https://github.com/alessioalex/get-folder-size) 是慢的，所以用
go 实现了递归获取文件夹大小，能跑在 nodejs 中。

具体可见 issue 👉
[get-folder-size/issues/22](https://github.com/alessioalex/get-folder-size/issues/22)

<br />

## 使用

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
  getFolderSizeWasm,
} from "go-get-folder-size";

const base = "./"; // 你想要获取的目录

await getFolderSizeBin(base); // 二进制 go，最快

await getFolderSize(base); // 原生 node

await getFolderSizeWasm(base); // Wasm go，最慢 🥵
```

### go

#### install

```shell
go install github.com/markthree/go-get-folder-size
```

#### cli

```shell
go-get-folder-size
```

#### program

```shell
go get github.com/markthree/go-get-folder-size
```

```go
package main

import (
	getFolderSize "github.com/markthree/go-get-folder-size/src"
)

func main() {
	size, err := getFolderSize.Parallel("./") // 并发计算，超级快

  size2 := getFolderSize.LooseParallel("./") // 有时我们可能会遇到不可访问的文件，我们可以使用 loose 来忽略它们
}
```

<br />

##### IPC

适用于多路径

```ts
import { createGetFolderSizeBinIpc } from "go-get-folder-size";

const { getFolderSizeWithIpc, close } = createGetFolderSizeBinIpc();

Promise.all([
  getFolderSizeWithIpc("./"),
  getFolderSizeWithIpc("../"),
])
  .then((vs) => console.log(vs))
  .finally(close); // 手动退出是必需的
```


## loose

有时我们可能会遇到不可访问的文件，我们可以使用 `loose` 来忽略它们

### cli

```shell
go-get-folder-size --loose
```

### program

```ts
import {
  getFolderSize,
  getFolderSizeBin,
  getFolderSizeWasm,
} from "go-get-folder-size";

const base = "./"; // 你想要获取的目录
const pretty = false; // 人类可读的方式
const loose = true;

await getFolderSizeBin(base, pretty, { loose }); // Binary go, fastest

await getFolderSize(base, pretty, { loose }); // native node

await getFolderSizeWasm(base, pretty, { loose }); // Wasm go，slowest
```

<br />

## 提示

- 目前该包被使用在组织内的本地项目管理器中，首次获取项目大小优化到 `1s` 内 👉
  [x-pm](https://github.com/dishait/x-pm)

<br />

## 灵感来源

[esbuild](https://github.com/evanw/esbuild)

<br />

## License

Made with [markthree](https://github.com/markthree)

Published under
[MIT License](https://github.com/markthree/go-get-folder-size/blob/main/LICENSE).

<br />
