import { useRouter } from "@/Router";
import { CommonSingleManager } from "@/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@/hooks/useLocale";
import { useT } from "@/hooks/useT";
import { PublicJoinKeyEntity } from "src/sdk/fireback";
import { useGetPublicJoinKeyByUniqueId } from "src/sdk/fireback/modules/workspaces/useGetPublicJoinKeyByUniqueId";

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
