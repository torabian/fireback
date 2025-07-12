import { useContext } from "react";
import { useGetUrwQuery } from "../../sdk/modules/abac/useGetUrwQuery";
import { usePresenter } from "./SelectWorkspace.presenter";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

export const SelectWorkspaceScreen = () => {
  const { s } = usePresenter();
  const { query: queryWorkspaces } = useGetUrwQuery({
    queryOptions: { cacheTime: 50 },
    query: {},
  });
  const items = queryWorkspaces.data?.data?.items || [];
  const { selectedUrw, selectUrw } = useContext(RemoteQueryContext);

  return (
    <div className="signin-form-container">
      <div className="mb-4">
        <h1 className="h3">{s.selectWorkspaceTitle}</h1>
        <p className="text-muted">{s.selectWorkspace}</p>
      </div>

      {items.map((workspace) => (
        <div key={workspace.uniqueId} className="mb-4">
          <h2 className="h5">{workspace.name}</h2>
          <div className="d-flex flex-wrap gap-2 mt-2">
            {workspace.roles.map((role) => (
              <button
                key={role.uniqueId}
                className="btn btn-outline-primary w-100"
                onClick={() =>
                  selectUrw({
                    workspaceId: workspace.uniqueId,
                    roleId: role.uniqueId,
                  })
                }
              >
                Select ({role.name})
              </button>
            ))}
          </div>
        </div>
      ))}
    </div>
  );
};
