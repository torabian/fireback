import { WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";
import { type strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { type strings } from "./strings/translations";

export const columns = (s: typeof strings, uiS: typeof uiStrings) => [
  {
    name: WorkspaceDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 100,
  },
  {
    name: WorkspaceDto.Fields.name,
    title: s.name,
    width: 200,
  },
];
