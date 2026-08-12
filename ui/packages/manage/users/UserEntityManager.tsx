import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

import { useUserGetActionQuery } from "@fireback/manage/sdk/abac/UserGetAction";
import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useUserCreateAction } from "@fireback/manage/sdk/abac/UserCreateAction";
import { UserDto } from "@fireback/manage/sdk/abac/UserDto";
import { useUserUpdateAction } from "@fireback/manage/sdk/abac/UserUpdateAction";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { UserEditForm } from "./UserEditForm";

export const UserEntityManager = ({ data }: DtoEntity<UserDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<UserDto>
  >({
    data,
  });
  const s = useS(strings);

  const getSingleHook = useUserGetActionQuery({
    params: { uniqueId },
  });

  const postHook = useUserCreateAction({});

  const patchHook = useUserUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(UserNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        UserNavigation.single(response.data?.item?.uniqueId, locale)
      }
      Form={UserEditForm}
      onEditTitle={s.editUser}
      onCreateTitle={s.newUser}
      data={data}
    />
  );
};
