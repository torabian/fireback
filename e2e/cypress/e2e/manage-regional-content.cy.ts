import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/manage/regional-content end to end through the real screens - fixes
// verified here (see RegionalContentEditForm.tsx/RegionalContentEntityManager.tsx/
// RegionalContentSingleScreen.tsx and Abac.emi.yml's own comments for the full context):
//   1. The "Key group" select used to render with zero options at all
//      (createQuerySource() was called with no items) - there was no way to actually
//      pick SMS_OTP/EMAIL_OTP from the UI, even though keyGroup is validate:"required" on
//      the backend. Test #2 below is the regression guard: it drives that exact select.
//   2. The "Region" field used to be hardcoded to the literal value "global" and
//      readonly - there was no way to target a specific region from the UI at all (the
//      whole point of this entity), and editing an existing record showed that fake
//      value instead of whatever was actually stored. Test #2/#3 below create and edit
//      real, different region values through the UI.
//   3. Backend: the DB-level uniqueness on (languageId, keyGroup) didn't include region,
//      so two rows for the same language+keyGroup could never coexist even for different
//      regions - contradicting the whole premise of "regional" content (see
//      Abac.emi.yml's region field comment). Test #3 is the regression guard: it creates
//      a second row with the same keyGroup+languageId as #2 but a different region, and
//      expects it to succeed; test #4 confirms the *same* triple is still rejected.
//   4. RegionalContentEntityManager's onFinishUriResolver read response.data?.uniqueId
//      instead of response.data?.item?.uniqueId (the actual response shape) - fixed;
//      not directly observable here since CommonEntityManager prefers real browser
//      history over that resolved URI when reached via in-app navigation anyway (see
//      manage-capabilities.cy.ts's own comment on the same pattern), but the create/edit
//      flows below exercise the same code path.
//   5. RegionalContentSingleScreen was missing "Content"/"Key group" entirely from its
//      field list - arguably the two most important fields on this entity. Test #2 below
//      checks both now render.
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

function visitForm(url: string) {
  cy.visit(url);
  cy.get(".content-section fieldset").should("not.be.disabled");
}

// Selects an option from one of this app's FormSelect fields (react-select/async under
// the hood - see fireback-ui/components/forms/form-select/FormSelect.tsx). Types the
// option's label into the control's own search input, waits for loadOptions' async
// filtering to actually narrow the menu down to it, then clicks that option directly -
// a bare Enter isn't reliable here, it can confirm whatever was highlighted before the
// async filter landed instead.
function selectFormSelectOption(labelText: string, optionText: string) {
  cy.contains(".mb-3", labelText).find(".form-control").click();
  cy.contains(".mb-3", labelText)
    .find("input[role=combobox]")
    .type(optionText, { delay: 50 });
  cy.contains(".react-select-menu-area [role=option]", optionText, {
    timeout: 5000,
  }).click();
}

