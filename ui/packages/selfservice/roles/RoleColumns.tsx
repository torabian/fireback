import { type QueryArchiveColumn } from "@fireback/ui-core/types/QueryArchiveColumn";
import { RoleDto } from "@fireback/selfservice/sdk/abac/RoleDto";
import { type strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { type strings } from "./strings/translations";

export const columns = (
  s: typeof strings,
  uiS: typeof uiStrings,
): QueryArchiveColumn[] => [
  {
    name: RoleDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 200,
  },
  {
    name: RoleDto.Fields.name,
    title: s.role.name,
    width: 200,
  },
];
