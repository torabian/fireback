export const msalConfig = {
  auth: {
    clientId: process.env.NEXT_PUBLIC_MSAL_CLIENT_ID || "",
    authority: process.env.NEXT_PUBLIC_MSAL_AUTHORITY || "",
    redirectUri: process.env.NEXT_PUBLIC_REDIRECT_URI || "",
    postLogoutRedirectUri: process.env.NEXT_PUBLIC_LOGOUT_REDIRECT_URI || "",
  },
};

// Add here scopes for id token to be used at MS Identity Platform endpoints.
export const loginRequest = {
  scopes: ["User.Read"],
};

// Add here the endpoints for MS Graph API services you would like to use.
export const graphConfig = {
  graphMeEndpoint: process.env.NEXT_PUBLIC_GRAPH_ME_URI,
};
