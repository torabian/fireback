import { useT } from "@/hooks/useT";

import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { DriveNavigationTools } from "src/sdk/fireback/modules/drive/drive-navigation-tools";
import { useGetDrive } from "src/sdk/fireback/modules/drive/useGetDrive";
import { columns } from "./DriveColumns";

export const DriveList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetDrive}
        uniqueIdHrefHandler={(uniqueId: string) =>
          DriveNavigationTools.single(uniqueId)
        }
      ></CommonListManager>
    </>
  );
};
