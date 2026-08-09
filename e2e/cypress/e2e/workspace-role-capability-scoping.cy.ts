import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

interface CapabilityItem {
  uniqueId: string;
  name: string;
  description: string;
}

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

interface BrowseResponse<T> {
  data: {
    items: T[];
    totalItems: number;
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

// Root's own CLI account, created by withFirebackServer() (see support/setup.js's
// "auth --in-root=true --value test@agent.com ... --password 123321" call) - reused
// here to log the very same account into the manage UI.
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

// The two capability "domains" this test scopes a role/workspace type around. See the
// capability-creation `it` below for why the actual stored uniqueId isn't the bare
// "students.*"/"teachers.*" wildcard the domain name might suggest.
const STUDENTS_CAPABILITY = "root.students.*";
const TEACHERS_CAPABILITY = "root.teachers.*";

describe("Workspace role capability scoping (capabilities -> roles -> workspace types -> signup -> whoami)", () => {
  withFirebackServer();

  // Kept as a single flat list of `it`s (no nested `describe`) - see
  // workspace-type.cy.ts/passport-method.cy.ts for why: a nested describe's tests run
  // after every direct `it` in this describe, including endFirebackServer's "should stop
  // fireback" call below, which would tear the server down before these ever ran.

  it("should be able to create the passport method needed for signup.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
  });

  // --- Capabilities, created directly in the database via the CLI ---

  it("should be able to store a students capability and a teachers capability in the database via the CLI.", () => {
    // Two things about the exact uniqueId stored here are not what the "students.*"/
    // "teachers.*" domain names above might suggest, both discovered by actually
    // exercising this end to end rather than assumed:
    //
    // 1. RoleActions.go's filterRolePermissions only lets a caller assign a capability
    //    to a role if the caller's OWN granted accesses already cover it (Permissions.go's
    //    MeetsCheck does a substring check against the caller's perms with "*" stripped -
    //    e.g. "root.*" stripped to "root." must be a substring of the capability being
    //    assigned). Every access root ever holds is "root."-prefixed - a bare top-level
    //    "students.*"/"teachers.*" (no "root." prefix) would never satisfy that, and role
    //    create fails with RoleNeedsOneCapability even for the root CLI user. So both
    //    rows live under "root." instead.
    // 2. CapabilitiesTreeAction (which the manage UI's RolePermissionTree checkbox tree
    //    below renders from) only re-attaches a ".*" wildcard suffix to a node when it
    //    actually finds real children nested under it in the capability catalog - a
    //    catalog row that's ITSELF already suffixed with ".*" but has no further children
    //    (e.g. storing "root.students.*" verbatim) renders as a plain "students" leaf
    //    checkbox with the ".*" silently dropped. Storing one granular leaf permission
    //    ("root.students.manage") instead makes "students" a real, auto-materialized
    //    parent node in that tree, and checking it in the UI actually produces the
    //    "root.students.*" wildcard entry in capabilitiesListId - see the login/role
    //    creation `it` below.
    cy.task(
      "exec",
      ` abac internal capability-c --unique-id root.students.manage --name "Manage students"`,
    ).then((content: string) => {
      const capability = (JSON.parse(content) as SingleItemResponse<CapabilityItem>)
        .data.item;
      expect(capability.uniqueId).to.equal("root.students.manage");
    });

    cy.task(
      "exec",
      ` abac internal capability-c --unique-id root.teachers.manage --name "Manage teachers"`,
    ).then((content: string) => {
      const capability = (JSON.parse(content) as SingleItemResponse<CapabilityItem>)
        .data.item;
      expect(capability.uniqueId).to.equal("root.teachers.manage");
    });
  });

  // --- Roles, created through the manage React app itself (not the CLI) ---

  it("should log into the manage UI as root and create a role per capability by checking it in the permission tree, then sign out.", () => {
    // The RolePermissionTree/react-select bits below aren't reachable in the narrow
    // mobile viewport login.cy.js uses for the welcome/signup flow - the sidebar (and
    // its "Sign out" link) is only rendered inline, without a hamburger drawer, at
    // desktop widths (see deviceInformation.tsx's isMobileView, width < 1024).
    cy.viewport(1366, 900);

    Cypress.on("uncaught:exception", () => {
      // Same as login.cy.js's account-creation test - don't let an unrelated uncaught
      // exception in the app fail this test.
      return false;
    });

    cy.visit(ui("/manage/#/en/welcome"));
    cy.get("#using-email").should("exist").click();
    cy.get("#value-input").type(ROOT_EMAIL);
    cy.get("#submit-form").click({ force: true });

    // Root's passport already exists (created via the CLI's "auth --in-root=true" in
    // withFirebackServer), so - unlike login.cy.js's brand new account - this lands on
    // the "enter password" screen instead of "complete your account".
    cy.get("h1").should("have.text", "Enter Password");
    cy.get("#password-input").type(ROOT_PASSWORD);
    cy.get("#submit-form").click({ force: true });

    // Root only ever has the one "root" workspace, so WithSelfServiceRoutes.tsx's
    // auto-select skips the workspace picker and lands straight on the dashboard.
    cy.url().should("include", "/dashboard");

    createRoleWithCapability("studentRole", "students");
    createRoleWithCapability("teacherRole", "teachers");

    // Signing out goes through a plain window.confirm() (CurrentUser.tsx) - Cypress
    // auto-accepts that by default, no stub needed.
    cy.contains(".nav-link", "Sign out").click();
    cy.get("h1", { timeout: 10000 }).should("have.text", "Welcome back");
  });

  // --- Workspace types, one per role, created via the CLI (same mechanism already
  // covered end to end by workspace-type.cy.ts) ---

  let studentWorkspaceTypeId = "";
  let teacherWorkspaceTypeId = "";

  it("should be able to look the two roles created through the UI back up via the CLI, and create a workspace type for each.", () => {
    cy.task("exec", ` role browse`).then((content: string) => {
      const roles = (JSON.parse(content) as BrowseResponse<RoleItem>).data.items;

      const studentRole = roles.find((r) => r.name === "studentRole");
      const teacherRole = roles.find((r) => r.name === "teacherRole");

      expect(studentRole, "studentRole created through the UI").to.exist;
      expect(teacherRole, "teacherRole created through the UI").to.exist;

      // This is the real, end to end proof that checking the "students"/"teachers"
      // checkbox in the manage UI actually produced the intended wildcard capability -
      // see the long comment on the capability-creation `it` above.
      expect(studentRole!.capabilitiesListId).to.deep.equal([
        STUDENTS_CAPABILITY,
      ]);
      expect(teacherRole!.capabilitiesListId).to.deep.equal([
        TEACHERS_CAPABILITY,
      ]);

      cy.task(
        "exec",
        ` ws workspaceType-c --title Students --slug /students --role-id ${studentRole!.uniqueId}`,
      ).then((wtContent: string) => {
        const workspaceType = (
          JSON.parse(wtContent) as SingleItemResponse<WorkspaceTypeItem>
        ).data.item;
        studentWorkspaceTypeId = workspaceType.uniqueId;
        expect(workspaceType.roleId).to.equal(studentRole!.uniqueId);
      });

      cy.task(
        "exec",
        ` ws workspaceType-c --title Teachers --slug /teachers --role-id ${teacherRole!.uniqueId}`,
      ).then((wtContent: string) => {
        const workspaceType = (
          JSON.parse(wtContent) as SingleItemResponse<WorkspaceTypeItem>
        ).data.item;
        teacherWorkspaceTypeId = workspaceType.uniqueId;
        expect(workspaceType.roleId).to.equal(teacherRole!.uniqueId);
      });
    });
  });

  // --- Public signup, one user per workspace type, then whoami to prove each user's
  // workspace role capability is limited to just their own domain ---

  let studentToken = "";
  let studentWorkspaceId = "";
  let teacherToken = "";
  let teacherWorkspaceId = "";

  it("should be able to sign up a student and a teacher, each selecting their own workspace type.", () => {
    cy.request({
      method: "POST",
      url: SIGNUP_URL,
      body: {
        value: "student@test.com",
        type: "email",
        password: "testpass123",
        firstName: "Stu",
        lastName: "Dent",
        workspaceTypeId: studentWorkspaceTypeId,
      },
    }).then((response: Cypress.Response<SignupResponse>) => {
      expect(response.status).to.equal(200);
      const session = response.body.data.item.session;
      studentToken = session.token;
      studentWorkspaceId = session.userWorkspaces[0].workspaceId;
    });

    cy.request({
      method: "POST",
      url: SIGNUP_URL,
      body: {
        value: "teacher@test.com",
        type: "email",
        password: "testpass123",
        firstName: "Tea",
        lastName: "Cher",
        workspaceTypeId: teacherWorkspaceTypeId,
      },
    }).then((response: Cypress.Response<SignupResponse>) => {
      expect(response.status).to.equal(200);
      const session = response.body.data.item.session;
      teacherToken = session.token;
      teacherWorkspaceId = session.userWorkspaces[0].workspaceId;
    });
  });

  it("whoami should report the student's workspace as scoped to only the students capability.", () => {
    cy.request({
      method: "GET",
      url: WHOAMI_URL,
      headers: { authorization: studentToken },
    }).then((response: Cypress.Response<WhoamiResponse>) => {
      expect(response.status).to.equal(200);
      const workspace = response.body.data.item.workspaces.find(
        (w) => w.uniqueId === studentWorkspaceId,
      );
      expect(workspace).to.exist;
      expect(workspace!.capabilities).to.deep.equal([STUDENTS_CAPABILITY]);
      expect(workspace!.roles[0].capabilities).to.deep.equal([
        STUDENTS_CAPABILITY,
      ]);
      // The limiting half of "limited": the teachers capability must not have leaked in.
      expect(workspace!.capabilities).to.not.include(TEACHERS_CAPABILITY);
    });
  });

  it("whoami should report the teacher's workspace as scoped to only the teachers capability.", () => {
    cy.request({
      method: "GET",
      url: WHOAMI_URL,
      headers: { authorization: teacherToken },
    }).then((response: Cypress.Response<WhoamiResponse>) => {
      expect(response.status).to.equal(200);
      const workspace = response.body.data.item.workspaces.find(
        (w) => w.uniqueId === teacherWorkspaceId,
      );
      expect(workspace).to.exist;
      expect(workspace!.capabilities).to.deep.equal([TEACHERS_CAPABILITY]);
      expect(workspace!.roles[0].capabilities).to.deep.equal([
        TEACHERS_CAPABILITY,
      ]);
      expect(workspace!.capabilities).to.not.include(STUDENTS_CAPABILITY);
    });
  });

  endFirebackServer();
});

// Kept at module level (not nested inside the describe) - same as
// login.cy.js's validateAppMenuEntity.
function createRoleWithCapability(roleName: string, capabilityLeaf: string) {
  cy.visit(ui("/manage/#/en/selfservice/role/new"));

  // Scoped to Layout.tsx's ".content-section" (the routed page content) rather than
  // the whole document - Navbar.tsx's reactive-search box also renders an
  // `<input class="form-control">`, so an unscoped `cy.get("input.form-control")`
  // would match both it and the role name field.
  cy.get(".content-section").within(() => {
    // The role name is the only plain text input on this screen - the capability tree
    // below it is built entirely out of checkboxes (RolePermissionTree.tsx's
    // Checkbox.tsx renders a real <input type="checkbox" class="form-check-input">),
    // and neither FormText call in RoleEditForm.tsx sets an id.
    cy.get("input.form-control").clear().type(roleName);

    // The tree is fully expanded in the DOM at all times (PermissionTree.tsx renders
    // every level's children unconditionally, no expand/collapse toggle), so the
    // "students"/"teachers" node - auto-materialized as the parent of the
    // "root.students.manage"/"root.teachers.manage" leaf created via the CLI above -
    // is already there to check without navigating into "root" first.
    cy.contains("label", capabilityLeaf)
      .find('input[type="checkbox"]')
      .check({ force: true });

    // CommonEntityManager.tsx renders no visible "Save" button - submitting goes
    // through a hidden `<button type="submit" class="d-none">` inside the form, the
    // same element a Ctrl+S keyboard shortcut (useCommonCrudActions) would trigger.
    cy.get("form button[type=submit]").click({ force: true });
  });

  cy.wait(1000);
}
