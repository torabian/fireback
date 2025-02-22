import { useDeleteRole } from "../../sdk/modules/workspaces/useDeleteRole";
import { useGetCapabilities } from "../../sdk/modules/workspaces/useGetCapabilities";

export function DeleteRoleTest() {
  const { submit } = useDeleteRole();

  const onComplete = () => {
    submit({ uniqueId: "asd" })
      .then((res) => {
        console.log(res.data?.rowsAffected);
      })
      .catch((err) => {});
  };

  return <ul></ul>;
}
