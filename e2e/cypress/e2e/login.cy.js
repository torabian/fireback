import "cypress-real-events";
import "cypress-real-events/support";

const binary = "/Users/ali/work/fireback/app";
let configUniqueId = "";

const initialSigninActions = () => {
  cy.viewport(400, 500); // Set the window size dynamically

  cy.visit("http://localhost:3000/#/en/welcome");
  cy.get("#using-email").should("exist").as("btn").click({ force: true });
  cy.url().should("include", "/auth/email");
  cy.get("h1").should("have.text", "Continue with Email");

  // Check if the go back works just fine.
  cy.get("#back-to-general-step").should("exist").click({ force: true });
  cy.url().should("match", /\/welcome$/);

  // go to the email and complete the flow
  cy.get("#using-email").should("exist").as("btn").click({ force: true });
  cy.get("#value-input").type("admin"); // Fill the input with "admin"
  cy.wait(500);
};

describe("Logging in with the signin", () => {
  describe("Login with the email address needs to be working", () => {
    it("get the config", () => {
      cy.task("exec", `${binary} ws config q`).then((output) => {
        const result = JSON.parse(output);
        expect(typeof result.data.items[0]?.uniqueId).to.equal("string");
        configUniqueId = result.data.items[0]?.uniqueId;
        console.log(configUniqueId);
      });
    });

    // it("disable the recaptcha", () => {
    //   cy.task(
    //     "exec",
    //     `${binary} ws config u --uid ${configUniqueId} --enable-recaptcha2=false`
    //   );
    // });

    // it("test the flow without recaptcha 2", () => {
    //   initialSigninActions();

    //   // let's make sure if we input wrong password it would complain
    //   cy.get("#password-input").type("admin2"); // Fill the input with "admin"
    //   cy.get("#submit-form").click({ force: true }); // Submit the form

    //   cy.get(".basic-error-box").should(
    //     "contain",
    //     "This passport is not available."
    //   );
    //   cy.get("#password-input").clear().type("admin"); // Fill the input with "admin"
    //   cy.screenshot();
    //   cy.get("#back-to-general-step").should("exist");
    //   cy.get("#submit-form").click({ force: true }); // Submit the form

    //   // check if app has navigated to /users screen as the entry point, since project has no dashboard
    //   cy.url().should("match", /\/users$/);
    // });

    it("enable the recaptcha with testing keys to see if it works correctly.", () => {
      // The values are provided by google in public for v2
      const serverKey = "6LeIxAcTAAAAAGG-vFI1TnRWxMZNFuojJ4WifJWe";
      const clientKey = "6LeIxAcTAAAAAJcZVRqyHh71UMIEGNQ_MXjiZKhI";

      cy.task(
        "exec",
        `${binary} ws config u --uid ${configUniqueId} --enable-recaptcha2=true --recaptcha2-server-key=${serverKey} --recaptcha2-client-key=${clientKey}`
      );
    });

    it("test the flow with recaptcha 2 enabled on test env", () => {
      initialSigninActions();

      // cy.get("#submit-form").click({ force: true }); // Submit the form
    });
  });
});
