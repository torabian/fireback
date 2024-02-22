/** @type {import('next').NextConfig} */

const p = require("path");

// const manifest =
//   "." +
//   p.sep +
//   p.join("targets", process.env.TARGET_APP.replace(" ", ""), "manifest.json");

// const app = require(manifest);

const {
  getMultiTenantEnv,
  assertEnvironment,
} = require(`.${p.sep}nextjs-env-common`);

const nextConfig = {
  reactStrictMode: true,
  env: {},
  typescript: {
    ignoreBuildErrors: true,
  },
  eslint: {
    // Warning: This allows production builds to successfully complete even if
    // your project has ESLint errors.
    ignoreDuringBuilds: true,
  },
  experimental: {
    // appDir: true,
  },
  distDir: "build",
  images: {
    unoptimized: true,
  },
  env: {
    // ...app,
    // ...getMultiTenantEnv(app),

    // Since we also build the app using cra, react native, we set which is the builder
    RUNNING_ON_NEXT: true,
  },
};

// validates and kills if env is missing
assertEnvironment(nextConfig.env);

module.exports = nextConfig;
