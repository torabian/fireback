import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useT } from "@/modules/fireback/hooks/useT";
import { useRoleCreateAction } from "@/modules/fireback/sdk/abac/RoleCreateAction";
import { useRoleGetActionQuery } from "@/modules/fireback/sdk/abac/RoleGetAction";
import { useRoleUpdateAction } from "@/modules/fireback/sdk/abac/RoleUpdateAction";
import { RoleDto } from "@/modules/fireback/sdk/abac/RoleDto";
import { RoleNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { RoleEditForm } from "./RoleEditForm";

export const RoleEntityManager = ({ data }: DtoEntity<RoleDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<RoleDto>
  >({
    data,
  });
  const t = useT();

  const getSingleHook = useRoleGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useRoleCreateAction({});

  const patchHook = useRoleUpdateAction({ params: { uniqueId } });

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
        router.goBackOrDefault(RoleNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        RoleNavigation.single(response.data?.uniqueId, locale)
      }
      Form={RoleEditForm}
      onEditTitle={t.fb.editRole}
      onCreateTitle={t.fb.newRole}
      data={data}
    />
  );
};
