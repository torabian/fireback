import "cypress-real-events";
import "cypress-real-events/support";
import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

describe("Logging in with the signin", () => {
  withFirebackServer();

  describe("Login with the email address needs to be working", () => {
    it("on a fresh install, there should be no authentication available at all.", () => {
      cy.viewport(400, 750); // Set the window size dynamically
      cy.visit(ui("/en/welcome"));
      cy.wait(1000);
      cy.get("h1").should("have.text", "Authentication Currently Unavailable");
    });

    it("on creation of the passport method, both type and region need to be provided.", () => {
      cy.task("execSupress", ` passport method c`).then((content) => {
        expect(content).contain('"type, region"');
      }).ca;
    });

    it("only global needs to be an option.", () => {
      cy
        .task(
          "execSupress",
          ` passport method c --region unknownregion --type email`
        )
        .then((content) => {
          expect(content).contain("ValidationFailedOnSomeFields");
          expect(content).contain("oneof");
          expect(content).contain(`"errorParam": "global",`);
        }).ca;
    });

    it("should only accept email and phone are enabled as type.", () => {
      cy.task(
        "execSupress",
        ` passport method c --region global --type email2`
      ).then((content) => {
        expect(content).contain("ValidationFailedOnSomeFields");
        expect(content).contain("oneof");
        expect(content).contain(`"errorParam": "email phone google",`);
      });
    });

    it("should be able to create email method in database.", () => {
      cy.task("exec", ` passport method c --region global --type email`);
    });

    it("should be able to create phone method in database.", () => {
      cy.task("exec", ` passport method c --region global --type phone`);
    });

    it("get the passport methods", () => {
      cy.task("exec", ` passport method q`).then((content) => {
        const res = JSON.parse(content);
        expect(res.data.items.length).to.equal(2);
      });
    });

    Cypress.on("uncaught:exception", (err, runnable) => {
      return false;
    });

    it("should be able to create an account", () => {
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

      cy.get("h1").should("have.text", "Registeration not possible.");

      // cy.get("#first-name-input").type("Ali");
      // cy.get("#last-name-input").type("Torabi");

      // cy.get("#password-input").type("123321");
      // cy.get("#password-repeat-input").type("123321");

      // cy.get("#submit-form").click({ force: true }); // Submit the form
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

    it("entering email address now should allow user to create account using form.", () => {
      cy.viewport(400, 750); // Set the window size dynamically

      cy.visit(ui("/en/welcome"));
      cy.wait(10000);

      cy.get("#using-email").should("exist").as("btn").click({ force: true });
      cy.url().should("include", "/auth/email");
      cy.get("h1").should("have.text", "Continue with Email");
      cy.wait(1000);

      // cy.get("#value-input").type("test@test.com"); // Fill the input with "admin"
      // cy.wait(500);

      // cy.get("#submit-form").click({ force: true }); // Submit the form

      // cy.wait(500);

      // cy.get("h1").should("have.text", "Complete your account");

      // cy.get("#first-name-input").type("Ali");
      // cy.get("#last-name-input").type("Torabi");

      // cy.get("#password-input").type("123321");
      // cy.get("#password-repeat-input").type("123321");

      // cy.get("#submit-form").click({ force: true }); // Submit the form

      // cy.wait(2000);
    });

    endFirebackServer();
  });
});
