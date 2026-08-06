import { FormSelect } from "../../../components/forms/form-select/FormSelect";
import { FormText } from "../../../components/forms/form-text/FormText";
import { type EntityFormProps } from "../../../definitions/definitions";
import { createQuerySource } from "../../../hooks/useAsQuery";
import { useS } from "../../../hooks/useS";
import { PassportEntity } from "../../../sdk/modules/abac/PassportEntity";
import { getPassportTypes } from "./PassportCommon";
import { strings } from "./strings/translations";

export const PassportEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PassportEntity>>) => {
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
              field: PassportEntity.Fields.type,
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
              setFieldValue(PassportEntity.Fields.value, value, false)
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
