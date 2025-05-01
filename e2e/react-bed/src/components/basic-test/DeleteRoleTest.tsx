import { useDeleteRole } from "../../sdk/modules/abac/useDeleteRole";
import { useGetCapabilities } from "../../sdk/modules/fireback/useGetCapabilities";

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
