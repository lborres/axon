import { defineConfig } from "eslint/config";

import { baseConfig } from "@alima/eslint-config/base";

export default defineConfig(
  {
    ignores: [".nitro/**", ".output/**", ".tanstack/**"],
  },
  baseConfig,
);
