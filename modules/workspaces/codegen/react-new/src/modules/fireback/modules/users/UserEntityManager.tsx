import { useCommonEntityManager } from "../../hooks/useCommonEntityManager";

import { useGetUserByUniqueId } from "../../sdk/modules/workspaces/useGetUserByUniqueId";
import { usePatchUser } from "../../sdk/modules/workspaces/usePatchUser";
import { usePostUser } from "../../sdk/modules/workspaces/usePostUser";

import {
  CommonEntityManager,
  DtoEntity,
} from "../../components/entity-manager/CommonEntityManager";
import { UserEditForm } from "./UserEditForm";
import { UserEntity } from "../../sdk/modules/workspaces/UserEntity";

export const UserEntityManager = ({ data }: DtoEntity<UserEntity>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<UserEntity>
  >({
    data,
  });

  const getSingleHook = useGetUserByUniqueId({
    query: { uniqueId, deep: true },
  });

  const postHook = usePostUser({
    queryClient,
  });

  const patchHook = usePatchUser({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(UserEntity.Navigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        UserEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={UserEditForm}
      onEditTitle={t.user.editUser}
      onCreateTitle={t.user.newUser}
      data={data}
    />
  );
};
