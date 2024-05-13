"use strict";

const os = require("os");
const path = require("path");
const fs = require("fs");
const childProcess = require("child_process");

const BINARY_DISTRIBUTIONS = [
  { packageName: "@fireback/cli-darwin", subpath: "bin/fireback" },
  { packageName: "@fireback/cli-linux-x64", subpath: "bin/fireback" },
  { packageName: "@fireback/cli-linux-i686", subpath: "bin/fireback" },
  { packageName: "@fireback/cli-linux-arm64", subpath: "bin/fireback" },
  { packageName: "@fireback/cli-linux-arm", subpath: "bin/fireback" },
  { packageName: "@fireback/cli-win32-x64", subpath: "bin/fireback.exe" },
  { packageName: "@fireback/cli-win32-i686", subpath: "bin/fireback.exe" },
];

/**
 * This convoluted function resolves the path to the manually downloaded fallback
 * `fireback` binary in a way that can't be analysed by @vercel/nft.
 *
 * Without this, the binary can be detected as an asset and included by bundlers
 * that use @vercel/nft.
 *
 * @returns {string} The path to the fireback binary
 */
function getFallbackBinaryPath() {
  const parts = [];
  parts.push(__dirname);
  parts.push("..");
  parts.push(`fireback${process.platform === "win32" ? ".exe" : ""}`);
  return path.resolve(...parts);
}

function getDistributionForThisPlatform() {
  const arch = os.arch();
  const platform = os.platform();

  let packageName = undefined;
  if (platform === "darwin") {
    if (arch === "x64") {
      packageName = "@fireback/cli-darwin-x64";
    }
  } else if (platform === "linux" || platform === "freebsd") {
    switch (arch) {
      case "x64":
        packageName = "@fireback/cli-linux-x64";
        break;
      case "x86":
      case "ia32":
        packageName = "@fireback/cli-linux-i686";
        break;
      case "arm64":
        packageName = "@fireback/cli-linux-arm64";
        break;
      case "arm":
        packageName = "@fireback/cli-linux-arm";
        break;
    }
  } else if (platform === "win32") {
    switch (arch) {
      case "x64":
      // Windows arm64 can run x64 binaries
      case "arm64":
        packageName = "@fireback/cli-win32-x64";
        break;
      case "x86":
      case "ia32":
        packageName = "@fireback/cli-win32-i686";
        break;
    }
  }

  let subpath = undefined;
  switch (platform) {
    case "win32":
      subpath = "bin/fireback.exe";
      break;
    case "darwin":
    case "linux":
    case "freebsd":
      subpath = "bin/fireback";
      break;
    default:
      subpath = "bin/fireback";
      break;
  }

  return { packageName, subpath };
}

/**
 * Throws an error with a message stating that Fireback CLI doesn't support the current platform.
 *
 * @returns {never} nothing. It throws.
 */
function throwUnsupportedPlatformError() {
  throw new Error(
    `Unsupported operating system or architecture! Fireback CLI does not work on this architecture.

Fireback CLI supports:
- Darwin (macOS)
- Linux and FreeBSD on x64, x86, ia32, arm64, and arm architectures
- Windows x64, x86, and ia32 architectures`
  );
}

/**
 * Tries to find the installed Fireback CLI binary - either by looking into the relevant
 * optional dependencies or by trying to resolve the fallback binary.
 *
 * @returns {string} The path to the fireback binary
 */
