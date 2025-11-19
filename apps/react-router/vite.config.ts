import { fileURLToPath, URL } from "node:url";
import tailwindcss from "@tailwindcss/vite";
import { devtools } from "@tanstack/devtools-vite";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import viteReact from "@vitejs/plugin-react";
import { defineConfig } from "vite";

// https://vitejs.dev/config/
export default defineConfig({
  build: {
    outDir: "../server/dist",
    emptyOutDir: true,
    minify: true,
    rollupOptions: {
      output: {
        manualChunks(id) {
          // React core
          if (id.includes("react-dom")) {
            return "react-dom";
          }
          if (id.includes("react")) {
            return "react";
          }

          // TanStack libraries
          if (id.includes("@tanstack/react-router")) {
            return "router";
          }
          if (id.includes("@tanstack/react-query")) {
            return "query";
          }

          // UI components (shadcn/radix)
          if (
            id.includes("@radix-ui") ||
            id.includes("cmdk") ||
            id.includes("lucide-react")
          ) {
            return "ui";
          }

          // Monaco Editor (large dependency)
          if (
            id.includes("monaco-editor") ||
            id.includes("@monaco-editor/react")
          ) {
            return "monaco";
          }

          // Lexical Editor
          if (id.includes("lexical") || id.includes("@lexical/react")) {
            return "lexical";
          }
        },
      },
    },
  },

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
    devtools(),
    tanstackRouter({
      target: "react",
      autoCodeSplitting: true,
    }),
    viteReact(),
    tailwindcss(),
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },
});
