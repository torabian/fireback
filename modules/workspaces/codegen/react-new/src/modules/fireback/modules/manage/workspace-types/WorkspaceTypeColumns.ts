import { WorkspaceTypeEntity } from "../../../sdk/modules/workspaces/WorkspaceTypeEntity";
import { enTranslations } from "../../../translations/en";

export const columns = (t: typeof enTranslations) => [
  {
    name: "uniqueId",
    title: t.table.uniqueId,
    width: 200,
  },
  {
    name: "title",
    title: t.wokspaces.title,
    width: 200,
    getCellValue: (entity: WorkspaceTypeEntity) => entity.title,
  },
  {
    name: "slug",
    slug: t.wokspaces.slug,
    width: 200,
    getCellValue: (entity: WorkspaceTypeEntity) => entity.slug,
  },
];
