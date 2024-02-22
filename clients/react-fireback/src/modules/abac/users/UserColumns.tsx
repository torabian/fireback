import { enTranslations } from "@/translations/en";
import { UserEntityFields } from "src/sdk/fireback/modules/workspaces/user-fields";

export const columns = (t: typeof enTranslations) => [
  {
    name: UserEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: UserEntityFields.firstName,
    title: t.users.firstName,
    width: 200,
  },
  {
    name: UserEntityFields.lastName,
    title: t.users.lastName,
    width: 200,
  },
];
