import { CommonSingleManager } from "{{ .FirebackUiDir }}/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "{{ .FirebackUiDir }}/components/general-entity-view/GeneralEntityView";
import { useCommonEntityManager } from "{{ .FirebackUiDir }}/hooks/useCommonEntityManager";
import { useGet{{ .Template }}ByUniqueId } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}ByUniqueId";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

export const {{ .Template }}SingleScreen = () => {
  const { uniqueId, queryClient } = useCommonEntityManager<Partial<any>>({});

  const getSingleHook = useGet{{ .Template }}ByUniqueId({ query: { uniqueId } });
  var d: {{ .Template }}Entity | undefined = getSingleHook.query.data?.data;
  const t = useS(strings);
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
      
              {{ if or (eq .Type "string") (eq .Type "text") (eq .Type "int64") (eq .Type "float64") (eq .Type "bool") }}
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
