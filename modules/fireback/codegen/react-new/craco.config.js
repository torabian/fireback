const path = require("path");
const { getLoader, loaderByName } = require("@craco/craco");

const alias = {
  "@": path.resolve(__dirname, "src"),
  "@apps": path.resolve(__dirname, "src", "apps", process.env.TARGET_APP),
};
const appConfigs = {
  // anotherapp: {
  //   port: 3001,
  // },
};

const options = {
  devServer: {
    ...appConfigs[process.env.TARGET_APP],
    port: 3670,
    client: {
      overlay: {
        runtimeErrors: (error) => {
          if (
            error?.message ===
            "ResizeObserver loop completed with undelivered notifications."
          ) {
            console.error(error);
            return false;
          }
          return true;
        },
      },
    },
  },
  environment: {},
  webpack: {
    alias,
    configure: (webpackConfig) => {
      webpackConfig.module.rules = [
        ...webpackConfig.module.rules,
        {
          test: /\.(tsx?|scss)$/,
          exclude: /node_modules/,
          use: [
            {
              loader: "ifdef-loader",
              options: {
                TARGET_APP: process.env.TARGET_APP,
                REACT_APP_SUPPORTED_LANGUAGES:
                  process.env.REACT_APP_SUPPORTED_LANGUAGES,
                // env: process.env,
                // ...process.env,
              },
            },
          ],
        },
      ];

      const { isFound, match: fileLoaderMatch } = getLoader(
        webpackConfig,
        (rule) => rule.type === "asset/resource"
      );

      if (!isFound) {
        throw {
          message: `Can't find file-loader in the ${context.env} webpack config!`,
        };
      }

      fileLoaderMatch.loader.exclude.push(/\.ya?ml$/);

      const yamlLoader = {
        use: "yaml-loader",
        test: /\.(ya?ml)$/,
      };
      webpackConfig.module.rules.push(yamlLoader);

      return webpackConfig;
    },
  },
  typescript: {
    compilerOptions: {
      paths: {
        "@/*": ["./src/*"],
        "@apps/*": ["./src/apps/" + process.env.TARGET_APP + "/*"],
      },
    },
  },
};

module.exports = options;
