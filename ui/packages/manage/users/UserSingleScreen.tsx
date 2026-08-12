import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { UserDto } from "@fireback/manage/sdk/abac/UserDto";
import { UserNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { useUserGetActionQuery } from "@fireback/manage/sdk/abac/UserGetAction";
import { UserPassportList } from "./UserPassportsList";
import { UserPhotoThumbnail } from "./UserPhotoThumbnail";

export const UserSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useUserGetActionQuery({ params: { uniqueId } });
  var d: UserDto | undefined = getSingleHook.data?.data?.item;
  usePageTitle(d?.firstName || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(UserNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.photo,
              copyableContent: d?.photo || "",
              elem: (
                <UserPhotoThumbnail
                  photo={d?.photo}
                  style={{
                    width: "64px",
                    height: "64px",
                    borderRadius: "50%",
                    objectFit: "cover",
                  }}
                />
              ),
            },
            {
              label: s.firstName,
              elem: d?.firstName,
            },
            {
              label: s.lastName,
              elem: d?.lastName,
            },
            {
              label: s.phoneNumber,
              elem: d?.phoneNumber,
            },
            {
              label: s.jobTitle,
              elem: d?.jobTitle,
            },
            {
              label: s.company,
              elem: d?.company,
            },
            {
              label: s.bio,
              elem: d?.bio,
            },
            {
              label: s.addressLine1,
              elem: d?.primaryAddress?.addressLine1,
            },
            {
              label: s.addressLine2,
              elem: d?.primaryAddress?.addressLine2,
            },
            {
              label: s.cityName,
              elem: d?.primaryAddress?.city,
            },
          ]}
        />

        <UserPassportList userId={uniqueId} />
      </CommonSingleManager>
    </>
  );
};
