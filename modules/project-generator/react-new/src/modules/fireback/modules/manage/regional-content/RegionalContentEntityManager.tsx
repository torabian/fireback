import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  type DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { RegionalContentForm } from "./RegionalContentEditForm";
import { useRegionalContentGetActionQuery } from "@/modules/fireback/sdk/abac/RegionalContentGetAction";
import { useRegionalContentCreateAction } from "@/modules/fireback/sdk/abac/RegionalContentCreateAction";
import { useRegionalContentUpdateAction } from "@/modules/fireback/sdk/abac/RegionalContentUpdateAction";
import { RegionalContentDto } from "@/modules/fireback/sdk/abac/RegionalContentDto";
import { RegionalContentNavigation } from "@/modules/fireback/sdk/navigation/AbacNavigation";
import { useS } from "@/modules/fireback/hooks/useS";
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
          RegionalContentNavigation.query(undefined, locale)
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
