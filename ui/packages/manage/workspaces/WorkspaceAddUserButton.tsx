import { useQueryClient } from "@tanstack/react-query";
import { useOverlay } from "@fireback/ui-core/components/overlay/OverlayProvider";
import { useS } from "@fireback/ui-core/hooks/useS";
import { httpErrorHanlder } from "@fireback/ui-core/hooks/api";
import { Toast } from "@fireback/ui-core/hooks/toast";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import { useAddUserToWorkspaceAction } from "@fireback/manage/sdk/abac/AddUserToWorkspaceAction";
import { AddUserToWorkspaceActionReq } from "@fireback/manage/sdk/abac/AddUserToWorkspaceAction";
import { WorkspaceAddUserDrawer } from "./WorkspaceAddUserDrawer";
import { type WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";

// Per-row "Add user" button rendered via WorkspaceColumns.tsx's synthetic "actions"
// column (DatatableColumn.getCellValue can return any React node, not just a plain
// value - see PaginateUtils.tsx's castColumns). Opens WorkspaceAddUserDrawer to collect
// {userId, roleId}, then submits AddUserToWorkspaceAction (root only) against this row's
// workspace - the same openDrawer(...).promise.then(...).catch(httpErrorHanlder)
// division of labor as UserPassportsList.tsx's SetPasswordDrawer.
export const WorkspaceAddUserButton = ({
  workspace,
}: {
  workspace: WorkspaceDto;
}) => {
  const s = useS(strings);
  const cs = useS(coreStrings);
  const { openDrawer } = useOverlay();
  const queryClient = useQueryClient();
  const { mutateAsync } = useAddUserToWorkspaceAction({});

  const onClick = () => {
    openDrawer<{ userId: string; roleId: string }>((props) => (
      <WorkspaceAddUserDrawer {...props} workspaceId={workspace.uniqueId!} />
    ))
      .promise.then(({ type, data }) => {
        if (type !== "resolved" || !data) {
          return;
        }
        return mutateAsync(
          AddUserToWorkspaceActionReq.with({
            userId: data.userId,
            workspaceId: workspace.uniqueId!,
            roleId: data.roleId,
          }),
        ).then(() => {
          Toast(s.userAddedToWorkspace, { type: "success" });
          queryClient.invalidateQueries("*fireback.UserWorkspaceEntity");
          queryClient.invalidateQueries("*fireback.WorkspaceRoleEntity");
        });
      })
      .catch((err) => httpErrorHanlder(err, cs));
  };

  return (
    <button
      type="button"
      className="btn btn-sm btn-outline-primary"
      onClick={(e) => {
        e.stopPropagation();
        onClick();
      }}
    >
      {s.addUser}
    </button>
  );
};
