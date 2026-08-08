import { WorkspaceDto } from "../../../sdk/abac/WorkspaceDto";
import { enTranslations } from "../../../translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: WorkspaceDto.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: WorkspaceDto.Fields.name,
    title: t.wokspaces.name,
    width: 200,
  },
];
