import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers 5 of Abac.emi.yml's CRUD entities: PendingWorkspaceInvite, PhoneConfirmation,
// Preference, PublicAuthentication, PublicJoinKey. None of these have a dedicated
// manage-UI admin screen except PublicJoinKey (selfservice/public-join-keys, which
// wires its own hooks correctly - checked, no repeat of the PassportMethodList bug), so
// all five are exercised directly over HTTP, same as the other CRUD-entity specs in
// this pass.
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

interface SingleItemResponse<T> {
  data: { item: T };
}
interface ListResponse<T> {
  data: { items: T[] };
}
interface SigninResponse {
  data: { item: { session: { token: string } } };
}

function setShared(key: string, value: string) {
  return cy.task("setShared", { key, value });
}
function sharedValues<K extends string>(
  keys: K[],
): Cypress.Chainable<Record<K, string>> {
  return cy.task("getSharedState", keys);
}

function rootHeaders(token: string) {
  return { authorization: token, "workspace-id": "root" };
}

describe("Abac: PendingWorkspaceInvite, PhoneConfirmation, Preference, PublicAuthentication, PublicJoinKey", () => {
  withFirebackServer();

  it("should sign in as root over the API for the checks below.", () => {
    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: { value: ROOT_EMAIL, password: ROOT_PASSWORD },
    }).then((response: Cypress.Response<SigninResponse>) => {
      expect(response.status).to.equal(200);
      setShared("rootToken", response.body.data.item.session.token);
    });
  });

  it("PendingWorkspaceInvite CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      const value = `checkendpointtests-pwi-${Date.now()}@example.com`;
      cy.request({
        method: "POST",
        url: ui("/pendingWorkspaceInvite"),
        headers: rootHeaders(rootToken),
        body: { value, type: "email", workspaceName: "checkendpointtests workspace" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/pendingWorkspaceInvite/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "POST",
          url: ui("/pendingWorkspaceInvite/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("PhoneConfirmation CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/phoneConfirmation"),
        headers: rootHeaders(rootToken),
        body: { phoneNumber: `+1555${Date.now() % 10000000}`, status: "pending" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; status: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/phoneConfirmation/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/phoneConfirmation/${id}`),
          headers: rootHeaders(rootToken),
          body: { status: "confirmed" },
        }).then((update: Cypress.Response<SingleItemResponse<{ status: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.status).to.equal("confirmed");
        });

        cy.request({
          method: "POST",
          url: ui("/phoneConfirmation/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("Preference CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/preference"),
        headers: rootHeaders(rootToken),
        body: { timezone: "America/New_York" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "PATCH",
          url: ui(`/preference/${id}`),
          headers: rootHeaders(rootToken),
          body: { timezone: "Europe/Warsaw" },
        }).then((update: Cypress.Response<SingleItemResponse<{ timezone: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.timezone).to.equal("Europe/Warsaw");
        });

        cy.request({
          method: "GET",
          url: ui(`/preference/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/preference/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("PublicAuthentication CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      const value = `checkendpointtests-pa-${Date.now()}@example.com`;
      cy.request({
        method: "POST",
        url: ui("/publicAuthentication"),
        headers: rootHeaders(rootToken),
        body: { passportValue: value, status: "pending", otp: "123456", workspaceId: "root" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/publicAuthentication/browse"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "POST",
          url: ui("/publicAuthentication/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("PublicJoinKey CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/role"),
        headers: rootHeaders(rootToken),
        body: {
          name: `checkendpointtests public-join-key role ${Date.now()}`,
          capabilitiesListId: ["root.abac.email-confirmation.query"],
        },
      }).then((roleResponse: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(roleResponse.status).to.equal(200);
        const roleId = roleResponse.body.data.item.uniqueId;

        cy.request({
          method: "POST",
          url: ui("/publicJoinKey"),
          headers: rootHeaders(rootToken),
          body: { roleId, workspaceId: "root" },
        }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; roleId: string }>>) => {
          expect(response.status).to.equal(200);
          const id = response.body.data.item.uniqueId;
          expect(response.body.data.item.roleId).to.equal(roleId);

          cy.request({
            method: "GET",
            url: ui("/publicJoinKey/browse"),
            headers: rootHeaders(rootToken),
          }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
            expect(browse.status).to.equal(200);
            expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
          });

          cy.request({
            method: "POST",
            url: ui("/publicJoinKey/delete"),
            headers: rootHeaders(rootToken),
            body: { uniqueIds: [id] },
          }).then((del) => {
            expect(del.status).to.equal(200);
          });
        });

        cy.request({
          method: "POST",
          url: ui("/role/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [roleId] },
        });
      });
    });
  });

  endFirebackServer();
});
