import { useT } from "../../../hooks/useT";

import { CommonListManager } from "../../../components/entity-manager/CommonListManager";
import { useGetFiles } from "../../../sdk/modules/abac/useGetFiles";
import { columns } from "./DriveColumns";
import { FileEntity } from "../../../sdk/modules/abac/FileEntity";

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
