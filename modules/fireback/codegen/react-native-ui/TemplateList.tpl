import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./{{ .Template }}Columns";
import { {{ .Template }}Entity } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/{{ .Template}}Entity";
import { useGet{{ .Template }}s } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGet{{ .Template }}s";
import { useDelete{{ .Template }} } from "@/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useDelete{{ .Template }}";

export const {{ .Template }}List = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGet{{ .Template }}s}
        uniqueIdHrefHandler={(uniqueId: string) =>
          {{ .Template }}Entity.Navigation.single(uniqueId)
        }
        deleteHook={useDelete{{ .Template }}}
      ></CommonListManager>
    </>
  );
};
