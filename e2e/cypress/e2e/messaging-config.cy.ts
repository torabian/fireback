import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers the manage UI's messaging-config screen (ui/packages/manage/messaging-config -
// see MessagingConfigEditForm.tsx/MessagingConfigSingleScreen.tsx): the form actually
// renders with all five fields, each dropdown can be selected and the save round-trips
// through the single, global MessagingConfig row (see MessagingConfigActions.go), and -
// the whole reason FormSelect grew a "nullable" prop - a previously-selected dropdown
// can be explicitly cleared back to null and that null is what actually gets persisted,
// not just left alone. See GeneralEntityView.tsx's own field.elem === null branch (a
// bold "Not specified") vs field.elem === undefined ("-") for why the single screen -
// not the edit form's own transient state - is what these assertions read back from:
// only a real, persisted `null` renders that text, so it's the strongest proof the value
// actually cleared server-side rather than just looking empty in the browser.
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

interface SingleItemResponse<T> {
  data: { item: T };
}

function setShared(key: string, value: string) {
  return cy.task("setShared", { key, value });
}
function sharedValues<K extends string>(
  keys: K[],
): Cypress.Chainable<Record<K, string>> {
  return cy.task("getSharedState", keys);
}

// Same session-caching login helper messaging-providers.cy.ts uses - cy.session()
// resets in-browser state on first use per spec run, so every it() below re-invokes
// this rather than logging in once at the top.
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

// Selects an option from one of this app's FormSelect fields (react-select/async - see
// ui-core/components/forms/form-select/FormSelect.tsx). Types optionText into the
// control's own search input, waits for loadOptions' async filtering to actually narrow
// the menu down to it, then clicks that option directly.
function selectFormSelectOption(labelText: string, optionText: string) {
  cy.contains(".mb-3", labelText).find(".form-control").click();
  cy.contains(".mb-3", labelText)
    .find("input[role=combobox]")
    .type(optionText, { delay: 50 });
  cy.contains(".react-select-menu-area [role=option]", optionText, {
    timeout: 5000,
  }).click();
}

// Clears a nullable FormSelect field back to null. react-select clears its current
// single selection when Backspace is pressed on its own search input while that input
// is empty - the same interaction a real user hits reaching for the field's "x" - so
// this is what actually exercises FormSelect's `nullable` support end to end, rather
// than depending on react-select's internal (unstyled, hash-classed) clear-indicator
// markup.
function clearFormSelectOption(labelText: string) {
  cy.contains(".mb-3", labelText).find(".form-control").click();
  cy.contains(".mb-3", labelText)
    .find("input[role=combobox]")
    .type("{backspace}")
    // Backspace clears the selection but, since the input still has focus
    // and keeps whatever loadOptions last resolved, react-select leaves its
    // menu open afterwards. That floating menu overlaps whatever field comes
    // next in the DOM, so the next clearFormSelectOption/selectFormSelectOption's
    // own .form-control click lands on this stale menu instead - Escape closes
    // it (without touching the value that was just cleared) the same way a
    // real user tabbing away would.
    .type("{esc}");
}

function submitForm() {
  cy.get(".action-menu").contains("Save").click({ force: true });
}

// fieldValue reads GeneralEntityView's data-test-id="<label>" cell on the single screen
// - the source of truth for whether a value actually persisted (as opposed to whatever
// the edit form happens to be showing client-side).
function fieldValue(label: string) {
  return cy.get(`[data-test-id="${label}"]`);
}

