import {
  CommonEntityManager,
  type DtoEntity,
} from "@fireback/ui-core/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "@fireback/ui-core/hooks/useCommonEntityManager";
import { useS } from "@fireback/ui-core/hooks/useS";
import { usePassportMethodCreateAction } from "@fireback/manage/sdk/abac/PassportMethodCreateAction";
import { usePassportMethodGetActionQuery } from "@fireback/manage/sdk/abac/PassportMethodGetAction";
import { usePassportMethodUpdateAction } from "@fireback/manage/sdk/abac/PassportMethodUpdateAction";
import { PassportMethodDto } from "@fireback/manage/sdk/abac/PassportMethodDto";
import { PassportMethodNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { PassportMethodForm } from "./PassportMethodEditForm";
import { strings } from "./strings/translations";

export const PassportMethodEntityManager = ({
  data,
}: DtoEntity<PassportMethodDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<PassportMethodDto>
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
          PassportMethodNavigation.query(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        PassportMethodNavigation.single(response.data?.uniqueId, locale)
      }
      Form={PassportMethodForm}
      onEditTitle={s.passportMethods.editPassportMethod}
      onCreateTitle={s.passportMethods.newPassportMethod}
      data={data}
    />
  );
};
