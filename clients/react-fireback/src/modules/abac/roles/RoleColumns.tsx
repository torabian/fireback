import { QueryArchiveColumn } from "@/definitions/common";
import { enTranslations } from "@/translations/en";
import { RoleEntityFields } from "src/sdk/fireback/modules/workspaces/role-fields";

export const columns = (t: typeof enTranslations): QueryArchiveColumn[] => [
  {
    name: RoleEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: RoleEntityFields.name,
    title: t.role.name,
    width: 200,
  },
];
