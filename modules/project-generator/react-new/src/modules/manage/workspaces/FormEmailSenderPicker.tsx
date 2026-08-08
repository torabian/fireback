import { EmailSenderDto } from "../../sdk/messaging/EmailSenderDto";
import { useContext } from "react";
import { RemoteQueryContext } from "../../sdk/core/react-tools";

export function FormEmailSenderPicker({
  value,
  onChange,
}: {
  value: EmailSenderDto;
  onChange: (entity: EmailSenderDto) => void;
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
        labelFn={(t: EmailSenderDto) =>
          [t?.fromName, t.fromEmailAddress].join(" ")
        }
      /> */}
    </>
  );
}
