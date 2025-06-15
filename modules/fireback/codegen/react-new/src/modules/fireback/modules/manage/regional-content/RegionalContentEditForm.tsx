import { EntityFormProps } from "@/modules/fireback/definitions/definitions";
import { RemoteQueryContext } from "@/modules/fireback/sdk/core/react-tools";
import { useContext } from "react";
import { RegionalContentEntity } from "@/modules/fireback/sdk/modules/abac/RegionalContentEntity";
import { FormText } from "@/modules/fireback/components/forms/form-text/FormText";
import { FormSelect } from "@/modules/fireback/components/forms/form-select/FormSelect";
import { useS } from "@/modules/fireback/hooks/useS";
import { strings } from "./strings/translations";
import { FormRichText } from "@/modules/fireback/components/forms/form-richtext/FormRichText";
import { createQuerySource } from "@/modules/fireback/hooks/useAsQuery";

export const RegionalContentForm = ({
  form,
  isEditing,
}: EntityFormProps<RegionalContentEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  const keyGroupSource = createQuerySource(
    RegionalContentEntity.definition.fields
      .find((field) => field.name === "keyGroup")
      .of.map((item) => {
        return {
          label: item.k,
          value: item.k,
        };
      })
  );

  return (
    <>
      <FormSelect
        keyExtractor={(t) => t.value}
        formEffect={{
          form,
          field: RegionalContentEntity.Fields.keyGroup,
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
          setFieldValue(RegionalContentEntity.Fields.content, value, false)
        }
        errorMessage={errors.content}
        label={s.regionalContents.content}
        hint={s.regionalContents.contentHint}
      />

      <FormText
        value={"global"}
        readonly
        onChange={(value) =>
          setFieldValue(RegionalContentEntity.Fields.region, value, false)
        }
        errorMessage={errors.region}
        label={s.regionalContents.region}
        hint={s.regionalContents.regionHint}
      />
      <FormText
        value={values.title}
        onChange={(value) =>
          setFieldValue(RegionalContentEntity.Fields.title, value, false)
        }
        errorMessage={errors.title}
        label={s.regionalContents.title}
        hint={s.regionalContents.titleHint}
      />
      <FormText
        value={values.languageId}
        onChange={(value) =>
          setFieldValue(RegionalContentEntity.Fields.languageId, value, false)
        }
        errorMessage={errors.languageId}
        label={s.regionalContents.languageId}
        hint={s.regionalContents.languageIdHint}
      />
    </>
  );
};
