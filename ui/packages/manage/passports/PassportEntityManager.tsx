import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";

import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { PassportEditForm } from "./PassportEditForm";
import { usePassportGetActionQuery } from "@fireback/manage/sdk/abac/PassportGetAction";
import { usePassportCreateAction } from "@fireback/manage/sdk/abac/PassportCreateAction";
import { usePassportUpdateAction } from "@fireback/manage/sdk/abac/PassportUpdateAction";
import { PassportDto } from "@fireback/manage/sdk/abac/PassportDto";
import { PassportNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";

export const PassportEntityManager = ({ data }: DtoEntity<PassportDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<PassportDto>
  >({
    data,
  });

  const getSingleHook = useGetPassportByUniqueId({
    query: { uniqueId, deep: true },
  });

  const postHook = usePassportCreateAction({});

  const patchHook = usePassportUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      getSingleHook={getSingleHook}
      patchHook={patchHook}
      onCancel={() => {
        router.goBackOrDefault(PassportNavigation.query(undefined, locale));
      }}
      onFinishUriResolver={(response, locale) =>
        PassportNavigation.single(response.data?.uniqueId, locale)
      }
      Form={PassportEditForm}
      onEditTitle={s.editPassport}
      onCreateTitle={s.newPassport}
      data={data}
    />
  );
};
