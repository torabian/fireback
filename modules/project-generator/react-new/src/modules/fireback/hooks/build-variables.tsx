const UntypedEnvVariables = (import.meta as any).env.BUILD_VARIABLES

export const BUILD_VARIABLES = {
  /**
   * The location of the api endpoint which will be used to communicate as backend.
   */
  REMOTE_SERVICE: UntypedEnvVariables.VITE_REMOTE_SERVICE,

  /**
   * Public url of the project.
   */
  PUBLIC_URL: UntypedEnvVariables.PUBLIC_URL,

  /**
   * Default route when the app opens it will be redirected to.
   * can be the authentication location
   */
  DEFAULT_ROUTE: UntypedEnvVariables.VITE_DEFAULT_ROUTE,

  /**
   * Target application to be built.
   * In this react project, from same source code we can build different applications
   * with completely differnt entry points, hence modules, packages, and files
   * even might be completely different.
   * This is useful when building apps which have multiple sections, such as user panel,
   * admin panel. Or can be used as different branding for white labeling.
   */
  TARGET_APP: UntypedEnvVariables.TARGET_APP,

  /**
   * Enables the react app to show the os login button.
   * You can enable this on projects which are desktop, so the user doesn't
   * necessary have to be online, and can authenticate with user account on his computer
   */
  ALLOW_OS_LOGIN: UntypedEnvVariables.VITE_ALLOW_OS_LOGIN,

  /**
   * Supported languages in the app, which would import the corresponding
   * translation file.
   */
  SUPPORTED_LANGUAGES: UntypedEnvVariables.VITE_SUPPORTED_LANGUAGES,

  /**
   * Changes the router mechanism. Hash router it useful, when you want
   * to deploy the application into github pages for example, or when
   * using with electron.js there is no need to keep actual route to user.
   * Only changes the primary window, the 2nd+n windows are all memory router.
   */
  USE_HASH_ROUTER: UntypedEnvVariables.VITE_USE_HASH_ROUTER,

  /**
   * Project includes a mock server which is written in js.
   * Enabling this, would allow react.js app to work with that mock server,
   * instead of trying to connect the remote service.
   */
  INACCURATE_MOCK_MODE: UntypedEnvVariables.VITE_INACCURATE_MOCK_MODE,

  /**
   * Defautl application title, used in index.html and can be useful other places
   * when adding extra titles.
   */
  TITLE: UntypedEnvVariables.VITE_TITLE,

  /**
   * Enable when demoing the react app into github pages.
   */
  GITHUB_DEMO: UntypedEnvVariables.VITE_GITHUB_DEMO,

  /**
   * This is used in few places, but never tested.
   * It's forcing application to always be in an speific locale
   */
  FORCED_LOCALE: UntypedEnvVariables.VITE_FORCED_LOCALE,

  /**
   * This is also to prevent user be able to change the interface theme.
   */
  FORCE_APP_THEME: UntypedEnvVariables.VITE_FORCE_APP_THEME,

  /**
   * The route, where user will be navigated upon signout from
   * the ui interface
   */
  NAVIGATE_ON_SIGNOUT: UntypedEnvVariables.VITE_NAVIGATE_ON_SIGNOUT
}
