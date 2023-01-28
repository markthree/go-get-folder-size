import { defineConfig } from 'vite'
import { name } from './package.json'
import { builtinModules } from 'module'

export default defineConfig({
	build: {
		outDir: 'npm',
		emptyOutDir: false,
		lib: {
			name,
			formats: ['cjs', 'es'],
			entry: ['./src/index.ts', './src/cli.ts'],
			fileName(f, n) {
				if (f === 'cjs') {
					return `${n}.cjs`
				}
				if (f === 'es') {
					return `${n}.mjs`
				}
				return `${n}.js`
			}
		},
		rollupOptions: {
			external: [
				...builtinModules,
				...builtinModules.map(v => `node:${v}`)
			]
		}
	}
})
