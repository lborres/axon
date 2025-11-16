import { defineConfig } from "eslint/config";

import { baseConfig } from "@axon/eslint-config/base";
import { reactConfig } from "@axon/eslint-config/react";

export default defineConfig(
  {
    ignores: ["dist/**"],
  },
  baseConfig,
  reactConfig,
);
