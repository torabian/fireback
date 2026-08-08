import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { usePassportMethodCreateAction } from "../../../sdk/abac/PassportMethodCreateAction";
import { usePassportMethodGetActionQuery } from "../../../sdk/abac/PassportMethodGetAction";
import { usePassportMethodUpdateAction } from "../../../sdk/abac/PassportMethodUpdateAction";
import { PassportMethodDto } from "../../../sdk/abac/PassportMethodDto";
import { PassportMethodNavigation } from "../../../sdk/navigation/AbacNavigation";
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
