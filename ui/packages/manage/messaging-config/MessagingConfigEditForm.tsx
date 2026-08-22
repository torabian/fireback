import { type EntityFormProps } from "@fireback/ui-core/types/EntityManagement";
import { useS } from "@fireback/ui-core/hooks/useS";
import { MessagingConfigDto } from "@fireback/manage/sdk/messaging/MessagingConfigDto";
import { strings } from "./strings/translations";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { FormSelect } from "@fireback/ui-core/components/forms/form-select/FormSelect";
import { useEmailProvidersQuerySource } from "@fireback/ui-core/hooks/useEmailProvidersQuerySource";
import { useGsmProvidersQuerySource } from "@fireback/ui-core/hooks/useGsmProvidersQuerySource";
import { useRegionalContentsQuerySource } from "@fireback/ui-core/hooks/useRegionalContentsQuerySource";
import { type RegionalContentDto } from "@fireback/manage/sdk/abac/RegionalContentDto";

// Same fix WorkspaceConfigEditForm.tsx carries for its own copy of these three
// dropdowns: keyGroup + languageId always identify a regionalContent row
// meaningfully, even for rows (every OTP default - see RegionalContentDefaults.go)
// that never set a title at all.
function regionalContentOptionLabel(item: RegionalContentDto) {
  const detail = item.title || item.content?.slice(0, 40) || item.region;
  return `${item.keyGroup} - ${item.languageId}${detail ? ` (${detail})` : ""}`;
}

export const MessagingConfigForm = ({
  form,
  isEditing,
}: EntityFormProps<MessagingConfigDto>) => {
  const { errors } = form;
  const s = useS(strings);
  return (
    <PageSection
      title={s.messagingConfigs.notificationsSectionTitle}
      description={s.messagingConfigs.notificationsSectionDescription}
    >
      <FormSelect
        keyExtractor={(t) => t.uniqueId}
        formEffect={{
          form,
          field: MessagingConfigDto.Fields.generalEmailProviderId,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        fnLabelFormat={(e) => `${e.type} (${e.uniqueId})`}
        querySource={useEmailProvidersQuerySource}
        nullable
        errorMessage={errors.generalEmailProviderId}
        label={s.messagingConfigs.generalEmailProviderLabel}
        hint={s.messagingConfigs.generalEmailProviderHint}
      />

      <FormSelect
        keyExtractor={(t) => t.uniqueId}
        formEffect={{
          form,
          field: MessagingConfigDto.Fields.generalGsmProviderId,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        fnLabelFormat={(e) => `${e.type} (${e.uniqueId})`}
        querySource={useGsmProvidersQuerySource}
        nullable
        errorMessage={errors.generalGsmProviderId}
        label={s.messagingConfigs.generalGsmProviderLabel}
        hint={s.messagingConfigs.generalGsmProviderHint}
      />

      <FormSelect
        keyExtractor={(t) => t.uniqueId}
        formEffect={{
          form,
          field: MessagingConfigDto.Fields.inviteToWorkspaceContentId,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        fnLabelFormat={regionalContentOptionLabel}
        querySource={useRegionalContentsQuerySource}
        nullable
        errorMessage={errors.inviteToWorkspaceContentId}
        label={s.messagingConfigs.inviteToWorkspaceContentLabel}
        hint={s.messagingConfigs.inviteToWorkspaceContentHint}
      />

      <FormSelect
        keyExtractor={(t) => t.uniqueId}
        formEffect={{
          form,
          field: MessagingConfigDto.Fields.emailOtpContentId,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        fnLabelFormat={regionalContentOptionLabel}
        querySource={useRegionalContentsQuerySource}
        nullable
        errorMessage={errors.emailOtpContentId}
        label={s.messagingConfigs.emailOtpContentLabel}
        hint={s.messagingConfigs.emailOtpContentHint}
      />
      <FormSelect
        keyExtractor={(t) => t.uniqueId}
        formEffect={{
          form,
          field: MessagingConfigDto.Fields.smsOtpContentId,
          beforeSet(item) {
            return item.uniqueId;
          },
        }}
        fnLabelFormat={regionalContentOptionLabel}
        querySource={useRegionalContentsQuerySource}
        nullable
        errorMessage={errors.smsOtpContentId}
        label={s.messagingConfigs.smsOtpContentLabel}
        hint={s.messagingConfigs.smsOtpContentHint}
      />
    </PageSection>
  );
};
