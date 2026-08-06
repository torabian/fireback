import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import { useS } from "@/modules/fireback/hooks/useS";
import { usePassportMethodCreateAction } from "@/modules/fireback/sdk/abac/PassportMethodCreateAction";
import { usePassportMethodGetActionQuery } from "@/modules/fireback/sdk/abac/PassportMethodGetAction";
import { usePassportMethodUpdateAction } from "@/modules/fireback/sdk/abac/PassportMethodUpdateAction";
import { PassportMethodEntity } from "@/modules/fireback/sdk/modules/abac/PassportMethodEntity";
import { PassportMethodForm } from "./PassportMethodEditForm";
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

  const getSingleHook = usePassportMethodGetActionQuery({
    params: { uniqueId },
  });
  const postHook = usePassportMethodCreateAction();
  const patchHook = usePassportMethodUpdateAction({ params: { uniqueId } });

  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          PassportMethodEntity.Navigation.query(undefined, locale),
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