function getBinaryPath() {
  if (process.env.FIREBACK_BINARY_PATH) {
    return process.env.FIREBACK_BINARY_PATH;
  }

  const { packageName, subpath } = getDistributionForThisPlatform();

  if (packageName === undefined) {
    throwUnsupportedPlatformError();
  }

  let fallbackBinaryPath = getFallbackBinaryPath();
  if (fs.existsSync(fallbackBinaryPath)) {
    // Since the fallback got installed, the optional dependencies likely didn't get installed, so we just default to the fallback.
    return fallbackBinaryPath;
  }

  let compatibleBinaryPath;
  try {
    compatibleBinaryPath = require.resolve(`${packageName}/${subpath}`);
  } catch (e) {
    const otherInstalledDistribution = BINARY_DISTRIBUTIONS.find(
      ({ packageName, subpath }) => {
        try {
          require.resolve(`${packageName}/${subpath}`);
          return true;
        } catch (e) {
          return false;
        }
      }
    );

    // These error messages are heavily inspired by esbuild's error messages: https://github.com/evanw/esbuild/blob/f3d535262e3998d845d0f102b944ecd5a9efda57/lib/npm/node-platform.ts#L150
    if (otherInstalledDistribution) {
      throw new Error(`Fireback CLI binary for this platform/architecture not found!

The "${otherInstalledDistribution.packageName}" package is installed, but for the current platform, you should have the "${packageName}" package installed instead. This usually happens if the "@fireback/cli" package is installed on one platform (for example Windows or MacOS) and then the "node_modules" folder is reused on another operating system (for example Linux in Docker).

To fix this, avoid copying the "node_modules" folder, and instead freshly install your dependencies on the target system. You can also configure your package manager to install the right package. For example, yarn has the "supportedArchitectures" feature: https://yarnpkg.com/configuration/yarnrc/#supportedArchitecture.`);
    } else {
      throw new Error(`Fireback CLI binary for this platform/architecture not found!

Looking for ${packageName}, but it's not there.

It seems like none of the "@fireback/cli" package's optional dependencies got installed. Please make sure your package manager is configured to install optional dependencies. If you are using npm to install your dependencies, please don't set the "--no-optional", "--ignore-optional", or "--omit=optional" flags. Fireback CLI needs the "optionalDependencies" feature in order to install its binary.`);
    }
  }

  return compatibleBinaryPath;
}

/**
 * Will be used as the binary path when defined with `mockBinaryPath`.
 * @type {string | undefined}
 */
let mockedBinaryPath;

/**
 * Overrides the default binary path with a mock value, useful for testing.
 *
 * @param {string} mockPath The new path to the mock fireback binary
 * @deprecated This was used in tests internally and will be removed in the next major version.
 */
// TODO(v3): Remove this function
function mockBinaryPath(mockPath) {
  mockedBinaryPath = mockPath;
}

/**
 * The javascript type of a command line option.
 * @typedef {'array'|'string'|'boolean'|'inverted-boolean'} OptionType
 */

/**
 * Schema definition of a command line option.
 * @typedef {object} OptionSchema
 * @prop {string} param The flag of the command line option including dashes.
 * @prop {OptionType} type The value type of the command line option.
 */

/**
 * Schema definition for a command.
 * @typedef {Object.<string, OptionSchema>} OptionsSchema
 */

/**
 * Serializes command line options into an arguments array.
 *
 * @param {OptionsSchema} schema An options schema required by the command.
 * @param {object} options An options object according to the schema.
 * @returns {string[]} An arguments array that can be passed via command line.
 */
function serializeOptions(schema, options) {
  return Object.keys(schema).reduce((newOptions, option) => {
    const paramValue = options[option];
    if (paramValue === undefined || paramValue === null) {
      return newOptions;
    }

    const paramType = schema[option].type;
    const paramName = schema[option].param;

    if (paramType === "array") {
      if (!Array.isArray(paramValue)) {
        throw new Error(`${option} should be an array`);
      }

      return newOptions.concat(
        paramValue.reduce(
          (acc, value) => acc.concat([paramName, String(value)]),
          []
        )
      );
    }

    if (paramType === "boolean") {
      if (typeof paramValue !== "boolean") {
        throw new Error(`${option} should be a bool`);
      }

      const invertedParamName = schema[option].invertedParam;

      if (paramValue && paramName !== undefined) {
        return newOptions.concat([paramName]);
      }

      if (!paramValue && invertedParamName !== undefined) {
        return newOptions.concat([invertedParamName]);
      }

      return newOptions;
    }

    return newOptions.concat(paramName, paramValue);
  }, []);
}

