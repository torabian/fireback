import { EntityFormProps } from "@/definitions/definitions";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useContext } from "react";
import { AppMenuEntity } from "src/sdk/fireback/modules/workspaces/AppMenuEntity";
import { FormText } from "@/components/forms/form-text/FormText";
import { FormEntitySelect3 } from "@/components/forms/form-select/FormEntitySelect3";
import { useGetCapabilities } from "@/sdk/fireback/modules/workspaces/useGetCapabilities";
export const AppMenuForm = ({
  form,
  isEditing,
}: EntityFormProps<AppMenuEntity>) => {
  const { options } = useContext(RemoteQueryContext);
  const { values, setValues, setFieldValue, errors } = form;
  const t = useT();
  return (
    <>
      <FormText
        value={values.href}
        onChange={(value) =>
          setFieldValue(AppMenuEntity.Fields.href, value, false)
        }
        errorMessage={errors.href}
        label={t.appMenus.href}
        hint={t.appMenus.hrefHint}
      />
      <FormText
        value={values.icon}
        onChange={(value) =>
          setFieldValue(AppMenuEntity.Fields.icon, value, false)
        }
        errorMessage={errors.icon}
        label={t.appMenus.icon}
        hint={t.appMenus.iconHint}
      />
      <FormText
        value={values.label}
        onChange={(value) =>
          setFieldValue(AppMenuEntity.Fields.label, value, false)
        }
        errorMessage={errors.label}
        label={t.appMenus.label}
        hint={t.appMenus.labelHint}
      />
      <FormText
        value={values.activeMatcher}
        onChange={(value) =>
          setFieldValue(AppMenuEntity.Fields.activeMatcher, value, false)
        }
        errorMessage={errors.activeMatcher}
        label={t.appMenus.activeMatcher}
        hint={t.appMenus.activeMatcherHint}
      />
      <FormText
        value={values.applyType}
        onChange={(value) =>
          setFieldValue(AppMenuEntity.Fields.applyType, value, false)
        }
        errorMessage={errors.applyType}
        label={t.appMenus.applyType}
        hint={t.appMenus.applyTypeHint}
      />
      <FormEntitySelect3
        formEffect={{ form, field: AppMenuEntity.Fields.capability$ }}
        useQuery={useGetCapabilities}
        label={t.appMenus.capability}
        hint={t.appMenus.capabilityHint}
      />
    </>
  );
};
