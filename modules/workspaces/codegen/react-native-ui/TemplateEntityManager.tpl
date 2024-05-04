import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { {{ .Template }}Form } from "./{{ .Template }}EditForm";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useGet{{ .Template }}ByUniqueId } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}ByUniqueId";
import { usePost{{ .Template }} } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePost{{ .Template }}";
import { usePatch{{ .Template }} } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePatch{{ .Template }}";

export const {{ .Template }}EntityManager = ({ data }: DtoEntity<{{ .Template }}Entity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<{{ .Template }}Entity>
  >({
    data,
  });

  const getSingleHook = useGet{{ .Template }}ByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePost{{ .Template }}({
    queryClient,
  });

  const patchHook = usePatch{{ .Template }}({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          {{ .Template }}Entity.Navigation.query(undefined, locale)
        );
      } }
      onFinishUriResolver={(response, locale) =>
        {{ .Template }}Entity.Navigation.single(response.data?.uniqueId, locale)
      }
      Form={ {{ .Template }}Form }
      onEditTitle={t.{{ .templates }}.edit{{ .Template }} }
      onCreateTitle={t.{{ .templates }}.new{{ .Template }} }
      data={data}
    />
  );
};
