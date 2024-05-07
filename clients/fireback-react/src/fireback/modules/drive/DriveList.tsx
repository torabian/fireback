import { useT } from "@/fireback/hooks/useT";

import { CommonListManager } from "@/fireback/components/entity-manager/CommonListManager";
import { useGetFiles } from "@/sdk/fireback/modules/workspaces/useGetFiles";
import { columns } from "./DriveColumns";
import { FileEntity } from "@/sdk/fireback/modules/workspaces/FileEntity";

export const DriveList = () => {
  const t = useT();

  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetFiles}
        uniqueIdHrefHandler={(uniqueId: string) =>
          FileEntity.Navigation.single(uniqueId)
        }
      />
    </>
  );
};
