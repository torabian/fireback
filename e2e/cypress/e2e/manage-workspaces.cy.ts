import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/manage/workspaces end to end through the real screens - there was no
// UI-driven coverage of Workspace before this (see
// userworkspace-workspace-workspaceconfig-workspacerole.cy.ts's own top comment: "None
// of these four have a dedicated manage-UI admin screen, so all are exercised directly
// over HTTP" - stale by the time this was written, the screen already existed):
//   1. The "Workspaces" entry (abac.Menu, synced into app_menu_entities as part of
//      abac's own migration - see modules/abac/Menu.go) exists, correctly scoped.
//   2. A workspace can be created via the UI and shows up in the list.
//   3. It can be deleted again via the list's real selection + Delete flow.
//   4. Trying to delete the root workspace itself - rejected server-side
//      (WorkspaceAwareDeleteAction's root check, WorkspaceActions.go; the actual CRUD
//      guarantee is Go-tested directly in modules/abac/tests/workspace_test.go) - now
//      actually surfaces an error toast instead of silently doing nothing. Before the
//      httpErrorHanlder/ToastContainer fix (hooks/api.ts, apps/core/EssentialApp.tsx),
//      useDatatableFiltering's deleteItems() (CommonDataTable.tsx) had no error
//      handling at all on a rejected delete - not just here, on every CommonArchive
//      screen's delete action - and separately, react-toastify was never actually
//      given a mounted <ToastContainer/> anywhere in the app, so literally no toast of
//      any kind (success or error) ever rendered, root delete or otherwise.
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

// Most action-menu entries (New/Edit/Delete, unlike Save/Cancel) are registered with an
// icon (see CommonArchiveManager/CommonSingleManager's useNewAction/useEditAction, and
// useDatatableFiltering's deleteAction) - ActionMenuItem then renders an icon-only
// `<img title=".." alt="..">` instead of a `<span>{label}</span>`, so the label is never
// part of the element's visible text content and a text-based cy.contains can't find
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

