import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// The end-to-end "admin invites, invitee accepts" flow the fix in
// InviteToWorkspaceActionImplementation.go/AcceptInviteActionImplementation.go was for:
// an admin (a non-root user whose only special power is the workspace-invite create
// capability - not root's implicit wildcard) invites a new person by email into their
// workspace with a role through the UI, and that person - once they have an account of
// their own - logs in, sees the pending invitation, and accepts it.
//
// Root's own CLI account, created by withFirebackServer().
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

const ADMIN_EMAIL = "admin-checkendpointtests@example.com";
const ADMIN_PASSWORD = "checkendpointtests-pass-123";

const INVITEE_EMAIL = "invitee-checkendpointtests@example.com";
const INVITEE_PASSWORD = "checkendpointtests-pass-123";

const INVITEE_ROLE_NAME = "checkendpointtests invitee role";

interface SingleItemResponse<T> {
  data: { item: T };
}
interface SignupResponse {
  data: {
    item: {
      session: {
        token: string;
        userWorkspaces: { workspaceId: string }[];
      };
    };
  };
}
interface WhoamiResponse {
  data: {
    item: {
      workspaces: { uniqueId: string; name: string }[];
    };
  };
}

// cy.session() (used by loginAs, below) reloads the whole spec runner the
// first time each named session is established, which resets any in-browser
// JS state - plain `let`s declared at describe-scope, and even Cypress.env()
// overrides set at runtime - back to its initial value. That makes both
// unsafe for anything that needs to be read in a *later* it() than the next
// loginAs() call: the value looks correctly set right up until that reload,
// then silently reads back empty/undefined. setShared/sharedValues (backed
// by the setShared/getSharedState tasks in cypress.config.js) stash the
// value in the plugin's Node process instead, which nothing in the browser
// can reload out from under this spec.
function setShared(key: string, value: string) {
  return cy.task("setShared", { key, value });
}
function sharedValues<K extends string>(
  keys: K[],
): Cypress.Chainable<Record<K, string>> {
  return cy.task("getSharedState", keys);
}

function loginAs(sessionKey: string, email: string, password: string) {
  cy.viewport(1366, 900);
  Cypress.on("uncaught:exception", () => false);

  cy.session(sessionKey, () => {
    // Only the "email" passport method is registered, so /welcome auto-redirects
    // straight to the email-entry screen instead of showing a method-choice screen.
    cy.visit(ui("/manage/#/welcome"));
    cy.get("#value-input", { timeout: 10000 }).type(email);
    cy.get("#submit-form").click({ force: true });
    cy.get("h1", { timeout: 10000 }).should("have.text", "Enter Password");
    cy.get("#password-input").type(password);
    cy.get("#submit-form").click({ force: true });
    cy.wait(1000);
    // Root (via withFirebackServer()'s `auth --in-root=true --workspace-type-id root`)
    // ends up belonging to two workspaces and hits a picker; the admin/invitee accounts
    // this spec signs up each pick exactly one workspaceTypeId, so they never do - this
    // is only here so loginAs stays reusable for root too, not because it is expected
    // to trigger for the admin/invitee calls below.
    cy.get("body").then(($b) => {
      if ($b.text().includes("Select workspace")) {
        cy.contains("button", "Root Access").click({ force: true });
      }
    });
    cy.url({ timeout: 10000 }).should("include", "/dashboard");
  });
}

function visitForm(url: string) {
  cy.visit(url);
  cy.get(".content-section fieldset").should("not.be.disabled");
}

function selectFormSelectOption(labelText: string, optionText: string) {
  cy.contains(".mb-3", labelText).find(".form-control").click();
  cy.contains(".mb-3", labelText)
    .find("input[role=combobox]")
    .type(optionText, { delay: 50 });
  cy.contains(".react-select-menu-area [role=option]", optionText, {
    timeout: 5000,
  }).click();
}

function submitForm() {
  cy.get(".action-menu").contains("Save").click({ force: true });
}

