import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "../../../../fireback-ui/components/page-title/PageTitle";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useS } from "../../../../fireback-ui/hooks/useS";
import { strings } from "./strings/translations";
import { UserDto } from "../../../sdk/abac/UserDto";
import { UserNavigation } from "../../../sdk/navigation/AbacNavigation";
import { useUserGetActionQuery } from "../../../sdk/abac/UserGetAction";
import { UserPassportList } from "./UserPassportsList";

export const UserSingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useUserGetActionQuery({ params: { uniqueId } });
  var d: UserDto | undefined = getSingleHook.query.data?.data;
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
              label: s.firstName,
              elem: d?.firstName,
            },
            {
              label: s.lastName,
              elem: d?.lastName,
            },
          ]}
        />

        <UserPassportList userId={uniqueId} />
      </CommonSingleManager>
    </>
  );
};
