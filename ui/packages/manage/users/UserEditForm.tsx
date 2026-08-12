import { FormText } from "@fireback/ui-core/components/forms/form-text/FormText";
import { BaseFormElement } from "@fireback/ui-core/components/forms/base-form-element/BaseFormElement";
import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { useS } from "@fireback/ui-core/hooks/useS";
import { UserDto } from "@fireback/manage/sdk/abac/UserDto";
import { strings } from "./strings/translations";
import {
  FileUploader,
  UploaderConfigProvider,
} from "@fireback/ui-core/components/resumable-uploader";
import { useStorageUploaderConfig } from "@fireback/ui-core/hooks/useStorageUploaderConfig";
import { useAuthentication } from "@fireback/auth-client/AuthenticationContext";

export const UserEditForm = ({
  form,
  isEditing,
}: EntityFormProps<Partial<UserDto>>) => {
  const { values, setFieldValue, errors, setValues } = form;
  const { headers } = useAuthentication();

  const s = useS(strings);
  const uploaderConfig = useStorageUploaderConfig({
    validateFile: [{ mimeStartsWith: "image/", maxSize: 5 * 1024 * 1024 }],
    headers: {
      authorization: headers.authorization,
      ["workspace-id"]: headers["workspace-id"],
    },
  });

  return (
    <>
      <div className="row">
        <div className="col-md-12">
          <BaseFormElement label={s.photo} hint={s.photoHint}>
            <UploaderConfigProvider config={uploaderConfig}>
              <FileUploader
                multiple={false}
                value={values?.photo || null}
                onChange={(file) =>
                  setFieldValue(UserDto.Fields.photo, file?.id ?? "", false)
                }
              />
            </UploaderConfigProvider>
          </BaseFormElement>
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.firstName}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.firstName, value, false)
            }
            autoFocus={!isEditing}
            errorMessage={errors?.firstName}
            label={s.firstName}
            hint={s.inviteFirstNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.lastName}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.lastName, value, false)
            }
            errorMessage={errors?.lastName}
            label={s.lastName}
            hint={s.inviteLastNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.city}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.primaryAddress.city, value, false)
            }
            errorMessage={errors?.primaryAddress?.city}
            label={s.cityName}
            hint={s.cityNameHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine1}
            onChange={(value) =>
              setFieldValue(
                UserDto.Fields.primaryAddress.addressLine1,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine1}
            label={s.addressLine1}
            hint={s.addressLine1Hint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.primaryAddress?.addressLine2}
            onChange={(value) =>
              setFieldValue(
                UserDto.Fields.primaryAddress.addressLine2,
                value,
                false,
              )
            }
            errorMessage={errors?.primaryAddress?.addressLine2}
            label={s.addressLine2}
            hint={s.addressLine2Hint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.phoneNumber}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.phoneNumber, value, false)
            }
            errorMessage={errors?.phoneNumber}
            label={s.phoneNumber}
            hint={s.phoneNumberHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.jobTitle}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.jobTitle, value, false)
            }
            errorMessage={errors?.jobTitle}
            label={s.jobTitle}
            hint={s.jobTitleHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.company}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.company, value, false)
            }
            errorMessage={errors?.company}
            label={s.company}
            hint={s.companyHint}
          />
        </div>
        <div className="col-md-12">
          <FormText
            value={values?.bio}
            onChange={(value) =>
              setFieldValue(UserDto.Fields.bio, value, false)
            }
            errorMessage={errors?.bio}
            label={s.bio}
            hint={s.bioHint}
          />
        </div>
      </div>
    </>
  );
};
