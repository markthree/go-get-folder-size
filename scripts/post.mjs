import { platform } from 'os'
import { existsSync } from 'fs'
import { chmod } from 'fs/promises'

const isWindows = platform() === 'win32'

// Obtain permissions silently
async function setBinPermissions() {
	try {
		if (!isWindows) {
			const { detectDefaultBinPath } = await import(
				'../npm/bin.mjs'
			)
			const binPath = detectDefaultBinPath()
			const binPathExists = existsSync(binPath)
			if (binPathExists) {
				await chmod(binPath, '755')
			}
		}
	} catch (error) {}
}

setBinPermissions()
