const { execSync } = require("node:child_process");
const { copyFile } = require("node:fs/promises");

const wasmExec = execSync("go env GOROOT").toString().replace(/\n/g, "") +
  "/misc/wasm/wasm_exec.js";

copyFile(wasmExec, "./wasm/wasm_exec.js");
