import { useContext } from "react";
import { FormEntitySelect3 } from "../../components/forms/form-select/FormEntitySelect3";
import { FormText } from "../../components/forms/form-text/FormText";
import { EntityFormProps } from "../../definitions/definitions";
import { useS } from "../../hooks/useS";
import { RemoteQueryContext } from "../../sdk/core/react-tools";
import { PassportEntity } from "../../sdk/modules/workspaces/PassportEntity";
import { useGetPassportMethods } from "../../sdk/modules/workspaces/useGetPassportMethods";
import { strings } from "./strings/translations";
import { PassportMethodEntity } from "../../sdk/modules/workspaces/PassportMethodEntity";
import { castArrayToUseQuery } from "../../hooks/castArrayToUseQuery";
import { getPassportTypes } from "./PassportCommon";

export const PassportEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<PassportEntity>>) => {
  const { values, setFieldValue, errors, setValues } = form;
  const { options } = useContext(RemoteQueryContext);
  const s = useS(strings);

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <FormEntitySelect3
            onChange={(value) =>
              setFieldValue(PassportEntity.Fields.type, value.uniqueId)
            }
            useQuery={castArrayToUseQuery(getPassportTypes())}
            label="Type"
            hint="Passport methods which are available in this app"
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
