import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@fireback/ui-core/hooks/useLocale";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { PublicJoinKeyDto } from "@fireback/selfservice/sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { usePublicJoinKeyGetActionQuery } from "@fireback/selfservice/sdk/abac/PublicJoinKeyGetAction";

export const PublicJoinKeySingleScreen = () => {
  const router = useRouter();
  const s = useS(strings);
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetPublicJoinKeyByUniqueId({
    query: { uniqueId },
  });

  var d: PublicJoinKeyDto | undefined = getSingleHook.data?.data?.item;

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(PublicJoinKeyNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.roleName,
              elem: d?.role?.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
