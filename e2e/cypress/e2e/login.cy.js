import "cypress-real-events";
import "cypress-real-events/support";

describe("Logging in with the signin", () => {
  it("Login with the email address needs to be working", () => {
    cy.viewport(400, 900); // Set the window size dynamically

    cy.visit("http://localhost:3000/#/en/welcome");
    cy.get("#using-email").should("exist").as("btn").click({ force: true });
    cy.url().should("include", "/auth/email");
    cy.get("h1").should("have.text", "Continue with Email");

    cy.screenshot();

    // Check if the go back works just fine.
    cy.get("#back-to-general-step").should("exist").click({ force: true });
    cy.url().should("match", /\/welcome$/);

    // go to the email and complete the flow
    cy.get("#using-email").should("exist").as("btn").click({ force: true });
    cy.get("#value-input").type("admin"); // Fill the input with "admin"
    cy.wait(500);

    cy.get("#submit-form").click({ force: true }); // Submit the form

    // let's make sure if we input wrong password it would complain
    cy.get("#password-input").type("admin2"); // Fill the input with "admin"
    cy.get("#submit-form").click({ force: true }); // Submit the form

    cy.get(".basic-error-box").should(
      "contain",
      "This passport is not available."
    );

    cy.get("#password-input").clear().type("admin"); // Fill the input with "admin"

    cy.screenshot();

    cy.get("#back-to-general-step").should("exist");

    cy.get("#submit-form").click({ force: true }); // Submit the form

    // check if app has navigated to /users screen as the entry point, since project has no dashboard
    cy.url().should("match", /\/users$/);

    cy.wait(500);

    cy.realPress("Tab");
    cy.realPress("ArrowRight");
    cy.realPress("ArrowRight");
  });
});
