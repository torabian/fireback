import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { type EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailProviderEntity } from "@/modules/fireback/sdk/modules/abac/EmailProviderEntity";

const placeholder = `
curl -X POST https://api.sendgrid.com/v3/mail/send \
  -H "Authorization: Bearer %SENDGRID_API_KEY%" \
  -H "Content-Type: application/json" \
  -d '{
    "personalizations": [
      {
        "to": [
          {
            "email": "%ToEmail%",
            "name": "%ToName%"
          }
        ],
        "subject": "%Subject%"
      }
    ],
    "from": {
      "email": "%FromEmail%",
      "name": "%FromName%"
    },
    "content": [
      {
        "type": "text/plain",
        "value": "%Content%"
      }
    ]
  }'
`;

export const EmailProviderEditForm = ({
  form,
  isEditing,
}: EntityFormProps<EmailProviderEntity>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();

  // =====================
  // Providers
  // =====================
  const emailProviders = [
    { label: "SendGrid", value: "sendgrid" },
    { label: "Mailgun", value: "mailgun" },
    { label: "Postmark", value: "postmark" },
    { label: "Resend", value: "resend" },
    { label: "Curl", value: "curl" },
    { label: "SMTP", value: "smtp" },
    { label: "Terminal (Debug)", value: "terminal" },
  ];

  const querySource = createQuerySource(emailProviders);

  // =====================
  // Dynamic fields config
  // =====================
  const providerFields: Record<
    string,
    {
      key: string;
      label: string;
      type?: string;
      description?: string;
      placeholder?: string;
    }[]
  > = {
    sendgrid: [{ key: "apiKey", label: "API Key" }],

    mailgun: [
      { key: "apiKey", label: "API Key" },
      { key: "domain", label: "Domain" },
    ],

    postmark: [{ key: "apiKey", label: "Server Token" }],
    curl: [
      {
        key: "curl",
        label: "Curl script",
        type: "textarea",
        placeholder,
        description: `Curl script, which would be called upon sending email. 
          Kindly beaware, this is semantic rather 
          than actual bash script, so use limited features and no extra bash calls.
          <br />
          Make sure, you put the secrets, templates, credentials, here, there will be no
          other place to use. While sending an email, following variables will be
          replaced in your curl message.
          <br />
          
          %FromName%  string <br />
          %FromEmail% string <br />
          %ToName%    string <br />
          %ToEmail%   string <br />
          %Subject%   string <br />
          %Content%   string (It will be escaped with \", so use double qoutes) <br />
          `,
      },
    ],

    resend: [{ key: "apiKey", label: "API Key" }],

    smtp: [
      { key: "host", label: "Host" },
      { key: "port", label: "Port" },
      { key: "user", label: "Username" },
      { key: "pass", label: "Password", type: "password" },
    ],

    terminal: [],
  };

  const currentFields = providerFields[values.type || ""] || [];

  return (
    <>
      {/* <pre>{JSON.stringify(values, null, 2)}</pre> */}
      {/* Provider Type */}

      <FormText
        value={values.title}
        onChange={(value) => setFieldValue(`title`, value, false)}
        label={"Title"}
        hint="Title of the email provider, to search and allocate easier."
        autoFocus={!isEditing}
      />

      <FormSelect
        formEffect={{
          form,
          field: EmailProviderEntity.Fields.type,
          beforeSet(item) {
            // reset config when switching provider
            setFieldValue("config", {});
            return item.value;
          },
        }}
        keyExtractor={(item) => item.value}
        querySource={querySource}
        errorMessage={errors.type}
        label={t.mailProvider.type}
        hint={t.mailProvider.typeHint}
      />

      {/* Dynamic Config Fields */}
      {currentFields.map((field, index) => {
        if (field.type === "textarea") {
          return (
            <FormRichText
              forceBasic
              height={300}
              key={field.key}
              placeholder={field.placeholder}
              value={values.config?.[field.key] || ""}
              hint={field.description}
              autoFocus={!isEditing && index === 0}
              onChange={(value) =>
                setFieldValue(`config.${field.key}`, value, false)
              }
              dir="ltr"
              label={field.label}
            />
          );
        }

        return (
          <FormText
            key={field.key}
            value={values.config?.[field.key] || ""}
            autoFocus={!isEditing && index === 0}
            onChange={(value) =>
              setFieldValue(`config.${field.key}`, value, false)
            }
            dir="ltr"
            label={field.label}
          />
        );
      })}
    </>
  );
};
