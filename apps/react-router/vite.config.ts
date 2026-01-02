import { execSync } from "node:child_process";
import { readFileSync } from "node:fs";
import { resolve } from "node:path";
import { fileURLToPath, URL } from "node:url";
import tailwindcss from "@tailwindcss/vite";
import { devtools } from "@tanstack/devtools-vite";
import { tanstackRouter } from "@tanstack/router-plugin/vite";
import viteReact from "@vitejs/plugin-react";
import { defineConfig } from "vite";

function getAppVersion() {
  try {
    const pkg = JSON.parse(
      readFileSync(resolve(__dirname, "package.json"), "utf-8"),
    );
    return pkg.version;
  } catch (e) {
    throw new Error("Unable to read package version");
  }
}

function getCommitHash() {
  try {
    return execSync("git rev-parse --short HEAD").toString().trim();
  } catch (e) {
    throw new Error("Unable to read commit hash");
  }
}

function getBuildTime() {
  return new Date().toISOString();
}

const versionInfo = {
  version: getAppVersion(),
  hash: getCommitHash(),
  buildTime: getBuildTime(),
};

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
          if (id.includes("react-dom") || id.includes("react")) {
            return "react";
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
    {
      name: "html-transform",
      transformIndexHtml(html) {
        return html
          .replace("%APP_VERSION%", versionInfo.version)
          .replace("%COMMIT_HASH%", versionInfo.hash)
          .replace("%BUILD_TIME%", versionInfo.buildTime);
      },
    },
  ],
  resolve: {
    alias: {
      "@": fileURLToPath(new URL("./src", import.meta.url)),
    },
  },

  define: {
    __CLIENT_VERSION__: JSON.stringify(versionInfo.version),
    __CLIENT_COMMIT_HASH__: JSON.stringify(versionInfo.hash),
    __CLIENT_BUILD_TIME__: JSON.stringify(versionInfo.buildTime),
  },
});
