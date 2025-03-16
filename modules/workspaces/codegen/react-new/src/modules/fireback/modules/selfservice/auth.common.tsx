export enum AuthMethod {
  Email = "email",
  Phone = "phone",
  Google = "google",
}

export interface AuthAvailableMethods {
  email: boolean;
  phone: boolean;
  google: boolean;
}
