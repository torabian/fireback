import "cypress-real-events";
import "cypress-real-events/support";
import { endFirebackServer, withFirebackServer } from "../support/setup";

// This test needs to run based on a existing project, it doesn't configurate it's own database setup
describe("Testing the app menu and cte functionality", () => {
  withFirebackServer();

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

/**
 * Recursively would validate menu items, and if there is problem will throw error
 * @param {R} item
 */
function validateAppMenuEntity(items) {
  for (const item of items) {
    expect(item.visibility).to.equal("A");
    expect(typeof item.uniqueId).to.equal("string");

    if (item.children?.length > 0) {
      validateAppMenuEntity(item.children);
    }
  }
}
