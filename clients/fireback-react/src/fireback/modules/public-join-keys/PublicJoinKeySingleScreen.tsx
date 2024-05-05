import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/fireback/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@/fireback/hooks/useLocale";
import { useT } from "@/fireback/hooks/useT";
import { useGetPublicJoinKeyByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetPublicJoinKeyByUniqueId";
import { PublicJoinKeyEntity } from "@/sdk/fireback/modules/workspaces/PublicJoinKeyEntity";

export const PublicJoinKeySingleScreen = () => {
  const router = useRouter();
  const t = useT();
  const uniqueId = router.query.uniqueId as string;
  const { locale } = useLocale();

  const getSingleHook = useGetPublicJoinKeyByUniqueId({
    query: { uniqueId },
  });

  var d: PublicJoinKeyEntity | undefined = getSingleHook.query.data?.data;

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(`/${locale}/publicjoinkey/edit/${uniqueId}`);
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
