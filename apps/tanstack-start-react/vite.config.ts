import tailwindcss from "@tailwindcss/vite";
import { tanstackStart } from "@tanstack/react-start/plugin/vite";
import viteReact from "@vitejs/plugin-react";
import { defineConfig } from "vite";
import viteTsConfigPaths from "vite-tsconfig-paths";

const config = defineConfig({
  server: {
    proxy: {
      "/api": {
        target: "http://localhost:8080", // Go API server
        changeOrigin: true,
        rewrite: (path) => path.replace(/^\/api/, "/api"),
        secure: false, // disable SSL check if you're using HTTP
      },
    },
  },

  plugins: [
    // this is the plugin that enables path aliases
    viteTsConfigPaths({
      projects: ["./tsconfig.json"],
    }),
    tailwindcss(),
    tanstackStart(),

    // WARN: react's vite plugin must come after start's vite plugin
    viteReact(),
  ],
});

export default config;
