import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/abac/interfacetools's "tableViewSizing" entity end to end - it has no
// dedicated admin CRUD screen either (see interfacetools-appmenu.cy.ts's header comment
// for why these three entities are tested over HTTP directly instead).
//
// Note on why this isn't driven through an actual data-grid column resize instead:
// CommonListManager.tsx addresses this entity by `queryHook.UKEY` - a static property
// the comment on QueryHook (in that same file) says is "not yet present on the newer
// generated action hooks". None of this repo's currently-generated
// use<Entity>BrowseActionQuery hooks (see e.g. sdk/abac/WorkspaceTypeBrowseAction.ts)
// actually define one, so every real call site's `queryHook.UKEY` is `undefined` today -
// the column-width-persistence feature has no working uniqueId to key off of in the
// current app. That's a pre-existing gap in the query-hook generator, not something an
// e2e spec can route around by clicking a real grid - so this spec proves the endpoints
// themselves (including the upsert-by-caller-chosen-key behavior the real feature would
// depend on once UKEY is wired up) work correctly over HTTP instead.
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

describe("interfacetools: tableViewSizing", () => {
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

  it("TableViewSizingCreate should reject a missing tableName.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/tableViewSizing"),
        headers: authHeaders(rootToken),
        body: { sizes: "[]" },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.not.equal(200);
      });
    });
  });

  it("TableViewSizingCreate should succeed with a tableName.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/tableViewSizing"),
        headers: authHeaders(rootToken),
        body: {
          tableName: "checkendpointtests-table",
          sizes: `[{"columnName":"uniqueId","width":120}]`,
        },
      }).then(
        (
          response: Cypress.Response<
            SingleItemResponse<{ uniqueId: string; tableName: string }>
          >,
        ) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.tableName).to.equal(
            "checkendpointtests-table",
          );
          setShared("tvsId", response.body.data.item.uniqueId);
        },
      );
    });
  });

  it("TableViewSizingBrowse should include the just-created record.", () => {
    sharedValues(["rootToken", "tvsId"]).then(({ rootToken, tvsId }) => {
      cy.request({
        method: "GET",
        url: ui("/tableViewSizing/browse"),
        headers: authHeaders(rootToken),
      }).then((response: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        const ids = response.body.data.items.map((i) => i.uniqueId);
        expect(ids).to.include(tvsId);
      });
    });
  });

  it("TableViewSizingGet should return the same record by id.", () => {
    sharedValues(["rootToken", "tvsId"]).then(({ rootToken, tvsId }) => {
      cy.request({
        method: "GET",
        url: ui(`/tableViewSizing/${tvsId}`),
        headers: authHeaders(rootToken),
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.uniqueId).to.equal(tvsId);
      });
    });
  });

  it("TableViewSizingUpdate should persist a sizes change on an existing record.", () => {
    sharedValues(["rootToken", "tvsId"]).then(({ rootToken, tvsId }) => {
      cy.request({
        method: "PATCH",
        url: ui(`/tableViewSizing/${tvsId}`),
        headers: authHeaders(rootToken),
        body: { sizes: `[{"columnName":"uniqueId","width":240}]` },
      }).then((response: Cypress.Response<SingleItemResponse<{ sizes: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.sizes).to.contain("240");
      });
    });
  });

  // The behavior CommonListManager.tsx's column-resize handler actually depends on
  // (see the file-level comment above): the very first PATCH for a caller-chosen key
  // that doesn't exist yet must create the row, not 404.
  it("TableViewSizingUpdate should upsert (create) when the given uniqueId doesn't exist yet.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      const upsertKey = "checkendpointtests-upsert-key";
      cy.request({
        method: "PATCH",
        url: ui(`/tableViewSizing/${upsertKey}`),
        headers: authHeaders(rootToken),
        body: {
          tableName: "checkendpointtests-upsert-table",
          sizes: `[{"columnName":"uniqueId","width":100}]`,
        },
      }).then((response) => {
        expect(response.status).to.equal(200);
      });

      cy.request({
        method: "GET",
        url: ui(`/tableViewSizing/${upsertKey}`),
        headers: authHeaders(rootToken),
      }).then((getResponse) => {
        expect(getResponse.status).to.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/tableViewSizing/delete"),
        headers: authHeaders(rootToken),
        body: { uniqueIds: [upsertKey] },
      });
    });
  });

  it("TableViewSizingAwareDeletePreview then TableViewSizingAwareDelete should remove the record.", () => {
    sharedValues(["rootToken", "tvsId"]).then(({ rootToken, tvsId }) => {
      cy.request({
        method: "GET",
        url: ui(`/tableViewSizing/delete-preview?uniqueIds=${tvsId}`),
        headers: authHeaders(rootToken),
      }).then((preview) => {
        expect(preview.status).to.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/tableViewSizing/delete"),
        headers: authHeaders(rootToken),
        body: { uniqueIds: [tvsId] },
      }).then((del) => {
        expect(del.status).to.equal(200);
      });

      cy.request({
        method: "GET",
        url: ui(`/tableViewSizing/${tvsId}`),
        headers: authHeaders(rootToken),
        failOnStatusCode: false,
      }).then((getAfterDelete) => {
        expect(getAfterDelete.status).to.not.equal(200);
      });
    });
  });

  endFirebackServer();
});
