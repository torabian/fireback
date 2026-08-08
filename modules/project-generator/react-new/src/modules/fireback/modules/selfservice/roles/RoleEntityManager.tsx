import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { useRoleCreateAction } from "../../../sdk/abac/RoleCreateAction";
import { useRoleGetActionQuery } from "../../../sdk/abac/RoleGetAction";
import { useRoleUpdateAction } from "../../../sdk/abac/RoleUpdateAction";
import { RoleDto } from "../../../sdk/abac/RoleDto";
import { RoleNavigation } from "../../../sdk/navigation/AbacNavigation";
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
