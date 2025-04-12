import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useT } from "@/modules/fireback/hooks/useT";
import { PublicJoinKeyEntity } from "@/modules/fireback/sdk/modules/abac/PublicJoinKeyEntity";
import { useGetPublicJoinKeyByUniqueId } from "@/modules/fireback/sdk/modules/abac/useGetPublicJoinKeyByUniqueId";

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
          router.push(PublicJoinKeyEntity.Navigation.edit(uniqueId));
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
