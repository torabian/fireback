import { ui, endFirebackServer, withFirebackServer } from "../support/setup";

// Covers Abac.emi.yml's alternate auth methods: ClassicPassportRequestOtp,
// ClassicPassportOtp, ConfirmClassicPassportTotp, OauthAuthenticate,
// OsLoginAuthenticate. All over HTTP, mirroring
// modules/abac/tests/alt_auth_methods_test.go's approach (including reading the OTP
// code back via /publicAuthentication/browse, and the TOTP secret back via
// /passport/browse, as root - there's no other observable way to retrieve either in a
// black-box test, since delivery goes through a "terminal" email provider that only
// logs). No real-UI click-through for the otp-entry/totp-setup screens here: generating
// a valid TOTP code needs an RFC 6238 implementation, and none of this repo's e2e
// tooling has one available (unlike the Go side, which already depends on
// github.com/pquerna/otp) - adding a new npm dependency just for this felt like more
// than this spec's marginal value justified, so the same coverage is proven at the API
// level instead, same as OauthAuthenticate/OsLoginAuthenticate already have to be
// (no UI drives either of those at all).
const ROOT_EMAIL = "test@agent.com";
const ROOT_PASSWORD = "123321";

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

function authHeaders(token: string) {
  return { authorization: token, "workspace-id": "root" };
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
              // ClassicSignupActionImplementation.go's completeClassicSignupProcess sets
              // a secure "authorization" cookie for the newly-signed-up account - unlike
              // Go's plain http.Client (no cookie jar by default), cy.request() behaves
              // like a browser and both keeps and resends it automatically. The server's
              // AuthorizeRequest (see withAuthorization.go) prefers that cookie over any
              // explicit Authorization header, so without clearing it here, every
              // root-authenticated call made after this signup would silently run as
              // this throwaway account instead of root.
              const token = response.body.data.item.session.token;
              return cy
                .clearCookie("authorization")
                .then(() => ({ email, token }));
            });
        });
    });
}

function configureEmailSending() {
  const senderAddr = `checkendpointtests-alt-auth-${Date.now()}@example.com`;
  cy.task(
    "exec",
    ` messaging email sender c --from-name Checkendpointtests --from-email-address ${senderAddr} --reply-to ${senderAddr} --nick-name checkendpointtests`,
  ).then((content: string) => {
    const senderId = (
      JSON.parse(content) as SingleItemResponse<{ uniqueId: string }>
    ).data.item.uniqueId;
    cy.task(
      "exec",
      ` messaging email provider c --type terminal --title checkendpointtests`,
    ).then((providerContent: string) => {
      const providerId = (
        JSON.parse(providerContent) as SingleItemResponse<{ uniqueId: string }>
      ).data.item.uniqueId;
      cy.task(
        "exec",
        ` notification create --general-email-provider-id ${providerId} --invite-to-workspace-sender-id ${senderId}`,
      );
    });
  });
}

function fetchOtpCode(value: string) {
  return sharedValues(["rootToken"]).then(({ rootToken }) =>
    cy
      .request({
        method: "GET",
        url: ui("/publicAuthentication/browse"),
        headers: authHeaders(rootToken),
      })
      .then(
        (
          response: Cypress.Response<
            ListResponse<{ passportValue: string; otp: string }>
          >,
        ) => {
          expect(response.status).to.equal(200);
          const match = response.body.data.items
            .filter((i) => i.passportValue === value)
            .pop();
          expect(match, `a publicAuthentication row for ${value}`).to.exist;
          return match!.otp;
        },
      ),
  );
}

