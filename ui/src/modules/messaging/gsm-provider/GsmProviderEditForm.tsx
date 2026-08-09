import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { createQuerySource } from "../../fireback-ui/hooks/useAsQuery";
import { useS } from "../../fireback-ui/hooks/useS";
import { GsmProviderDto } from "../../sdk/messaging/GsmProviderDto";
import { strings } from "./strings/translations";
export const GsmProviderForm = ({
  form,
  isEditing,
}: EntityFormProps<GsmProviderDto>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  const gsmTypes = createQuerySource([
    { uniqueId: "terminal", title: s.gsmProviders.typeTerminal },
    { uniqueId: "url", title: s.gsmProviders.typeUrl },
  ]);
  return (
    <>
      <FormText
        value={values.apiKey}
        onChange={(value) =>
          setFieldValue(GsmProviderDto.Fields.apiKey, value, false)
        }
        errorMessage={errors.apiKey}
        label={s.gsmProviders.apiKey}
        hint={s.gsmProviders.apiKeyHint}
      />
      <FormText
        value={values.mainSenderNumber}
        onChange={(value) =>
          setFieldValue(GsmProviderDto.Fields.mainSenderNumber, value, false)
        }
        errorMessage={errors.mainSenderNumber}
        label={s.gsmProviders.mainSenderNumber}
        hint={s.gsmProviders.mainSenderNumberHint}
      />

      <FormSelect
        querySource={gsmTypes}
        value={values.type}
        fnLabelFormat={(item) => item.title}
        keyExtractor={(item) => item.uniqueId}
        onChange={(value) =>
          setFieldValue(GsmProviderDto.Fields.type, value.uniqueId, false)
        }
        errorMessage={errors.type}
        label={s.gsmProviders.type}
        hint={s.gsmProviders.typeHint}
      />

      <FormText
        value={values.invokeUrl}
        onChange={(value) =>
          setFieldValue(GsmProviderDto.Fields.invokeUrl, value, false)
        }
        errorMessage={errors.invokeUrl}
        label={s.gsmProviders.invokeUrl}
        hint={s.gsmProviders.invokeUrlHint}
      />
      <FormText
        value={values.invokeBody}
        onChange={(value) =>
          setFieldValue(GsmProviderDto.Fields.invokeBody, value, false)
        }
        errorMessage={errors.invokeBody}
        label={s.gsmProviders.invokeBody}
        hint={s.gsmProviders.invokeBodyHint}
      />
    </>
  );
};
