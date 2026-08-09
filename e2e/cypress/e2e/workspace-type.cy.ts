import { endFirebackServer, withFirebackServer } from "../support/setup";

interface RoleItem {
  uniqueId: string;
  name: string;
  capabilitiesListId: string[];
}

interface WorkspaceTypeItem {
  uniqueId: string;
  title: string;
  slug: string;
  roleId: string;
}

interface SingleItemResponse<T> {
  data: {
    item: T;
  };
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

interface WhoamiWorkspace {
  uniqueId: string;
  capabilities: string[];
  roles: { name: string; capabilities: string[] }[];
}

interface WhoamiResponse {
  data: {
    item: {
      userId: string;
      workspaces: WhoamiWorkspace[];
    };
  };
}

const SIGNUP_URL = "http://localhost:7794/passports/signup/classic";
const WHOAMI_URL = "http://localhost:7794/whoami";
const NOTIFICATION_CONFIG_BROWSE_URL =
  "http://localhost:7794/notificationConfig/browse";
const CAPABILITIES_TREE_URL = "http://localhost:7794/capabilitiesTree";
const RESTRICTED_CAPABILITY = "root.manage.abac.notification-config.query";

describe("Workspace types", () => {
  withFirebackServer();

  // Kept as a single flat list of `it`s (no nested `describe`) - see
  // passport-method.cy.ts/login.cy.js for why: a nested describe's tests run
  // after every direct `it` in this describe, including endFirebackServer's
  // "should stop fireback" call below, which would tear the server down
  // before these ever ran.

  it("should be able to create the passport method needed for signup.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  // --- WorkspaceType validation ---

  it("should NOT be able to create a workspace type with no role.", () => {
    cy.task(
      "execSupress",
      ` ws workspaceType-c --title noRole --slug /no-role`,
    ).then((content: string) => {
      expect(content).to.contain("RoleIsNotAccessible");
    });
  });

  it("should NOT be able to create a workspace type with a role that doesn't exist.", () => {
    cy.task(
      "execSupress",
      ` ws workspaceType-c --title badRole --slug /bad-role --role-id does-not-exist`,
    ).then((content: string) => {
      expect(content).to.contain("RoleIsNotAccessible");
    });
  });

  it("should NOT be able to assign a role holding the root.* wildcard to a workspace type.", () => {
    cy.task(
      "exec",
      ` role create --name wildcardRole --capabilities-list-id '["root.*"]'`,
    ).then((content: string) => {
      const role = (JSON.parse(content) as SingleItemResponse<RoleItem>).data
        .item;

      cy.task(
        "execSupress",
        ` ws workspaceType-c --title wildcard --slug /wildcard --role-id ${role.uniqueId}`,
      ).then((updateContent: string) => {
        expect(updateContent).to.contain("RoleCannotBeSuperAdmin");
      });
    });
  });

  let restrictedRoleId = "";

  it("should be able to create a role restricted to a single capability, for the workspace type.", () => {
    cy.task(
      "exec",
      ` role create --name customerRole --capabilities-list-id '["${RESTRICTED_CAPABILITY}"]'`,
    ).then((content: string) => {
      const role = (JSON.parse(content) as SingleItemResponse<RoleItem>).data
        .item;
      restrictedRoleId = role.uniqueId;
      expect(role.capabilitiesListId).to.deep.equal([RESTRICTED_CAPABILITY]);
    });
  });

  it("workspace type slug must start with a slash.", () => {
    cy.task(
      "execSupress",
      ` ws workspaceType-c --title bad --slug no-slash --role-id ${restrictedRoleId}`,
    ).then((content: string) => {
      expect(content).to.contain("SlugMustStartWithSlash");
    });
  });

  it("workspace type slug can only contain lowercase letters and dashes.", () => {
    cy.task(
      "execSupress",
      ` ws workspaceType-c --title bad --slug /Has_Upper --role-id ${restrictedRoleId}`,
    ).then((content: string) => {
      expect(content).to.contain("SlugInvalidFormat");
    });
  });

  it("workspace type slug cannot be longer than 50 characters.", () => {
    const longSlug = "/" + "a".repeat(60);
    cy.task(
      "execSupress",
      ` ws workspaceType-c --title bad --slug ${longSlug} --role-id ${restrictedRoleId}`,
    ).then((content: string) => {
      expect(content).to.contain("SlugTooLong");
    });
  });

  let customerWorkspaceTypeId = "";

  it("should be able to create the customer workspace type.", () => {
    cy.task(
      "exec",
      ` ws workspaceType-c --title customer --slug /customer --role-id ${restrictedRoleId}`,
    ).then((content: string) => {
      const workspaceType = (
        JSON.parse(content) as SingleItemResponse<WorkspaceTypeItem>
      ).data.item;
      customerWorkspaceTypeId = workspaceType.uniqueId;
      expect(workspaceType.roleId).to.equal(restrictedRoleId);
    });
  });

  it("should NOT be able to create another workspace type with the same slug.", () => {
    cy.task(
      "exec",
      ` role create --name anotherRole --capabilities-list-id '["${RESTRICTED_CAPABILITY}"]'`,
    ).then((content: string) => {
      const role = (JSON.parse(content) as SingleItemResponse<RoleItem>).data
        .item;

      cy.task(
        "execSupress",
        ` ws workspaceType-c --title customer2 --slug /customer --role-id ${role.uniqueId}`,
      ).then((updateContent: string) => {
        expect(updateContent).to.contain("SlugAlreadyTaken");
      });
    });
  });

  it("should NOT be able to delete the seeded root workspace type.", () => {
    cy.task("execSupress", ` ws workspaceType-d --unique-ids root`).then(
      (content: string) => {
        expect(content).to.contain("CannotDeleteRootWorkspaceType");
      },
    );
  });

  // --- Different workspace types actually grant different, enforced
  // permissions - not just a fixed default role regardless of which type was
  // selected (see WorkspaceCoreFeatures.go's ResolveWorkspaceTypeRoleCapabilities
  // and Permissions.go's MeetsAccessLevel). ---

  let customerToken = "";
  let customerWorkspaceId = "";

  it("should be able to sign up selecting the customer workspace type.", () => {
    cy.request({
      method: "POST",
      url: SIGNUP_URL,
      body: {
        value: "workspacetype-customer@test.com",
        type: "email",
        password: "testpass123",
        firstName: "Cust",
        lastName: "Omer",
        workspaceTypeId: customerWorkspaceTypeId,
      },
    }).then((response: Cypress.Response<SignupResponse>) => {
      expect(response.status).to.equal(200);
      const session = response.body.data.item.session;
      customerToken = session.token;
      customerWorkspaceId = session.userWorkspaces[0].workspaceId;
    });
  });

  it("the signed-up user's own role should carry the workspace type's capabilities, not a fixed default.", () => {
    cy.request({
      method: "GET",
      url: WHOAMI_URL,
      headers: { authorization: customerToken },
    }).then((response: Cypress.Response<WhoamiResponse>) => {
      expect(response.status).to.equal(200);
      const workspace = response.body.data.item.workspaces.find(
        (w) => w.uniqueId === customerWorkspaceId,
      );
      expect(workspace).to.exist;
      expect(workspace!.capabilities).to.deep.equal([RESTRICTED_CAPABILITY]);
      expect(workspace!.roles[0].capabilities).to.deep.equal([
        RESTRICTED_CAPABILITY,
      ]);
    });
  });

  it("the signed-up user SHOULD be able to reach an action within the workspace type's capability.", () => {
    cy.request({
      method: "GET",
      url: NOTIFICATION_CONFIG_BROWSE_URL,
      headers: {
        authorization: customerToken,
        "Workspace-id": customerWorkspaceId,
      },
    }).then((response: Cypress.Response<unknown>) => {
      expect(response.status).to.equal(200);
    });
  });

  it("the signed-up user should NOT be able to reach an action outside the workspace type's capability.", () => {
    cy.request({
      method: "GET",
      url: CAPABILITIES_TREE_URL,
      failOnStatusCode: false,
      headers: {
        authorization: customerToken,
        "Workspace-id": customerWorkspaceId,
      },
    }).then((response: Cypress.Response<{ error: { message: string } }>) => {
      expect(response.status).to.equal(401);
      expect(response.body.error.message).to.equal("NotEnoughPermission");
    });
  });

  endFirebackServer();
});