describe("Messaging config - manage UI (select, save, and clear back to null)", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  it("should be able to set up a terminal email provider and gsm provider to pick from.", () => {
    cy.task(
      "exec",
      ` messaging email provider create --type terminal --title checkendpointtests-messaging-config`,
    ).then((content: string) => {
      const id = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      expect(id).to.be.a("string").and.not.be.empty;
      setShared("messagingConfigEmailProviderId", id);
    });

    cy.task(
      "exec",
      ` messaging gsm provider create --type terminal --main-sender-number +19998887772`,
    ).then((content: string) => {
      const id = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      expect(id).to.be.a("string").and.not.be.empty;
      setShared("messagingConfigGsmProviderId", id);
    });
  });

  it("should render the messaging config form with all five fields.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/messaging-config/edit"));
    cy.contains("h2", "Notifications").should("exist");
    for (const label of [
      "General Email Provider",
      "General Gsm Provider",
      "Invite To Workspace Content",
      "Email Otp Content",
      "SMS Otp Content",
    ]) {
      cy.contains(".mb-3", label).should("exist");
    }
  });

  it("should be able to select every field and save, and see it persisted on the single screen.", () => {
    loginAsRoot();

    sharedValues([
      "messagingConfigEmailProviderId",
      "messagingConfigGsmProviderId",
    ]).then(
      ({ messagingConfigEmailProviderId, messagingConfigGsmProviderId }) => {
        visitForm(ui("/manage/#/manage/messaging-config/edit"));

        selectFormSelectOption(
          "General Email Provider",
          messagingConfigEmailProviderId,
        );
        selectFormSelectOption(
          "General Gsm Provider",
          messagingConfigGsmProviderId,
        );
        // These three pick from RegionalContent rows abac seeds by default on every
        // migration (RegionalContentDefaults.go) - "en" always exists out of the box,
        // no CLI setup needed, and keyGroup+languageId+region("any") is unique so this
        // substring can only ever match its one intended row.
        selectFormSelectOption(
          "Invite To Workspace Content",
          "INVITE_TO_WORKSPACE - en",
        );
        selectFormSelectOption("Email Otp Content", "EMAIL_OTP - en");
        selectFormSelectOption("SMS Otp Content", "SMS_OTP - en");

        submitForm();
        cy.wait(1000);

        cy.visit(ui("/manage/#/manage/messaging-config"));
        fieldValue("General Email Provider").should(
          "contain.text",
          messagingConfigEmailProviderId,
        );
        fieldValue("General Gsm Provider").should(
          "contain.text",
          messagingConfigGsmProviderId,
        );
        fieldValue("Invite To Workspace Content")
          .invoke("text")
          .should("not.be.empty")
          .and("not.contain", "Not specified");
        fieldValue("Email Otp Content")
          .invoke("text")
          .should("not.be.empty")
          .and("not.contain", "Not specified");
        fieldValue("SMS Otp Content")
          .invoke("text")
          .should("not.be.empty")
          .and("not.contain", "Not specified");
      },
    );
  });

  it("should be able to clear a selected field back to null, and have that null actually persist.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/messaging-config/edit"));

    // Only "General Email Provider" is cleared here - the other four are left exactly
    // as the previous test saved them, so this also proves clearing one nullable field
    // doesn't clobber its siblings (the same one-field-at-a-time diffing
    // MessagingConfigEntityUpdateFn's IsSet() checks are meant to guarantee).
    clearFormSelectOption("General Email Provider");
    submitForm();
    cy.wait(1000);

    cy.visit(ui("/manage/#/manage/messaging-config"));
    fieldValue("General Email Provider").should(
      "contain.text",
      "Not specified",
    );

    sharedValues(["messagingConfigGsmProviderId"]).then(
      ({ messagingConfigGsmProviderId }) => {
        fieldValue("General Gsm Provider").should(
          "contain.text",
          messagingConfigGsmProviderId,
        );
      },
    );
    fieldValue("Invite To Workspace Content")
      .invoke("text")
      .should("not.contain", "Not specified");
  });

  it("should be able to clear every remaining field back to null in one save.", () => {
    loginAsRoot();
    visitForm(ui("/manage/#/manage/messaging-config/edit"));

    clearFormSelectOption("General Gsm Provider");
    clearFormSelectOption("Invite To Workspace Content");
    clearFormSelectOption("Email Otp Content");
    clearFormSelectOption("SMS Otp Content");
    submitForm();
    cy.wait(10000);

    cy.visit(ui("/manage/#/manage/messaging-config"));
    for (const label of [
      "General Email Provider",
      "General Gsm Provider",
      "Invite To Workspace Content",
      "Email Otp Content",
      "SMS Otp Content",
    ]) {
      fieldValue(label).should("contain.text", "Not specified");
    }
  });

  endFirebackServer();
});
