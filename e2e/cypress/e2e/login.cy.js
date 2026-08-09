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
      // "role c" (a nested two-word command) doesn't exist - Role's actions are
      // registered flat at the top level, so its CliShort alias "role-c" is what's
      // reachable (see PassportCli.go's hand-written nesting vs. every other
      // entity, which is flat). "--capabilities" is also stale: the field is now
      // "capabilitiesListId", a JSON array - and it can't be the "root.*" wildcard
      // this test originally used, since WorkspaceTypeActions.go's
      // ValidateTheWorkspaceTypeEntity explicitly rejects assigning a role with the
      // root.* wildcard capability to a workspace type (it would hand every user of
      // that workspace type super-admin powers). The response envelope is
      // {data: {item: {...}}} (GResponseSingleItem), not a bare object.
      cy.task(
        "exec",
        ` role-c --name testagentrole --capabilities-list-id '["root.manage.abac.notification-config.query"]'`,
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

    // let successfulInserts = 0;
    // let appMenuItems = [];

    // describe("Login with the email address needs to be working", () => {
    //   it("should be able to wipe the menu items", () => {
    //     cy.task("exec", `misc appmenu wipe`).then((content) => {
    //       expect(content).to.contain("of entities");
    //     });
    //   });

    //   it("should be able to generate back the menu items from seeder.", () => {
    //     cy.task("exec", `misc appmenu ssync`).then((content) => {
    //       console.log(100, content);
    //       let countFilesImported = 0;
    //       for (const line of content.split("\n")) {
    //         if (line.startsWith("Success")) {
    //           successfulInserts += +line.match(/Success (\d+)/)[1];
    //           expect(line).to.contain("Failure 0");
    //         }
    //         if (line.endsWith(".yml")) {
    //           countFilesImported++;
    //         }
    //       }

    //       expect(countFilesImported).to.equal(3);
    //     });
    //   });

    //   it("should be able to query the created content, and total count of data in database are equal to what we have imported", () => {
    //     cy.task("exec", `misc appmenu q`).then((content) => {
    //       const res = JSON.parse(content);
    //       expect(res.data.totalItems).to.equal(successfulInserts);
    //       appMenuItems = res.data.items;
    //     });
    //   });

    //   describe("testing the menu items content", () => {
    //     it("all menu items, should have visibility of A", () => {
    //       validateAppMenuEntity(appMenuItems);
    //     });
    //   });

    //   describe("cte operations", () => {
    //     it("should be able to query as a tree structure the menu items", () => {
    //       cy.task("exec", `misc appmenu cte`).then((content) => {
    //         console.log(content);
    //         const res = JSON.parse(content);

    //         // general validation
    //         validateAppMenuEntity(res.data.items);

    //         // There should be 3 items because there are 3 root items
    //         expect(res.data.totalItems).to.equal(3);

    //         // test if items are having children with more than 1 item in them.
    //         for (const item of res.data.items) {
    //           if (item.children && item.children.length > 0) {
    //             expect(item.children?.length).to.be.greaterThan(0);
    //           }
    //         }
    //       });
    //     });
    //   });

    //   describe("operate on the single menu item", () => {
    //     it("running update on the entity, should only affected the updated time.", () => {
    //       const item = appMenuItems[0];

    //       cy.task(
    //         "exec",
    //         `misc appmenu q --query "unique_id = ${item.uniqueId}"`,
    //       ).then((content) => {
    //         const res = JSON.parse(content);
    //         console.log(res);
    //       });
    //     });
    //   });
    // });

    endFirebackServer();
  });
});

function validateAppMenuEntity(items) {
  for (const item of items) {
    expect(item.visibility).to.equal("A");
    expect(typeof item.uniqueId).to.equal("string");

    if (item.children?.length > 0) {
      validateAppMenuEntity(item.children);
    }
  }
}
