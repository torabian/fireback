import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import {
  CommonEntityManager,
  DtoEntity,
} from "@/components/entity-manager/CommonEntityManager";
import { TemplateForm } from "./TemplateEditForm";
import { TemplateEntity } from "src/sdk/{{ .SdkDir }}";
import { useGetTemplateByUniqueId } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGetTemplateByUniqueId";
import { usePostTemplate } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePostTemplate";
import { usePatchTemplate } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePatchTemplate";

import { TemplateNavigationTools } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/xnavigation";

export const TemplateEntityManager = ({ data }: DtoEntity<TemplateEntity>) => {
  const { router, uniqueId, queryClient, t, locale } = useCommonEntityManager<
    Partial<TemplateEntity>
  >({
    data,
  });

  const getSingleHook = useGetTemplateByUniqueId({
    query: { uniqueId },
  });

  const postHook = usePostTemplate({
    queryClient,
  });

  const patchHook = usePatchTemplate({
    queryClient,
  });

  return (
    <CommonEntityManager
      postHook={postHook}
      patchHook={patchHook}
      getSingleHook={getSingleHook}
      onCancel={() => {
        router.goBackOrDefault(
          TemplateNavigationTools.query(undefined, locale)
        );
      }}
      onFinishUriResolver={(response, locale) =>
        TemplateNavigationTools.single(response.data?.uniqueId, locale)
      }
      Form={TemplateForm}
      onEditTitle={t.templates.editTemplate}
      onCreateTitle={t.templates.newTemplate}
      data={data}
    />
  );
};
