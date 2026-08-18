import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers ui/packages/ui-core/components/data-grid-list/DataGridList.tsx - the planned
// CommonListManager replacement - through its own demo screen
// (ui/src/apps/projectname/demo/data-grid-list/DataGridListDemo.tsx, registered at
// demo/data-grid-list in ApplicationRoutes.tsx). The demo's three examples are entirely
// client-side (a fake, deterministic in-memory dataset via useFakeCursorListQuery - see
// that file), so this suite is really exercising DataGridList/its hooks
// (useCursor/useColumnFilters/useInfiniteScroll/useRowSelectionDelete) against real
// browser behavior, not any backend endpoint:
//   1. Example 1 ("every emi field type as a column") renders every column type -
//      string/int/int64/float64/bool/complex(date)/enum/slice/one/collection/object/any.
//   2. Its column filter narrows what's actually shown (JsonLogic "contains", debounced).
//   3. Infinite scroll loads more rows, both page-scrolled (this app's real scroll
//      container is an inner ".content-container" div, not the window - see
//      useInfiniteScroll.ts's own doc comment on why that distinction matters) and boxed
//      inside its own fixed-height div (Example 2).
//   4. Row selection + delete (Example 3), ported from CommonListManager/
//      useDatatableFiltering, still works the same way (select -> action-menu "Delete" ->
//      confirm -> row gone).
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

