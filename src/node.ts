import prettyBytes from 'pretty-bytes'
import { readdir, lstat } from 'node:fs/promises'

export function slash(path) {
	return path.replace(/\\/g, '/')
}

export function zipSizes(sizes: number[]) {
	return sizes.reduce((total, size) => (total += size), 0)
}

export async function getFolderSize(
	base: string,
	pretty: false
): Promise<number>
export async function getFolderSize(
	base: string,
	pretty: true
): Promise<string>
export async function getFolderSize(
	base: string,
	pretty = false
) {
	const dirents = await readdir(base, {
		withFileTypes: true
	})
	if (dirents.length === 0) {
		return 0
	}

	const files = []
	const directorys = []

	for (const dirent of dirents) {
		if (dirent.isFile()) {
			files.push(dirent)
			continue
		}
		if (dirent.isDirectory()) {
			directorys.push(dirent)
		}
	}

	const sizes = await Promise.all(
		[
			files.map(async file => {
				const path = `${slash(base)}/${file.name}`
				const { size } = await lstat(path)
				return size
			}),
			directorys.map(directory => {
				const path = `${slash(base)}/${directory.name}`
				return getFolderSize(path, false)
			})
		].flat()
	)

	if (!pretty) {
		return zipSizes(sizes)
	}

	return prettyBytes(zipSizes(sizes))
}
