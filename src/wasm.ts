import "../wasm/polyfill.js";
import "../wasm/wasm_exec.js";
// @ts-ignore
import init from "../wasm/main.wasm?init";
import prettyBytes from "pretty-bytes";

// Not recommended. It may be slower than the native node
export async function getFolderSizeWasm(
  base: string,
  pretty?: false,
): Promise<number>;
export async function getFolderSizeWasm(
  base: string,
  pretty?: true,
): Promise<string>;
export async function getFolderSizeWasm(
  base: string,
  pretty = false,
) {
  const go = new global.Go();
  go.env = { base };
  const instance = await init(go.importObject);
  await go.run(instance);
  if (global.$folderSizeError) {
    throw new Error(global.$folderSizeError);
  }
  const size = global.$folderSize;
  global.$folderSize = null;
  if (pretty) {
    return prettyBytes(Number(size));
  }
  return Number(size);
}
