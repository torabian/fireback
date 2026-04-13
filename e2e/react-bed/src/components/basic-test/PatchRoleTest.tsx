import { CheckClassicPassportActionReq, useCheckClassicPassportAction } from "../../sdk/modules/abac/CheckClassicPassport";
import { usePatchRole } from "../../sdk/modules/abac/usePatchRole";

export function PatchRoleTest() {
  const { submit } = usePatchRole();
  submit({ name: "asd" }).then((res) => {
    console.log(res.data?.name);
  });

  const { mutateAsync } = useCheckClassicPassportAction({});

  mutateAsync(new CheckClassicPassportActionReq({ value: "adasd" })).then((x: any) => x.data?.data?.item?.otpInfo);

  return <ul></ul>;
}