describe("Manage: Workspaces - menu, create via UI, delete via UI", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
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

  it("the Workspaces entry exists in root's menu (seeded via abac.Menu, correctly scoped), and the list itself renders.", () => {
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
        const workspacesEntry = (rootGroup.children || []).find(
          (item: any) => item.uniqueId === "workspaces",
        );
        expect(workspacesEntry, '"Workspaces" menu entry present under "Root"').to
          .exist;
        expect(workspacesEntry.href).to.equal("/manage/workspaces");
        expect(workspacesEntry.capabilityId).to.equal(
          "root.manage.abac.workspace.query",
        );
      });
    });

    cy.visit(ui("/manage/#/en/manage/workspaces"));
    // "root" itself is always a row here (see RepairTheWorkspaces) - proof the list
    // (not just the URL) actually rendered.
    cy.contains("root").should("exist");
  });

  it("shows an error toast (and does not delete anything) when trying to delete the root workspace via the UI.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));

    cy.contains(".rdg-row", "root")
      .find("input.rdg-checkbox-input")
      .click({ force: true });
    clickIconAction("Delete");
    cy.contains(".confirm-drawer-container button", "Yes").click({
      force: true,
    });

    cy.contains(".Toastify__toast--error", "root workspace cannot be deleted", {
      timeout: 10000,
    }).should("exist");

    // Still there, both in the list and over the API.
    cy.contains("root").should("exist");
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui("/workspace/root"),
        headers: { authorization: rootToken, "workspace-id": "root" },
      }).then((response) => {
        expect(response.status).to.equal(200);
      });
    });
  });

  let workspaceUniqueId = "";
  const suffix = Date.now();
  const workspaceName = `CypressWorkspace${suffix}`;

  it("creates a workspace through the UI form, reachable via the list's own \"New\" action, and it shows up in the list.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/workspace/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Workspace name").type(workspaceName);

    // CommonEntityManager's onSubmit goes back to real browser history (this list
    // page, visited just before clicking "New") rather than the resolved
    // single-screen URI - same as manage-users.cy.ts/manage-capabilities.cy.ts's own
    // create steps - so capturing the created uniqueId from the POST response itself.
    cy.intercept("POST", "**/workspace").as("createWorkspace");
    submitForm();
    cy.wait("@createWorkspace").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.name).to.equal(workspaceName);
      workspaceUniqueId = item.uniqueId;
      expect(workspaceUniqueId).to.be.a("string").and.not.be.empty;
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/workspaces");
    cy.contains(workspaceName).should("exist");
  });

  it("deletes the workspace from the list via row selection + the Delete action, and it disappears.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));
    cy.contains(workspaceName).should("exist");

    // The grid is react-data-grid, not a plain <table> - rows are div.rdg-row, each
    // with its own input.rdg-checkbox-input selection checkbox as the first cell.
    cy.contains(".rdg-row", workspaceName)
      .find("input.rdg-checkbox-input")
      .click({ force: true });
    clickIconAction("Delete");
    cy.contains(".confirm-drawer-container button", "Yes").click({
      force: true,
    });
    cy.contains(workspaceName).should("not.exist");

    // Belt-and-suspenders confirmation straight from the API, since the UI's own
    // delete flow has no success/failure feedback to assert on (see this file's top
    // comment).
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui(`/workspace/${workspaceUniqueId}`),
        headers: { authorization: rootToken, "workspace-id": "root" },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.not.equal(200);
      });
    });
  });

  // Selects an option from one of this app's FormSelect fields (react-select/async
  // under the hood) - same helper as manage-regional-content.cy.ts's own
  // selectFormSelectOption, kept local here since spec files in this suite each
  // already redeclare their own copies of shared helpers (loginAsRoot, fieldInput,
  // ...) rather than importing them.
  function selectFormSelectOption(labelText: string, optionText: string) {
    cy.contains(".mb-3", labelText).find(".form-control").click();
    cy.contains(".mb-3", labelText)
      .find("input[role=combobox]")
      .type(optionText, { delay: 50 });
    cy.contains(".react-select-menu-area [role=option]", optionText, {
      timeout: 5000,
    }).click();
  }

  it("rejects creating a workspace with a blank name.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/workspace/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    // Name deliberately left blank.
    cy.intercept("POST", "**/workspace").as("createWorkspaceRejected");
    submitForm();
    cy.wait("@createWorkspaceRejected").then((interception) => {
      expect(interception.response?.statusCode).to.not.equal(200);
    });
    // Still on the form - nothing got created.
    cy.url({ timeout: 10000 }).should("include", "/manage/workspace/new");
  });

  const parentSuffix = Date.now();
  const parentName = `CypressParent${parentSuffix}`;
  const childName = `CypressChild${parentSuffix}`;
  const customUniqueId = `cypress-custom-id-${parentSuffix}`;
  let parentUniqueId = "";

  it("creates a parent workspace with a self-chosen unique id, and it's honored.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/workspace/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Workspace name").type(parentName);
    fieldInput("Unique id").type(customUniqueId);

    cy.intercept("POST", "**/workspace").as("createParentWorkspace");
    submitForm();
    cy.wait("@createParentWorkspace").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.uniqueId).to.equal(customUniqueId);
      parentUniqueId = item.uniqueId;
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/workspaces");
    cy.contains(parentName).should("exist");
  });

  it("editing the parent workspace no longer offers a unique id field at all.", () => {
    loginAsRoot();
    cy.then(() => {
      cy.visit(ui(`/manage/#/en/manage/workspace/edit/${parentUniqueId}`));
    });
    cy.get(".content-section fieldset").should("not.be.disabled");
    cy.contains(".mb-3", "Unique id").should("not.exist");
  });

  it("creates a child workspace, picking the parent workspace via the Parent Record select, and it's saved with the right parentId.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/workspaces"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/workspace/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    fieldInput("Workspace name").type(childName);
    selectFormSelectOption("Parent Record", parentName);

    cy.intercept("POST", "**/workspace").as("createChildWorkspace");
    submitForm();
    cy.wait("@createChildWorkspace").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.parentId).to.equal(parentUniqueId);
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/workspaces");
    cy.contains(childName).should("exist");
  });

  endFirebackServer();
});
