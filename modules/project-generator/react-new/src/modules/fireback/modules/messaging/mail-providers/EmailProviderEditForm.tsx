import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { type EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";
import { useS } from "@/modules/fireback/hooks/useS";
import { useT } from "@/modules/fireback/hooks/useT";
import { EmailProviderDto } from "@/modules/fireback/sdk/messaging/EmailProviderDto";
import { strings } from "./strings/translations";

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
}: EntityFormProps<EmailProviderDto>) => {
  const { values, setFieldValue, errors } = form;
  const t = useT();
  const s = useS(strings);

  // =====================
  // Providers
  // =====================
  const emailProviders = [
    { label: s.emailProviders.typeSendgrid, value: "sendgrid" },
    { label: s.emailProviders.typeMailgun, value: "mailgun" },
    { label: s.emailProviders.typePostmark, value: "postmark" },
    { label: s.emailProviders.typeResend, value: "resend" },
    { label: s.emailProviders.typeCurl, value: "curl" },
    { label: s.emailProviders.typeSmtp, value: "smtp" },
    { label: s.emailProviders.typeTerminal, value: "terminal" },
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
    sendgrid: [{ key: "apiKey", label: s.emailProviders.fieldApiKey }],

    mailgun: [
      { key: "apiKey", label: s.emailProviders.fieldApiKey },
      { key: "domain", label: s.emailProviders.fieldDomain },
    ],

    postmark: [{ key: "apiKey", label: s.emailProviders.fieldServerToken }],
    curl: [
      {
        key: "curl",
        label: s.emailProviders.fieldCurlScript,
        type: "textarea",
        placeholder,
        description: s.emailProviders.fieldCurlScriptDescription,
      },
    ],

    resend: [{ key: "apiKey", label: s.emailProviders.fieldApiKey }],

    smtp: [
      { key: "host", label: s.emailProviders.fieldHost },
      { key: "port", label: s.emailProviders.fieldPort },
      { key: "user", label: s.emailProviders.fieldUsername },
      { key: "pass", label: s.emailProviders.fieldPassword, type: "password" },
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
        label={s.emailProviders.title}
        hint={s.emailProviders.titleHint}
        autoFocus={!isEditing}
      />

      <FormSelect
        formEffect={{
          form,
          field: EmailProviderDto.Fields.type,
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
