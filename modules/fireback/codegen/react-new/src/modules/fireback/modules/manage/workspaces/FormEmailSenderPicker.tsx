import { EmailSenderEntity } from "../../../sdk/modules/abac/EmailSenderEntity";
import { useContext } from "react";
import { RemoteQueryContext } from "../../../sdk/core/react-tools";

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
      {/* <FormEntitySelect
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
