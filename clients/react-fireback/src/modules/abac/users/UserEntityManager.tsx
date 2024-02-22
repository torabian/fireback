import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";

import { UserEntity } from "src/sdk/fireback";
import { useGetUserByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetUserByUniqueId";
import { usePatchUser } from "src/sdk/fireback/modules/workspaces/usePatchUser";
import { usePostUser } from "src/sdk/fireback/modules/workspaces/usePostUser";
import { UserNavigationTools } from "src/sdk/fireback/modules/workspaces/user-navigation-tools";

import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { UserEditForm } from "./UserEditForm";

export const UserEntityManager = ({ data }: DtoEntity<UserEntity>) => {
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<UserEntity>
  >({
    data,
  });

  const getSingleHook = useGetUserByUniqueId({
    query: { uniqueId },
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
        router.goBackOrDefault(UserNavigationTools.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        UserNavigationTools.single(response.data?.uniqueId, locale)
      }
      Form={UserEditForm}
      onEditTitle={t.user.editUser}
      onCreateTitle={t.user.newUser}
      data={data}
    />
  );
};
