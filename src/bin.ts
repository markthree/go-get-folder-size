import { execa } from 'execa'
import { resolve } from 'node:path'
import { readdir } from 'node:fs/promises'
import {
	arch as _arch,
	platform as _platform
} from 'node:os'

let defaultBinPath = ''

export function inferVersion() {
	const platform = _platform()
	if (!/win32|linux|darwin/.test(platform)) {
		throw new Error(`${platform} is not support`)
	}

	const arch = _arch()

	if (!/amd64_v1|arm64|386|x64/.test(arch)) {
		throw new Error(`${arch} is not support`)
	}

	return `${platform === 'win32' ? 'windows' : platform}_${
		arch === 'x64' ? 'amd64_v1' : arch
	}`
}

export async function detectDefaultBinPath() {
	if (defaultBinPath) {
		return defaultBinPath
	}

	const dist = resolve(
		__dirname,
		'../dist',
		`go-get-folder-size_${inferVersion()}`
	)

	const [binPath] = await readdir(dist)
	defaultBinPath = resolve(dist, binPath)
	return defaultBinPath
}

interface Options {
	binPath?: string
}

export async function getFolderSizeBin(
	base: string,
	pretty?: false,
	options?: Options
): Promise<number>
export async function getFolderSizeBin(
	base: string,
	pretty?: true,
	options?: Options
): Promise<string>
export async function getFolderSizeBin(
	base: string,
	pretty = false,
	options: Options = {}
) {
	const { binPath = await detectDefaultBinPath() } = options

	const args = pretty ? ['-p'] : []

	try {
		const { stdout } = await execa(binPath, args, {
			cwd: base
		})

		if (pretty) {
			return stdout
		}

		return Number(stdout)
	} catch (error) {
		console.log(error)
	}
}
