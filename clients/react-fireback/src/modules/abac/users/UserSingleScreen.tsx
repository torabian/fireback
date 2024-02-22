import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { usePageTitle } from "@/components/page-title/PageTitle";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { UserEntity } from "src/sdk/fireback";
import { useGetUserByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetUserByUniqueId";
import { UserNavigationTools } from "src/sdk/fireback/modules/workspaces/user-navigation-tools";

export const UserSingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetUserByUniqueId({ query: { uniqueId } });
  var d: UserEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.firstName || "");

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(UserNavigationTools.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.users.firstName,
              elem: d?.firstName,
            },
            {
              label: t.users.lastName,
              elem: d?.lastName,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
