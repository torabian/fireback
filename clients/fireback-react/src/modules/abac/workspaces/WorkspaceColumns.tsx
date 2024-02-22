import { enTranslations } from "@/translations/en";
import { WorkspaceEntityFields } from "src/sdk/fireback/modules/workspaces/workspace-fields";

export const columns = (t: typeof enTranslations) => [
  {
    name: WorkspaceEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: WorkspaceEntityFields.name,
    title: t.wokspaces.name,
    width: 200,
  },
];
