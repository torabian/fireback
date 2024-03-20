import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./TemplateColumns";
import { TemplateNavigationTools } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/xnavigation";
import { useGetTemplates } from "src/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useGetTemplates";
import { useDeleteTemplate } from "@/sdk/{{ .SdkDir }}/modules/{{ .ModuleDir }}/useDeleteTemplate";

export const TemplateList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetTemplates}
        uniqueIdHrefHandler={(uniqueId: string) =>
          TemplateNavigationTools.single(uniqueId)
        }
        deleteHook={useDeleteTemplate}
      ></CommonListManager>
    </>
  );
};
