import '../wasm/polyfill.js'
import '../wasm/wasm_exec.js'
import init from '../wasm/main.wasm?init'

const go = new global.Go()

init(go.importObject).then(instance => {
	return go.run(instance)
})
