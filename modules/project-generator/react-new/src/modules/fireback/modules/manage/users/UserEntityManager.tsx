import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { UserEditForm } from "./UserEditForm";
import { useUserGetActionQuery } from "../../../sdk/abac/UserGetAction";
import { useUserCreateAction } from "../../../sdk/abac/UserCreateAction";
import { useUserUpdateAction } from "../../../sdk/abac/UserUpdateAction";
import { UserDto } from "../../../sdk/abac/UserDto";
import { UserNavigation } from "../../../sdk/navigation/AbacNavigation";

export const UserEntityManager = ({ data }: DtoEntity<UserDto>) => {
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<UserDto>
  >({
    data,
  });
  const s = useS(strings);

  const getSingleHook = useGetUserByUniqueId({
    query: { uniqueId, deep: true },
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
        UserNavigation.single(response.data?.uniqueId, locale)
      }
      Form={UserEditForm}
      onEditTitle={s.editUser}
      onCreateTitle={s.newUser}
      data={data}
    />
  );
};
