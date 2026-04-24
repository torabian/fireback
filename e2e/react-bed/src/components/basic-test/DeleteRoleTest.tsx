import { usePostRoleRemove } from "../../sdk/modules/abac/usePostRoleRemove";

export function DeleteRoleTest() {
  const { submit } = usePostRoleRemove();

  const onComplete = () => {
    submit({ uniqueId: "asd" })
      .then((res) => {
        console.log(res.data?.rowsAffected);
      })
      .catch((err) => {});
  };

  return <ul></ul>;
}
