import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers Abac.emi.yml's core session lifecycle actions: ClassicSignup, CheckClassicPassport,
// ClassicSignin, Whoami, ChangePassword, Signout. Two real UI flows exist here (signin and
// change-password - see selfservice/ClassicSigninPassword.screen.tsx and
// selfservice/ChangePassword.screen.tsx) and are driven for real; signup/whoami/validation
// edge cases are exercised directly over HTTP, same as the interfacetools specs.
//
// Note on Signout: fireback-ui/auth/AuthenticationProvider.tsx's signout() and
// fireback-ui/hooks/authContext.tsx's signout() (the manage app's own) both only clear
// local session state - neither actually calls POST /passport/signout. So clicking the
// UI's "Sign out" button proves the client-side logout works, but doesn't exercise the
// backend action at all - that's covered separately below over HTTP.
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

interface SingleItemResponse<T> {
  data: { item: T };
}
interface SignupResponse {
  data: {
    item: {
      session: {
        token: string;
        userWorkspaces: { workspaceId: string }[];
      };
    };
  };
}

function setShared(key: string, value: string) {
  return cy.task("setShared", { key, value });
}
function sharedValues<K extends string>(
  keys: K[],
): Cypress.Chainable<Record<K, string>> {
  return cy.task("getSharedState", keys);
}

function loginAs(sessionKey: string, email: string, password: string) {
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session(sessionKey, () => {
    cy.visit(ui("/manage/#/welcome"));
    cy.get("#value-input", { timeout: 10000 }).type(email);
    cy.get("#submit-form").click({ force: true });
    cy.get("h1", { timeout: 10000 }).should("have.text", "Enter Password");
    cy.get("#password-input").type(password);
    cy.get("#submit-form").click({ force: true });
    cy.wait(1000);
    cy.get("body").then(($b) => {
      if ($b.text().includes("Select workspace")) {
        cy.contains("button", "Root Access").click({ force: true });
      }
    });
    cy.url({ timeout: 10000 }).should("include", "/dashboard");
  });
}

