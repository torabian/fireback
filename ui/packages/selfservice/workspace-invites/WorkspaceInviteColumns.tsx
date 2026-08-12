import { type DatatableColumn } from "@fireback/ui-core/types/DatatableColumn";
import { WorkspaceInviteDto } from "@fireback/selfservice/sdk/abac/WorkspaceInviteDto";
import { type strings as uiStrings } from "@fireback/ui-core/components/strings/translations";
import { type strings } from "./strings/translations";

export const columns = (
  s: typeof strings,
  uiS: typeof uiStrings,
): DatatableColumn[] => [
  {
    name: WorkspaceInviteDto.Fields.uniqueId,
    title: uiS.table.uniqueId,
    width: 100,
  },
  {
    name: "firstName",
    title: s.firstName,
    width: 100,
  },
  {
    name: "lastName",
    title: s.lastName,
    width: 100,
  },
  {
    name: "phoneNumber",
    title: s.phoneNumber,
    width: 100,
  },
  {
    name: "email",
    title: s.email,
    width: 100,
  },
  {
    name: "role_id",
    title: s.roleLabel,
    width: 100,
    getCellValue: (invite?: WorkspaceInviteDto) => invite?.role?.name,
  },
];
