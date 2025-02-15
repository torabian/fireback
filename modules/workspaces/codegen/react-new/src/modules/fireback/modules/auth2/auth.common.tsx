export enum AuthMethod {
  Email = "email",
  Phone = "phone",
}

export interface AuthAvailableMethods {
  email: boolean;
  phone: boolean;
  google: boolean;
}
