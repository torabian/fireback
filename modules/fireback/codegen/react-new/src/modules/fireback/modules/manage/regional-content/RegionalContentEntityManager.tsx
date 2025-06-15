import { useCommonEntityManager } from "@/modules/fireback/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/modules/fireback/components/entity-manager/CommonEntityManager";
import { RegionalContentForm } from "./RegionalContentEditForm";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
import { useGetRegionalContentByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetRegionalContentByUniqueId";
import { usePostRegionalContent } from "@/modules/fireback/sdk/modules/abac/usePostRegionalContent";
import { usePatchRegionalContent } from "@/modules/fireback/sdk/modules/abac/usePatchRegionalContent";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const RegionalContentEntityManager = ({
  data,
}: DtoEntity<RegionalContentEntity>) => {
  const s = useS(strings);
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
    Partial<RegionalContentEntity>
  >({
    data,
  });
  const getSingleHook = useGetRegionalContentByUniqueId({
    query: { uniqueId },
  });
  const postHook = usePostRegionalContent({
    queryClient,
  });
  const patchHook = usePatchRegionalContent({
    queryClient,
  });
  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          RegionalContentEntity.Navigation.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        RegionalContentEntity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={RegionalContentForm}
      onEditTitle={s.regionalContents.editRegionalContent}
      onCreateTitle={s.regionalContents.newRegionalContent}
      data={data}
    />
  );
};
