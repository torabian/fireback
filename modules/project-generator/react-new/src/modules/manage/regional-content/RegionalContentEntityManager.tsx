import { useCommonEntityManager } from "../../fireback-ui/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  type DtoEntity,
} from "../../fireback-ui/components/entity-manager/CommonEntityManager";
import { RegionalContentForm } from "./RegionalContentEditForm";
import { useRegionalContentGetActionQuery } from "../../sdk/abac/RegionalContentGetAction";
import { useRegionalContentCreateAction } from "../../sdk/abac/RegionalContentCreateAction";
import { useRegionalContentUpdateAction } from "../../sdk/abac/RegionalContentUpdateAction";
import { RegionalContentDto } from "../../sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "../../sdk/navigation/AbacNavigation";
import { useS } from "../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentEntityManager = ({
  data,
}: DtoEntity<RegionalContentDto>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<RegionalContentDto>
  >({
    data,
  });
  const getSingleHook = useRegionalContentGetActionQuery({
    params: { uniqueId },
  });
  const postHook = useRegionalContentCreateAction({});
  const patchHook = useRegionalContentUpdateAction({ params: { uniqueId } });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          RegionalContentNavigation.query(undefined, locale),
        );
      }}
      onFinishUriResolver={(response, locale) =>
        RegionalContentNavigation.single(response.data?.uniqueId, locale)
      }
      Form={RegionalContentForm}
      onEditTitle={s.regionalContents.editRegionalContent}
      onCreateTitle={s.regionalContents.newRegionalContent}
      data={data}
    />
  );
};
