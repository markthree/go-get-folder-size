import { bench, describe } from 'vitest'
import {
	createGetFolderSizeBinIpc,
	getFolderSize,
	getFolderSizeBin,
	getFolderSizeWasm
} from './npm'

const { getFolderSizeWithIpc } = createGetFolderSizeBinIpc()

describe('basic', () => {
	const base = `./`
	bench('getFolderSize', async () => {
		await getFolderSize(base)
	})

	bench('getFolderSizeBin', async () => {
		await getFolderSizeBin(base)
	})

	bench('getFolderSizeWithIpc', async () => {
		await getFolderSizeWithIpc(base)
	})

	bench('getFolderSizeWasm', async () => {
		await getFolderSizeWasm(base)
	})
})
