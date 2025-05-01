import { usePatchRole } from "../../sdk/modules/abac/usePatchRole";
import { usePostWorkspacePassportCheck } from "../../sdk/modules/abac/usePostWorkspacePassportCheck";

export function PatchRoleTest() {
  const { submit } = usePatchRole();
  submit({ name: "asd" }).then((res) => {
    console.log(res.data?.name);
  });

  const { submit: submit2 } = usePostWorkspacePassportCheck({});

  submit2({ value: "adasd" }).then((x) => x.data?.otpInfo);

  return <ul></ul>;
}
