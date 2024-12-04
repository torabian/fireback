import { QueryArchiveColumn } from "../../definitions/common";
import { RoleEntity } from "../../sdk/modules/workspaces/RoleEntity";
import { enTranslations } from "../../translations/en";

export const columns = (t: typeof enTranslations): QueryArchiveColumn[] => [
  {
    name: RoleEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: RoleEntity.Fields.name,
    title: t.role.name,
    width: 200,
  },
];
