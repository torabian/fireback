import { EmailSenderEntity } from "@/sdk/fireback/modules/workspaces/EmailSenderEntity";
import { useContext } from "react";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";

export function FormEmailSenderPicker({
  value,
  onChange,
}: {
  value: EmailSenderEntity;
  onChange: (entity: EmailSenderEntity) => void;
}) {
  const { options } = useContext(RemoteQueryContext);

  return (
    <>
      {/* <FormEntitySelect2
        fnLoadOptions={async (keyword) => {
          return (
            (
              await EmailSenderActions.fn(options)
                .query(`name %"${keyword}"%`)
                .getEmailSenders()
            ).data?.items || []
          );
        }}
        value={value}
        onChange={onChange}
        labelFn={(t: EmailSenderEntity) =>
          [t?.fromName, t.fromEmailAddress].join(" ")
        }
      /> */}
    </>
  );
}
