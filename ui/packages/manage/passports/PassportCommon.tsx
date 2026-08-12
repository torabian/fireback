import { type KeyValue } from "@fireback/ui-core/types/KeyValue";
import { strings } from "./strings/translations";

export const getPassportOptions = (s: typeof strings): KeyValue[] => {
  return [
    {
      label: s.optionEmailPassword,
      value: "EmailPassword",
    },
    {
      label: s.optionPhoneNumber,
      value: "PhoneNumber",
    },
  ];
};

export const getPasswordOptions = (s: typeof strings): KeyValue[] => {
  return [
    {
      label: s.passwordByEmail,
      value: "ByEmail",
    },
    {
      label: s.passwordShowPassword,
      value: "ShowPassword",
    },
  ];
};

export const getPassportTypes = (s: typeof strings) => {
  return [
    {
      name: s.typePhone,
      uniqueId: "phonenumber",
    },
    { name: s.typeEmail, uniqueId: "email" },
  ];
};
