import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { PageSection } from "@/modules/fireback/components/page-section/PageSection";
import { usePageTitle } from "@/modules/fireback/components/page-title/PageTitle";
import { useLocale } from "@/modules/fireback/hooks/useLocale";
import { useT } from "@/modules/fireback/hooks/useT";
import { useEffect, useState } from "react";
import { useQueryClient } from "react-query";
import { RolePermissionTree } from "./RolePermissionTree";
import { useGetRoleByUniqueId } from "../../sdk/modules/workspaces/useGetRoleByUniqueId";
import { RoleEntity } from "../../sdk/modules/workspaces/RoleEntity";

export const RoleSingleScreen = () => {
  const router = useRouter();
  const queryClient = useQueryClient();
  const uniqueId = router.query.uniqueId as string;
  const t = useT();
  const { locale } = useLocale();
  const [value, setValue] = useState<string[]>([]);

  const getSingleHook = useGetRoleByUniqueId({
    query: { uniqueId, deep: true },
  });
  var d: RoleEntity | undefined = getSingleHook.query.data?.data;
  usePageTitle(d?.name || "");

  useEffect(() => {
    setValue(d?.capabilities?.map((t) => t.uniqueId || "") as any);
  }, [d?.capabilities]);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(RoleEntity.Navigation.edit(uniqueId, locale));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: t.role.name,
              elem: d?.name,
            },
          ]}
        />

        <PageSection title={t.role.permissions} className="mt-3">
          <RolePermissionTree value={value} />
        </PageSection>
      </CommonSingleManager>
    </>
  );
};
