import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers 5 of Abac.emi.yml's CRUD entities: RegionalContent, Role, Token, User,
// UserProfile. Role (selfservice/roles) and User (manage/users) and RegionalContent
// (manage/regional-content) each have a dedicated manage-UI screen - checked, all three
// wire their queryHook/deleteHook to the matching entity's own SDK hooks, no repeat of
// the PassportMethodList copy-paste bug. Token and UserProfile have no dedicated screen,
// so (like PendingWorkspaceInvite/PhoneConfirmation/etc in the previous batch) all five
// are exercised directly over HTTP here.
//
// User/Token/Role/UserProfile Create/Update/AwareDelete/AwareDeletePreview all have
// AllowOnRoot:true (see the matching *_test.go Go tests), so every write below runs in
// the literal "root" workspace, same as the Go suite.
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

describe("Abac: RegionalContent, Role, Token, User, UserProfile", () => {
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

  it("RegionalContent CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      const keyGroup = `checkendpointtests-rc-${Date.now()}`;
      cy.request({
        method: "POST",
        url: ui("/regionalContent"),
        headers: rootHeaders(rootToken),
        body: { content: "checkendpointtests content", region: "any", languageId: "en", keyGroup },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; content: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/regionalContent/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/regionalContent/${id}`),
          headers: rootHeaders(rootToken),
          body: { content: "checkendpointtests content updated" },
        }).then((update: Cypress.Response<SingleItemResponse<{ content: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.content).to.equal("checkendpointtests content updated");
        });

        cy.request({
          method: "POST",
          url: ui("/regionalContent/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("Role CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/role"),
        headers: rootHeaders(rootToken),
        body: {
          name: `checkendpointtests role ${Date.now()}`,
          capabilitiesListId: ["root.abac.email-confirmation.query"],
        },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; name: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/role/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        // RoleOptionalDto.CapabilitiesListId has no "was this key present" tracking, so
        // an update must always resupply it (see role_crud_test.go's TestRoleUpdate_HTTP_
        // Succeeds comment) or it's rejected by RoleNeedsOneCapability.
        cy.request({
          method: "PATCH",
          url: ui(`/role/${id}`),
          headers: rootHeaders(rootToken),
          body: {
            name: `checkendpointtests role renamed ${Date.now()}`,
            capabilitiesListId: ["root.abac.email-confirmation.query"],
          },
        }).then((update: Cypress.Response<SingleItemResponse<{ name: string }>>) => {
          expect(update.status).to.equal(200);
        });

        cy.request({
          method: "GET",
          url: ui(`/role/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/role/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("Token CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/token"),
        headers: rootHeaders(rootToken),
        body: { token: `checkendpointtests-token-${Date.now()}` },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; token: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        // Regression guard for TokenCreateAction's workspace-stamping fix.
        cy.request({
          method: "GET",
          url: ui("/token/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/token/${id}`),
          headers: rootHeaders(rootToken),
          body: { token: `checkendpointtests-token-updated-${Date.now()}` },
        }).then((update: Cypress.Response<SingleItemResponse<{ token: string }>>) => {
          expect(update.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/token/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("User CRUD should work end to end, and reject a missing required field.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      // Regression guard for UserCreateAction now actually calling
      // fireback.CommonStructValidatorPointer (firstName/lastName are validate:"required").
      cy.request({
        method: "POST",
        url: ui("/user"),
        headers: rootHeaders(rootToken),
        body: { firstName: "OnlyFirstName" },
        failOnStatusCode: false,
      }).then((rejected) => {
        expect(rejected.status).to.not.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/user"),
        headers: rootHeaders(rootToken),
        body: { firstName: "Checkendpointtests", lastName: `User${Date.now()}` },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; firstName: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.firstName).to.equal("Checkendpointtests");
        const id = response.body.data.item.uniqueId;

        // Regression guard for the UserActionCreate shared-helper workspace-stamping fix
        // that (before the fix) meant NO user, anywhere, ever showed up in this listing.
        cy.request({
          method: "GET",
          url: ui("/user/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/user/${id}`),
          headers: rootHeaders(rootToken),
          body: { title: "checkendpointtests title" },
        }).then((update) => {
          expect(update.status).to.equal(200);
        });

        cy.request({
          method: "GET",
          url: ui(`/user/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/user/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  it("UserProfile CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/userProfile"),
        headers: rootHeaders(rootToken),
        body: { firstName: "Checkendpointtests", lastName: "Profile" },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; firstName: string }>>) => {
        expect(response.status).to.equal(200);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/userProfile/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/userProfile/${id}`),
          headers: rootHeaders(rootToken),
          body: { firstName: "Renamed" },
        }).then((update: Cypress.Response<SingleItemResponse<{ firstName: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.firstName).to.equal("Renamed");
        });

        cy.request({
          method: "GET",
          url: ui(`/userProfile/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/userProfile/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });
    });
  });

  endFirebackServer();
});
