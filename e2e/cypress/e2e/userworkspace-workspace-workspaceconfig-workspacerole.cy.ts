import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers 4 of Abac.emi.yml's CRUD entities: UserWorkspace, Workspace, WorkspaceConfig,
// WorkspaceRole (WorkspaceType has its own dedicated spec - workspace-type.cy.ts - and
// WorkspaceInvite its own too - workspace-invite.cy.ts). None of these four have a
// dedicated manage-UI admin screen, so all are exercised directly over HTTP, same as the
// other CRUD-entity specs in this pass.
//
// Workspace/WorkspaceType/WorkspaceConfig Create/Update/AwareDelete/AwareDeletePreview
// all have AllowOnRoot:true, and UserWorkspace/WorkspaceRole browse resolve by the
// caller's own userId/workspace_id respectively (see the matching *_test.go Go tests'
// comments) - every write below runs in the literal "root" workspace, same as the Go
// suite.
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
interface WhoamiResponse {
  data: { item: { userId: string } };
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

// letters base-26-encodes n into a lowercase a-z string (no digits/dashes) - workspace
// type slugs are restricted to lowercase letters and dashes only, same reasoning as
// workspace_invite_accept_test.go's lettersFromInt (Go side).
function letters(n: number): string {
  if (n <= 0) return "a";
  let s = "";
  while (n > 0) {
    s = String.fromCharCode(97 + (n % 26)) + s;
    n = Math.floor(n / 26);
  }
  return s;
}

describe("Abac: UserWorkspace, Workspace, WorkspaceConfig, WorkspaceRole", () => {
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

    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "GET",
        url: ui("/whoami"),
        headers: rootHeaders(rootToken),
      }).then((whoami: Cypress.Response<WhoamiResponse>) => {
        expect(whoami.status).to.equal(200);
        setShared("rootUserId", whoami.body.data.item.userId);
      });
    });
  });

  // createWorkspaceChain builds Role -> WorkspaceType -> Workspace, same dependency
  // order as workspace_test.go's createSampleWorkspace, and yields the ids the rest of
  // this spec chains off.
  function createWorkspaceChain(rootToken: string) {
    return cy
      .request({
        method: "POST",
        url: ui("/role"),
        headers: rootHeaders(rootToken),
        body: {
          name: `checkendpointtests ws-chain role ${Date.now()}`,
          capabilitiesListId: ["root.abac.email-confirmation.query"],
        },
      })
      .then((roleRes: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
        expect(roleRes.status).to.equal(200);
        const roleId = roleRes.body.data.item.uniqueId;

        return cy
          .request({
            method: "POST",
            url: ui("/workspaceType"),
            headers: rootHeaders(rootToken),
            body: {
              title: "checkendpointtests ws-chain type",
              slug: `/checkendpointtests-wsc-${letters(Date.now())}`,
              roleId,
            },
          })
          .then((wtRes: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
            expect(wtRes.status).to.equal(200);
            const typeId = wtRes.body.data.item.uniqueId;

            return cy
              .request({
                method: "POST",
                url: ui("/workspace"),
                headers: rootHeaders(rootToken),
                body: { name: "checkendpointtests ws-chain workspace", typeId },
              })
              .then((wsRes: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
                expect(wsRes.status).to.equal(200);
                return {
                  roleId,
                  typeId,
                  workspaceId: wsRes.body.data.item.uniqueId,
                };
              });
          });
      });
  }

  function deleteWorkspaceChain(
    rootToken: string,
    chain: { roleId: string; typeId: string; workspaceId: string },
  ) {
    cy.request({
      method: "POST",
      url: ui("/workspace/delete"),
      headers: rootHeaders(rootToken),
      body: { uniqueIds: [chain.workspaceId] },
    });
    cy.request({
      method: "POST",
      url: ui("/workspaceType/delete"),
      headers: rootHeaders(rootToken),
      body: { uniqueIds: [chain.typeId] },
    });
    cy.request({
      method: "POST",
      url: ui("/role/delete"),
      headers: rootHeaders(rootToken),
      body: { uniqueIds: [chain.roleId] },
    });
  }

  it("Workspace CRUD should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      createWorkspaceChain(rootToken).then((chain) => {
        cy.request({
          method: "GET",
          url: ui("/workspace/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(chain.workspaceId);
        });

        cy.request({
          method: "PATCH",
          url: ui(`/workspace/${chain.workspaceId}`),
          headers: rootHeaders(rootToken),
          body: { name: "checkendpointtests ws-chain workspace renamed" },
        }).then((update: Cypress.Response<SingleItemResponse<{ name: string }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.name).to.equal(
            "checkendpointtests ws-chain workspace renamed",
          );
        });

        cy.request({
          method: "GET",
          url: ui(`/workspace/delete-preview?uniqueIds=${chain.workspaceId}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        deleteWorkspaceChain(rootToken, chain);
      });
    });
  });

  it("UserWorkspace CRUD should work end to end.", () => {
    sharedValues(["rootToken", "rootUserId"]).then(({ rootToken, rootUserId }) => {
      createWorkspaceChain(rootToken).then((chain) => {
        cy.request({
          method: "POST",
          url: ui("/userWorkspace"),
          headers: rootHeaders(rootToken),
          body: { userId: rootUserId, workspaceId: chain.workspaceId },
        }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; workspaceId: string }>>) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.workspaceId).to.equal(chain.workspaceId);
          const id = response.body.data.item.uniqueId;

          // Root's own userId is required here: UserWorkspaceActions.go's security uses
          // ResolveStrategy "user", so browse/query filters by user_id = the caller's
          // own userId, not by workspace - a row attached to some other user would
          // never show up in root's own browse.
          cy.request({
            method: "GET",
            url: ui("/userWorkspace/browse?itemsPerPage=1000"),
            headers: rootHeaders(rootToken),
          }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
            expect(browse.status).to.equal(200);
            expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
          });

          cy.request({
            method: "GET",
            url: ui(`/userWorkspace/${id}`),
            headers: rootHeaders(rootToken),
          }).then((get) => {
            expect(get.status).to.equal(200);
          });

          cy.request({
            method: "GET",
            url: ui(`/userWorkspace/delete-preview?uniqueIds=${id}`),
            headers: rootHeaders(rootToken),
          }).then((preview) => {
            expect(preview.status).to.equal(200);
          });

          cy.request({
            method: "POST",
            url: ui("/userWorkspace/delete"),
            headers: rootHeaders(rootToken),
            body: { uniqueIds: [id] },
          }).then((del) => {
            expect(del.status).to.equal(200);
          });
        });

        deleteWorkspaceChain(rootToken, chain);
      });
    });
  });

  it("WorkspaceRole CRUD should work end to end.", () => {
    sharedValues(["rootToken", "rootUserId"]).then(({ rootToken, rootUserId }) => {
      createWorkspaceChain(rootToken).then((chain) => {
        cy.request({
          method: "POST",
          url: ui("/userWorkspace"),
          headers: rootHeaders(rootToken),
          body: { userId: rootUserId, workspaceId: chain.workspaceId },
        }).then((uwRes: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
          expect(uwRes.status).to.equal(200);
          const userWorkspaceId = uwRes.body.data.item.uniqueId;

          cy.request({
            method: "POST",
            url: ui("/role"),
            headers: rootHeaders(rootToken),
            body: {
              name: `checkendpointtests wr role ${Date.now()}`,
              capabilitiesListId: ["root.abac.email-confirmation.query"],
            },
          }).then((roleRes: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>) => {
            expect(roleRes.status).to.equal(200);
            const roleId = roleRes.body.data.item.uniqueId;

            cy.request({
              method: "POST",
              url: ui("/workspaceRole"),
              headers: rootHeaders(rootToken),
              body: { userWorkspaceId, roleId },
            }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; roleId: string }>>) => {
              expect(response.status).to.equal(200);
              expect(response.body.data.item.roleId).to.equal(roleId);
              const id = response.body.data.item.uniqueId;

              // Regression guard for WorkspaceRoleCreateAction's workspace-stamping fix.
              cy.request({
                method: "GET",
                url: ui("/workspaceRole/browse?itemsPerPage=1000"),
                headers: rootHeaders(rootToken),
              }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
                expect(browse.status).to.equal(200);
                expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
              });

              cy.request({
                method: "GET",
                url: ui(`/workspaceRole/delete-preview?uniqueIds=${id}`),
                headers: rootHeaders(rootToken),
              }).then((preview) => {
                expect(preview.status).to.equal(200);
              });

              cy.request({
                method: "POST",
                url: ui("/workspaceRole/delete"),
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

          cy.request({
            method: "POST",
            url: ui("/userWorkspace/delete"),
            headers: rootHeaders(rootToken),
            body: { uniqueIds: [userWorkspaceId] },
          });
        });

        deleteWorkspaceChain(rootToken, chain);
      });
    });
  });

  it("WorkspaceConfig CRUD, and the /workspace-config/distinct endpoints, should work end to end.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/workspaceConfig"),
        headers: rootHeaders(rootToken),
        body: { enableTotp: true },
      }).then((response: Cypress.Response<SingleItemResponse<{ uniqueId: string; enableTotp: boolean }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.enableTotp).to.equal(true);
        const id = response.body.data.item.uniqueId;

        cy.request({
          method: "GET",
          url: ui("/workspaceConfig/browse?itemsPerPage=1000"),
          headers: rootHeaders(rootToken),
        }).then((browse: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
          expect(browse.status).to.equal(200);
          expect(browse.body.data.items.map((i) => i.uniqueId)).to.include(id);
        });

        cy.request({
          method: "GET",
          url: ui(`/workspaceConfig/delete-preview?uniqueIds=${id}`),
          headers: rootHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/workspaceConfig/delete"),
          headers: rootHeaders(rootToken),
          body: { uniqueIds: [id] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      });

      // The two /workspace-config/distinct endpoints always operate on the single, real
      // workspace_id='root' row shared by the whole disposable server this spec started
      // (see withFirebackServer()) - captured and restored, same reasoning as
      // workspace_config_test.go's Go equivalents, though here it's a throwaway server
      // torn down at the end of this spec either way.
      cy.request({
        method: "GET",
        url: ui("/workspace-config/distinct"),
        headers: rootHeaders(rootToken),
      }).then((original: Cypress.Response<SingleItemResponse<{ enableTotp: boolean | null }>>) => {
        expect(original.status).to.equal(200);
        const originalEnableTotp = !!original.body.data.item.enableTotp;
        const toggled = !originalEnableTotp;

        cy.request({
          method: "PATCH",
          url: ui("/workspace-config/distinct"),
          headers: rootHeaders(rootToken),
          body: { enableTotp: toggled },
        }).then((update: Cypress.Response<SingleItemResponse<{ enableTotp: boolean }>>) => {
          expect(update.status).to.equal(200);
          expect(update.body.data.item.enableTotp).to.equal(toggled);
        });

        cy.request({
          method: "GET",
          url: ui("/workspace-config/distinct"),
          headers: rootHeaders(rootToken),
        }).then((reGet: Cypress.Response<SingleItemResponse<{ enableTotp: boolean }>>) => {
          expect(reGet.status).to.equal(200);
          expect(reGet.body.data.item.enableTotp).to.equal(toggled);
        });

        cy.request({
          method: "PATCH",
          url: ui("/workspace-config/distinct"),
          headers: rootHeaders(rootToken),
          body: { enableTotp: originalEnableTotp },
        });
      });
    });
  });

  endFirebackServer();
});
