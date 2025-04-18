import { QueryArchiveColumn } from "@/modules/fireback/definitions/common";
import { RoleEntity } from "@/modules/fireback/sdk/modules/abac/RoleEntity";
import { enTranslations } from "@/modules/fireback/translations/en";

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
