import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/manage/capabilities end to end through the real screens:
//   1. A capability can be created via the UI with an admin-chosen uniqueId (not just
//      an auto-generated one) - CapabilityCreateAction (CapabilityActions.go) already
//      accepted a caller-supplied uniqueId on the backend before this change; what was
//      missing was a way to actually type one in from the form, and a friendly
//      "id already taken" check ahead of the raw gorm:"unique" constraint violation.
//   2. Two capabilities created this way both persist their uniqueId/name/description
//      correctly (checked via the list).
//   3. Creating a third with a uniqueId that's already taken is rejected with a
//      field-scoped validation error under the Id field, not a raw database error.
//   4. Editing an existing capability shows its id but does not allow changing it
//      (CapabilityEntityUpdateFn, generated, never reads input.UniqueId at all - this
//      guards the form side of that same guarantee).
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

function loginAsRoot() {
  // Sidebar/content-section layout (and the form's field labels used below) only
  // render at desktop widths - see deviceInformation.tsx's isMobileView.
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session("root-login", () => {
    // Only the "email" passport method is registered (see the "should be able to
    // create the passport method" step below), so /welcome auto-redirects straight to
    // the email-entry screen instead of showing a method-choice screen first.
    cy.visit(ui("/manage/#/en/welcome"));
    cy.get("#value-input", { timeout: 10000 }).type(ROOT_EMAIL);
    cy.get("#submit-form").click({ force: true });
    cy.get("h1").should("have.text", "Enter Password");
    cy.get("#password-input").type(ROOT_PASSWORD);
    cy.get("#submit-form").click({ force: true });
    cy.wait(1000);
    cy.url({ timeout: 10000 }).should("include", "/dashboard");
  });
}

// The visible "Save" action-menu entry (top-right) - triggers the same submit
// CommonEntityManager's own hidden `<button type="submit" class="d-none">` does.
function submitForm() {
  cy.get(".action-menu").contains("Save").click({ force: true });
}

// Most action-menu entries (New/Edit, unlike Save/Cancel) are registered with an icon
// (see CommonArchiveManager/CommonSingleManager's useNewAction/useEditAction) -
// ActionMenuItem then renders an icon-only `<img title=".." alt="..">` instead of a
// `<span>{label}</span>`, so the label is never part of the element's visible text
// content and a text-based cy.contains can't find it. Match on the img's title/alt.
function clickIconAction(label: string) {
  // Both the desktop and mobile action-menu variants (Navbar.tsx/Layout.tsx) render at
  // once (one hidden via CSS depending on viewport) - .first() picks the desktop one,
  // which matches the 1366x900 viewport set in loginAsRoot().
  cy.get(`.action-menu img[title="${label}"]`).first().click({ force: true });
}

function fieldInput(labelText: string) {
  return cy.contains(".mb-3", labelText).find("input");
}

function visitForm(url: string) {
  cy.visit(url);
  cy.get(".content-section fieldset").should("not.be.disabled");
}

