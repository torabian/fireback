import { type QueryArchiveColumn } from "@/modules/fireback/definitions/common";
import { RoleDto } from "@/modules/fireback/sdk/abac/RoleDto";
import { enTranslations } from "@/modules/fireback/translations/en";

export const columns = (t: typeof enTranslations): QueryArchiveColumn[] => [
  {
    name: RoleDto.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: RoleDto.Fields.name,
    title: t.role.name,
    width: 200,
  },
];
