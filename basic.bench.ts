import { bench, describe } from "vitest";
import { getFolderSize, getFolderSizeBin, getFolderSizeWasm } from "./npm";

describe("basic", () => {
  const base = `./`;
  bench("getFolderSize", async () => {
    await getFolderSize(base);
  });

  bench("getFolderSizeBin", async () => {
    await getFolderSizeBin(base);
  });

  bench("getFolderSizeWasm", async () => {
    await getFolderSizeWasm(base);
  });
});
