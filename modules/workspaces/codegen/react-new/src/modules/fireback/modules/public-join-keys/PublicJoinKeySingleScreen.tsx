import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";
import { useGetPublicJoinKeyByUniqueId } from "../../sdk/modules/workspaces/useGetPublicJoinKeyByUniqueId";
import { PublicJoinKeyEntity } from "../../sdk/modules/workspaces/PublicJoinKeyEntity";

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
