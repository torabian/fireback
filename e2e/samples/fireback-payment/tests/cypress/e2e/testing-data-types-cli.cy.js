import "cypress-real-events";
import "cypress-real-events/support";

describe("Logging in with the signin", () => {
  it("create the test agent as root access", () => {
    cy.task("dbcon");
    cy.task(
      "exec",
      ` passport new --in-root=true --value test@${Math.random()
        .toString()
        .replace(
          ".",
          ""
        )}agent.com --workspace-type-id root --type email --password 123321 --first-name testagent --last-name testagent`
    );
  });

  it("should get the fireback version", () => {
    cy.task("exec", `version`).then((x) => {
      expect(x).to.contain("Written with love by Ali Tor");
    });
  });

  it("should contain all of the data types that we used on the help seciton of create product", () => {
    cy.task("exec", `product product c --help`).then((x) => {
      // console.log(x.split("\n"));
    });
  });

  it("should create a sample product", () => {
    const body = {
      name: "Sample Product",
      sku: "pair",
    };
    cy.task("exec", `product product c ${castToCli(body)}`).then((x) => {
      cy.task("log", JSON.parse(x));
    });
  });

  let toUpdateEntityUniqueId;

  it("should be able to query the created product", () => {
    cy.task("exec", `product product q`).then((x) => {
      const items = JSON.parse(x).data.items;
      toUpdateEntityUniqueId = items[0].uniqueId;
    });
  });

  it("should update the the product", () => {
    const body = {
      name: "Product edited",
      sku: null,
    };

    const cmd = `product product u --uid ${toUpdateEntityUniqueId} ${castToCli(
      body
    )}`;

    cy.task("exec", cmd).then((x) => {
      const res = JSON.parse(x);

      expect(res.sku).to.equal(body.sku);
      expect(res.name).to.equal(body.name);
    });
  });

  it("name cannot be null", () => {
    const body = {
      name: null,
    };

    const cmd = `product product u --uid ${toUpdateEntityUniqueId} ${castToCli(
      body
    )}`;

    cy.task("exec", cmd).then((x) => {
      const res = JSON.parse(x);
      // non-nullable string field will treat null as a value, not an action
      expect(res.name).to.equal("null");
    });
  });

  it("should be able to export the products", () => {
    cy.task("exec", `product product export --file test.json`).then((x) => {
      cy.task("log", x);
    });
  });

  // it("should end the session", () => {
  //   endFirebackServer();
  // });
});

/**
 * Coverts a javascript object to --var=value string format.
 * If the body is a nested object, it would go through them and prefix it with parent
 * object name. Make sure all the string out values are dashed even if the object name is camel case.
 * @param {C} body
 */
function castToCli(body, prefix = "") {
  return Object.entries(body)
    .flatMap(([key, value]) => {
      const dashedKey = (prefix + key).replace(
        /[A-Z]/g,
        (m) => `-${m.toLowerCase()}`
      );

      if (typeof value === "object" && value !== null) {
        return castToCli(value, dashedKey + "-");
      }

      const formattedValue =
        typeof value === "string" ? `"${value.replace(/"/g, '\\"')}"` : value;

      return `--${dashedKey}=${formattedValue}`;
    })
    .join(" ");
}
