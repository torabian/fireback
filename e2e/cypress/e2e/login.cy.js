describe("Logging in with the signin", () => {
  it("Login with the email address needs to be working", () => {
    cy.visit("http://localhost:3000/#/en/welcome");
    cy.get("#using-email").should("exist").as("btn").click({ force: true });
    cy.url().should("include", "/auth/email");
    cy.get("h1").should("have.text", "Continue with Email");

    // Check if the go back works just fine.
    cy.get("#back-to-general-step").should("exist").click({ force: true });
    cy.url().should("match", /\/welcome$/);

    // go to the email and complete the flow
    cy.get("#using-email").should("exist").as("btn").click({ force: true });
    cy.get("#email-input").type("admin"); // Fill the input with "admin"
    cy.wait(500);
    cy.get("#submit-form").click({ force: true }); // Submit the form
    cy.get("#password-input").type("admin"); // Fill the input with "admin"
    cy.get("#submit-form").click({ force: true }); // Submit the form

    // check if app has navigated to /users screen as the entry point, since project has no dashboard
    cy.url().should("match", /\/users$/);
  });
});
