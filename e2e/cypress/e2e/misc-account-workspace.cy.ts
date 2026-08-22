import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers Abac.emi.yml's remaining standalone (non-CRUD) actions: CreateWorkspace,
// QueryUserRoleWorkspaces, QueryWorkspaceTypesPublicly, UserInvitations, UserPassports,
// GsmSendSms. None of these have a dedicated admin-CRUD-style screen, so - same as
// core-session.cy.ts/alt-auth-methods.cy.ts - they're driven directly over HTTP.
//
// This spec is also the regression guard for two crash-on-every-call bugs this pass
// fixed:
//  - CreateWorkspaceAndAssignUser (WorkspaceCoreFeatures.go) called q.Tx.Model(...)
//    directly - q.Tx is nil whenever CreateWorkspaceAction calls it outside a
//    transaction, so every POST /workspaces/create panicked the request.
//  - GsmSendSMSUsingNotificationConfig (GsmProviderActions.go) dereferenced a nil
//    *NotificationConfigEntity before its own nil check, so every POST /gsm/send/sms
//    panicked whenever no NotificationConfig row existed yet (the default state on any
//    fresh install).
interface SingleItemResponse<T> {
  data: { item: T };
}
interface ListResponse<T> {
  data: { items: T[] };
}
interface SignupResponse {
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

function signupFreshAccount(label: string) {
  return cy
    .task(
      "exec",
      ` role create --name "checkendpointtests ${label} role" --capabilities-list-id '["root.abac.email-confirmation.query"]'`,
    )
    .then((content: string) => {
      const roleId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      return cy
        .task(
          "exec",
          ` ws type c  --title "checkendpointtests ${label} type" --slug /checkendpointtests-${label} --role-id ${roleId}`,
        )
        .then((wtContent: string) => {
          const workspaceTypeId = (
            JSON.parse(wtContent) as SingleItemResponse<{ uniqueId: string }>
          ).data.item.uniqueId;
          const email = `checkendpointtests-${label}-${Date.now()}@example.com`;
          return cy
            .request({
              method: "POST",
              url: ui("/passports/signup/classic"),
              body: {
                value: email,
                type: "email",
                password: "checkendpointtests-pass-123",
                firstName: "Checkendpointtests",
                lastName: label,
                workspaceTypeId,
              },
            })
            .then((response: Cypress.Response<SignupResponse>) => {
              expect(response.status).to.equal(200);
              const token = response.body.data.item.session.token;
              // See alt-auth-methods.cy.ts's identical comment: signup sets an
              // "authorization" cookie that cy.request() (unlike Go's cookie-jar-less
              // http.Client) would otherwise keep resending, silently overriding any
              // explicit Authorization header used for a *different* account later in
              // the same test.
              return cy
                .clearCookie("authorization")
                .then(() => ({ email, token }));
            });
        });
    });
}

describe("Abac: misc standalone account/workspace actions", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for signup.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  it("CreateWorkspace should succeed and link the creator to it, and reject an unauthenticated call.", () => {
    cy.request({
      method: "POST",
      url: ui("/workspaces/create"),
      body: { name: "checkendpointtests anon workspace" },
      failOnStatusCode: false,
    }).then((anon) => {
      expect(anon.status).to.not.equal(200);
    });

    signupFreshAccount("create-workspace").then(({ token }) => {
      cy.request({
        method: "POST",
        url: ui("/workspaces/create"),
        headers: { authorization: token },
        body: { name: "checkendpointtests new workspace" },
      }).then(
        (
          response: Cypress.Response<
            SingleItemResponse<{ workspaceId: string }>
          >,
        ) => {
          expect(response.status).to.equal(200);
          const newWorkspaceId = response.body.data.item.workspaceId;
          expect(newWorkspaceId, "workspaceId in the response").to.be.a(
            "string",
          ).and.not.be.empty;

          cy.request({
            method: "GET",
            url: ui("/urw/query"),
            headers: { authorization: token },
          }).then(
            (urw: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
              expect(urw.status).to.equal(200);
              const ids = urw.body.data.items.map((w) => w.uniqueId);
              expect(ids).to.include(newWorkspaceId);
            },
          );
        },
      );
    });
  });

  it("QueryUserRoleWorkspaces should include the signup's own workspace.", () => {
    signupFreshAccount("urw-query").then(({ token }) => {
      cy.request({
        method: "GET",
        url: ui("/urw/query"),
        headers: { authorization: token },
      }).then(
        (response: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.items.length).to.be.greaterThan(0);
        },
      );
    });
  });

  it("QueryWorkspaceTypesPublicly should include a custom type and exclude root, with no auth.", () => {
    cy.task(
      "exec",
      ` role create --name "checkendpointtests public-types role" --capabilities-list-id '["root.abac.email-confirmation.query"]'`,
    ).then((content: string) => {
      const roleId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      cy.task(
        "exec",
        ` ws type c  --title "checkendpointtests public-types type" --slug /checkendpointtests-public-types --role-id ${roleId}`,
      ).then((wtContent: string) => {
        const customId = (
          JSON.parse(wtContent) as SingleItemResponse<{ uniqueId: string }>
        ).data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/workspace/public/types"),
        }).then(
          (response: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
            expect(response.status).to.equal(200);
            const ids = response.body.data.items.map((w) => w.uniqueId);
            expect(ids).to.include(customId);
            expect(ids).to.not.include("root");
          },
        );
      });
    });
  });

  it("UserInvitations should succeed (even with none pending) and reject an unauthenticated call.", () => {
    signupFreshAccount("user-invitations").then(({ token }) => {
      cy.request({
        method: "GET",
        url: ui("/users/invitations"),
        headers: { authorization: token },
      }).then((response) => {
        expect(response.status).to.equal(200);
      });
    });

    cy.request({
      method: "GET",
      url: ui("/users/invitations"),
      failOnStatusCode: false,
    }).then((anon) => {
      expect(anon.status).to.not.equal(200);
    });
  });

  it("UserPassports should return the signed-up account's own passport.", () => {
    signupFreshAccount("user-passports").then(({ email, token }) => {
      cy.request({
        method: "GET",
        url: ui("/user/passports"),
        headers: { authorization: token },
      }).then(
        (
          response: Cypress.Response<
            ListResponse<{ value: string; type: string }>
          >,
        ) => {
          expect(response.status).to.equal(200);
          const values = response.body.data.items.map((p) => p.value);
          expect(values).to.include(email);
        },
      );
    });
  });

  // Regression guard for the nil-config crash - see this spec's header comment.
  it("GsmSendSms should succeed (falling back to a terminal queue with no gsm provider configured) and reject missing fields.", () => {
    cy.request({
      method: "POST",
      url: ui("/gsm/send/sms"),
      body: { toNumber: "+15550000000", body: "checkendpointtests sms" },
    }).then((response) => {
      expect(response.status).to.equal(200);
    });

    cy.request({
      method: "POST",
      url: ui("/gsm/send/sms"),
      body: { body: "checkendpointtests sms" },
      failOnStatusCode: false,
    }).then((missingToNumber) => {
      expect(missingToNumber.status).to.not.equal(200);
    });

    cy.request({
      method: "POST",
      url: ui("/gsm/send/sms"),
      body: { toNumber: "+15550000000" },
      failOnStatusCode: false,
    }).then((missingBody) => {
      expect(missingBody.status).to.not.equal(200);
    });
  });

  endFirebackServer();
});
