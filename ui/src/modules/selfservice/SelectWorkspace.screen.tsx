import { usePresenter } from "./SelectWorkspace.presenter";
import { useAuthentication } from "../fireback-ui/auth/AuthenticationContext";
import {
  QueryUserRoleWorkspacesActionRes,
  useQueryUserRoleWorkspacesActionQuery,
} from "../sdk/abac/QueryUserRoleWorkspacesAction";
import type { MArray } from "../sdk/sdk/common/operators";

export const SelectWorkspaceScreen = () => {
  const { s } = usePresenter();
  const queryUrw = useQueryUserRoleWorkspacesActionQuery({
    cacheTime: 50,
  });

  const items = queryUrw.data?.data?.items || [];
  const { selectWorkspace } = useAuthentication();

  return (
    <div className="signin-form-container">
      <div className="mb-4">
        <h1 className="h3">{s.selectWorkspaceTitle}</h1>
        <p className="text-muted">{s.selectWorkspace}</p>
      </div>

      {(items || []).map((workspace) => (
        <div key={workspace.uniqueId} className="mb-4">
          <h2 className="h5">{workspace.name}</h2>
          <div className="d-flex flex-wrap gap-2 mt-2">
            {(workspace.roles as MArray<any>).get().map((role) => (
              <button
                key={role.uniqueId}
                className="btn btn-outline-primary w-100"
                onClick={() =>
                  selectWorkspace({
                    workspaceId: workspace.uniqueId,
                    roleId: role.uniqueId,
                  })
                }
              >
                {s.select} ({role.name})
              </button>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};
