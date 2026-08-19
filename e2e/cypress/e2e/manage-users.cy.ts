import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers the "Users" section of the manage UI (modules/manage/users) end to end, driven
// through the real screens rather than raw HTTP - unlike
// regionalcontent-role-token-user-userprofile.cy.ts, which already covers User's CRUD
// actions over HTTP directly:
//   1. The "Users" entry (seeded by
//      modules/abac/interfacetools/seeders/AppMenu/fireback-manage-menu.yml) exists in
//      root's menu, correctly scoped. See the big comment further down for why this
//      doesn't also click it from the rendered sidebar - a real, separate, pre-existing
//      bug found while writing this spec.
//   2. The list renders.
//   3. A user can be created via the "New" form (UserEditForm.tsx), including its
//      primaryAddress sub-fields and the phoneNumber/jobTitle/company/bio fields added
//      alongside them.
//   4. That user can be edited afterwards - specifically its address fields - which is
//      a regression guard for a real bug: UserDto/UserDto.PrimaryAddress use genuine ES
//      private (#) class fields, and CommonEntityManager.tsx used to hand Formik the
//      nested primaryAddress value as still that raw class instance (only the top level
//      went through toJSON()). Formik/lodash's setFieldValue clones every object along a
//      dotted path (e.g. "primaryAddress.city") to update it - cloning a private-fielded
//      class instance that way throws ("Cannot read/write private member ... from an
//      object whose class did not declare it"), which made the address fields
//      effectively uneditable. Fixed by deep-plainifying the fetched item
//      (JSON.parse(JSON.stringify(...))) before handing it to Formik.
//   5. The user can be removed again via the list's real selection + Delete flow.
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

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
  // Sidebar/content-section layout (and the form's field labels used below) only
  // render at desktop widths - see deviceInformation.tsx's isMobileView.
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session("root-login", () => {
    // Only the "email" passport method is registered (see the "should be able to
    // create the passport method" step below), so /welcome auto-redirects straight to
    // the email-entry screen instead of showing a method-choice screen first.
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

// The visible "Save" action-menu entry (top-right) - triggers the same submit
// CommonEntityManager's own hidden `<button type="submit" class="d-none">` does.
function submitForm() {
  cy.get(".action-menu").contains("Save").click({ force: true });
}

// Most action-menu entries (New/Edit/Delete, unlike Save/Cancel) are registered with an
// icon (see CommonArchiveManager/CommonSingleManager's useNewAction/useEditAction, and
// useDatatableFiltering's deleteAction) - ActionMenuItem then renders an icon-only
// `<img title=".." alt="..">` instead of a `<span>{label}</span>`, so the label is never
// part of the element's visible text content and a text-based `cy.contains` can't find
// it. Match on the img's title/alt instead.
function clickIconAction(label: string) {
  // Both the desktop and mobile action-menu variants (Navbar.tsx/Layout.tsx) render at
  // once (one hidden via CSS depending on viewport) - .first() picks the desktop one,
  // which matches the 1366x900 viewport set in loginAsRoot().
  cy.get(`.action-menu img[title="${label}"]`).first().click({ force: true });
}

function fieldInput(labelText: string) {
  return cy.contains(".mb-3", labelText).find("input");
}

describe("Manage: Users - menu visibility, list, create, edit (incl. address), delete", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  // The sidebar's "Users" entry isn't hardcoded - it's a seeded AppMenu row (see
  // modules/abac/interfacetools/seeders/AppMenu/fireback-manage-menu.yml, rendered via
  // useRemoteMenuResolver/cte-app-menus - fireback-ui/hooks/useRemoteMenuResolver.tsx).
  // withFirebackServer() only creates the root account, it doesn't import any seed
  // data, so without this step the appMenu table is empty and the sidebar shows
  // nothing but the workspace switcher (same "seeders" step login.cy.js's welcome-page
  // flow already relies on).
  it("should be able to generate the menu items from seeders.", () => {
    cy.task("exec", ` seeders`).then((content) => {
      const successLines = content
        .split("\n")
        .filter((line) => line.startsWith("Success"));
      expect(successLines.length).to.be.greaterThan(0);
      for (const line of successLines) {
        expect(line).to.contain("Failure 0");
      }
    });
  });

  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  it("should sign in as root over the API too, to check the seeded menu directly below.", () => {
    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: { value: ROOT_EMAIL, password: ROOT_PASSWORD },
    }).then((response: Cypress.Response<SigninResponse>) => {
      expect(response.status).to.equal(200);
      setShared("rootToken", response.body.data.item.session.token);
    });
  });

  // NOTE ON THE SIDEBAR ITSELF: this doesn't click through the "Users" sidebar entry to
  // reach the list, even though that's the literal ask ("I have a menu item called
  // Users... make sure the list is visible"). Root here ends up on a *single* workspace
  // ("root") straight after login, which resolveSelectedWorkspace/AuthenticationProvider.
  // tsx auto-selects via `session.workspaces[0]` - and that fallback path has no roleId
  // to give (mapRawSessionToAuthenticationSession builds `workspaces` directly off the
  // signin response's userWorkspaces, which carries no role info at all - only
  // WorkspaceRoleEntity does, and nothing cross-references it there). Every
  // capability-gated AppMenu row (this seed file's entire "Root" group, "Users" included
  // - see useRemoteMenuResolver.tsx's userMeetsAccess2, which requires a matching roleId)
  // is therefore invisible on a plain page load for *any* single-workspace user, root
  // included - confirmed directly against a live server: /cte-app-menus and /urw/query
  // both come back with the right data (checked below), but the sidebar renders nothing
  // under "Root" regardless. This reproduces on every existing manage-UI entity, not just
  // Users, and every other admin-screen spec in this suite (see
  // messaging-providers.cy.ts's visitForm/its own comments) already routes around it by
  // navigating straight to the screen's URL instead of clicking the sidebar - the same
  // thing this spec's other its() do. Flagged to the user rather than fixed here, since
  // it's a pre-existing, unrelated session/authorization bug, not part of what was asked.
  it("the Users entry exists in root's menu (seeded, correctly scoped), and the users list itself renders.", () => {
    loginAsRoot();

    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui("/cte-app-menus"),
        headers: { authorization: rootToken, "workspace-id": "root" },
      }).then((response) => {
        const rootGroup = response.body.data.items.find(
          (item: any) => item.uniqueId === "root-actions",
        );
        expect(rootGroup, '"Root" menu group present').to.exist;
        const usersEntry = (rootGroup.children || []).find(
          (item: any) => item.uniqueId === "users",
        );
        expect(usersEntry, '"Users" menu entry present under "Root"').to.exist;
        expect(usersEntry.href).to.equal("/manage/users");
        expect(usersEntry.capabilityId).to.equal("root.manage.abac.user.query");
      });
    });

    cy.visit(ui("/manage/#/manage/users"));
    // The list's own rendered grid - proof the datatable (not just the URL) loaded, and
    // it already has at least the root account itself in it (root's own UserEntity row,
    // created by withFirebackServer()'s `auth --in-root=true ... --first-name testagent
    // --last-name testagent`).
    cy.contains("testagent").should("exist");
  });

  let userUniqueId = "";
  const suffix = Date.now();
  const lastName = `CypressUser${suffix}`;
  const editedLastName = `CypressUserEdited${suffix}`;

  it('creates a user through the UI form, filling in the address and the new profile fields, reachable via the list\'s own "New" action.', () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/manage/users"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/user/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("First name").type("Cypress");
    fieldInput("Last name").type(lastName);
    fieldInput("Address Line 1").type("742 Evergreen Terrace");
    fieldInput("Address Line 2").type("Apt 4B");
    fieldInput("City name").type("Springfield");
    fieldInput("Phone Number").type("+15551234567");
    fieldInput("Job Title").type("QA Engineer");
    fieldInput("Company").type("Acme Corp");
    fieldInput("Bio").type("Created by a Cypress test.");

    // UserEntityManager's onFinishUriResolver would send a successful create to the
    // single screen at /manage/user/<uniqueId> - but CommonEntityManager's onSubmit
    // calls router.goBackOrDefault(...) instead, which prefers real browser history
    // (this list page, just visited before clicking "New") over that resolved URI, so
    // submitting actually lands back on the list rather than the single screen.
    // Capturing the created uniqueId straight from the POST /user response instead of
    // relying on either destination.
    cy.intercept("POST", "**/user").as("createUser");
    submitForm();
    cy.wait("@createUser").then((interception) => {
      userUniqueId = interception.response?.body?.data?.item?.uniqueId;
      expect(userUniqueId).to.be.a("string").and.not.be.empty;
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/users");
    cy.contains(lastName).should("exist");

    // userUniqueId is only actually assigned once the intercept's .then() above runs -
    // building this visit's URL eagerly here (e.g. a template literal in the very next
    // statement) would still see it as "", since that .then() is itself queued rather
    // than awaited. Read it inside a .then() too, so it runs in order after the intercept
    // callback has set it.
    cy.then(() => {
      cy.visit(ui(`/manage/#/manage/user/${userUniqueId}`));
    });
    cy.get('[data-test-id="First name"]').should("contain.text", "Cypress");
    cy.get('[data-test-id="Last name"]').should("contain.text", lastName);
    cy.get('[data-test-id="Address Line 1"]').should(
      "contain.text",
      "742 Evergreen Terrace",
    );
    cy.get('[data-test-id="Address Line 2"]').should("contain.text", "Apt 4B");
    cy.get('[data-test-id="City name"]').should("contain.text", "Springfield");
    cy.get('[data-test-id="Phone Number"]').should(
      "contain.text",
      "+15551234567",
    );
    cy.get('[data-test-id="Job Title"]').should("contain.text", "QA Engineer");
    cy.get('[data-test-id="Company"]').should("contain.text", "Acme Corp");
    cy.get('[data-test-id="Bio"]').should(
      "contain.text",
      "Created by a Cypress test.",
    );
  });

  it("the new user shows up in the list.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/manage/users"));
    cy.contains(lastName).should("exist");
  });

  it("edits the user - including its address fields - via the single screen's Edit action, and the change persists. Regression guard for the primaryAddress private-field clone bug.", () => {
    loginAsRoot();
    cy.visit(ui(`/manage/#/manage/user/${userUniqueId}`));
    clickIconAction("Edit");
    cy.url({ timeout: 10000 }).should(
      "include",
      `/manage/user/edit/${userUniqueId}`,
    );
    cy.get(".content-section fieldset").should("not.be.disabled");

    // The form should already be pre-filled with what was just created - if this ever
    // regresses to the private-field clone bug, editing any single field below silently
    // wipes the sibling address fields instead of updating just the one changed.
    fieldInput("Last name").should("have.value", lastName);
    fieldInput("Address Line 1").should("have.value", "742 Evergreen Terrace");
    fieldInput("City name").should("have.value", "Springfield");

    fieldInput("Last name").clear().type(editedLastName);
    fieldInput("Address Line 1").clear().type("221B Baker Street");
    fieldInput("City name").clear().type("London");
    fieldInput("Job Title").clear().type("Senior QA Engineer");

    // addressLine2/phoneNumber/company/bio are deliberately left untouched here - if the
    // private-field clone bug ever came back, editing only city/addressLine1/lastName
    // above would already have silently blanked these out from Formik's state.
    submitForm();

    cy.url({ timeout: 10000 }).should(
      "include",
      `/manage/user/${userUniqueId}`,
    );
    cy.get('[data-test-id="Last name"]').should("contain.text", editedLastName);
    cy.get('[data-test-id="Address Line 1"]').should(
      "contain.text",
      "221B Baker Street",
    );
    cy.get('[data-test-id="City name"]').should("contain.text", "London");
    cy.get('[data-test-id="Job Title"]').should(
      "contain.text",
      "Senior QA Engineer",
    );
    // Untouched fields must have survived the edit.
    cy.get('[data-test-id="Address Line 2"]').should("contain.text", "Apt 4B");
    cy.get('[data-test-id="Phone Number"]').should(
      "contain.text",
      "+15551234567",
    );
    cy.get('[data-test-id="Company"]').should("contain.text", "Acme Corp");
  });

  it("the list reflects the edited name and address.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/manage/users"));
    cy.contains(editedLastName).should("exist");
    cy.contains(lastName).should("not.exist");
    cy.contains("221B Baker Street").should("exist");
  });

  it("deletes the user from the list via row selection + the Delete action, and it disappears.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/manage/users"));
    // The grid is react-data-grid, not a plain <table> - rows are div.rdg-row, each
    // with its own div.rdg-checkbox-input selection checkbox as the first cell.
    cy.contains(".rdg-row", editedLastName)
      .find("input.rdg-checkbox-input")
      .click({ force: true });
    clickIconAction("Delete");
    cy.contains(".confirm-drawer-container button", "Yes").click({
      force: true,
    });
    cy.contains(editedLastName).should("not.exist");
  });

  endFirebackServer();
});
