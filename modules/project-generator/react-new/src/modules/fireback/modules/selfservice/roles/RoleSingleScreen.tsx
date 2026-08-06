import { CommonSingleManager } from "@/modules/fireback/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@/modules/fireback/components/general-entity-view/GeneralEntityView";
import { PageSection } from "@/modules/fireback/components/page-section/PageSection";
import { usePageTitle } from "@/modules/fireback/hooks/authContext";
import { useRouter } from "@/modules/fireback/hooks/useRouter";
import { useT } from "@/modules/fireback/hooks/useT";
import { RoleEntity } from "@/modules/fireback/sdk/modules/abac/RoleEntity";
import { useEffect, useState } from "react";
import { RolePermissionTree } from "./RolePermissionTree";
import { useRoleGetActionQuery } from "@/modules/fireback/sdk/abac/RoleGetAction";

export const RoleSingleScreen = () => {
  const router = useRouter();
  const uniqueId = router.query.uniqueId as string;
  const t = useT();
  const [value, setValue] = useState<string[]>([]);

  const getSingleHook = useRoleGetActionQuery({
    params: { uniqueId },
  });

  var d = getSingleHook.data?.data.item;
  usePageTitle(d?.name || "");

  useEffect(() => {
    if (Array.isArray(d?.capabilitiesListId)) {
      setValue(d?.capabilitiesListId);
    }
  }, [d?.capabilitiesListId]);

  return (
    <>
      <CommonSingleManager
        editEntityHandler={() => {
          router.push(RoleEntity.Navigation.edit(uniqueId));
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
