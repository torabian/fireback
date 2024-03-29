import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { useT } from "@/hooks/useT";
import { RoleEditForm } from "./RoleEditForm";
import { useGetRoleByUniqueId } from "@/sdk/fireback/modules/workspaces/useGetRoleByUniqueId";
import { usePostRole } from "@/sdk/fireback/modules/workspaces/usePostRole";
import { usePatchRole } from "@/sdk/fireback/modules/workspaces/usePatchRole";
import { RoleEntity } from "@/sdk/fireback/modules/workspaces/RoleEntity";

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
        router.goBackOrDefault(RoleEntity.Navigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        RoleEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={RoleEditForm}
      onEditTitle={t.fb.editRole}
      onCreateTitle={t.fb.newRole}
      data={data}
    />
  );
};
