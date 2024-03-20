import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { TemplateEntity } from "src/sdk/{{ .SdkDir }}";
import { useGetTemplateByUniqueId } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGetTemplateByUniqueId";
import { TemplateNavigationTools } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/xnavigation";

export const TemplateSingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});

  const getSingleHook = useGetTemplateByUniqueId({ query: { uniqueId } });
  var d: TemplateEntity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push(TemplateNavigationTools.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              <% for (let field of fields) { let name = field.name ; %>
                <% if (field.type.includes("string")) {  %>
                  {
                    elem: d?.<%- name %>,
                    label: t.templates.<%- name %>,
                  },    
                <% } %>
              <% } %>
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};
