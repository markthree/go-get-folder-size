import { resolve } from "node:path";
import type { Dirent } from "node:fs";
import prettyBytes from "pretty-bytes";
import { lstat, readdir } from "node:fs/promises";

export function zipSizes(sizes: number[]) {
  return sizes.reduce((total, size) => (total += size), 0);
}

export async function getFolderSize(
  base: string,
  pretty?: false,
): Promise<number>;
export async function getFolderSize(
  base: string,
  pretty?: true,
): Promise<string>;
export async function getFolderSize(
  base: string,
  pretty = false,
) {
  const dirents = await readdir(base, {
    withFileTypes: true,
  });
  if (dirents.length === 0) {
    return 0;
  }

  const files: Dirent[] = [];
  const directorys: Dirent[] = [];

  for (const dirent of dirents) {
    if (dirent.isFile()) {
      files.push(dirent);
      continue;
    }
    if (dirent.isDirectory()) {
      directorys.push(dirent);
    }
  }

  const sizes = await Promise.all(
    [
      files.map(async (file) => {
        const path = resolve(base, file.name);
        const { size } = await lstat(path);
        return size;
      }),
      directorys.map((directory) => {
        const path = resolve(base, directory.name);
        return getFolderSize(path, false);
      }),
    ].flat(),
  );

  if (!pretty) {
    return zipSizes(sizes);
  }

  return prettyBytes(zipSizes(sizes));
}
