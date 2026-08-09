import { endFirebackServer, withFirebackServer } from "../support/setup";

interface PassportMethodItem {
  uniqueId: string;
  type: string;
  region: string;
  clientKey: string;
  workspaceId: string | null;
  userId: string | null;
  createdAt: string;
  updatedAt: string;
}

interface SingleItemResponse<T> {
  data: {
    item: T;
  };
}

interface BrowseResponse<T> {
  data: {
    items: T[];
    totalItems: number;
  };
}

interface AvailableMethodsState {
  email: boolean;
  phone: boolean;
  google: boolean;
  facebook: boolean;
}

const AVAILABLE_METHODS_URL = "http://localhost:7794/passports/available-methods";
const PASSPORT_METHOD_URL = "http://localhost:7794/passportMethod";

describe("Passport method management", () => {
  withFirebackServer();

  // Kept as a single flat list of `it`s (no nested `describe`) - a nested describe's
  // tests run after every direct `it` in this describe, including endFirebackServer's
  // "should stop fireback" call below, which would tear the server down before these
  // ever ran.

  it("available passport methods should be reachable with no authentication, and report nothing available before any method exists.", () => {
    cy.request({
      method: "GET",
      url: AVAILABLE_METHODS_URL,
      // Deliberately no Authorization header - this endpoint must work unauthenticated,
      // since it's what an unauthenticated visitor's welcome page relies on.
    }).then((response: Cypress.Response<SingleItemResponse<AvailableMethodsState>>) => {
      expect(response.status).to.equal(200);
      expect(response.body.data.item.email).to.equal(false);
      expect(response.body.data.item.phone).to.equal(false);
    });
  });

  it("should be able to create an email method in database.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("should be able to create a phone method in database.", () => {
    cy.task("exec", ` passport method create --region global --type phone`);
  });

  it("available passport methods should now report email and phone as available, still with no authentication.", () => {
    cy.request({
      method: "GET",
      url: AVAILABLE_METHODS_URL,
    }).then((response: Cypress.Response<SingleItemResponse<AvailableMethodsState>>) => {
      expect(response.status).to.equal(200);
      expect(response.body.data.item.email).to.equal(true);
      expect(response.body.data.item.phone).to.equal(true);
    });
  });

  it("should NOT be able to create a passport method with a (type, region) pair that already exists.", () => {
    // "email" + "global" was already created above - creating it again must fail.
    cy.task(
      "execSupress",
      ` passport method create --region global --type email`,
    ).then((content: string) => {
      expect(content).to.contain("PassportMethodAlreadyExists");
    });
  });

  it("should NOT be able to update a passport method into a (type, region) pair another row already occupies.", () => {
    // A fresh "facebook"/"global" row, then renaming it onto "email"/"global" -
    // already taken by the row created above - must fail the same way a fresh
    // Create would.
    cy.task(
      "exec",
      ` passport method create --region global --type facebook`,
    ).then((content: string) => {
      const facebookMethod = (
        JSON.parse(content) as SingleItemResponse<PassportMethodItem>
      ).data.item;

      cy.task(
        "execSupress",
        ` passport method update --pp-uniqueId ${facebookMethod.uniqueId} --type email`,
      ).then((updateContent: string) => {
        expect(updateContent).to.contain("PassportMethodAlreadyExists");
      });
    });
  });

  it("should NOT be able to create a passport method while not in the root workspace, even with a valid root token.", () => {
    cy.task("exec", ` config cli-token get`).then((tokenOutput: string) => {
      const token = tokenOutput.trim();

      cy.request({
        method: "POST",
        url: PASSPORT_METHOD_URL,
        failOnStatusCode: false,
        headers: {
          authorization: token,
          // A workspace-id the root token has no real standing in at all (not
          // just "not root") now gets rejected one gate earlier than before:
          // Permissions.go's MeetsAccessLevel used to always report "meets",
          // regardless of the caller's actual capabilities, so every request
          // sailed through it and only the dedicated AllowOnRoot check (root
          // workspace required) ever caught this case, as 400
          // ActionOnlyInRoot. Now that MeetsAccessLevel actually enforces its
          // verdict, a workspace the caller isn't even a member of fails the
          // capability check first, as 401 NotEnoughPermission - still a
          // rejection either way, just a different (now correct) reason.
          "Workspace-id": "not-root",
        },
        body: {
          type: "google",
          region: "global",
        },
      }).then((response: Cypress.Response<{ error: { message: string } }>) => {
        expect(response.status).to.equal(401);
        expect(response.body.error.message).to.equal("NotEnoughPermission");
      });
    });
  });

  it("should be able to browse passport methods, and see every method created so far.", () => {
    cy.task("exec", ` passport method browse`).then((content: string) => {
      const res = JSON.parse(content) as BrowseResponse<PassportMethodItem>;
      const types = res.data.items.map((item) => item.type).sort();
      expect(types).to.deep.equal(["email", "facebook", "phone"]);
    });
  });

  let googleMethodUniqueId = "";

  it("should be able to create a google method, in order to delete it next.", () => {
    cy.task(
      "exec",
      ` passport method create --region global --type google --client-key test-client-key`,
    ).then((content: string) => {
      const googleMethod = (
        JSON.parse(content) as SingleItemResponse<PassportMethodItem>
      ).data.item;
      googleMethodUniqueId = googleMethod.uniqueId;
      expect(googleMethod.type).to.equal("google");
    });
  });

  it("should be able to delete the google passport method, and no longer see it in browse.", () => {
    cy.task(
      "exec",
      ` passport method delete --unique-ids ${googleMethodUniqueId}`,
    ).then(() => {
      cy.task("exec", ` passport method browse`).then((content: string) => {
        const res = JSON.parse(content) as BrowseResponse<PassportMethodItem>;
        const ids = res.data.items.map((item) => item.uniqueId);
        expect(ids).to.not.include(googleMethodUniqueId);
      });
    });
  });

  it("available passport methods should no longer report google as available after it was deleted.", () => {
    cy.request({
      method: "GET",
      url: AVAILABLE_METHODS_URL,
    }).then((response: Cypress.Response<SingleItemResponse<AvailableMethodsState>>) => {
      expect(response.status).to.equal(200);
      expect(response.body.data.item.google).to.equal(false);
    });
  });

  endFirebackServer();
});