function loginAsRoot() {
  // Sidebar/content-section layout only renders at desktop widths.
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session("root-login", () => {
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

// Both the desktop and mobile action-menu variants render at once (one hidden via CSS
// depending on viewport) - .first() picks the desktop one, matching loginAsRoot()'s
// 1366x900 viewport. Icon-only actions (Delete, New, ...) render as
// `<img title=".." alt="..">`, never plain text - see manage-workspaces.cy.ts's own
// clickIconAction, copied here verbatim.
function clickIconAction(label: string) {
  cy.get(`.action-menu img[title="${label}"]`).first().click({ force: true });
}

// Every example on the demo page is `<div><h2>{title}</h2>...<DataGridList/></div>` (or,
// for Example 2, a `.data-grid-list-demo__boxed` wrapper around DataGridList as that same
// sibling) - scoping through the h2's own parent is robust against the exact number/order
// of `.rdg` grids on the page changing, unlike an index-based `.rdg` selector would be.
function example(headingText: string) {
  return cy.contains("h2", headingText).parent();
}

// Repeatedly scrolls `container` towards its own bottom, pausing between steps so each
// intersection has a chance to fire its own fetchNextPage - a single scrollTo("bottom")
// only ever catches one page: by the time it lands, freshly-appended rows have already
// moved the "bottom" further down again.
function scrollRepeatedly(container: Cypress.Chainable<JQuery<HTMLElement>>, times = 8) {
  for (let i = 0; i < times; i++) {
    container.scrollTo("bottom", { ensureScrollable: false });
    cy.wait(400);
  }
}

describe("Manage: DataGridList demo - all emi column types, filtering, infinite scroll, delete", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should log into the manage UI as root.", () => {
    loginAsRoot();
  });

  it("Example 1 renders a column for every emi field type, with real values.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const grid = () => example("Every emi field type").find(".rdg");

    grid().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    // string/int/int64/float64/bool/complex(date) - all filterable, so their header
    // renders an <input> (placeholder = the column title) rather than plain text.
    example("Every emi field type").within(() => {
      cy.get('input[placeholder="uniqueId (string)"]').should("exist");
      cy.get('input[placeholder="title (string)"]').should("exist");
      cy.get('input[placeholder="age (int)"]').should("exist");
      cy.get('input[placeholder="views (int64)"]').should("exist");
      cy.get('input[placeholder="price (float64)"]').should("exist");
      cy.get('input[placeholder="active (bool)"]').should("exist");
      cy.get('input[placeholder="createdAt (complex: PlainTime)"]').should("exist");
      cy.get('input[placeholder="status (enum)"]').should("exist");
    });

    // complex(TString) is its own case: a button (not a plain input) that opens
    // TStringFilterDrawer - covered end to end in its own test below.
    example("Every emi field type")
      .find('button.data-table-filter-tstring[title="label (complex: TString)"]')
      .should("exist");

    // slice/one/collection/object/any - not filterable, so the header shows plain text.
    // react-data-grid column-virtualizes horizontally too, so these - further right
    // than the viewport width fits - don't even exist in the DOM until the grid
    // itself is scrolled right.
    grid().scrollTo("right", { ensureScrollable: false });
    example("Every emi field type").within(() => {
      cy.contains(".rdg-header-row", "tags (slice/array<string>)").should("exist");
      cy.contains(".rdg-header-row", "owner (one)").should("exist");
      cy.contains(".rdg-header-row", "collaborators (collection)").should("exist");
      cy.contains(".rdg-header-row", "metadata (object/map)").should("exist");
      cy.contains(".rdg-header-row", "note (any)").should("exist");
    });
    grid().scrollTo("left", { ensureScrollable: false });

    // The rendered *values* prove each type's getCellValue actually worked, not just
    // that the column exists: a price formatted as currency, a real status badge, and
    // at least one boolean glyph (✓ or —) in the active column.
    grid().contains(".rdg-row", "Item #1 -").within(() => {
      cy.get(".emi-demo-badge").should("exist");
      cy.contains(/^\$\d+\.\d{2}$/).should("exist");
      // label's getCellValue joins every locale's text with " · " - proof the
      // TString value round-trips through all its languages, not just one.
      cy.contains(/ · /).should("exist");
    });
    grid()
      .find(".rdg-cell")
      .contains(/^(✓|—)$/)
      .should("exist");
  });

  it("Example 1's column filter narrows the grid to matching rows only (debounced JsonLogic).", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const grid = () => example("Every emi field type").find(".rdg");
    grid().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    example("Every emi field type")
      .find('input[placeholder="title (string)"]')
      .should("have.length", 1)
      .click()
      .type("Rocket", { delay: 60 });

    // Debounce is 500ms (useColumnFilters.ts) - give it real time to settle, then every
    // currently-rendered row's title cell must contain what was typed.
    cy.wait(1200);
    grid().find(".rdg-row").should("have.length.greaterThan", 0);
    grid()
      .find(".rdg-row")
      .each(($row) => {
        expect($row.text()).to.include("Rocket");
      });
  });

  it("Example 1's TString column opens a drawer to filter by per-language text.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const grid = () => example("Every emi field type").find(".rdg");
    grid().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    const filterButton = () =>
      example("Every emi field type").find(
        'button.data-table-filter-tstring[title="label (complex: TString)"]',
      );

    filterButton().should("contain.text", "Filter...").click();

    // TStringFilterDrawer.tsx: one FormText per BUILD_VARIABLES.SUPPORTED_LANGUAGES
    // locale ("fa,en" for this build) - heading is the column's own title. Each
    // FormText renders its own <label class="form-label">EN</label> sibling to its
    // <input> (BaseFormElement.tsx) - scope through that label's own parent (not
    // the whole drawer, which contains one input per locale) to land on the right one.
    cy.contains(".confirm-drawer-container h2", "label (complex: TString)").should(
      "exist",
    );
    const enField = () =>
      cy.contains(".confirm-drawer-container label.form-label", "EN").parent();
    enField().find("input").type("Rocket", { delay: 60 });
    cy.contains(".confirm-drawer-container button", "Save").click({ force: true });

    // Debounce is 500ms (useColumnFilters.ts).
    cy.wait(1200);
    grid()
      .find(".rdg-row")
      .should("have.length.greaterThan", 0)
      .each(($row) => {
        expect($row.text()).to.include("Rocket");
      });

    // The button itself now summarizes what's active, and reopening the drawer
    // shows the same value back - both backed by useColumnFilters' own
    // structuredValues (not the header cell's own local state, which react-data-grid
    // doesn't guarantee survives every re-render).
    filterButton().should("contain.text", "Rocket").click();
    enField().find("input").should("have.value", "Rocket");

    // Clearing it and saving again lifts the filter entirely.
    enField().find("input").clear();
    cy.contains(".confirm-drawer-container button", "Save").click({ force: true });
    cy.wait(1200);
    filterButton().should("contain.text", "Filter...");
  });

  it("Example 1 loads more rows as the page scrolls (infinite scroll, page-level).", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const grid = () => example("Every emi field type").find(".rdg");
    grid().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    grid()
      .find(".rdg-row")
      .its("length")
      .then((initialCount) => {
        // This app's real scroll container is an inner .content-container div, not
        // the window (see useInfiniteScroll.ts's doc comment) - scrolling that is
        // what a user's mouse wheel actually moves too.
        scrollRepeatedly(cy.get(".content-container"));
        grid()
          .find(".rdg-row")
          .its("length")
          .should("be.greaterThan", initialCount);
      });
  });

  it("Example 2 (boxed in its own scrollable div) loads more rows on its own scroll, independent of the page.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const grid = () => example("Contained in a scrollable box").find(".rdg");
    grid().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    grid()
      .find(".rdg-row")
      .its("length")
      .then((initialCount) => {
        scrollRepeatedly(
          example("Contained in a scrollable box").find(".data-grid-list-demo__boxed"),
        );
        grid()
          .find(".rdg-row")
          .its("length")
          .should("be.greaterThan", initialCount);
      });
  });

  it("Example 3: selecting a row and confirming Delete removes it and updates the deleted count.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    const scope = () => example("Selection + delete");
    scope().find(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);

    scope().contains("0 rows deleted so far").should("exist");

    scope()
      .find(".rdg-row")
      .first()
      .invoke("text")
      .then((firstRowText) => {
        scope()
          .find(".rdg-row")
          .first()
          .find('input[type="checkbox"]')
          .click({ force: true });

        clickIconAction("Delete");
        cy.contains(".confirm-drawer-container button", "Yes").click({ force: true });

        scope().contains("1 row deleted so far", { timeout: 10000 }).should("exist");
        scope().find(".rdg-row").first().should(($row) => {
          expect($row.text()).to.not.equal(firstRowText);
        });
      });
  });

  it("shows a loading indicator (with a Cancel option) while the very first page is still loading.", () => {
    loginAsRoot();
    cy.visit(ui("/manage/#/en/demo/data-grid-list"));

    // The fake dataset's own simulated latency (~500-900ms, see
    // useFakeCursorListQuery.ts) is what this races against - Cypress's default
    // command retries are fast enough to usually still catch it.
    cy.get(".loading-progress").should("exist");
    cy.contains(".loading-progress button", "Cancel").should("exist");

    // ... and it goes away once real rows arrive.
    cy.get(".rdg-row", { timeout: 15000 }).should("have.length.greaterThan", 0);
  });

  endFirebackServer();
});