describe("Abac: alternate auth methods (otp, totp, oauth, os-login)", () => {
  withFirebackServer();

  it("should be able to create the passport method needed for signup, and sign in as root for API calls.", () => {
    cy.task("exec", ` passport method create --region global --type email`);
    cy.request({
      method: "POST",
      url: ui("/passports/signin/classic"),
      body: { value: ROOT_EMAIL, password: ROOT_PASSWORD },
    }).then((response: Cypress.Response<SignupResponse>) => {
      expect(response.status).to.equal(200);
      setShared("rootToken", response.body.data.item.session.token);
    });
    configureEmailSending();
  });

  it("ClassicPassportRequestOtp should succeed, then block an immediate repeat request.", () => {
    signupFreshAccount("otp-request").then(({ email }) => {
      cy.request({
        method: "POST",
        url: ui("/workspace/passport/request-otp"),
        body: { value: email },
      }).then((first) => {
        expect(first.status).to.equal(200);
      });

      cy.request({
        method: "POST",
        url: ui("/workspace/passport/request-otp"),
        body: { value: email },
        failOnStatusCode: false,
      }).then((second) => {
        expect(second.status).to.not.equal(200);
      });
    });
  });

  it("ClassicPassportOtp should sign in with the real code and reject a wrong one.", () => {
    signupFreshAccount("otp-confirm").then(({ email }) => {
      cy.request({
        method: "POST",
        url: ui("/workspace/passport/request-otp"),
        body: { value: email },
      });

      fetchOtpCode(email).then((code) => {
        cy.request({
          method: "POST",
          url: ui("/workspace/passport/otp"),
          body: { value: email, otp: code },
        }).then((response: Cypress.Response<SignupResponse>) => {
          expect(response.status).to.equal(200);
          expect(response.body.data.item.session.token).to.be.a("string").and
            .not.be.empty;
        });
      });

      cy.request({
        method: "POST",
        url: ui("/workspace/passport/otp"),
        body: { value: email, otp: "000000" },
        failOnStatusCode: false,
      }).then((wrong) => {
        expect(wrong.status).to.not.equal(200);
      });
    });
  });

  // ConfirmClassicPassportTotp's success path needs a real RFC 6238 code generated from
  // the passport's own totpSecret - see this spec's header comment for why that's Go-
  // side only (modules/abac/tests/alt_auth_methods_test.go's
  // TestConfirmClassicPassportTotp_HTTP_Succeeds). A wrong code is still enough to
  // prove the endpoint itself is reachable and rejects correctly from the browser.
  it("ConfirmClassicPassportTotp should reject a wrong code for an account with no totp secret set up.", () => {
    signupFreshAccount("totp-confirm").then(({ email }) => {
      cy.request({
        method: "POST",
        url: ui("/passport/totp/confirm"),
        body: {
          value: email,
          password: "checkendpointtests-pass-123",
          totpCode: "000000",
        },
        failOnStatusCode: false,
      }).then((response) => {
        expect(response.status).to.not.equal(200);
      });
    });
  });

  it("OauthAuthenticate should reject an unsupported service and an invalid google token.", () => {
    cy.request({
      method: "POST",
      url: ui("/passport/via-oauth"),
      body: { service: "not-a-real-provider", token: "whatever" },
      failOnStatusCode: false,
    }).then((response) => {
      expect(response.status).to.not.equal(200);
    });

    cy.request({
      method: "POST",
      url: ui("/passport/via-oauth"),
      body: {
        service: "google",
        token: `checkendpointtests-not-real-${Date.now()}`,
      },
      failOnStatusCode: false,
    }).then((response) => {
      expect(response.status).to.not.equal(200);
    });
  });

  // OsLoginAuthenticate covers the AbacTest crash-fix in
  // OsLoginAuthenticateActionImplementation.go - before the fix, this call panicked the
  // whole server process on every single request (a nil *fireback.IError returned
  // directly as the built-in error interface), so this is a real regression guard, not
  // just a coverage checkbox.
  it("OsLoginAuthenticate should succeed and return a session token.", () => {
    cy.request({
      method: "GET",
      url: ui("/passports/os/login"),
    }).then(
      (response: Cypress.Response<SingleItemResponse<{ token: string }>>) => {
        expect(response.status).to.equal(200);
        expect(response.body.data.item.token).to.be.a("string").and.not.be
          .empty;
      },
    );
  });

  endFirebackServer();
});
