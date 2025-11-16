import { defineConfig } from "eslint/config";

import { baseConfig } from "@axon/eslint-config/base";
import { reactConfig } from "@axonn/eslint-config/react";

export default defineConfig(
  {
    ignores: [".nitro/**", ".output/**", ".tanstack/**"],
  },
  baseConfig,
  reactConfig,
);
