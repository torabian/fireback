import { useCommonEntityManager } from "{{ .FirebackUiDir }}/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "{{ .FirebackUiDir }}/components/entity-manager/CommonEntityManager";
import { {{ .Template }}Form } from "./{{ .Template }}EditForm";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useGet{{ .Template }}ByUniqueId } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}ByUniqueId";
import { usePost{{ .Template }} } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePost{{ .Template }}";
import { usePatch{{ .Template }} } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePatch{{ .Template }}";
import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

export const {{ .Template }}EntityManager = ({ data }: DtoEntity<{{ .Template }}Entity>) => {
  const s = useS(strings);
  
  const { router, uniqueId, queryClient, locale } = useCommonEntityManager<
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
      onEditTitle={s.{{ .templates }}.edit{{ .Template }} }
      onCreateTitle={s.{{ .templates }}.new{{ .Template }} }
      data={data}
    />
  );
};
