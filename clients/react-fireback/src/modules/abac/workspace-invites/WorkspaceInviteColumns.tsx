import { DatatableColumn } from "@/definitions/definitions";
import { enTranslations } from "@/translations/en";
import { WorkspaceInviteEntity } from "src/sdk/fireback";
import { WorkspaceEntityFields } from "src/sdk/fireback/modules/workspaces/workspace-fields";

export const columns = (t: typeof enTranslations): DatatableColumn[] => [
  {
    name: WorkspaceEntityFields.uniqueId,
    title: t.table.uniqueId,
    width: 100,
  },
  {
    name: "firstName",
    title: t.wokspaces.invite.firstName,
    width: 100,
  },
  {
    name: "lastName",
    title: t.wokspaces.invite.lastName,
    width: 100,
  },
  {
    name: "phoneNumber",
    title: t.wokspaces.invite.phoneNumber,
    width: 100,
  },
  {
    name: "email",
    title: t.wokspaces.invite.email,
    width: 100,
  },
  {
    name: "role_id",
    title: t.wokspaces.invite.role,
    width: 100,
    getCellValue: (invite?: WorkspaceInviteEntity) => invite?.role?.name,
  },
];
