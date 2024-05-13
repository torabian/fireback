"use strict";

const pkgInfo = require("../package.json");
const helper = require("./helper");
const Releases = require("./releases");

/**
 * Interface to and wrapper around the `fireback` executable.
 *
 * Commands are grouped into namespaces. See the respective namespaces for more
 * documentation. To use this wrapper, simply create an instance and call methods:
 *
 * @example
 * const cli = new FirebackCli();
 * console.log(FirebackCli.getVersion());
 *
 * @example
 * const cli = new FirebackCli('path/to/custom/fireback.properties');
 * const release = await cli.releases.proposeVersion());
 * console.log(release);
 */
class FirebackCli {
  /**
   * Creates a new `FirebackCli` instance.
   *
   * If the `configFile` parameter is specified, configuration located in the default
   * location and the value specified in the `FIREBACK_PROPERTIES` environment variable is
   * overridden.
   *
   * @param {string} [configFile] Relative or absolute path to the configuration file.
   * @param {Object} [options] More options to pass to the CLI
   */
  constructor(configFile, options) {
    if (typeof configFile === "string") {
      this.configFile = configFile;
    }
    this.options = options || { silent: false };
    this.releases = new Releases({ ...this.options, configFile });
  }

  /**
   * Returns the version of the installed `fireback` binary.
   * @returns {string}
   */
  static getVersion() {
    return pkgInfo.version;
  }

  /**
   * Returns an absolute path to the `fireback` binary.
   * @returns {string}
   */
  static getPath() {
    return helper.getPath();
  }

  /**
   * See {helper.execute} docs.
   * @param {string[]} args Command line arguments passed to `fireback`.
   * @param {boolean} live We inherit stdio to display `fireback` output directly.
   * @returns {Promise.<string>} A promise that resolves to the standard output.
   */
  execute(args, live) {
    return helper.execute(
      args,
      live,
      this.options.silent,
      this.configFile,
      this.options
    );
  }
}

module.exports = FirebackCli;
