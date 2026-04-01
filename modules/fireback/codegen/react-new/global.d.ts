declare module "*.css";
declare module "*.scss";

interface Window {
    BUILD_VARIABLES?: {
        REMOTE_SERVICE?: string;
        PUBLIC_URL?: string;
        DEFAULT_ROUTE?: string;
        TARGET_APP?: string;
        ALLOW_OS_LOGIN?: string;
        SUPPORTED_LANGUAGES?: string;
        USE_HASH_ROUTER?: string;
        INACCURATE_MOCK_MODE?: string;
        TITLE?: string;
        GITHUB_DEMO?: string;
        FORCED_LOCALE?: string;
        FORCE_APP_THEME?: string;
        NAVIGATE_ON_SIGNOUT?: string;
    };
}