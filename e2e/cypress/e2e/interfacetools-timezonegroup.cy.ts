import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/abac/interfacetools's "timezoneGroup" entity end to end - it has no
// dedicated admin CRUD screen either (see interfacetools-appmenu.cy.ts's header comment
// for why these three entities are tested over HTTP directly instead), and unlike
// appMenu/tableViewSizing it isn't referenced by any hand-written front-end code at all
// today (only its own generated SDK bindings under ui/src/modules/sdk/interfacetools
// exist) - so there's no real UI flow to additionally exercise here.
//
// Root's own CLI account, created by withFirebackServer().
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

function authHeaders(token: string) {
  return { authorization: token, "workspace-id": "root" };
}

describe("interfacetools: timezoneGroup", () => {
  withFirebackServer();

  it("should be able to sign in as root over the API to get a bearer token for the rest of this spec.", () => {
    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: { value: ROOT_EMAIL, password: ROOT_PASSWORD },
    }).then((response: Cypress.Response<SigninResponse>) => {
      expect(response.status).to.equal(200);
      const token = response.body.data.item.session.token;
      expect(token).to.be.a("string").and.not.be.empty;
      setShared("rootToken", token);
    });
  });

  it("TimezoneGroupCreate should succeed, and be rejected without auth.", () => {
    cy.request({
      method: "POST",
      url: ui("/timezoneGroup"),
      body: { title: "checkendpointtests no-auth" },
      failOnStatusCode: false,
    }).then((noAuth) => {
      expect(noAuth.status).to.not.equal(200);
    });

    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/timezoneGroup"),
        headers: authHeaders(rootToken),
        body: { title: "checkendpointtests timezone" },
      }).then(
        (
          response: Cypress.Response<SingleItemResponse<{ uniqueId: string; title: string }>>,
        ) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.title).to.equal("checkendpointtests timezone");
          setShared("tzId", response.body.data.item.uniqueId);
        },
      );
    });
  });

  it("TimezoneGroupBrowse should include the just-created record.", () => {
    sharedValues(["rootToken", "tzId"]).then(({ rootToken, tzId }) => {
      cy.request({
        method: "GET",
        url: ui("/timezoneGroup/browse"),
        headers: authHeaders(rootToken),
      }).then((response: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const ids = response.body.data.items.map((i) => i.uniqueId);
        expect(ids).to.include(tzId);
      });
    });
  });

  it("TimezoneGroupGet should return the same record by id.", () => {
    sharedValues(["rootToken", "tzId"]).then(({ rootToken, tzId }) => {
      cy.request({
        method: "GET",
        url: ui(`/timezoneGroup/${tzId}`),
        headers: authHeaders(rootToken),
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.uniqueId).to.equal(tzId);
      });
    });
  });

  it("TimezoneGroupUpdate should persist a title change.", () => {
    sharedValues(["rootToken", "tzId"]).then(({ rootToken, tzId }) => {
      cy.request({
        method: "PATCH",
        url: ui(`/timezoneGroup/${tzId}`),
        headers: authHeaders(rootToken),
        body: { title: "checkendpointtests renamed" },
      }).then((response: Cypress.Response<SingleItemResponse<{ title: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.title).to.equal("checkendpointtests renamed");
      });
    });
  });

  it("TimezoneGroupAwareDeletePreview then TimezoneGroupAwareDelete should remove the record.", () => {
    sharedValues(["rootToken", "tzId"]).then(({ rootToken, tzId }) => {
      cy.request({
        method: "GET",
        url: ui(`/timezoneGroup/delete-preview?uniqueIds=${tzId}`),
        headers: authHeaders(rootToken),
      }).then((preview) => {
        expect(preview.status).to.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/timezoneGroup/delete"),
        headers: authHeaders(rootToken),
        body: { uniqueIds: [tzId] },
      }).then((del) => {
        expect(del.status).to.equal(200);
      });

      cy.request({
        method: "GET",
        url: ui(`/timezoneGroup/${tzId}`),
        headers: authHeaders(rootToken),
        failOnStatusCode: false,
      }).then((getAfterDelete) => {
        expect(getAfterDelete.status).to.not.equal(200);
      });
    });
  });

  endFirebackServer();
});
