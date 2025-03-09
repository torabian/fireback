/**
 * Setup a fireback server, lifts it, and creates a root account.
 * Make sure you include this in the first describe, and it's the only describe
 * of the test case because it has afterEach to stop the server.
 */

export function ui(affix) {
  return `http://localhost:${Cypress.env("PORT")}/#${affix}`;
}

export function withFirebackServer() {
  Cypress.on("fail", (err) => {
    console.error("Test failed, stopping Fireback...");
    cy.task("stopFireback");
    throw err; // Re-throw error so Cypress still fails the test
  });

  it("create the test agent as root access", () => {
    // cy.task("dbcon");
    cy.task(
      "exec",
      ` passport new --in-root=true --value test@${Math.random()
        .toString()
        .replace(
          ".",
          ""
        )}agent.com --workspace-type-id root --type email --password 123321 --first-name testagent --last-name testagent`
    ).then((x) => {});
    cy.task("startFireback");
    cy.wait(2500);
  });
}

export function endFirebackServer() {
  it("should stop fireback", () => {
    cy.task("stopFireback");
  });
}
