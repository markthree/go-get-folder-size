import { execa } from "execa";
import prettyBytes from "pretty-bytes";
import { arch as _arch, platform as _platform } from "node:os";

import { fileURLToPath } from "node:url";
import { dirname, resolve } from "node:path";

const _dirname = typeof __dirname !== "undefined"
  ? __dirname
  : dirname(fileURLToPath(import.meta.url));

let defaultBinPath = "";

export function inferVersion() {
  const platform = _platform();
  if (!/win32|linux|darwin/.test(platform)) {
    throw new Error(`${platform} is not support`);
  }

  const arch = _arch();

  if (!/amd64_v1|arm64|386|x64/.test(arch)) {
    throw new Error(`${arch} is not support`);
  }

  return `${platform === "win32" ? "windows" : platform}_${
    arch === "x64" ? "amd64_v1" : arch
  }`;
}

export function detectBinName(version = inferVersion()) {
  return `go-get-folder-size${version.startsWith("windows") ? ".exe" : ""}`;
}

export function detectDefaultBinPath() {
  if (defaultBinPath) {
    return defaultBinPath;
  }

  const version = inferVersion();
  const name = detectBinName(version);
  defaultBinPath = resolve(
    _dirname,
    `../dist/go-get-folder-size_${version}/${name}`,
  );
  return defaultBinPath;
}

interface Options {
  binPath?: string;
}

export async function getFolderSizeBin(
  base: string,
  pretty?: false,
  options?: Options,
): Promise<number>;
export async function getFolderSizeBin(
  base: string,
  pretty?: true,
  options?: Options,
): Promise<string>;
export async function getFolderSizeBin(
  base: string,
  pretty = false,
  options: Options = {},
) {
  const { binPath = detectDefaultBinPath() } = options;

  const { stdout, stderr } = await execa(binPath, {
    cwd: base,
  });

  if (stderr) {
    throw stderr;
  }

  if (pretty) {
    return prettyBytes(Number(stdout));
  }

  return Number(stdout);
}

export function createGetFolderSizeBinIpc(options: Options = {}) {
  const { binPath = detectDefaultBinPath() } = options;

  let tasks = new Map<
    string,
    { pretty: boolean; resolve: Function; reject: Function }
  >();

  const go = execa(binPath, {
    env: {
      ipc: String(true),
    },
  });

  function close() {
    if (!go.killed) {
      go.cancel();
      tasks.clear();
      tasks = null;
    }
  }

  function send(base: string) {
    go.stdin.write(`${base},`);
  }

  go.stdout.on("data", (item: string) => {
    const [base, size] = String(item).split(",");
    const { pretty, resolve } = tasks.get(base);
    resolve(pretty ? prettyBytes(Number(size)) : Number(size));
    tasks.delete(base);
  });

  go.stderr.on("data", (item: string) => {
    const [base, ...error] = String(item).split(",");
    console.log("error", base);
    const { reject } = tasks.get(base);
    reject(error.toString());
    tasks.delete(base);
  });

  process.once("exit", () => {
    close();
  });

  async function getFolderSizeWithIpc(
    base: string,
    pretty?: false,
  ): Promise<number>;
  async function getFolderSizeWithIpc(
    base: string,
    pretty?: true,
  ): Promise<string>;
  async function getFolderSizeWithIpc(
    base: string,
    pretty = false,
  ): Promise<number | string> {
    return new Promise((resolve, reject) => {
      tasks.set(base, { pretty, resolve, reject });
      send(base);
    });
  }

  return {
    close,
    getFolderSizeWithIpc,
  };
}
