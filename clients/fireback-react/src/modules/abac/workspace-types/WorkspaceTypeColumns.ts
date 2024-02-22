import { enTranslations } from "@/translations/en";
import { WorkspaceTypeEntity } from "src/sdk/fireback";

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
  {
    name: "role",
    title: t.wokspaces.role,
    width: 200,
    getCellValue: (entity: WorkspaceTypeEntity) => entity.role?.name,
  },
];
