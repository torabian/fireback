import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers the three root-only passport admin actions added to the user single screen
// (modules/manage/users/UserPassportsList.tsx) - driven through the real screen, not
// raw HTTP, for the actions themselves (setup/assertions use cy.request as root,
// matching manage-workspaces.cy.ts's own root-delete-rejected test):
//   1. "Set password" - opens a drawer, types a new password, and it actually takes
//      effect (SetPassportPasswordAction) - the new password signs in afterward.
//   2. "Disable 2FA" - only rendered when the passport actually has a totpSecret
//      (DisablePassportTotpAction), and clears it for real.
//   3. "Send reset email" - triggers the same OTP-email plumbing
//      ClassicPassportRequestOtpAction uses (SendPassportResetEmailAction), and the
//      emailed code genuinely completes a self-service reset on the real
//      ui/packages/selfservice "reset-password" screen
//      (CompletePassportPasswordResetAction), including signing the user back in.
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

const TARGET_EMAIL = `cypress-target-${Date.now()}@example.com`;
const TARGET_OLD_PASSWORD = "OldPass123";
const TARGET_NEW_PASSWORD = "BrandNewPass456";

interface SigninResponse {
  data: { item: { session: { token: string } } };
}

function setShared(key: string, value: string) {
  return cy.task("setShared", { key, value });
}
function sharedValues<K extends string>(
  keys: K[],
): Cypress.Chainable<Record<K, string>> {
  return cy.task("getSharedState", keys);
}

function loginAsRoot() {
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session("root-login", () => {
    cy.visit(ui("/manage/#/welcome"));
    cy.get("#value-input", { timeout: 10000 }).type(ROOT_EMAIL);
    cy.get("#submit-form").click({ force: true });
    cy.get("h1").should("have.text", "Enter Password");
    cy.get("#password-input").type(ROOT_PASSWORD);
    cy.get("#submit-form").click({ force: true });
    cy.wait(1000);
    cy.url({ timeout: 10000 }).should("include", "/dashboard");
  });
}

function signin(value: string, password: string) {
  return cy.request({
    method: "POST",
    url: ui("/passports/signin/classic"),
    failOnStatusCode: false,
    body: { value, password },
  });
}

function asRoot(rootToken: string) {
  return { authorization: rootToken, "workspace-id": "root" };
}

