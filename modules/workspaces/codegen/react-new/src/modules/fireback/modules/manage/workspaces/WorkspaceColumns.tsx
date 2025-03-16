import { WorkspaceEntity } from "../../../sdk/modules/workspaces/WorkspaceEntity";
import { enTranslations } from "../../../translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: WorkspaceEntity.Fields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: WorkspaceEntity.Fields.name,
    title: t.wokspaces.name,
    width: 200,
  },
];
