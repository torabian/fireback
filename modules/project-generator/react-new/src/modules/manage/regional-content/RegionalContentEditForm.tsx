import { useContext } from "react";
import { FormRichText } from "../../fireback-ui/components/forms/form-richtext/FormRichText";
import { FormSelect } from "../../fireback-ui/components/forms/form-select/FormSelect";
import { FormText } from "../../fireback-ui/components/forms/form-text/FormText";
import { createQuerySource } from "../../fireback-ui/hooks/useAsQuery";
import { useS } from "../../fireback-ui/hooks/useS";
import { type EntityFormProps } from "../../fireback-ui/types/EntityManagement";
import { RegionalContentDto } from "../../sdk/abac/RegionalContentDto";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { strings } from "./strings/translations";

export const RegionalContentForm = ({
  form,
  isEditing,
}: EntityFormProps<RegionalContentDto>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  const keyGroupSource = createQuerySource(
    RegionalContentDto.definition.fields
      .find((field) => field.name === "keyGroup")
      .of.map((item) => {
        return {
          label: item.k,
          value: item.k,
        };
      }),
  );

  return (
    <>
      <FormSelect
        keyExtractor={(t) => t.value}
        formEffect={{
          form,
          field: RegionalContentDto.Fields.keyGroup,
          beforeSet(item) {
            return item.value;
          },
        }}
        querySource={keyGroupSource}
        errorMessage={errors.keyGroup}
        label={s.regionalContents.keyGroup}
        hint={s.regionalContents.keyGroupHint}
      />
      <FormRichText
        value={values.content}
        forceRich={values.keyGroup === "EMAIL_OTP"}
        forceBasic={values.keyGroup === "SMS_OTP"}
        onChange={(value) =>
          setFieldValue(RegionalContentDto.Fields.content, value, false)
        }
        errorMessage={errors.content}
        label={s.regionalContents.content}
        hint={s.regionalContents.contentHint}
      />

      <FormText
        value={"global"}
        readonly
        onChange={(value) =>
          setFieldValue(RegionalContentDto.Fields.region, value, false)
        }
        errorMessage={errors.region}
        label={s.regionalContents.region}
        hint={s.regionalContents.regionHint}
      />
      <FormText
        value={values.title}
        onChange={(value) =>
          setFieldValue(RegionalContentDto.Fields.title, value, false)
        }
        errorMessage={errors.title}
        label={s.regionalContents.title}
        hint={s.regionalContents.titleHint}
      />
      <FormText
        value={values.languageId}
        onChange={(value) =>
          setFieldValue(RegionalContentDto.Fields.languageId, value, false)
        }
        errorMessage={errors.languageId}
        label={s.regionalContents.languageId}
        hint={s.regionalContents.languageIdHint}
      />
    </>
  );
};
