import {
  useWorkspaceBrowseActionQuery,
  WorkspaceBrowseActionQueryParams,
} from "@fireback/manage/sdk/abac/WorkspaceBrowseAction";
import { type WorkspaceDto } from "@fireback/manage/sdk/abac/WorkspaceDto";
import { type UseRemoteQuery } from "@fireback/ui-core/types/remoteQuery";

// Adapts useWorkspaceBrowseActionQuery into the {query, items, keyExtractor} shape
// FormSelect's querySource prop expects - same pattern as ui-core's own
// useRegionalContentsQuerySource/useEmailProvidersQuerySource, kept local to
// ui/packages/manage (workspace admin management is manage-only, unlike those
// cross-app entities) - used by WorkspaceEditForm.tsx for the "parent workspace"
// picker.
export const useWorkspacesQuerySource = (params?: UseRemoteQuery) => {
  const query = useWorkspaceBrowseActionQuery({
    qs: new WorkspaceBrowseActionQueryParams({
      itemsPerPage: params?.query?.itemsPerPage ?? 200,
      searchPhrase: params?.query?.searchPhrase || undefined,
    }),
  });
  const items = ((query.data as any)?.data?.items ?? []) as WorkspaceDto[];
  return {
    query: query as any,
    items,
    keyExtractor: (item: WorkspaceDto) => item.uniqueId,
  };
};
