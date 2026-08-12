# fireback-ui packages

Everything under this directory is a separate npm package, wired together with
[npm workspaces](https://docs.npmjs.com/cli/v10/using-npm/workspaces) (see the root
`package.json`'s `"workspaces": ["packages/*"]`).

| package | what it is | depends on |
|---|---|---|
| `@fireback/js-remote-ctx` | shared runtime every generated SDK module references: fetch/WebSocket/SSE wrappers, response envelopes, React hooks (`useFetchx`/`useSse`/`useWebSocketX`) | — |
| `@fireback/auth-client` | standalone fireback self-service auth client: session/workspace state, `AuthenticationProvider`/`useAuthentication`, self-service login redirect + session capture. No fireback UI dependencies - meant to be pulled into a brand new React project on its own | — |
| `@fireback/complexes` | TypeScript counterparts of fireback's Go `complexes` package (`TString`, …) - plain value types, no dependencies at all (not even React) | — |
| `@fireback/ui-core` | hooks, component library ("the core") - carries its own generated `sdk/` (abac, messaging, interfacetools, reactivesearch, eventbus) | `@fireback/auth-client`, `@fireback/js-remote-ctx` |
| `@fireback/messaging` | mail/GSM provider admin screens - carries its own generated `sdk/messaging` | `@fireback/ui-core`, `@fireback/js-remote-ctx` |
| `@fireback/selfservice` | sign-in/up, passports, OTP/TOTP, workspace selection - carries its own generated `sdk/` (abac, messaging) | `@fireback/ui-core`, `@fireback/auth-client`, `@fireback/js-remote-ctx` |
| `@fireback/manage` | administration screens (capabilities, users, workspaces, …) - carries its own generated `sdk/` (abac, messaging) | `@fireback/ui-core`, `@fireback/auth-client`, `@fireback/messaging`, `@fireback/js-remote-ctx` |
| `@fireback/mobile-kit` | mobile dashboard kit | `@fireback/ui-core` |
| `@fireback/styles` | shared stylesheets (base + apple-family themes) | — |
| `@fireback/enterprise-shell` | enterprise application shell every app boots through - `EssentialApp`/`EssentialRouter`, panel & sidebar layout, `WithFireback`/`WithSelfServiceRoutes` provider wiring | `@fireback/ui-core`, `@fireback/auth-client`, `@fireback/manage`, `@fireback/mobile-kit`, `@fireback/selfservice`, `@fireback/styles`, `@fireback/js-remote-ctx` |

## No shared `@fireback/sdk` - each package carries its own generated SDK

There used to be one shared `@fireback/sdk` package holding every backend module's
generated actions/DTOs, imported by whichever package needed it. That's gone now -
every `<package>/sdk/<module>/` folder above (`ui-core/sdk/abac`, `manage/sdk/abac`,
`selfservice/sdk/abac`, ...) is generated straight into that package by `make defs`,
via a dedicated `- compiler: js` target per consumer in the owning backend module's
`*.emi.yml` (see `modules/abac/Abac.emi.yml`, `modules/abac/messaging/Messaging.emi.yml`,
etc.). When more than one package needs the same module (`abac`'s SDK, for instance,
is used by `ui-core`, `manage`, and `selfservice`), each one gets its own independent
copy - the duplication is intentional, not drift: it means every package is
self-contained (no `@fireback/sdk` cross-package dependency edge to carry around) and
`make defs` regenerates every copy identically in one pass, so they can't go stale
relative to each other. Only `@fireback/js-remote-ctx` (the runtime the generated code
calls into) is still a single shared package - see its own note in `Makefile`'s
`defs-sdk` target for why that one stays centralized.

None of these packages ship a build step or a `dist/` — they're raw `.ts`/`.tsx`/`.css`
source, compiled by whichever app imports them. That's required, not just simpler:
`@fireback/ui-core` reads `BUILD_VARIABLES` (from `import.meta.env`, injected by the
*consuming app's* `vite.config.ts` per build mode) and uses
`vite-plugin-conditional-compiler` `#if` macros, both of which only resolve correctly
when compiled as part of the final app's own Vite build.

## Consuming from another project (no registry)

These packages aren't published to any npm registry. There are two ways to pull them
into a downstream project instead: install a released tarball (no submodule, nothing
to clone), or vendor the whole repo as a git submodule (needed if you want
`@fireback/ui-core`/`selfservice`/`manage`/`messaging`, which read `BUILD_VARIABLES`
injected by *this* repo's own `vite.config.ts` - see the "don't ship a dist/" note
above).

### Option A: install a released tarball

Every tagged release (see `.github/workflows/fireback-build.yml`'s
`build-ui-packages`/`deploy_github_release` jobs, or run `make ui-packages-pack`
locally) attaches one plain npm tarball per package - `fireback-ui-core-0.1.0.tgz`,
`fireback-js-remote-ctx-0.1.0.tgz`, etc. - to the GitHub release as regular assets.
npm can install straight from a tarball URL, so a brand new project just needs:

```json
{
  "dependencies": {
    "@fireback/js-remote-ctx": "https://github.com/torabian/fireback/releases/download/<tag>/fireback-js-remote-ctx-0.1.0.tgz",
    "@fireback/ui-core": "https://github.com/torabian/fireback/releases/download/<tag>/fireback-ui-core-0.1.0.tgz"
  }
}
```

then `npm install`. List every package you need as its own tarball URL, same as the
`file:` example below - npm resolves a package's own `"@fireback/js-remote-ctx": "*"`
dependency against whichever copy your own `package.json` already pins, it won't try
(and fail) to fetch it from the registry. `@fireback/js-remote-ctx` on its own is the
only fully registry-free, framework-agnostic package (see the table above); everything
built on `ui-core` still needs Option B to actually build.

To build the tarballs yourself instead of using a release: `make ui-packages-pack` at
the repo root writes them to `artifacts/fireback-packages/`, and
`npm install ./path/to/fireback-ui-core-0.1.0.tgz` works identically to a URL.

### Option B: vendor via git submodule

1. Add this repo as a submodule, pinned to a tag/commit:

   ```sh
   git submodule add git@github.com:torabian/fireback.git vendor/fireback
   # or, if you only want the ui half tracked at a known-good commit:
   git -C vendor/fireback checkout <tag-or-sha>
   ```

2. In the downstream project's `package.json`, depend on whichever packages you
   actually need via `file:`:

   ```json
   {
     "dependencies": {
       "@fireback/js-remote-ctx": "file:vendor/fireback/ui/packages/js-remote-ctx",
       "@fireback/ui-core": "file:vendor/fireback/ui/packages/ui-core",
       "@fireback/selfservice": "file:vendor/fireback/ui/packages/selfservice",
       "@fireback/manage": "file:vendor/fireback/ui/packages/manage",
       "@fireback/messaging": "file:vendor/fireback/ui/packages/messaging",
       "@fireback/styles": "file:vendor/fireback/ui/packages/styles"
     }
   }
   ```

3. Run `npm install` **at the downstream project's root**. npm symlinks each `file:`
   target into `node_modules/@fireback/*` and installs the dependencies each package
   declares in its own `package.json` (React, Formik, date libs, etc.).

4. To pick up updates: bump the submodule pointer (`git -C vendor/fireback pull` /
   `git submodule update --remote`, then commit the new pointer) and run
   `npm install` again.

**Never run `npm install` inside the vendored submodule itself** — only at the
downstream project's root. That keeps a single hoisted `node_modules`, so there's one
copy of React and the symlinked packages resolve their peer deps correctly.

Every cross-module import inside these packages already uses the bare `@fireback/...`
specifier (never a relative `../` reach into a sibling package) - including each
package's own generated `sdk/` folder (e.g. a file in `manage/` imports its own SDK as
`@fireback/manage/sdk/abac/RoleDto`, not a relative path) - so a downstream project can
vendor a subset — e.g. just `@fireback/ui-core` on its own — without dragging in
`manage`/`selfservice` if it doesn't need them.
