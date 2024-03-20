import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "@/hooks/useCommonEntityManager";
import { useT } from "@/hooks/useT";
import { useGet{{ .Template }}ByUniqueId } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}ByUniqueId";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";

export const {{ .Template }}SingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});

  const getSingleHook = useGet{{ .Template }}ByUniqueId({ query: { uniqueId } });
  var d: {{ .Template }}Entity | undefined = getSingleHook.query.data?.data;
  const t = useT();
  // usePageTitle(`${d?.name}`);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={({ locale, router }) => {
          router.push({{ .Template }}Entity.Navigation.edit(uniqueId, locale));
        {{"}}"}}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={
            [
              {{ range .e.CompleteFields }}
      
              {{ if or (eq .Type "string") (eq .Type "text") (eq .Type "int64") (eq .Type "float64") }}
              {
                elem: d?.{{ .Name }},
                label: t.{{ $.templates}}.{{ .Name }},
              },    
              {{ end }}
              
              {{ end }}
              
             
            ]
          }
        />
      </CommonSingleManager>
    </>
  );
};