describe("Workspace invite: an admin (capability-driven, not root) invites someone, and they accept", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for login/signup.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  // --- Admin setup: a role carrying only the workspace-invite create capability, a
  // workspace type bound to it, and the email-sending config InviteToWorkspaceAction
  // needs (a terminal EmailProvider - SendMail's terminal case just logs and returns
  // nil, no real SMTP needed - plus an EmailSender and a NotificationConfig wiring both
  // for the root workspace, which is where the admin's own workspace-invite membership
  // will resolve to). All via the CLI, as root, mirroring
  // workspace-role-capability-scoping.cy.ts's own setup style. ---

  it("should be able to set up an admin role/workspace type, an invitee-signup role, and email sending as root.", () => {
    // The capability's CompleteKey (see NewCrudPermissionSet/workspaceInvitePerms in
    // WorkspaceInviteActions.go): "root.modules" + ".abac." + "workspace-invite" +
    // ".create". This - not root's own "root.*" wildcard - is what
    // InviteToWorkspaceActionImplementation.go's fix actually gates the action on.
    // "root.modules.abac.role.*" is also granted so the admin can create a role of
    // their own to assign on the invite below (roles are workspace-scoped - see
    // RoleBrowseAction/RoleCreateAction - so one root creates here, in root's own
    // workspace, wouldn't even show up in the admin's own workspace's role picker).
    cy.task(
      "exec",
      ` role create --name "checkendpointtests admin role" --capabilities-list-id '["root.modules.abac.workspace-invite.create", "root.modules.abac.role.create", "root.modules.abac.role.query"]'`,
    ).then((content: string) => {
      const roleId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;

      cy.task(
        "exec",
        ` ws workspaceType-c --title "checkendpointtests admin type" --slug /checkendpointtests-admin --role-id ${roleId}`,
      ).then((wtContent: string) => {
        setShared(
          "adminWorkspaceTypeId",
          (JSON.parse(wtContent) as SingleItemResponse<{ uniqueId: string }>)
            .data.item.uniqueId,
        );
      });
    });

    cy.task(
      "exec",
      ` role create --name "checkendpointtests invitee own role" --capabilities-list-id '["root.abac.email-confirmation.query"]'`,
    ).then((content: string) => {
      const roleId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      cy.task(
        "exec",
        ` ws workspaceType-c --title "checkendpointtests invitee type" --slug /checkendpointtests-invitee --role-id ${roleId}`,
      ).then((wtContent: string) => {
        setShared(
          "inviteeOwnWorkspaceTypeId",
          (JSON.parse(wtContent) as SingleItemResponse<{ uniqueId: string }>)
            .data.item.uniqueId,
        );
      });
    });

    cy.task(
      "exec",
      ` messaging emailSender-c --from-name Checkendpointtests --from-email-address checkendpointtests-sender@example.com --reply-to checkendpointtests-sender@example.com --nick-name checkendpointtests`,
    ).then((content: string) => {
      const senderId = (
        JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;

      cy.task(
        "exec",
        ` messaging emailProvider-c --type terminal --title checkendpointtests`,
      ).then((providerContent: string) => {
        const providerId = (
          JSON.parse(providerContent) as SingleItemResponse<{
            uniqueId: string;
          }>
        ).data.item.uniqueId;

        cy.task(
          "exec",
          ` notification create --general-email-provider-id ${providerId} --invite-to-workspace-sender-id ${senderId}`,
        );
      });
    });
  });

  it("should be able to sign up the admin account against the admin workspace type.", () => {
    sharedValues(["adminWorkspaceTypeId"]).then(({ adminWorkspaceTypeId }) => {
      cy.request({
        method: "POST",
        url: ui("/passports/signup/classic"),
        body: {
          value: ADMIN_EMAIL,
          type: "email",
          password: ADMIN_PASSWORD,
          firstName: "Checkendpointtests",
          lastName: "Admin",
          workspaceTypeId: adminWorkspaceTypeId,
        },
      }).then((response: Cypress.Response<SignupResponse>) => {
        expect(response.status).to.equal(200);
        setShared("adminToken", response.body.data.item.session.token);
        const adminWorkspaceId =
          response.body.data.item.session.userWorkspaces[0].workspaceId;
        setShared("adminWorkspaceId", adminWorkspaceId);
        expect(adminWorkspaceId).to.be.a("string").and.not.be.empty;
      });
    });
  });

  it("should be able to, as the admin, create the role to assign on the invite.", () => {
    // Done over the API rather than the UI - this is setup for the scenario being
    // tested (an invite needs *some* role to assign), not itself part of what's being
    // proven. The admin's own capabilities only cover granting
    // "root.modules.abac.workspace-invite.create" (see filterRolePermissions in
    // RoleActions.go - a caller can only grant a capability their own access already
    // covers), so that's what this new role gets too.
    sharedValues(["adminToken", "adminWorkspaceId"]).then(
      ({ adminToken, adminWorkspaceId }) => {
        cy.request({
          method: "POST",
          url: ui("/role"),
          headers: {
            authorization: adminToken,
            "workspace-id": adminWorkspaceId,
          },
          body: {
            name: INVITEE_ROLE_NAME,
            capabilitiesListId: ["root.modules.abac.workspace-invite.create"],
          },
        }).then(
          (
            response: Cypress.Response<
              SingleItemResponse<{ uniqueId: string }>
            >,
          ) => {
            expect(response.status).to.equal(200);
            expect(response.body.data.item.uniqueId).to.be.a("string").and.not
              .be.empty;
          },
        );
      },
    );
  });

  // --- The admin logs into the manage UI and sends the invite ---

  it("should log into the manage UI as the admin and invite the invitee by email with a role.", () => {
    loginAs("admin-login", ADMIN_EMAIL, ADMIN_PASSWORD);
    cy.intercept("POST", "**/workspace/invite").as("createInvite");

    visitForm(ui("/manage/#/selfservice/workspace-invite/new"));
    cy.contains(".mb-3", "First name").find("input").type("Checkendpointtests");
    cy.contains(".mb-3", "Last name").find("input").type("Invitee");
    cy.contains(".mb-3", "Email address").find("input").type(INVITEE_EMAIL);
    selectFormSelectOption("Role", INVITEE_ROLE_NAME);
    submitForm();
    // Confirming the request itself succeeded (rather than also waiting on whatever
    // page transition follows, which router.goBackOrDefault doesn't guarantee stays
    // client-side) is enough here - the interception is the actual proof the invite
    // (email/SMS send included) went through; the rest of this spec re-authenticates
    // as the invitee anyway, which navigates fresh regardless of where this leaves off.
    cy.wait("@createInvite").its("response.statusCode").should("equal", 200);
  });

  // --- The invitee signs up their own account, then logs in and accepts ---

  it("should be able to sign up a separate account for the invitee's email address.", () => {
    sharedValues(["inviteeOwnWorkspaceTypeId"]).then(
      ({ inviteeOwnWorkspaceTypeId }) => {
        cy.request({
          method: "POST",
          url: ui("/passports/signup/classic"),
          body: {
            value: INVITEE_EMAIL,
            type: "email",
            password: INVITEE_PASSWORD,
            firstName: "Checkendpointtests",
            lastName: "Invitee",
            workspaceTypeId: inviteeOwnWorkspaceTypeId,
          },
        }).then((response: Cypress.Response<SignupResponse>) => {
          expect(response.status).to.equal(200);
          const inviteeToken = response.body.data.item.session.token;
          setShared("inviteeToken", inviteeToken);
          setShared(
            "inviteeOwnWorkspaceId",
            response.body.data.item.session.userWorkspaces[0].workspaceId,
          );
          expect(inviteeToken).to.be.a("string").and.not.be.empty;
        });
      },
    );
  });

  it("should log into the manage UI as the invitee, see the pending invitation, and accept it.", () => {
    loginAs("invitee-login", INVITEE_EMAIL, INVITEE_PASSWORD);

    cy.visit(ui("/manage/#/selfservice/user-invitations"));
    cy.wait(1000);

    // The invite is matched to the invitee purely by their email's passport (see
    // UserInvitationsAction.go/queries/UserInvitations.vsql), not any stored userId -
    // so it shows up here even though nothing linked this account to it explicitly.
    cy.contains(INVITEE_ROLE_NAME, { timeout: 10000 }).should("exist");
    cy.contains(".btn-success", "Accept").click({ force: true });
    cy.contains(".modal-footer button", "Yes", { timeout: 10000 }).click({
      force: true,
    });
    // UserInvitationList.tsx's onAccept calls window.alert() on success, which Cypress
    // auto-accepts - nothing further to interact with here.
    cy.wait(1000);
  });

  it("whoami should confirm the invitee now belongs to the admin's workspace, in addition to their own.", () => {
    sharedValues([
      "inviteeToken",
      "adminWorkspaceId",
      "inviteeOwnWorkspaceId",
    ]).then(({ inviteeToken, adminWorkspaceId, inviteeOwnWorkspaceId }) => {
      cy.request({
        method: "GET",
        url: ui("/whoami"),
        headers: { authorization: inviteeToken },
      }).then((response: Cypress.Response<WhoamiResponse>) => {
        expect(response.status).to.equal(200);
        const workspaceIds = response.body.data.item.workspaces.map(
          (w) => w.uniqueId,
        );
        expect(workspaceIds).to.include(adminWorkspaceId);
        expect(workspaceIds).to.include(inviteeOwnWorkspaceId);
      });
    });
  });

  endFirebackServer();
});
