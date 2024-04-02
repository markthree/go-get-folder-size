import { resolve } from "node:path";
import { runtime } from "std-env";
import prettyBytes from "pretty-bytes";
import { lstat, readdir } from "node:fs/promises";

interface Options {
  /**
   * @default false
   */
  loose?: boolean;
}

async function strictGetFileSize(path: string) {
  const { size } = await lstat(path);
  return size;
}

async function looseGetFileSize(path: string) {
  try {
    const size = await strictGetFileSize(path);
    return size;
  } catch (error) {
    return 0;
  }
}

export async function getFolderSize(
  base: string,
  pretty?: false,
  options?: Options,
): Promise<number>;
export async function getFolderSize(
  base: string,
  pretty?: true,
  options?: Options,
): Promise<string>;
export async function getFolderSize(
  base: string,
  pretty = false,
  options?: Options,
) {
  const { loose = false } = options || {};
  let total = 0;
  const sumTotal = (size: number) => total += size;
  const getFileSize = loose ? looseGetFileSize : strictGetFileSize;

  // bun (use recursive)
  if (runtime === "bun") {
    const dirents = await readdir(base, {
      recursive: true,
      withFileTypes: true,
    });

    if (dirents.length === 0) {
      return mayBeWithPrettyBytes();
    }

    const promises = dirents.map(async (dirent) => {
      if (!dirent.isFile()) {
        return;
      }
      const size = await getFileSize(resolve(base, dirent.name));
      sumTotal(size);
    });

    await Promise.all(promises);

    return mayBeWithPrettyBytes();
  }

  const dirents = await readdir(base, {
    withFileTypes: true,
  });

  if (dirents.length === 0) {
    return mayBeWithPrettyBytes();
  }

  const promises: Array<Promise<number>> = [];

  for (const dirent of dirents) {
    if (dirent.isFile()) {
      const path = resolve(base, dirent.name);
      promises.push(
        getFileSize(path).then(sumTotal),
      );
      continue;
    }
    if (dirent.isDirectory()) {
      const path = resolve(base, dirent.name);
      promises.push(
        getFolderSize(path, false, options).then(sumTotal),
      );
    }
  }

  await Promise.all(promises);

  return mayBeWithPrettyBytes();

  function mayBeWithPrettyBytes() {
    return pretty ? prettyBytes(total) : total;
  }
}
