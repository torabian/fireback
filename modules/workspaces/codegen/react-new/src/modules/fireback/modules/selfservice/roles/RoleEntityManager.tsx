import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { RoleEntity } from "@/modules/fireback/sdk/modules/workspaces/RoleEntity";
import { useGetRoleByUniqueId } from "@/modules/fireback/sdk/modules/workspaces/useGetRoleByUniqueId";
import { usePatchRole } from "@/modules/fireback/sdk/modules/workspaces/usePatchRole";
import { usePostRole } from "@/modules/fireback/sdk/modules/workspaces/usePostRole";
import { RoleEditForm } from "./RoleEditForm";

export const RoleEntityManager = ({ data }: DtoEntity<RoleEntity>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<RoleEntity>
  >({
    data,
  });
  const t = useT();

  const getSingleHook = useGetRoleByUniqueId({
    query: { uniqueId },
    queryOptions: {
      enabled: !!uniqueId,
    },
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
