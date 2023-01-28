import mri from 'mri'
import { getFolderSizeBin } from './bin'

function printUsage() {
	console.log(` 	go-get-folder-size 
	
	Get the size of a folder by recursively iterating through all its sub(files && folders). Use go, so high-speed

	usage:
		go-get-folder-size [options]

	options:
		-h, --help            check help
		-p, --pretty          pretty bytes (default true)
		-b, --base            target base path (default cwd)\n`)
}

async function main() {
	const _argv = process.argv.slice(2)
	const argv = mri(_argv, {
		default: {
			help: false,
			pretty: true,
			base: process.cwd()
		},
		string: ['base'],
		boolean: ['pretty', 'help'],
		alias: {
			h: ['help'],
			b: ['base'],
			p: ['pretty']
		}
	})

	if (argv.help) {
		printUsage()
	} else {
		const size = await getFolderSizeBin(
			argv.base,
			argv.pretty
		)
		console.log(size)
	}
}

main()
