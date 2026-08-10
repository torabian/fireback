import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers 5 of Abac.emi.yml's CRUD entities: Capability, EmailConfirmation,
// NotificationConfig, Passport, PassportMethod. PassportMethod has the most complete
// manage-UI admin screen (see manage/passport-method) and is driven for real below -
// it's also the regression guard for a real bug this pass fixed:
// manage/passport-method/PassportMethodList.tsx wired the *Passport* entity's delete
// hook (usePassportAwareDeleteAction) onto the PassportMethod list instead of its own
// (usePassportMethodAwareDeleteAction) - a copy-paste mistake that meant an admin's
// delete click there tried to delete a Passport row that doesn't exist, silently
// leaving the actual passportMethod row in place. EmailConfirmation/NotificationConfig
// have no admin screen at all, and Capability/Passport are simpler root-only config/
// catalog entities, so those three are exercised directly over HTTP instead, same as
// the other specs in this pass.
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

interface SingleItemResponse<T> {
  data: { item: T };
}
interface ListResponse<T> {
  data: { items: T[] };
}
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

function rootHeaders(token: string) {
  return { authorization: token, "workspace-id": "root" };
}

function loginAs(sessionKey: string, email: string, password: string) {
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session(sessionKey, () => {
    cy.visit(ui("/manage/#/en/welcome"));
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

function visitForm(url: string) {
  cy.visit(url);
  cy.get(".content-section fieldset").should("not.be.disabled");
}

function selectFormSelectOption(labelText: string, optionText: string) {
  cy.contains(".mb-3", labelText).find(".form-control").click();
  cy.contains(".mb-3", labelText)
    .find("input[role=combobox]")
    .type(optionText, { delay: 50 });
  cy.contains(".react-select-menu-area [role=option]", optionText, {
    timeout: 5000,
  }).click();
}

function submitForm() {
  cy.get(".action-menu").contains("Save").click({ force: true });
}

describe("Abac: Capability, EmailConfirmation, NotificationConfig, Passport, PassportMethod", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login, and sign in as root over the API for the HTTP checks below.", () => {
    cy.task("exec", ` passport method create --region global --type email`);

    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: { value: ROOT_EMAIL, password: ROOT_PASSWORD },
    }).then((response: Cypress.Response<SigninResponse>) => {
      expect(response.status).to.equal(200);
      setShared("rootToken", response.body.data.item.session.token);
    });
  });

  // --- PassportMethod: real UI create + browse, plus the delete-hook regression guard ---

  it("should create a passport method through the real admin form and see it in the list.", () => {
    loginAs("root-login", ROOT_EMAIL, ROOT_PASSWORD);

    visitForm(ui("/manage/#/en/manage/passport-method/new"));
    selectFormSelectOption("Type", "Google");
    cy.contains(".mb-3", "Region").find("input").clear().type("global");
    submitForm();
    cy.wait(500);

    cy.visit(ui("/manage/#/en/manage/passport-methods"));
    cy.contains("google", { timeout: 10000 }).should("exist");
  });

  it("deleting a passport method through the admin UI should actually delete the passport method (not silently target the wrong entity).", () => {
    loginAs("root-login", ROOT_EMAIL, ROOT_PASSWORD);

    // A fresh, disposable row ("facebook", not "google") so this doesn't depend on the
    // previous test's row or collide with its type+region duplicate-check.
    visitForm(ui("/manage/#/en/manage/passport-method/new"));
    selectFormSelectOption("Type", "Facebook");
    cy.contains(".mb-3", "Region").find("input").clear().type("global");
    submitForm();
    cy.wait(500);

    cy.visit(ui("/manage/#/en/manage/passport-methods"));
    cy.contains("facebook", { timeout: 10000 }).should("exist");
    // The devexpress grid's cells aren't plain <td>/<tr> here (confirmed by both a
    // parents("tr") and a contains("td", ...) miss against a row visibly on screen), so
    // rather than keep guessing its actual DOM shape, use the checkbox column directly:
    // the header's own "select all" checkbox is first, so the last checkbox belongs to
    // the last data row - "facebook", created most recently of the three rows seeded by
    // this point (email via ensureEmailPassportMethod, google by the previous test).
    cy.get('input[type="checkbox"]').last().click({ force: true });
    // ActionMenuItem (ActionMenu.tsx) renders an icon action as a bare <img alt/title>,
    // not text content - ".action-menu" alone has no "Delete" text to match against
    // (that only happens on form-page actions like "Save", which have no icon set).
    // Two matches are expected here (a duplicate desktop/mobile ActionMenu render -
    // see deviceInformation.tsx's isMobileView elsewhere in this codebase); the first
    // is the one actually visible/interactive at this viewport size.
    cy.get('.action-menu img[alt="Delete"]').first().click({ force: true });
    // This confirm dialog (commonDialogs()/confirmModal in
    // useDatatableFiltering.tsx) isn't the bootstrap ".modal-footer" one other specs'
    // accept flows use (e.g. workspace-invite.cy.ts's invite-accept) - a plain
    // "button, Yes" match covers it.
    cy.contains("button", "Yes", { timeout: 10000 }).click({ force: true });
    cy.wait(1000);

    // Before the fix, this delete called usePassportAwareDeleteAction - POST
    // /passport/delete with the *passportMethod's* uniqueId, which matches no Passport
    // row, so the passportMethod row itself was never actually removed and this
    // assertion would have failed.
    cy.contains("facebook").should("not.exist");
  });

  // --- Capability, EmailConfirmation, NotificationConfig, Passport: direct over HTTP ---

  it("Capability CRUD should work end to end, and reject a non-root create.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/capability"),
        headers: { authorization: rootToken, "workspace-id": "not-root" },
        body: { name: "checkendpointtests non-root capability" },
        failOnStatusCode: false,
      }).then((nonRoot) => {
        expect(nonRoot.status).to.not.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/capability"),
        headers: rootHeaders(rootToken),
        body: { name: "checkendpointtests capability", description: "checkendpointtests" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; name: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/capability/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/capability/${id}`),
          headers: rootHeaders(rootToken),
          body: { name: "checkendpointtests renamed capability" },
        }).then((update) => {
          expect(update.status).to.equal(200);
        });

        cy.request({
          method: "GET",
          url: ui(`/capability/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/capability/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("EmailConfirmation CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      const email = `checkendpointtests-ec-${Date.now()}@example.com`;
      cy.request({
        method: "POST",
        url: ui("/emailConfirmation"),
        headers: rootHeaders(rootToken),
        body: { email, status: "pending", key: "checkendpointtests" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/emailConfirmation/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/emailConfirmation/${id}`),
          headers: rootHeaders(rootToken),
          body: { status: "confirmed" },
        }).then((update: Cypress.Response<SingleItemResponse<{ status: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.status).to.equal("confirmed");
        });

        cy.request({
          method: "POST",
          url: ui("/emailConfirmation/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("NotificationConfig CRUD should work end to end (Create/AwareDelete on a distinct row, Update upserting by workspace).", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/notificationConfig"),
        headers: rootHeaders(rootToken),
        body: { acceptLanguage: `checkendpointtests-${Date.now()}` },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/notificationConfig/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "POST",
          url: ui("/notificationConfig/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });

      const title = `checkendpointtests invite title ${Date.now()}`;
      cy.request({
        method: "PATCH",
        url: ui("/notificationConfig/checkendpointtests-any-placeholder"),
        headers: rootHeaders(rootToken),
        body: { inviteToWorkspaceTitle: title },
      }).then((update: Cypress.Response<SingleItemResponse<{ inviteToWorkspaceTitle: string }>>) => {
        expect(update.status).to.equal(200);
        expect(update.body.data.item.inviteToWorkspaceTitle).to.equal(title);
      });
    });
  });

  it("Passport CRUD should work end to end, and reject a missing value.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/passport"),
        headers: rootHeaders(rootToken),
        body: { type: "email" },
        failOnStatusCode: false,
      }).then((missingValue) => {
        expect(missingValue.status).to.not.equal(200);
      });

      const value = `checkendpointtests-passport-${Date.now()}@example.com`;
      cy.request({
        method: "POST",
        url: ui("/passport"),
        headers: rootHeaders(rootToken),
        body: { type: "email", value },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/passport/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "POST",
          url: ui("/passport/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  endFirebackServer();
});
