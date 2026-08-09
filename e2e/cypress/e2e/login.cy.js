import "cypress-real-events";
import "cypress-real-events/support";
import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

describe("Logging in with the signin", () => {
  withFirebackServer();
  Cypress.on("log:added", (log) => {
    console.log(log);
  });

  describe("Login with the email address needs to be working", () => {
    it("on a fresh install, there should be no authentication available at all.", () => {
      cy.viewport(400, 750); // Set the window size dynamically
      cy.visit(ui("/manage/#/en/welcome"));
      cy.wait(1000);
      cy.get("h1").should("have.text", "Authentication Currently Unavailable");
    });

    // Passport method CRUD, uniqueness/root-only enforcement, and the public
    // available-methods API all moved to passport-method.cy.ts. This file still
    // needs email+phone methods enabled as its own fixture, to test the welcome
    // page's behavior once authentication is available.
    it("should be able to create email and phone passport methods.", () => {
      cy.task("exec", ` passport method create --region global --type email`);
      cy.task("exec", ` passport method create --region global --type phone`);
    });

    it("get the data of the public", () => {
      cy.request(
        "GET",
        "http://localhost:7794/passports/available-methods",
      ).then((response) => {
        cy.task("log", response.body);
        expect(response.body.data.item.email).to.equal(true);
        expect(response.body.data.item.phone).to.equal(true);
      });
    });

    it("should show welcome back when it's email and phone enabled.", () => {
      cy.viewport(400, 750); // Set the window size dynamically
      cy.visit(ui("/manage/#/en/welcome"));
      cy.wait(1000);
      cy.get("h1").should("have.text", "Welcome back");
    });

    it("on creation of the passport method, both type and region need to be provided.", () => {
      // The old CLI shorthand ("passport method c") and error shape
      // ('"type, region"' as one combined string) are both stale - "c" doesn't
      // resolve to anything under the current nested CLI structure, and
      // validation now reports structured per-field errors (see
      // fireback.CommonStructValidatorPointer) instead of one joined string.
      cy.task("execSupress", ` passport method create`).then((content) => {
        expect(content).to.contain("ValidationFailedOnSomeFields");
        expect(content).to.contain('"location": "type"');
        expect(content).to.contain('"location": "region"');
      });
    });

    let roleId = "";
    it("should be able to create a role in order to assign it into the workspace type.", () => {
      // "role c" doesn't exist, but Role's actions are now nested under their own
      // "role" group (AbacModule.go), so "role create" is what's reachable there.
      // "--capabilities" is also stale: the field is now "capabilitiesListId", a
      // JSON array - and it can't be the "root.*" wildcard this test originally
      // used, since WorkspaceTypeActions.go's ValidateTheWorkspaceTypeEntity
      // explicitly rejects assigning a role with the root.* wildcard capability to
      // a workspace type (it would hand every user of that workspace type
      // super-admin powers). The response envelope is {data: {item: {...}}}
      // (GResponseSingleItem), not a bare object.
      cy.task(
        "exec",
        ` role create --name testagentrole --capabilities-list-id '["root.manage.abac.notification-config.query"]'`,
      ).then((res) => {
        roleId = JSON.parse(res).data.item.uniqueId;
        expect(roleId).to.be.a("string").and.not.be.empty;
      });
    });

    it("should be able to create a workspace type", () => {
      // "ws type c" doesn't exist either - WorkspaceType's actions are registered
      // flat inside the "ws"/"workspace" group (WorkspaceCliCommands in
      // WorkspaceCli.go), not under their own "type" sub-group, so
      // "workspaceType-c" (its CliShort) is what's reachable there. The slug also
      // now has an enforced format: must start with "/" (see
      // ValidateTheWorkspaceTypeEntity/WorkspaceTypeActions.go) - "customer" alone
      // fails validation, "/customer" doesn't.
      cy.task(
        "exec",
        ` ws workspaceType-c --title customer --slug /customer --role-id ${roleId}`,
      );
    });

    it("should be able to create an account", () => {
      Cypress.on("uncaught:exception", (err, runnable) => {
        // returning false here prevents Cypress from
        // failing the test
        return false;
      });

      cy.viewport(400, 750); // Set the window size dynamically

      cy.visit(ui("/manage/#/en/welcome"));
      cy.get("#using-email").should("exist").click();
      cy.url().should("include", "/selfservice/email");
      cy.get("h1").should("have.text", "Continue with Email");
      cy.wait(1000);

      // // Check if the go back works just fine.
      cy.get("#go-back-button").should("exist").click({ force: true });
      cy.url().should("match", /\/welcome$/);

      cy.get("#using-email").should("exist").as("btn").click({ force: true });
      cy.get("#value-input").type("test@test.com"); // Fill the input with "admin"
      cy.wait(500);

      cy.get("#submit-form").click({ force: true }); // Submit the form

      cy.wait(500);

      cy.get("h1").should("have.text", "Complete your account");

      cy.get("#first-name-input").type("Ali");
      cy.get("#last-name-input").type("Torabi");

      cy.get("#password-input").type("123321");

      cy.get("#submit-form").click({ force: true }); // Submit the form

      cy.wait(500);
    });

    // The old Module3-generated "misc appmenu wipe/ssync/q/cte" CLI commands don't
    // exist anymore - AppMenu was migrated to the standard Emi-generated CRUD set
    // (browse/get/create/update/delete-preview/delete) plus one restored custom
    // command, "cte-app-menus" (see CteAppMenusActionImplementation.go), and moved
    // from "misc" to its own nested "interface"/"intf"/"if" group
    // (InterfaceToolsModule.go), so it's "interface appmenu ..." now. Kept as flat
    // `it`s (not nested describes) for the same reason as above - a nested
    // describe's tests run after every direct `it` in this describe, including
    // endFirebackServer's "should stop fireback" call.
    let appMenuTree = [];

    it("should be able to generate the menu items from seeders.", () => {
      // "misc appmenu ssync" (scoped to just this entity) has no replacement -
      // the only seeder-import command left is the global "seeders", which
      // imports every module's seed data in one pass (time zones, passport
      // methods, and the 3 appmenu yml files), not just appmenu's. So this only
      // checks that nothing failed, rather than counting exactly 3 imported files.
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

    it("should be able to query the app menu items as a tree structure.", () => {
      // "misc appmenu q" (flat query) is NOT usable here: AppMenuSyncSeeders
      // stamps every seeded row workspaceId="system" (see the yml files under
      // interfacetools/seeders/AppMenu/), but a CLI-authenticated root user's
      // Browse scope only covers workspaces they belong to ("root"), so Browse
      // silently returns 0 items even though the rows exist (confirmed directly
      // in the database). Single-item Get/Update aren't scoped the same way and
      // do reach "system" rows - and "cte-app-menus" doesn't go through Browse's
      // scoping at all (AppMenuCte.vsql has no workspace filter), so it's what
      // this test uses, which conveniently is also the direct replacement for the
      // old "cte" command.
      cy.task("exec", ` interface appmenu cte-app-menus`).then((content) => {
        const res = JSON.parse(content);
        appMenuTree = res.data.items;

        // There should be 3 items because there are 3 root items
        expect(res.data.totalItems).to.equal(3);

        validateAppMenuEntity(appMenuTree);

        // test if items are having children with more than 1 item in them.
        for (const item of appMenuTree) {
          if (item.children && item.children.length > 0) {
            expect(item.children?.length).to.be.greaterThan(0);
          }
        }
      });
    });

    it("should be able to look up a single app menu item by its uniqueId.", () => {
      const item = appMenuTree[0];

      cy.task(
        "exec",
        ` interface appmenu get --pp-uniqueId ${item.uniqueId}`,
      ).then((content) => {
        const res = JSON.parse(content);
        expect(res.data.item.uniqueId).to.equal(item.uniqueId);
      });
    });

    endFirebackServer();
  });
});

function validateAppMenuEntity(items) {
  // AppMenuEntity no longer has a "visibility" field (that check was dropped along
  // with it) - what's left to validate structurally is the uniqueId shape and the
  // recursive children nesting.
  for (const item of items) {
    expect(typeof item.uniqueId).to.equal("string");

    if (item.children?.length > 0) {
      validateAppMenuEntity(item.children);
    }
  }
}
