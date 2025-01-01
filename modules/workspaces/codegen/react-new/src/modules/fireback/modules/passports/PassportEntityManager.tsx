import { useCommonEntityManager } from "../../hooks/useCommonEntityManager";

import { useGetPassportByUniqueId } from "../../sdk/modules/workspaces/useGetPassportByUniqueId";
import { usePatchPassport } from "../../sdk/modules/workspaces/usePatchPassport";
import { usePostPassport } from "../../sdk/modules/workspaces/usePostPassport";

import {
  CommonEntityManager,
  DtoEntity,
} from "../../components/entity-manager/CommonEntityManager";
import { PassportEditForm } from "./PassportEditForm";
import { PassportEntity } from "../../sdk/modules/workspaces/PassportEntity";
import { useS } from "../../hooks/useS";
import { strings } from "./strings/translations";

export const PassportEntityManager = ({ data }: DtoEntity<PassportEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
    Partial<PassportEntity>
  >({
    data,
  });

  const getSingleHook = useGetPassportByUniqueId({
    query: { uniqueId, deep: true },
  });

  const postHook = usePostPassport({
    queryClient,
  });

  const patchHook = usePatchPassport({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(
          PassportEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PassportEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PassportEditForm}
      onEditTitle={s.editPassport}
      onCreateTitle={s.newPassport}
      data={data}
    />
  );
};
