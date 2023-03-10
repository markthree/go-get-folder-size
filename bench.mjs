import {
	getFolderSize,
	getFolderSizeBin,
	createGetFolderSizeBinIpc
} from './npm/index.mjs'

const { getFolderSizeWithIpc, close } =
	createGetFolderSizeBinIpc()

const base = '../'

const nodeStartTime = Date.now()
const nodeResult = await getFolderSize(base, true)
const nodeDuration = Date.now() - nodeStartTime

const goBinStartTime = Date.now()
const goBinResult = await getFolderSizeBin(base, true)
const goBinDuration = Date.now() - goBinStartTime

const goIpcStartTime = Date.now()
const goIpcResult = await getFolderSizeWithIpc(base, true)
const goIpcDuration = Date.now() - goIpcStartTime

close()

console.log(
	`node  - duration: ${
		nodeDuration / 1000
	}s result: ${nodeResult}`
)
console.log(
	`goBin - duration: ${
		goBinDuration / 1000
	}s result: ${goBinResult}`
)
console.log(
	`goIpc - duration: ${
		goIpcDuration / 1000
	}s result: ${goIpcResult}`
)

console.log('\n')

console.log(
	'goBin vs node -',
	(nodeDuration / goBinDuration).toFixed(2) + ' ↑'
)

console.log(
	'goIpc vs node -',
	(nodeDuration / goIpcDuration).toFixed(2) + ' ↑'
)
