import { CommonListManager } from "{{ .FirebackUiDir }}/components/entity-manager/CommonListManager";
import { columns } from "./{{ .Template }}Columns";
import { {{ .Template }}Entity } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useGet{{ .Template }}s } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}s";
import { usePost{{ .Template }}Remove } from "{{ .SdkDir }}/modules/{{ .ModuleDir }}/usePost{{ .Template }}Remove";
import { useS } from "{{ .FirebackUiDir }}/hooks/useS";
import { strings } from "./strings/translations";

export const {{ .Template }}List = () => {
  const s = useS(strings);

  return (
    <>
      <CommonListManager
        columns={columns(s)}
        queryHook={useGet{{ .Template }}s}
        uniqueIdHrefHandler={(uniqueId: string) =>
          {{ .Template }}Entity.Navigation.single(uniqueId)
        }
        deleteHook={usePost{{ .Template }}Remove}
      ></CommonListManager>
    </>
  );
};
