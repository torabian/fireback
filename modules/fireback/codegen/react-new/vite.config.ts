import react from '@vitejs/plugin-react'
import fs from 'fs'

import { resolve } from 'path'
import { defineConfig } from 'vite'
import tsconfigPaths from 'vite-tsconfig-paths'



// https://vite.dev/config/
export default defineConfig(({ command, mode }) => {
  console.log('Vite command:', command) // 'serve' or 'build'
  console.log('Current mode:', mode)    // 'development', 'production', or custom mode

  // Path to the folder
  const bv = mode.split('/')
  const variation = bv.pop()
  const jsonFolderPath = resolve(__dirname, `./src/apps/${bv.join('/')}/build-variables/${variation}.json`)
  let build_variables = {};

  // Check if folder exists
  if (fs.existsSync(jsonFolderPath)) {
    console.log("Build variation selected:", jsonFolderPath)
    build_variables = require(jsonFolderPath)
  } else {
    console.log("Build variation does not exists:", jsonFolderPath)
  }

  return {
    define: {
      BUILD_VARIABLES: build_variables,
      __BUILD_VARIABLES__: build_variables
    },
    plugins: [
      tsconfigPaths(),
      react({
        exclude: [
          'node_modules/**', // exclude everything else in node_modules
        ],
        babel: {
          plugins: [
            ['@babel/plugin-proposal-decorators', { version: 'legacy' }],
            ['@babel/plugin-proposal-class-properties', { loose: true }],
            ['@babel/plugin-transform-private-methods', { loose: true }],
            ['@babel/plugin-transform-private-property-in-object', { loose: true }],
          ],
        },
      }),
    ],
    resolve: {
      alias: [
        { find: "@/", replacement: resolve(__dirname, "./src/") }
      ]
    },
    // Optional: change envDir based on mode
    envDir: mode === 'staging' ? './config/staging-envs' : './config/env',
  }
})