/**
 * Serializes the command and its options into an arguments array.
 *
 * @param {string} command The literal name of the command.
 * @param {OptionsSchema} [schema] An options schema required by the command.
 * @param {object} [options] An options object according to the schema.
 * @returns {string[]} An arguments array that can be passed via command line.
 */
function prepareCommand(command, schema, options) {
  return command.concat(serializeOptions(schema || {}, options || {}));
}

/**
 * Returns the absolute path to the `fireback` binary.
 * @returns {string}
 */
function getPath() {
  return mockedBinaryPath !== undefined ? mockedBinaryPath : getBinaryPath();
}

/**
 * Runs `fireback` with the given command line arguments.
 *
 * Use {@link prepareCommand} to specify the command and add arguments for command-
 * specific options. For top-level options, use {@link serializeOptions} directly.
 *
 * The returned promise resolves with the standard output of the command invocation
 * including all newlines. In order to parse this output, be sure to trim the output
 * first.
 *
 * If the command failed to execute, the Promise rejects with the error returned by the
 * CLI. This error includes a `code` property with the process exit status.
 *
 * @example
 * const output = await execute(['--version']);
 * expect(output.trim()).toBe('fireback x.y.z');
 *
 * @param {string[]} args Command line arguments passed to `fireback`.
 * @param {boolean} live We inherit stdio to display `fireback` output directly.
 * @param {boolean} silent Disable stdout for silents build (CI/Webpack Stats, ...)
 * @param {string} [configFile] Relative or absolute path to the configuration file.
 * @param {Object} [config] More configuration to pass to the CLI
 * @returns {Promise.<string>} A promise that resolves to the standard output.
 */
async function execute(args, live, silent, configFile, config = {}) {
  const env = { ...process.env };
  if (configFile) {
    env.FIREBACK_PROPERTIES = configFile;
  }
  if (config.url) {
    env.FIREBACK_URL = config.url;
  }
  if (config.authToken) {
    env.FIREBACK_AUTH_TOKEN = config.authToken;
  }
  if (config.apiKey) {
    env.FIREBACK_API_KEY = config.apiKey;
  }
  if (config.dsn) {
    env.FIREBACK_DSN = config.dsn;
  }
  if (config.org) {
    env.FIREBACK_ORG = config.org;
  }
  if (config.project) {
    env.FIREBACK_PROJECT = config.project;
  }
  if (config.vcsRemote) {
    env.FIREBACK_VCS_REMOTE = config.vcsRemote;
  }
  if (config.customHeader) {
    env.CUSTOM_HEADER = config.customHeader;
  } else if (config.headers) {
    const headers = Object.entries(config.headers).flatMap(([key, value]) => [
      "--header",
      `${key}:${value}`,
    ]);
    args = [...headers, ...args];
  }
  return new Promise((resolve, reject) => {
    if (live === true) {
      const output = silent ? "ignore" : "inherit";
      const pid = childProcess.spawn(getPath(), args, {
        env,
        // stdin, stdout, stderr
        stdio: ["ignore", output, output],
      });
      pid.on("exit", () => {
        resolve();
      });
    } else {
      childProcess.execFile(getPath(), args, { env }, (err, stdout) => {
        if (err) {
          reject(err);
        } else {
          resolve(stdout);
        }
      });
    }
  });
}

function getProjectFlagsFromOptions({ projects = [] } = {}) {
  return projects.reduce((flags, project) => flags.concat("-p", project), []);
}

module.exports = {
  execute,
  getPath,
  getProjectFlagsFromOptions,
  mockBinaryPath,
  prepareCommand,
  serializeOptions,
  getDistributionForThisPlatform,
  throwUnsupportedPlatformError,
  getFallbackBinaryPath,
};