describe("Manage: passport admin actions (set password, disable 2FA, send reset email)", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should be able to generate the menu items from seeders.", () => {
    cy.task("exec", ` seeders`);
  });

  it("should sign in as root over the API, to drive setup for the rest of this spec.", () => {
    signin(ROOT_EMAIL, ROOT_PASSWORD).then((response) => {
      expect(response.status).to.equal(200);
      setShared(
        "rootToken",
        (response.body as SigninResponse).data.item.session.token,
      );
    });
  });

  // Establishes the cy.session("root-login", ...) used by every loginAsRoot() call
  // below, here rather than letting the first later call establish it implicitly.
  // cy.session() reloads the whole spec runner the *first* time a given session name
  // is set up (see manage-users.cy.ts's own top comment on this) - which re-executes
  // this file's module-level code, including the `Date.now()` in TARGET_EMAIL/etc.
  // above, silently giving them a *different* value than whatever was already sent to
  // the server. Doing that reload now, before any Date.now()-based value is actually
  // used, means every test below (this file never calls Date.now() again) sees one
  // single, stable value throughout.
  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  // Creates the target account this whole spec operates on, entirely over the (root-
  // authenticated) API - not via the CLI's `auth` command (which withFirebackServer()
  // itself uses for root): `auth` persists whichever identity it just created as the
  // *CLI's own* .env CLI_TOKEN/CLI_WORKSPACE, so a second call for a non-root target
  // account would silently swap every later cy.task("exec", ...) in this spec from
  // root over to that unprivileged account instead. Passport's own generic `create`
  // stores --password as-is, unhashed (PassportCreateAction.go), so a bare passport is
  // created first and then given a real bcrypt-hashed password through
  // SetPassportPasswordAction itself - which doubles as coverage that the action also
  // works as a first-time password set, not only a reset.
  it("creates the target account this whole spec operates on.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/user"),
        headers: asRoot(rootToken),
        body: { firstName: "Cypress", lastName: "Target" },
      }).then((userResponse) => {
        const userId = userResponse.body.data.item.uniqueId;
        expect(userId).to.be.a("string").and.not.be.empty;
        setShared("targetUserId", userId);

        cy.request({
          method: "POST",
          url: ui("/passport"),
          headers: asRoot(rootToken),
          body: {
            type: "email",
            value: TARGET_EMAIL,
            userId,
            confirmed: true,
          },
        }).then((passportResponse) => {
          const passportId = passportResponse.body.data.item.uniqueId;
          expect(passportId).to.be.a("string").and.not.be.empty;
          setShared("targetPassportId", passportId);

          cy.request({
            method: "POST",
            url: ui("/passport/set-password"),
            headers: asRoot(rootToken),
            body: { uniqueId: passportId, password: TARGET_OLD_PASSWORD },
          }).then((setPasswordResponse) => {
            expect(setPasswordResponse.status).to.equal(200);
          });
        });
      });
    });
  });

  it("signing in with the old password works, before anything below changes it.", () => {
    signin(TARGET_EMAIL, TARGET_OLD_PASSWORD).then((response) => {
      expect(response.status).to.equal(200);
    });
  });

  it("navigates to the target user's single screen and finds the passport, with Set password and Send reset email but not Disable 2FA (no totpSecret yet).", () => {
    loginAsRoot();
    sharedValues(["targetUserId"]).then(({ targetUserId }) => {
      cy.visit(ui(`/manage/#/manage/user/${targetUserId}`));
    });
    cy.contains(TARGET_EMAIL).should("exist");
    cy.contains("button", "Set password").should("exist");
    cy.contains("button", "Send reset email").should("exist");
    cy.contains("button", "Disable 2FA").should("not.exist");
  });

  it("sets a new password via the drawer, and it actually takes effect.", () => {
    loginAsRoot();
    sharedValues(["targetUserId"]).then(({ targetUserId }) => {
      cy.visit(ui(`/manage/#/manage/user/${targetUserId}`));
    });

    cy.contains("button", "Set password").click({ force: true });
    cy.contains(".confirm-drawer-container", "Set password").within(() => {
      cy.get("input[type='password']").type(TARGET_NEW_PASSWORD);
      cy.contains("button", "Save").click({ force: true });
    });

    cy.contains(".Toastify__toast--success", "Password updated", {
      timeout: 10000,
    }).should("exist");

    signin(TARGET_EMAIL, TARGET_OLD_PASSWORD).then((response) => {
      expect(response.status).to.not.equal(200);
    });
    signin(TARGET_EMAIL, TARGET_NEW_PASSWORD).then((response) => {
      expect(response.status).to.equal(200);
    });
  });

  it("gives the passport a totpSecret directly over the API (simulating a confirmed 2FA setup), and Disable 2FA now shows up and clears it.", () => {
    sharedValues(["rootToken", "targetPassportId"]).then(
      ({ rootToken, targetPassportId }) => {
        cy.request({
          method: "PATCH",
          url: ui(`/passport/${targetPassportId}`),
          headers: asRoot(rootToken),
          body: { totpSecret: "CYPRESSFAKESECRET", totpConfirmed: true },
        }).then((response) => {
          expect(response.status).to.equal(200);
        });

        loginAsRoot();
        sharedValues(["targetUserId"]).then(({ targetUserId }) => {
          cy.visit(ui(`/manage/#/manage/user/${targetUserId}`));
        });

        cy.contains("button", "Disable 2FA").click({ force: true });
        cy.contains(".confirm-drawer-container button", "Yes").click({
          force: true,
        });

        cy.contains(
          ".Toastify__toast--success",
          "Two-factor authentication disabled",
          { timeout: 10000 },
        ).should("exist");

        cy.request({
          method: "GET",
          url: ui(`/passport/${targetPassportId}`),
          headers: asRoot(rootToken),
        }).then((response) => {
          expect(response.body.data.item.totpSecret).to.equal("");
        });
      },
    );
  });

  it("sends a reset email, and the code it contains genuinely completes a self-service reset (including signing the user back in) on the real reset-password screen.", () => {
    loginAsRoot();
    sharedValues(["targetUserId"]).then(({ targetUserId }) => {
      cy.visit(ui(`/manage/#/manage/user/${targetUserId}`));
    });

    cy.contains("button", "Send reset email").click({ force: true });
    cy.contains(".confirm-drawer-container button", "Yes").click({
      force: true,
    });
    cy.contains(".Toastify__toast--success", "Reset email sent", {
      timeout: 10000,
    }).should("exist");

    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui("/publicAuthentication/browse"),
        headers: asRoot(rootToken),
      }).then((response) => {
        const items = response.body.data.items as Array<{
          passportValue: string;
          otp: string;
        }>;
        const matching = items.filter(
          (item) => item.passportValue === TARGET_EMAIL,
        );
        const otp = matching[matching.length - 1].otp;
        expect(otp).to.be.a("string").and.not.be.empty;

        const resetPassword = `EvenNewer${Date.now()}`;

        cy.visit(
          ui(
            `/selfservice/#/selfservice/reset-password?value=${encodeURIComponent(TARGET_EMAIL)}`,
          ),
        );
        cy.get("#value-input", { timeout: 10000 }).should(
          "have.value",
          TARGET_EMAIL,
        );
        cy.get("#otp-input").type(otp);
        cy.get("#password-input").type(resetPassword);
        cy.get("#submit-form").click({ force: true });

        // useCompleteAuth's onComplete (auth.common.tsx) lands on the self-service
        // app's own default route (VITE_DEFAULT_ROUTE, "/{locale}/passports") once
        // the session is set - not "/dashboard", which is the manage app's own
        // default and only relevant to loginAsRoot() above.
        cy.url({ timeout: 10000 }).should("include", "/selfservice/passports");

        cy.then(() => {
          signin(TARGET_EMAIL, resetPassword).then((signinResponse) => {
            expect(signinResponse.status).to.equal(200);
          });
        });
      });
    });
  });

  endFirebackServer();
});
