import { WorkspaceTypeDto } from "@fireback/manage/sdk/abac/WorkspaceTypeDto";
import { type strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { type strings } from "./strings/translations";

export const columns = (s: typeof strings, uiS: typeof uiStrings) => [
  {
    name: "uniqueId",
    title: uiS.table.uniqueId,
    width: 200,
  },
  {
    name: "title",
    title: s.title,
    width: 200,
    getCellValue: (entity: WorkspaceTypeDto) => entity.title,
  },
  {
    name: "slug",
    slug: s.slug,
    width: 200,
    getCellValue: (entity: WorkspaceTypeDto) => entity.slug,
  },
];
