import { FormEntitySelect2 } from "@/components/forms/form-select/FormEntitySelect2";
import { EmailSenderEntity } from "src/sdk/fireback";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { EmailSenderActions } from "src/sdk/fireback/modules/workspaces/email-sender-actions";
import { useContext } from "react";

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
      <FormEntitySelect2
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
      />
    </>
  );
}
