import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { createQuerySource } from "@fireback/ui-core/hooks/useAsQuery";
import { useS } from "@fireback/ui-core/hooks/useS";
import { PassportDto } from "@fireback/manage/sdk/abac/PassportDto";
import { getPassportTypes } from "./PassportCommon";
import { strings } from "./strings/translations";

export const PassportEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PassportDto>>) => {
  const { values, setFieldValue, errors, setValues } = form;
  const s = useS(strings);
  const passportTypesQuery = createQuerySource(getPassportTypes(s));

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormSelect
            formEffect={{
              form,
              field: PassportDto.Fields.type,
              beforeSet(item) {
                return item.uniqueId;
              },
            }}
            querySource={passportTypesQuery}
            label={s.type}
            hint={s.typeHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values.value}
            onChange={(value) =>
              setFieldValue(PassportDto.Fields.value, value, false)
            }
            autoFocus={!isEditing}
            label={s.value}
            hint={s.valueHint}
          />
        </div>
      </div>
    </>
  );
};
