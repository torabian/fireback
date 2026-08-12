import { EmailSenderDto } from "@fireback/manage/sdk/messaging/EmailSenderDto";

export function FormEmailSenderPicker({
  value,
  onChange,
}: {
  value: EmailSenderDto;
  onChange: (entity: EmailSenderDto) => void;
}) {
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
