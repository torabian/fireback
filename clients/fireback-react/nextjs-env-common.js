/**
 * assertEnvironments the env variables needed to be set before build
 */

const commonError = `If you are missing .env.local, make a copy from .env.production and place it in the root of the app`;

function getMultiTenantEnv(app) {
  if (!app || !app.env || !app.env[process.env.NODE_ENV]) {
    return {};
  }

  return app.env[process.env.NODE_ENV];
}

function assertEnvironment(appEnv) {
  const env = { ...process.env, ...(appEnv || {}) };
  // if (!env.NEXT_PUBLIC_GRAPH_ME_URI) {
  //   console.error(
  //     `NEXT_PUBLIC_GRAPH_ME_URI is necessary, set it using .env.local or targets/*/manifest.json`,
  //     commonError
  //   );
  //   process.exit();
  // }
}
module.exports = {
  assertEnvironment,
  getMultiTenantEnv,
};
