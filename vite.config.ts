import { defineConfig } from 'vite'
import { name } from './package.json'

export default defineConfig({
	build: {
		outDir: 'npm',
		emptyOutDir: false,
		lib: {
			name,
			formats: ['cjs', 'es'],
			entry: './src/index.ts',
			fileName(f) {
				if (f === 'cjs') {
					return `index.cjs`
				}
				if (f === 'es') {
					return 'index.mjs'
				}
				return 'index.js'
			}
		},
		rollupOptions: {
			external: ['node:fs/promises']
		}
	}
})
