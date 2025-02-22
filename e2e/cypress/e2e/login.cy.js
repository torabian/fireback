import "cypress-real-events";
import "cypress-real-events/support";

let binary = "/Users/ali/work/fireback/app";
let cwd = "/Users/ali/work/fireback";
let configUniqueId = "";
const isGitHubActions = !!process.env.GITHUB_ACTIONS;

if (isGitHubActions) {
  binary = "/home/runner/work/fireback/fireback";
  cwd = "/home/runner/work/fireback";
}

describe("Logging in with the signin", () => {
  describe("Login with the email address needs to be working", () => {
    console.log(12, binary, cwd);

    Cypress.on("fail", (err) => {
      console.error("Test failed, stopping Fireback...");
      cy.task("stopFireback");
      throw err; // Re-throw error so Cypress still fails the test
    });

    beforeEach(() => {
      cy.task("execCwd", cwd);
    });
    it("create a new database connection", () => {
      cy.task(
        "exec",
        `${binary} config db-name set /tmp/test-agent-${new Date().getTime()}.db`
      );
    });
    it("get db name", () => {
      cy.task(
        "exec",
        `${binary} config db-name get && ${binary} migration apply`
      );
    });

    it("create the test agent as root access", () => {
      cy.task(
        "exec",
        `${binary} passport new --in-root=true --value testagent --workspace-type-id root --type email --password 123321 --first-name testagent --last-name testagent`
      );
    });

    it("start the server", () => {
      cy.task("startFireback");
    });

    if (isGitHubActions) {
      it("on a fresh install, there should be no authentication available at all.", () => {
        cy.viewport(400, 750); // Set the window size dynamically

        cy.visit("http://localhost:4502/#/en/welcome");
        cy.wait(1000);
        cy.get("h1").should(
          "have.text",
          "Authentication Currently Unavailable"
        );
      });
    }

    it("on creation of the passport method, both type and region need to be provided.", () => {
      cy.task("execSupress", `${binary} passport method c`).then((content) => {
        expect(content).contain('"type, region"');
      }).ca;
    });

    it("only global needs to be an option.", () => {
      cy
        .task(
          "execSupress",
          `${binary} passport method c --region unknownregion --type email`
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
        `${binary} passport method c --region global --type email2`
      ).then((content) => {
        expect(content).contain("ValidationFailedOnSomeFields");
        expect(content).contain("oneof");
        expect(content).contain(`"errorParam": "email phone google",`);
      });
    });

    it("should be able to create email method in database.", () => {
      cy.task(
        "exec",
        `${binary} passport method c --region global --type email`
      );
    });

    it("should be able to create phone method in database.", () => {
      cy.task(
        "exec",
        `${binary} passport method c --region global --type phone`
      );
    });

    Cypress.on("uncaught:exception", (err, runnable) => {
      return false;
    });

    it("test the flow with recaptcha 2 enabled on test env", () => {
      cy.viewport(400, 750); // Set the window size dynamically

      cy.visit("http://localhost:4502/#/en/welcome");
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
        `${binary} role c --name testagentrole --capabilities "root/*"`
      ).then((res) => {
        console.log((roleId = JSON.parse(res).uniqueId));
      });
    });

    it("should be able to create a workspace name", () => {
      cy.task(
        "exec",
        `${binary} ws type c --title customer --slug customer --role-id ${roleId}`
      );
    });

    it("entering email address now should allow user to create account using form.", () => {
      cy.viewport(400, 750); // Set the window size dynamically

      cy.visit("http://localhost:4502/#/en/welcome");
      cy.wait(1000);
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

    it("should stop fireback", () => {
      cy.task("stopFireback");
    });
  });
});
