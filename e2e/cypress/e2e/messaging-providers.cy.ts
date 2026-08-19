import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/abac/messaging's admin screens (Gsm Provider / Email Provider / Email
// Sender "create" forms, under /manage/#/manage/<entity>/new - see
// modules/manage/ManageRoutes.tsx): required-field validation actually reaching the
// user (see CommonEntityManager.tsx's response.error handling - a bug that, until
// fixed, meant field errors never rendered anywhere in the app, only a raw untranslated
// error code) and successful creates actually showing up in the list (see
// GsmProviderCreateAction/EmailProviderCreateAction/EmailSenderCreateAction - a bug
// that, until fixed, meant every created row's workspaceId/userId stayed null, making
// it invisible to the workspace-scoped browse/list these screens use).
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

// Logs in once per spec run and caches the session (cookies/localStorage) via
// cy.session(), restoring it instantly on every later call instead of re-running the
// UI login flow - needed because Cypress resets browser state between `it()`s by
// default (testIsolation), and this spec's many small, single-purpose `it()`s (one per
// form/assertion, matching this repo's established convention - see
// workspace-role-capability-scoping.cy.ts) each need to already be logged in.
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
    // withFirebackServer()'s `auth --in-root=true --workspace-type-id root` creates a
    // personal "workspace" (from the "root" workspace *type*) in addition to elevating
    // into the literal "root" workspace, so root actually lands on a workspace picker
    // here rather than skipping straight to the dashboard. The messaging entities this
    // spec creates need workspace-scoped capabilities only the "root" workspace grants
    // (PERM_ROOT_GSM_PROVIDER_CREATE and friends, under "root.manage.*") - the other
    // workspace's seeded "Admin" role only covers "root.modules.*" - so it must be this
    // one, not just whichever picker option happens to be first.
    cy.get("body").then(($b) => {
      if ($b.text().includes("Select workspace")) {
        // Each option's button is labelled "Select (<role name>)" - root's own
        // membership in the literal "root" workspace always carries the seeded
        // "Root Access" role (see AbacModule's root bootstrap), so matching on that
        // is more specific/stable than matching workspace order or the bare word
        // "root" (which "workspace"'s own title/role text could also contain).
        cy.contains("button", "Root Access").click({ force: true });
      }
    });
    cy.url({ timeout: 10000 }).should("include", "/dashboard");
  });
}

// Visits a CommonEntityManager-based form page and waits for its fieldset to become
// interactive. It's disabled (see CommonEntityManager.tsx's `formWorking`) for a brief
// moment after every visit while its getSingleHook query resolves (fired even on a
// "new" form with no uniqueId yet - it 404s harmlessly) - typing into a field before
// that settles fails outright rather than retrying, since Cypress treats a disabled
// target as an intentional block rather than a transient state.
function visitForm(url: string) {
  cy.visit(url);
  cy.get(".content-section fieldset").should("not.be.disabled");
}

// Selects an option from one of this app's FormSelect fields (react-select/async
// under the hood - see fireback-ui/components/forms/form-select/FormSelect.tsx).
// Types the option's label into the control's own search input, waits for loadOptions'
// async filtering (see FormSelect.tsx's promiseOptions) to actually narrow the menu
// down to it, then clicks that option directly - a bare Enter isn't reliable here, it
// can confirm whatever was highlighted before the async filter landed instead (e.g.
// the list's first entry "SendGrid" while typing a same-prefix "SMTP").
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
  // The visible "Save" action-menu entry (top-right) is what a real user actually
  // clicks - it triggers the same submit CommonEntityManager's own hidden
  // `<button type="submit" class="d-none">`/Ctrl+S shortcut does.
  cy.get(".action-menu").contains("Save").click({ force: true });
}

describe("Messaging providers (Gsm/Email Provider/Sender) - create form validation and success", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  // --- Gsm Provider ---

  it("Gsm Provider: submitting the new-provider form empty shows required-field errors for mainSenderNumber and type.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/gsm-provider/new"));
    submitForm();
    cy.contains(".mb-3", "Main sender number")
      .find(".invalid-feedback")
      .should("have.text", "This field is required");
    cy.contains(".mb-3", "Type")
      .find(".invalid-feedback")
      .should("have.text", "This field is required");
  });

  it("Gsm Provider: filling required fields and submitting succeeds, and the new row appears in the list.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/gsm-provider/new"));
    cy.contains(".mb-3", "Main sender number")
      .find("input")
      .type("+19998887771");
    selectFormSelectOption("Type", "Terminal");
    submitForm();
    cy.wait(1000);

    cy.visit(ui("/manage/#/manage/gsm-providers"));
    cy.contains("+19998887771").should("exist");
  });

  // --- Email Provider ---

  it("Email Provider: submitting the new-provider form empty shows a required-field error for its Service Type.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/email-provider/new"));
    submitForm();
    cy.contains(".mb-3", "Service Type")
      .find(".invalid-feedback")
      .should("have.text", "This field is required");
  });

  it("Email Provider: switching the Service Type swaps in the right dynamic config fields.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/email-provider/new"));
    selectFormSelectOption("Service Type", "SMTP");
    cy.contains(".mb-3", "Host").should("exist");
    cy.contains(".mb-3", "Port").should("exist");
    cy.contains(".mb-3", "Username").should("exist");
    cy.contains(".mb-3", "Password").should("exist");
    // Switching again replaces the config fields rather than accumulating them.
    selectFormSelectOption("Service Type", "SendGrid");
    cy.contains(".mb-3", "API Key").should("exist");
    cy.contains(".mb-3", "Host").should("not.exist");
  });

  it("Email Provider: filling the required Service Type and submitting succeeds, and the new row appears in the list.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/email-provider/new"));
    cy.contains(".mb-3", "Title")
      .find("input")
      .type("checkendpointtests provider");
    selectFormSelectOption("Service Type", "Terminal (Debug)");
    submitForm();
    cy.wait(1000);

    cy.visit(ui("/manage/#/manage/email-providers"));
    cy.contains("checkendpointtests provider").should("exist");
  });

  // --- Email Sender ---

  it("Email Sender: submitting the new-sender form empty shows required-field errors for every field.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/email-sender/new"));
    submitForm();
    for (const label of [
      "From email address",
      "From name",
      "Nick name",
      "Reply to",
    ]) {
      cy.contains(".mb-3", label)
        .find(".invalid-feedback")
        .should("have.text", "This field is required");
    }
  });

  it("Email Sender: filling required fields and submitting succeeds, and the new row appears in the list.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/email-sender/new"));
    cy.contains(".mb-3", "From email address")
      .find("input")
      .type("checkendpointtests@example.com");
    cy.contains(".mb-3", "From name")
      .find("input")
      .type("Checkendpointtests Sender");
    cy.contains(".mb-3", "Nick name").find("input").type("checkendpointtests");
    cy.contains(".mb-3", "Reply to")
      .find("input")
      .type("checkendpointtests@example.com");
    submitForm();
    cy.wait(1000);

    cy.visit(ui("/manage/#/manage/email-senders"));
    cy.contains("checkendpointtests@example.com").should("exist");
  });

  endFirebackServer();
});
