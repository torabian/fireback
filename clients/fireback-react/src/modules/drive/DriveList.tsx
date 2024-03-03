import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { useGetFiles } from "@/sdk/fireback/modules/drive/useGetFiles";
import { DriveNavigationTools } from "src/sdk/fireback/modules/drive/drive-navigation-tools";
import { columns } from "./DriveColumns";

export const DriveList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetFiles}
        uniqueIdHrefHandler={(uniqueId: string) =>
          DriveNavigationTools.single(uniqueId)
        }
      ></CommonListManager>
    </>
  );
};