describe("Manage: Capabilities - create with a chosen id, id uniqueness, id locked on edit", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  const suffix = Date.now();
  const capOneId = `cypresstest.cap.one.${suffix}`;
  const capOneName = `Cypress Cap One ${suffix}`;
  const capOneDescription = "First capability created by a Cypress test.";
  const capTwoId = `cypresstest.cap.two.${suffix}`;
  const capTwoName = `Cypress Cap Two ${suffix}`;
  const capTwoDescription = "Second capability created by a Cypress test.";

  // Verified via each capability's own single screen (direct visit by its chosen id)
  // rather than the archive list: the capability catalog is seeded with the entire
  // permission tree (hundreds of rows - see the screenshot any manual run of this shows,
  // "root.*", "root.manage.abac.workspace.*", ...), and that grid is react-data-grid,
  // which virtualizes rows - only whatever's scrolled into view actually exists in the
  // DOM, so a freshly created row sorted somewhere in the middle of hundreds of others
  // isn't reliably findable via cy.contains on the list page. The single screen loads
  // just the one row by id, which is both simpler and a more precise check of "id/name/
  // description are stored correctly" than eyeballing a giant grid anyway.
  it("creates two capabilities through the UI, each with its own chosen id, and both persist id/name/description correctly.", () => {
    loginAsRoot();

    // --- Capability #1 ---
    cy.visit(ui("/manage/#/en/manage/capabilities"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/capability/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Id").type(capOneId);
    fieldInput("Name").type(capOneName);
    fieldInput("Description").type(capOneDescription);

    cy.intercept("POST", "**/capability").as("createCapOne");
    submitForm();
    cy.wait("@createCapOne").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.uniqueId).to.equal(capOneId);
      expect(item.name).to.equal(capOneName);
      expect(item.description).to.equal(capOneDescription);
    });

    // CommonEntityManager's onSubmit goes back to real browser history (this list page,
    // visited just before clicking "New") rather than the resolved single-screen URI -
    // same as manage-users.cy.ts's create step - so this is back on the list now.
    cy.url({ timeout: 10000 }).should("include", "/manage/capabilities");

    // --- Capability #2 ---
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/capability/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Id").type(capTwoId);
    fieldInput("Name").type(capTwoName);
    fieldInput("Description").type(capTwoDescription);

    cy.intercept("POST", "**/capability").as("createCapTwo");
    submitForm();
    cy.wait("@createCapTwo").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.uniqueId).to.equal(capTwoId);
      expect(item.name).to.equal(capTwoName);
      expect(item.description).to.equal(capTwoDescription);
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/capabilities");

    // Both capabilities' own single screens confirm id/name/description together, as
    // actually rendered by the UI (not just the raw API response above).
    cy.visit(ui(`/manage/#/en/manage/capability/${capOneId}`));
    cy.url().should("include", `/manage/capability/${capOneId}`);
    cy.contains(capOneName).should("exist");
    cy.contains(capOneDescription).should("exist");

    cy.visit(ui(`/manage/#/en/manage/capability/${capTwoId}`));
    cy.url().should("include", `/manage/capability/${capTwoId}`);
    cy.contains(capTwoName).should("exist");
    cy.contains(capTwoDescription).should("exist");
  });

  it("rejects creating a capability whose id is already taken, with a field error under Id (not a raw database error).", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/en/manage/capability/new"));

    fieldInput("Id").type(capOneId);
    fieldInput("Name").type("Should not be created");

    cy.intercept("POST", "**/capability").as("rejectedCreate");
    submitForm();
    cy.wait("@rejectedCreate").then((interception) => {
      expect(interception.response?.statusCode).to.equal(400);
    });

    cy.contains(".mb-3", "Id")
      .find(".invalid-feedback")
      .should("contain.text", "already used");
  });

  it("shows the id while editing an existing capability, but does not let it be changed.", () => {
    loginAsRoot();
    // Reached via the single screen's own "Edit" action (real in-app navigation, same
    // as an admin clicking through the UI) rather than a direct visit to the edit URL -
    // CommonEntityManager's onSubmit prefers going back in browser history over its
    // resolved onFinishUriResolver URI (see the create test's own comment above), and
    // createEntityNavigation's absolute paths ("/en/capability/...") don't carry the
    // "/manage" prefix ManageRoutes.tsx nests these routes under, so that resolved URI
    // is only actually reachable via history, not a fresh page load.
    cy.visit(ui(`/manage/#/en/manage/capability/${capOneId}`));
    clickIconAction("Edit");
    cy.url({ timeout: 10000 }).should(
      "include",
      `/manage/capability/edit/${capOneId}`,
    );
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Id").should("have.value", capOneId).and("be.disabled");

    const editedName = `${capOneName} Edited`;
    fieldInput("Name").clear().type(editedName);

    cy.intercept("PATCH", "**/capability/**").as("updateCapOne");
    submitForm();
    cy.wait("@updateCapOne").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      // The id survived the edit unchanged while the name did change.
      expect(item.uniqueId).to.equal(capOneId);
      expect(item.name).to.equal(editedName);
      expect(item.description).to.equal(capOneDescription);
    });

    // Back on the single screen for the same (unchanged) id.
    cy.url({ timeout: 10000 }).should("include", `/manage/capability/${capOneId}`);
    cy.contains(editedName).should("exist");
  });

  endFirebackServer();
});
