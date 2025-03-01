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
      cy.visit(ui("/en/welcome"));
      cy.wait(1000);
      cy.get("h1").should("have.text", "Authentication Currently Unavailable");
    });

    it("get the data of the public", () => {
      cy.request(
        "GET",
        "http://localhost:7793/passports/available-methods"
      ).then((response) => {
        cy.task("log", response.body);
        expect(response.body.data.email).to.equal(null);
        expect(response.body.data.phone).to.equal(null);
      });
    });

    it("should be able to create email method in database.", () => {
      cy.task("exec", ` passport method c --region global --type email`);
    });

    it("should be able to create phone method in database.", () => {
      cy.task("exec", ` passport method c --region global --type phone`);
    });

    it("get the data of the public", () => {
      cy.request(
        "GET",
        "http://localhost:7793/passports/available-methods"
      ).then((response) => {
        cy.task("log", response.body);
        expect(response.body.data.email).to.equal(true);
        expect(response.body.data.phone).to.equal(true);
      });
    });

    it("on a fresh install, there should be no authentication available at all.", () => {
      cy.viewport(400, 750); // Set the window size dynamically
      cy.visit(ui("/en/welcome"));
      cy.wait(1000);
      cy.get("h1").should("have.text", "Welcome back");
    });

    it("on creation of the passport method, both type and region need to be provided.", () => {
      cy.task("execSupress", ` passport method c`).then((content) => {
        expect(content).contain('"type, region"');
      }).ca;
    });

    let roleId = "";
    it("should be able to create a role in order to assign it into the workspace type.", () => {
      cy.task(
        "exec",
        ` role c --name testagentrole --capabilities "root/*"`
      ).then((res) => {
        console.log((roleId = JSON.parse(res).uniqueId));
      });
    });

    it("should be able to create a workspace name", () => {
      cy.task(
        "exec",
        ` ws type c --title customer --slug customer --role-id ${roleId}`
      );
    });

    it("should be able to create an account", () => {
      Cypress.on("uncaught:exception", (err, runnable) => {
        // returning false here prevents Cypress from
        // failing the test
        return false;
      });

      cy.viewport(400, 750); // Set the window size dynamically

      cy.visit(ui("/en/welcome"));
      cy.get("#using-email").should("exist").click();
      cy.url().should("include", "/auth/email");
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

    let successfulInserts = 0;
    let appMenuItems = [];

    describe("Login with the email address needs to be working", () => {
      it("should be able to wipe the menu items", () => {
        cy.task("exec", `misc appmenu wipe`).then((content) => {
          expect(content).to.contain("of entities");
        });
      });

      it("should be able to generate back the menu items from seeder.", () => {
        cy.task("exec", `misc appmenu ssync`).then((content) => {
          let countFilesImported = 0;
          for (const line of content.split("\n")) {
            if (line.startsWith("Success")) {
              successfulInserts += +line.match(/Success (\d+)/)[1];
              expect(line).to.contain("Failure 0");
            }
            if (line.endsWith(".yml")) {
              countFilesImported++;
            }
          }

          expect(countFilesImported).to.equal(4);
        });
      });

      it("should be able to query the created content, and total count of data in database are equal to what we have imported", () => {
        cy.task("exec", `misc appmenu q`).then((content) => {
          const res = JSON.parse(content);
          expect(res.data.totalItems).to.equal(successfulInserts);
          appMenuItems = res.data.items;
        });
      });

      describe("testing the menu items content", () => {
        it("all menu items, should have visibility of A", () => {
          validateAppMenuEntity(appMenuItems);
        });
      });

      describe("cte operations", () => {
        it("should be able to query as a tree structure the menu items", () => {
          cy.task("exec", `misc appmenu cte`).then((content) => {
            console.log(content);
            const res = JSON.parse(content);

            // general validation
            validateAppMenuEntity(res.data.items);

            // There should be 3 items because there are 3 root items
            expect(res.data.totalItems).to.equal(3);

            // test if items are having children with more than 1 item in them.
            for (const item of res.data.items) {
              expect(item.children.length).to.be.greaterThan(0);
            }
          });
        });
      });

      describe("operate on the single menu item", () => {
        it("running update on the entity, should only affected the updated time.", () => {
          const item = appMenuItems[0];

          cy.task(
            "exec",
            `misc appmenu q --query "unique_id = ${item.uniqueId}"`
          ).then((content) => {
            const res = JSON.parse(content);
            console.log(res);
          });
        });
      });
    });

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
