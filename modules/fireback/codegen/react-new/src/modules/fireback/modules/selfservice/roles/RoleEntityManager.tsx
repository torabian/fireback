import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { RoleEntity } from "@/modules/fireback/sdk/modules/abac/RoleEntity";
import { useGetRoleByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetRoleByUniqueId";
import { usePatchRole } from "@/modules/fireback/sdk/modules/abac/usePatchRole";
import { usePostRole } from "@/modules/fireback/sdk/modules/abac/usePostRole";
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
      beforeSubmit={(data) => {
        if (data.capabilities?.length > 0 && data.capabilitiesListId === null) {
          return {
            ...data,
            capabilitiesListId: data.capabilities.map((item) => item.uniqueId),
          };
        }
        return data;
      }}
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