describe("Abac: core session lifecycle (signup, check-passport, signin, whoami, change-password, signout)", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login/signup.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  // --- ClassicSignup + CheckClassicPassport, over HTTP (no fixed-workspaceType UI
  // screen exists in this embed to drive a custom workspaceTypeId through) ---

  it("CheckClassicPassport should offer create-with-password for a brand-new email.", () => {
    const value = `checkendpointtests-new-${Date.now()}@example.com`;
    cy.request({
      method: "POST",
      url: ui("/workspace/passport/check"),
      body: { value },
    }).then(
      (response: Cypress.Response<SingleItemResponse<{ next: string[] }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.next).to.include("create-with-password");
      },
    );
  });

  it("ClassicSignup should reject a missing workspaceTypeId, and succeed with one.", () => {
    cy.request({
      method: "POST",
      url: ui("/passports/signup/classic"),
      body: {
        value: `checkendpointtests-notype-${Date.now()}@example.com`,
        type: "email",
        password: "checkendpointtests-pass-123",
        firstName: "Checkendpointtests",
        lastName: "NoType",
      },
      failOnStatusCode: false,
    }).then((noType) => {
      expect(noType.status).to.not.equal(200);
    });

    cy.task(
      "exec",
      ` role create --name "checkendpointtests core-session role" --capabilities-list-id '["root.abac.email-confirmation.query"]'`,
    ).then((content: string) => {
      const roleId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      cy.task(
        "exec",
        ` ws type c  --title "checkendpointtests core-session type" --slug /checkendpointtests-core-session --role-id ${roleId}`,
      ).then((wtContent: string) => {
        const workspaceTypeId = (
          JSON.parse(wtContent) as SingleItemResponse<{ uniqueId: string }>
        ).data.item.uniqueId;

        const email = `checkendpointtests-signup-${Date.now()}@example.com`;
        setShared("sessionEmail", email);

        cy.request({
          method: "POST",
          url: ui("/passports/signup/classic"),
          body: {
            value: email,
            type: "email",
            password: "checkendpointtests-pass-123",
            firstName: "Checkendpointtests",
            lastName: "Session",
            workspaceTypeId,
          },
        }).then((response: Cypress.Response<SignupResponse>) => {
          expect(response.status).to.equal(200);
          setShared("sessionToken", response.body.data.item.session.token);
          expect(
            response.body.data.item.session.userWorkspaces,
          ).to.have.length.above(0);
        });
      });
    });
  });

  it("CheckClassicPassport should offer signin-with-password for the just-signed-up email.", () => {
    sharedValues(["sessionEmail"]).then(({ sessionEmail }) => {
      cy.request({
        method: "POST",
        url: ui("/workspace/passport/check"),
        body: { value: sessionEmail },
      }).then(
        (
          response: Cypress.Response<SingleItemResponse<{ next: string[] }>>,
        ) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.next).to.include(
            "signin-with-password",
          );
        },
      );
    });
  });

  // --- ClassicSignin, real UI flow (the same /welcome -> password screen every other
  // spec's loginAs drives), then Signout, real UI flow ---

  it("should log into the manage UI with the signed-up account, then sign out.", () => {
    sharedValues(["sessionEmail"]).then(({ sessionEmail }) => {
      loginAs(
        "core-session-login",
        sessionEmail,
        "checkendpointtests-pass-123",
      );

      cy.visit(ui("/manage/#/dashboard"));
      cy.url().should("include", "/dashboard");

      // Signout button - see selfservice/UserPassports.screen.tsx (root-level manage
      // navbar also renders a sign-out control, but this screen's is a plain, stable
      // button to target). Only clears local session state (see file header comment) -
      // proven here by the redirect back to a public/welcome screen.
      cy.visit(ui("/manage/#/selfservice/passports"));
      cy.contains("button", "Sign").click({ force: true });
      cy.url({ timeout: 10000 }).should("not.include", "/dashboard");
    });
  });

  // --- ChangePassword, real UI flow ---

  it("should log back in, change the password through the UI, and the new password should work.", () => {
    sharedValues(["sessionEmail"]).then(({ sessionEmail }) => {
      loginAs(
        "core-session-login-2",
        sessionEmail,
        "checkendpointtests-pass-123",
      );

      cy.visit(ui("/manage/#/selfservice/passports"));
      cy.contains("button", "Change", { timeout: 10000 }).click({
        force: true,
      });

      cy.url({ timeout: 10000 }).should("include", "/change-password/");
      cy.get("#password-input", { timeout: 10000 }).type(
        "checkendpointtests-new-pass-789",
      );
      cy.get("#password-input-2").type("checkendpointtests-new-pass-789");
      cy.get("#submit-form").should("not.be.disabled").click({ force: true });
      cy.wait(1000);
    });
  });

  it("the old password should now be rejected and the new one should sign in.", () => {
    sharedValues(["sessionEmail"]).then(({ sessionEmail }) => {
      cy.request({
        method: "POST",
        url: ui("/passports/signin/classic"),
        body: { value: sessionEmail, password: "checkendpointtests-pass-123" },
        failOnStatusCode: false,
      }).then((oldPassResp) => {
        expect(oldPassResp.status).to.not.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/passports/signin/classic"),
        body: {
          value: sessionEmail,
          password: "checkendpointtests-new-pass-789",
        },
      }).then((newPassResp: Cypress.Response<SignupResponse>) => {
        expect(newPassResp.status).to.equal(200);
        setShared("sessionToken", newPassResp.body.data.item.session.token);
      });
    });
  });

  // --- Whoami, Signout (backend action), and the remaining validation edge cases, all
  // direct over HTTP ---

  it("Whoami should return the session's own userId, and reject an unauthenticated call.", () => {
    sharedValues(["sessionToken"]).then(({ sessionToken }) => {
      cy.request({
        method: "GET",
        url: ui("/whoami"),
        headers: { authorization: sessionToken },
      }).then(
        (
          response: Cypress.Response<SingleItemResponse<{ userId: string }>>,
        ) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.userId).to.be.a("string").and.not.be
            .empty;
        },
      );

      cy.request({
        method: "GET",
        url: ui("/whoami"),
        failOnStatusCode: false,
      }).then((noAuth) => {
        expect(noAuth.status).to.not.equal(200);
      });
    });
  });

  it("Signout (the backend action, not just the UI's local logout) should succeed with a valid token.", () => {
    sharedValues(["sessionToken"]).then(({ sessionToken }) => {
      cy.request({
        method: "POST",
        url: ui("/passport/signout"),
        headers: { authorization: sessionToken },
        body: {},
      }).then((response: Cypress.Response<{ okay: boolean }>) => {
        expect(response.status).to.equal(200);
        expect(response.body.okay).to.equal(true);
      });
    });
  });

  it("ChangePassword should reject a too-short password and an unauthenticated call.", () => {
    sharedValues(["sessionToken"]).then(({ sessionToken }) => {
      cy.request({
        method: "POST",
        url: ui("/passport/change-password"),
        headers: { authorization: sessionToken },
        body: { password: "abc", uniqueId: "placeholder" },
        failOnStatusCode: false,
      }).then((tooShort) => {
        expect(tooShort.status).to.not.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/passport/change-password"),
        body: { password: "some-new-password-123", uniqueId: "placeholder" },
        failOnStatusCode: false,
      }).then((noAuth) => {
        expect(noAuth.status).to.not.equal(200);
      });
    });
  });

  it("ClassicSignin should reject a wrong password and an unknown passport value.", () => {
    sharedValues(["sessionEmail"]).then(({ sessionEmail }) => {
      cy.request({
        method: "POST",
        url: ui("/passports/signin/classic"),
        body: { value: sessionEmail, password: "definitely-wrong" },
        failOnStatusCode: false,
      }).then((wrongPass) => {
        expect(wrongPass.status).to.not.equal(200);
      });
    });

    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: {
        value: `checkendpointtests-nobody-${Date.now()}@example.com`,
        password: "whatever",
      },
      failOnStatusCode: false,
    }).then((unknown) => {
      expect(unknown.status).to.not.equal(200);
    });
  });

  endFirebackServer();
});
