import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers modules/abac/interfacetools's "appMenu" entity end to end - it has no
// dedicated admin CRUD screen (see InterfaceToolsModule.go's own doc comment: "They
// have no dedicated front-end management screens ... so only the backend moved here"),
// so every action is exercised directly over HTTP, the same way tools/checkendpointtests
// itself only cares that *some* test - Go or otherwise - covers each action. The one
// genuinely UI-driven check here is CteAppMenus/the sidebar at the bottom of this file:
// AppMenu rows are real, live sidebar configuration (see
// fireback-ui/hooks/useRemoteMenuResolver.tsx), so that part *is* a real front-end
// screen, just not an admin CRUD one.
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

describe("interfacetools: appMenu", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

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

  it("AppMenuCreate should succeed, and the record should be invisible without auth.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/appMenu"),
        body: {
          label: "checkendpointtests appmenu",
          href: "/checkendpointtests",
        },
        failOnStatusCode: false,
      }).then((noAuth) => {
        expect(noAuth.status).to.not.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/appMenu"),
        headers: authHeaders(rootToken),
        body: {
          label: "checkendpointtests appmenu",
          href: "/checkendpointtests",
        },
      }).then(
        (
          response: Cypress.Response<
            SingleItemResponse<{ uniqueId: string; href: string }>
          >,
        ) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.href).to.equal("/checkendpointtests");
          setShared("appMenuId", response.body.data.item.uniqueId);
        },
      );
    });
  });

  it("AppMenuBrowse should include the just-created record.", () => {
    sharedValues(["rootToken", "appMenuId"]).then(
      ({ rootToken, appMenuId }) => {
        cy.request({
          method: "GET",
          url: ui("/appMenu/browse?itemsPerPage=100"),
          headers: authHeaders(rootToken),
        }).then(
          (response: Cypress.Response<ListResponse<{ uniqueId: string }>>) => {
            expect(response.status).to.equal(200);
            const ids = response.body.data.items.map((i) => i.uniqueId);
            expect(ids).to.include(appMenuId);
          },
        );
      },
    );
  });

  it("AppMenuGet should return the same record by id.", () => {
    sharedValues(["rootToken", "appMenuId"]).then(
      ({ rootToken, appMenuId }) => {
        cy.request({
          method: "GET",
          url: ui(`/appMenu/${appMenuId}`),
          headers: authHeaders(rootToken),
        }).then(
          (
            response: Cypress.Response<
              SingleItemResponse<{ uniqueId: string }>
            >,
          ) => {
            expect(response.status).to.equal(200);
            expect(response.body.data.item.uniqueId).to.equal(appMenuId);
          },
        );
      },
    );
  });

  it("AppMenuUpdate should persist a field change.", () => {
    sharedValues(["rootToken", "appMenuId"]).then(
      ({ rootToken, appMenuId }) => {
        cy.request({
          method: "PATCH",
          url: ui(`/appMenu/${appMenuId}`),
          headers: authHeaders(rootToken),
          body: { href: "/checkendpointtests-renamed" },
        }).then(
          (
            response: Cypress.Response<SingleItemResponse<{ href: string }>>,
          ) => {
            expect(response.status).to.equal(200);
            expect(response.body.data.item.href).to.equal(
              "/checkendpointtests-renamed",
            );
          },
        );
      },
    );
  });

  it("AppMenuAwareDeletePreview then AppMenuAwareDelete should remove the record.", () => {
    sharedValues(["rootToken", "appMenuId"]).then(
      ({ rootToken, appMenuId }) => {
        cy.request({
          method: "GET",
          url: ui(`/appMenu/delete-preview?uniqueIds=${appMenuId}`),
          headers: authHeaders(rootToken),
        }).then((preview) => {
          expect(preview.status).to.equal(200);
        });

        cy.request({
          method: "POST",
          url: ui("/appMenu/delete"),
          headers: authHeaders(rootToken),
          body: { uniqueIds: [appMenuId] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });

        cy.request({
          method: "GET",
          url: ui(`/appMenu/${appMenuId}`),
          headers: authHeaders(rootToken),
          failOnStatusCode: false,
        }).then((getAfterDelete) => {
          expect(getAfterDelete.status).to.not.equal(200);
        });
      },
    );
  });

  it("cte-app-menus should nest a child row under its parent's children array.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/appMenu"),
        headers: authHeaders(rootToken),
        body: {
          label: "checkendpointtests cte parent",
          href: "/checkendpointtests/cte-parent",
        },
      }).then(
        (
          parentResp: Cypress.Response<
            SingleItemResponse<{ uniqueId: string }>
          >,
        ) => {
          const parentId = parentResp.body.data.item.uniqueId;
          setShared("cteParentId", parentId);

          cy.request({
            method: "POST",
            url: ui("/appMenu"),
            headers: authHeaders(rootToken),
            body: {
              label: "checkendpointtests cte child",
              href: "/checkendpointtests/cte-child",
              parentId,
            },
          }).then(
            (
              childResp: Cypress.Response<
                SingleItemResponse<{ uniqueId: string }>
              >,
            ) => {
              const childId = childResp.body.data.item.uniqueId;
              setShared("cteChildId", childId);

              cy.request({
                method: "GET",
                url: ui("/cte-app-menus"),
                headers: authHeaders(rootToken),
              }).then(
                (
                  cteResp: Cypress.Response<ListResponse<Record<string, any>>>,
                ) => {
                  expect(cteResp.status).to.equal(200);
                  const parentNode = cteResp.body.data.items.find(
                    (item) => item.uniqueId === parentId,
                  );
                  expect(parentNode, "parent node present at the top level").to
                    .exist;
                  const childIds = (parentNode?.children || []).map(
                    (c: any) => c.uniqueId,
                  );
                  expect(childIds).to.include(childId);
                },
              );
            },
          );
        },
      );
    });
  });

  it("cleans up the cte-app-menus fixture rows.", () => {
    sharedValues(["rootToken", "cteParentId", "cteChildId"]).then(
      ({ rootToken, cteParentId, cteChildId }) => {
        cy.request({
          method: "POST",
          url: ui("/appMenu/delete"),
          headers: authHeaders(rootToken),
          body: { uniqueIds: [cteChildId, cteParentId] },
        }).then((del) => {
          expect(del.status).to.equal(200);
        });
      },
    );
  });

  // --- The one genuinely UI-driven check: a no-capability appMenu row is real sidebar
  // configuration (see useRemoteMenuResolver.tsx/Sidebar.tsx's dataMenuToMenu -
  // visibilityCheck lets any logged-in user see an item with no capabilityId), so
  // logging in for real and finding it in the rendered sidebar proves CteAppMenus round
  // -trips through the actual front-end, not just its own HTTP response shape. ---

  it("a menu item with no capability requirement should render in the manage UI's sidebar after login.", () => {
    sharedValues(["rootToken"]).then(({ rootToken }) => {
      cy.request({
        method: "POST",
        url: ui("/appMenu"),
        headers: authHeaders(rootToken),
        body: {
          label: "checkendpointtests sidebar item",
          href: "/checkendpointtests/sidebar-item",
          icon: "home",
        },
      }).then(
        (
          response: Cypress.Response<SingleItemResponse<{ uniqueId: string }>>,
        ) => {
          setShared("sidebarMenuId", response.body.data.item.uniqueId);
        },
      );
    });

    // Sidebar/content-section layout only renders at desktop widths - see
    // deviceInformation.tsx's isMobileView.
    cy.viewport(1366, 900);
    Cypress.on("uncaught:exception", () => false);

    cy.session("root-login-appmenu", () => {
      cy.visit(ui("/manage/#/en/welcome"));
      cy.get("#value-input", { timeout: 10000 }).type(ROOT_EMAIL);
      cy.get("#submit-form").click({ force: true });
      cy.get("h1", { timeout: 10000 }).should("have.text", "Enter Password");
      cy.get("#password-input").type(ROOT_PASSWORD);
      cy.get("#submit-form").click({ force: true });
      cy.wait(1000);
      cy.get("body").then(($b) => {
        if ($b.text().includes("Select workspace")) {
          cy.contains("button", "Root Access").click({ force: true });
        }
      });
      cy.url({ timeout: 10000 }).should("include", "/dashboard");
    });

    cy.visit(ui("/manage/#/en/dashboard"));
    cy.contains(".nav-link-text", "checkendpointtests sidebar item", {
      timeout: 10000,
    }).should("exist");

    sharedValues(["rootToken", "sidebarMenuId"]).then(
      ({ rootToken, sidebarMenuId }) => {
        cy.request({
          method: "POST",
          url: ui("/appMenu/delete"),
          headers: authHeaders(rootToken),
          body: { uniqueIds: [sidebarMenuId] },
        });
      },
    );
  });

  endFirebackServer();
});
