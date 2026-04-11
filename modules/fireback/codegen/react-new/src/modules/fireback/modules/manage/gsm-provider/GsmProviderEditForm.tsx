import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { useContext } from "react";
import { GsmProviderEntity } from "@/modules/fireback/sdk/modules/abac/GsmProviderEntity";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
export const GsmProviderForm = ({
  form,
  isEditing,
}: EntityFormProps<GsmProviderEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);
  return (
    <>
        <FormText
          value={values.apiKey }
          onChange={(value) => setFieldValue(GsmProviderEntity.Fields.apiKey, value, false)}
          errorMessage={errors.apiKey }
          label={s.gsmProviders.apiKey }
          hint={s.gsmProviders.apiKeyHint}
        />
        <FormText
          value={values.mainSenderNumber }
          onChange={(value) => setFieldValue(GsmProviderEntity.Fields.mainSenderNumber, value, false)}
          errorMessage={errors.mainSenderNumber }
          label={s.gsmProviders.mainSenderNumber }
          hint={s.gsmProviders.mainSenderNumberHint}
        />
        {/*
          <FormText
            type="?"
            value={values.type }
            onChange={(value) => setFieldValue(GsmProviderEntity.Fields.type, value, false)}
            errorMessage={errors.type }
            label={s.gsmProviders.type }
            hint={s.gsmProviders.typeHint}
          />
         */}
        <FormText
          value={values.invokeUrl }
          onChange={(value) => setFieldValue(GsmProviderEntity.Fields.invokeUrl, value, false)}
          errorMessage={errors.invokeUrl }
          label={s.gsmProviders.invokeUrl }
          hint={s.gsmProviders.invokeUrlHint}
        />
        <FormText
          value={values.invokeBody }
          onChange={(value) => setFieldValue(GsmProviderEntity.Fields.invokeBody, value, false)}
          errorMessage={errors.invokeBody }
          label={s.gsmProviders.invokeBody }
          hint={s.gsmProviders.invokeBodyHint}
        />
    </>
  );
};