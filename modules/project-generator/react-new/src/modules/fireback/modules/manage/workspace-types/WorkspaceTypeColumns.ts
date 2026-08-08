import { WorkspaceTypeDto } from "../../../sdk/abac/WorkspaceTypeDto";
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
    getCellValue: (entity: WorkspaceTypeDto) => entity.title,
  },
  {
    name: "slug",
    slug: t.wokspaces.slug,
    width: 200,
    getCellValue: (entity: WorkspaceTypeDto) => entity.slug,
  },
];
