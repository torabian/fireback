import { useCommonEntityManager } from "../../../../fireback-ui/hooks/useCommonEntityManager";


import {
  CommonEntityManager,
  type DtoEntity,
} from "../../../../fireback-ui/components/entity-manager/CommonEntityManager";
import { PassportEditForm } from "./PassportEditForm";
import { usePassportGetActionQuery } from "../../../sdk/abac/PassportGetAction";
import { usePassportCreateAction } from "../../../sdk/abac/PassportCreateAction";
import { usePassportUpdateAction } from "../../../sdk/abac/PassportUpdateAction";
import { PassportDto } from "../../../sdk/abac/PassportDto";
import { PassportNavigation } from "../../../sdk/navigation/AbacNavigation";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";

export const PassportEntityManager = ({ data }: DtoEntity<PassportDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale, t } = useCommonEntityManager<
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
        router.goBackOrDefault(
          PassportNavigation.query(undefined, locale)
        );
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
