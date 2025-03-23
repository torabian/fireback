import { UserInvitationsQueryColumns } from "@/modules/fireback/sdk/modules/workspaces/UserInvitationsQueryColumns";
import { DatatableColumn } from "../../../definitions/definitions";
import { strings } from "./strings/translations";

export const userInvitationColumns = (
  s: typeof strings,
  onAccept: (invitationId: UserInvitationsQueryColumns) => void,
  onReject: (invitationId: UserInvitationsQueryColumns) => void
): DatatableColumn[] => [
  {
    name: "roleName",
    title: s.roleName,
    width: 100,
  },
  {
    name: "workspaceName",
    title: s.workspaceName,
    width: 100,
  },
  {
    name: "method",
    title: s.method,
    width: 100,
    getCellValue: (dto: any) => {
      return dto.type as any;
    },
  },
  {
    name: "value",
    title: s.passport,
    width: 100,
    getCellValue: (dto: any) => {
      return dto.value;
    },
  },
  {
    name: "actions",
    title: s.actions,
    width: 100,
    getCellValue: (dto) => {
      return (
        <>
          <button
            className="btn btn-sm btn-success"
            style={{ marginRight: "2px" }}
            onClick={(e) => {
              onAccept(dto);
            }}
          >
            {s.accept}
          </button>
          <button
            onClick={(e) => {
              onReject(dto);
            }}
            className="btn btn-sm btn-danger"
          >
            {s.reject}
          </button>
        </>
      ) as any;
    },
  },
];
