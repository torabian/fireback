import { KeyValue } from "@/definitions/definitions";
import { enTranslations } from "@/translations/en";

export const getPassportOptions = (t: typeof enTranslations): KeyValue[] => {
  return [
    {
      label: "Email and Password",
      value: "EmailPassword",
    },
    {
      label: "Phone number",
      value: "PhoneNumber",
    },
  ];
};

export const getPasswordOptions = (t: typeof enTranslations): KeyValue[] => {
  return [
    {
      label: "Send an email to user, containing their temporary password",
      value: "ByEmail",
    },
    {
      label: "Show me their temporary password",
      value: "ShowPassword",
    },
  ];
};
