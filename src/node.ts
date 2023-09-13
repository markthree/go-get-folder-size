import { resolve } from "node:path";
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
  const dirents = await readdir(base, {
    withFileTypes: true,
  });
  if (dirents.length === 0) {
    return 0;
  }

  const promises: Array<Promise<number>> = [];

  let total = 0;
  const sumTotal = (size: number) => total += size;
  const getFileSize = loose ? looseGetFileSize : strictGetFileSize;

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

  if (!pretty) {
    return total;
  }

  return prettyBytes(total);
}
