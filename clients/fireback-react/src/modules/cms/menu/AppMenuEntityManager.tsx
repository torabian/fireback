import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { AppMenuForm } from "./AppMenuEditForm";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
import { useGetAppMenuByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetAppMenuByUniqueId";
import { usePostAppMenu } from "src/sdk/fireback/modules/workspaces/usePostAppMenu";
import { usePatchAppMenu } from "src/sdk/fireback/modules/workspaces/usePatchAppMenu";
export const AppMenuEntityManager = ({ data }: DtoEntity<AppMenuEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<AppMenuEntity>
  >({
    data,
  });
  const getSingleHook = useGetAppMenuByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostAppMenu({
    queryClient,
  });
  const patchHook = usePatchAppMenu({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          AppMenuEntity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        AppMenuEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ AppMenuForm }
      onEditTitle={t.appMenus.editAppMenu }
      onCreateTitle={t.appMenus.newAppMenu }
      data={data}
    />
  );
};