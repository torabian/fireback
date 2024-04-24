import { useT } from "@/hooks/useT";
import { CommonListManager } from "@/components/entity-manager/CommonListManager";
import { columns } from "./AppMenuColumns";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
import { useGetAppMenus } from "src/sdk/fireback/modules/workspaces/useGetAppMenus";
import { useDeleteAppMenu } from "@/sdk/fireback/modules/workspaces/useDeleteAppMenu";
import { useGetCteAppMenus } from "@/sdk/fireback/modules/workspaces/useGetCteAppMenus";
export const AppMenuList = () => {
  const t = useT();
  return (
    <>
      <CommonListManager
        columns={columns(t)}
        queryHook={useGetCteAppMenus}
        uniqueIdHrefHandler={(uniqueId: string) =>
          AppMenuEntity.Navigation.single(uniqueId)
        }
        deleteHook={useDeleteAppMenu}
      ></CommonListManager>
    </>
  );
};
