import { type QueryArchiveColumn } from "../../../../fireback-ui/types/QueryArchiveColumn";
import { RoleDto } from "../../../sdk/abac/RoleDto";
import { enTranslations } from "../../../translations/en";

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
