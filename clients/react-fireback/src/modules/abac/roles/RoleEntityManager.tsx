import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { RoleEntity } from "src/sdk/fireback";
import { RoleNavigationTools } from "src/sdk/fireback/modules/workspaces/role-navigation-tools";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { useT } from "@/hooks/useT";
import { RoleEditForm } from "./RoleEditForm";
import { useGetRoleByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetRoleByUniqueId";
import { usePostRole } from "@/sdk/fireback/modules/workspaces/usePostRole";
import { usePatchRole } from "@/sdk/fireback/modules/workspaces/usePatchRole";

export const RoleEntityManager = ({ data }: DtoEntity<RoleEntity>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<RoleEntity>
  >({
    data,
  });
  const t = useT();

  const getSingleHook = useGetRoleByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostRole({
    queryClient,
  });

  const patchHook = usePatchRole({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(RoleNavigationTools.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        RoleNavigationTools.single(response.data?.uniqueId, locale)
      }
      Form={RoleEditForm}
      onEditTitle={t.fb.editRole}
      onCreateTitle={t.fb.newRole}
      data={data}
    />
  );
};
