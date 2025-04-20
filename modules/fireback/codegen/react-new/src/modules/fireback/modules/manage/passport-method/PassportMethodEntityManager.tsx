import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { PassportMethodForm } from "./PassportMethodEditForm";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { useGetPassportMethodByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetPassportMethodByUniqueId";
import { usePostPassportMethod } from "@/modules/fireback/sdk/modules/abac/usePostPassportMethod";
import { usePatchPassportMethod } from "@/modules/fireback/sdk/modules/abac/usePatchPassportMethod";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const PassportMethodEntityManager = ({
  data,
}: DtoEntity<PassportMethodEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<PassportMethodEntity>
  >({
    data,
  });
  const getSingleHook = useGetPassportMethodByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostPassportMethod({
    queryClient,
  });
  const patchHook = usePatchPassportMethod({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PassportMethodEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PassportMethodEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={PassportMethodForm}
      onEditTitle={s.passportMethods.editPassportMethod}
      onCreateTitle={s.passportMethods.newPassportMethod}
      data={data}
    />
  );
};
