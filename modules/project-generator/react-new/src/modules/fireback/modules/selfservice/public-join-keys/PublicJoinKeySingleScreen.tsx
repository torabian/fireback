import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { useLocale } from "../../../../fireback-ui/hooks/useLocale";
import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { PublicJoinKeyDto } from "../../../sdk/abac/PublicJoinKeyDto";
import { PublicJoinKeyNavigation } from "../../../sdk/navigation/AbacNavigation";
import { usePublicJoinKeyGetActionQuery } from "../../../sdk/abac/PublicJoinKeyGetAction";

export const PublicJoinKeySingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetPublicJoinKeyByUniqueId({
    query: { uniqueId },
  });

  var d: PublicJoinKeyDto | undefined = getSingleHook.query.data?.data;

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
              label: t.role.name,
              elem: d?.role?.name,
            },
          ]}
        />
      </CommonSingleManager>
    </>
  );
};
