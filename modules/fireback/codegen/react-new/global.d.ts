declare module "*.css";
declare module "*.scss";

interface Window {
    // This is correctly any.
    // Problem with this is, we never access process.env or import.meta.env directly
    // in this project, since the build system can change over time.
    // BUILD_VARIABLES are only accessed in build-variables.tsx,
    // and there it would become type safe.
    BUILD_VARIABLES?: any;
}