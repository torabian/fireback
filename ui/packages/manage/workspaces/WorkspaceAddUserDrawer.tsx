import { useState } from "react";
import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { useS } from "@fireback/ui-core/hooks/useS";
import { useUsersQuerySource } from "@fireback/ui-core/hooks/useUsersQuerySource";
import { type UserDto } from "@fireback/manage/sdk/abac/UserDto";
import { strings as coreStrings } from "@fireback/ui-core/components/strings/translations";
import { strings } from "./strings/translations";
import {
  useWorkspaceRolesQuerySource,
  type WorkspaceRoleOption,
} from "./useWorkspaceRolesQuerySource";

export interface WorkspaceAddUserDrawerResult {
  userId: string;
  roleId: string;
}

// A minimal drawer form (same render-prop shape as UserPassportsList.tsx's
// SetPasswordDrawer) collecting {userId, roleId} for WorkspaceAddUserButton to submit
// against AddUserToWorkspaceAction - this component only picks the values and resolves
// with them, the actual mutation/error-handling/toast happens at the call site, same
// division of responsibility as every other openDrawer(...).promise.then(...) caller in
// this app.
export const WorkspaceAddUserDrawer = ({
  close,
  resolve,
  workspaceId,
}: {
  close: () => void;
  resolve: (result?: WorkspaceAddUserDrawerResult) => void;
  workspaceId: string;
}) => {
  const s = useS(strings);
  const cs = useS(coreStrings);

  const [user, setUser] = useState<UserDto | undefined>();
  const [role, setRole] = useState<WorkspaceRoleOption | undefined>();

  const canSubmit = !!user?.uniqueId && !!role?.uniqueId;

  return (
    <div className="confirm-drawer-container p-3">
      <h2>{s.addUser}</h2>
      <p>{s.addUserToWorkspaceHint}</p>

      <FormSelect<UserDto, string>
        querySource={useUsersQuerySource}
        value={user}
        onChange={(item) => setUser(item)}
        keyExtractor={(item) => item?.uniqueId}
        fnLabelFormat={(item) =>
          [item?.firstName, item?.lastName].filter(Boolean).join(" ") ||
          item?.uniqueId ||
          ""
        }
        label={s.selectUser}
      />

      <FormSelect<WorkspaceRoleOption, string>
        querySource={(params) => useWorkspaceRolesQuerySource(workspaceId, params)}
        value={role}
        onChange={(item) => setRole(item)}
        keyExtractor={(item) => item?.uniqueId}
        fnLabelFormat={(item) => item?.name || item?.uniqueId || ""}
        label={s.selectRole}
      />

      <div>
        <button
          className="d-block w-100 btn btn-primary"
          disabled={!canSubmit}
          onClick={() =>
            resolve({ userId: user!.uniqueId!, roleId: role!.uniqueId! })
          }
        >
          {cs.common.save}
        </button>
        <button className="d-block w-100 btn" onClick={() => close()}>
          {cs.common.cancel}
        </button>
      </div>
    </div>
  );
};