describe("Manage: Regional Content - menu, keyGroup select, per-region uniqueness, edit, delete", () => {
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

  it("the Regional Content entry exists in root's menu (seeded via abac.Menu, correctly scoped), and the list itself renders.", () => {
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
        const regionalContentEntry = (rootGroup.children || []).find(
          (item: any) => item.uniqueId === "regional_content",
        );
        expect(
          regionalContentEntry,
          '"Regional Content" menu entry present under "Root"',
        ).to.exist;
        expect(regionalContentEntry.href).to.equal("/manage/regional-contents");
        expect(regionalContentEntry.capabilityId).to.equal(
          "root.manage.abac.regional-content.query",
        );
      });
    });

    cy.visit(ui("/manage/#/en/manage/regional-contents"));
    cy.get(".content-section").should("exist");
  });

  const suffix = Date.now();
  const languageId = "en";
  const keyGroupLabel = "SMS one-time password"; // forces the FormRichText's plain
  // <textarea> fallback (see RegionalContentEditForm's forceBasic), avoiding any
  // dependency on TinyMCE actually loading from its CDN for this test.
  const contentUs = `Cypress content US ${suffix}`;
  const contentEu = `Cypress content EU ${suffix}`;
  let idUs = "";
  let idEu = "";

  it('creates a regional content through the UI ("SMS one-time password", region "us"), with keyGroup actually selectable, and content/keyGroup both show on its own single screen.', () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/regional-contents"));
    clickIconAction("New");
    cy.url({ timeout: 10000 }).should("include", "/manage/regional-content/new");
    cy.get(".content-section fieldset").should("not.be.disabled");

    selectFormSelectOption("Key group", keyGroupLabel);
    cy.contains(".mb-3", "Content").find("textarea").type(contentUs);
    fieldInput("Region").clear().type("us");
    fieldInput("Language id").clear().type(languageId);

    cy.intercept("POST", "**/regionalContent").as("createRc");
    submitForm();
    cy.wait("@createRc").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.keyGroup).to.equal("SMS_OTP");
      expect(item.content).to.equal(contentUs);
      expect(item.region).to.equal("us");
      expect(item.languageId).to.equal(languageId);
      idUs = item.uniqueId;
      expect(idUs).to.be.a("string").and.not.be.empty;
    });

    cy.url({ timeout: 10000 }).should("include", "/manage/regional-contents");

    cy.then(() => {
      cy.visit(ui(`/manage/#/en/manage/regional-content/${idUs}`));
      cy.url().should("include", `/manage/regional-content/${idUs}`);
      // Regression guard: content/keyGroup used to be entirely absent from this
      // screen's field list.
      cy.contains(contentUs).should("exist");
      cy.contains("SMS_OTP").should("exist");
      cy.contains("us").should("exist");
      cy.contains(languageId).should("exist");
    });
  });

  it('allows creating a second regional content with the same key group and language, but a different region ("eu") - the per-region uniqueness fix.', () => {
    loginAsRoot();
    visitForm(ui("/manage/#/en/manage/regional-content/new"));

    selectFormSelectOption("Key group", keyGroupLabel);
    cy.contains(".mb-3", "Content").find("textarea").type(contentEu);
    fieldInput("Region").clear().type("eu");
    fieldInput("Language id").clear().type(languageId);

    cy.intercept("POST", "**/regionalContent").as("createRcEu");
    submitForm();
    cy.wait("@createRcEu").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      const item = interception.response?.body?.data?.item;
      expect(item.region).to.equal("eu");
      idEu = item.uniqueId;
      expect(idEu).to.be.a("string").and.not.be.empty;
      expect(idEu).to.not.equal(idUs);
    });
  });

  it("rejects creating a third regional content with the exact same key group, language, and region as an existing one.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/en/manage/regional-content/new"));

    selectFormSelectOption("Key group", keyGroupLabel);
    cy.contains(".mb-3", "Content").find("textarea").type("should not be created");
    fieldInput("Region").clear().type("us"); // same triple as the very first one created above
    fieldInput("Language id").clear().type(languageId);

    cy.intercept("POST", "**/regionalContent").as("rejectedCreate");
    submitForm();
    cy.wait("@rejectedCreate").then((interception) => {
      expect(interception.response?.statusCode).to.not.equal(200);
    });
  });

  it("edits the first regional content's content through the UI, and it persists.", () => {
    loginAsRoot();
    // Reached via the single screen's own "Edit" action (real in-app navigation) rather
    // than a direct visit to the edit URL - CommonEntityManager's onSubmit prefers going
    // back in browser history over its resolved onFinishUriResolver URI, same as
    // manage-capabilities.cy.ts's own edit test.
    cy.visit(ui(`/manage/#/en/manage/regional-content/${idUs}`));
    clickIconAction("Edit");
    cy.url({ timeout: 10000 }).should(
      "include",
      `/manage/regional-content/edit/${idUs}`,
    );
    cy.get(".content-section fieldset").should("not.be.disabled");

    const editedContent = `${contentUs} Edited`;
    cy.contains(".mb-3", "Content")
      .find("textarea")
      .clear()
      .type(editedContent);

    cy.intercept("PATCH", "**/regionalContent/**").as("updateRc");
    submitForm();
    cy.wait("@updateRc").then((interception) => {
      expect(interception.response?.statusCode).to.equal(200);
      expect(interception.response?.body?.data?.item?.content).to.equal(
        editedContent,
      );
    });

    cy.url({ timeout: 10000 }).should(
      "include",
      `/manage/regional-content/${idUs}`,
    );
    cy.contains(editedContent).should("exist");
  });

  it("deletes the second regional content from the list via row selection + the Delete action, and it disappears.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/manage/regional-contents"));
    cy.contains(contentEu).should("exist");

    // The grid is react-data-grid, not a plain <table> - rows are div.rdg-row, each with
    // its own input.rdg-checkbox-input selection checkbox as the first cell.
    cy.contains(".rdg-row", contentEu)
      .find("input.rdg-checkbox-input")
      .click({ force: true });
    clickIconAction("Delete");
    cy.contains(".confirm-drawer-container button", "Yes").click({
      force: true,
    });
    cy.contains(contentEu).should("not.exist");

    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui(`/regionalContent/${idEu}`),
        headers: { authorization: rootToken, "workspace-id": "root" },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.not.equal(200);
      });
    });
  });

  endFirebackServer();
});
