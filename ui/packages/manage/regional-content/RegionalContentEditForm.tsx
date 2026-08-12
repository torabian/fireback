import { FormRichText } from "@fireback/ui-core/components/forms/form-richtext/FormRichText";
import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { createQuerySource } from "@fireback/ui-core/hooks/useAsQuery";
import { useS } from "@fireback/ui-core/hooks/useS";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { RegionalContentDto } from "@fireback/manage/sdk/abac/RegionalContentDto";
import { strings } from "./strings/translations";

export const RegionalContentForm = ({
  form,
  isEditing,
}: EntityFormProps<RegionalContentDto>) => {
  const { values, setValues, setFieldValue, errors } = form;
  const s = useS(strings);

  // Bug fix: this used to be createQuerySource() with no items at all, so
  // the select below always rendered with zero options - there was no way
  // to actually pick a keyGroup from the UI (keyGroup is validate:"required"
  // on the backend, see Abac.emi.yml, so every create from this form failed
  // validation). Mirrors the enum values declared there (SMS_OTP/EMAIL_OTP).
  const keyGroupSource = createQuerySource([
    { value: "SMS_OTP", label: s.regionalContents.keyGroupSmsOtp },
    { value: "EMAIL_OTP", label: s.regionalContents.keyGroupEmailOtp },
  ]);

  return (
    <>
      <FormSelect
        keyExtractor={(t) => t.value}
        fnLabelFormat={(t) => t.label}
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

      {/* Bug fix: this used to be hardcoded to value={"global"} and
          readonly, so (a) there was no way to target a specific region from
          the UI at all - the whole point of a "regional" content entity -
          and (b) editing an existing record showed the fake "global" value
          instead of whatever region was actually stored, e.g. "us"/"eu"/
          "any". Now a normal editable field bound to the real value, same
          as every other field on this form. */}
      <FormText
        value={values.region}
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
