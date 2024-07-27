import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { UserForm } from "./UserEditForm";
import { UserEntity } from "src/sdk/fireback/modules/workspaces/UserEntity";
import { useGetUserByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetUserByUniqueId";
import { usePostUser } from "src/sdk/fireback/modules/workspaces/usePostUser";
import { usePatchUser } from "src/sdk/fireback/modules/workspaces/usePatchUser";
export const UserEntityManager = ({ data }: DtoEntity<UserEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
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
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          UserEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        UserEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ UserForm }
      onEditTitle={t.users.editUser }
      onCreateTitle={t.users.newUser }
      data={data}
    />
  );
};