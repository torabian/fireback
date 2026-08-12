import { CommonSingleManager } from "@fireback/ui-core/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "@fireback/ui-core/components/general-entity-view/GeneralEntityView";
import { PageSection } from "@fireback/ui-core/components/page-section/PageSection";
import { useRouter } from "@fireback/ui-core/hooks/useRouter";
import { useS } from "@fireback/ui-core/hooks/useS";
import { strings } from "./strings/translations";
import { RoleNavigation } from "@fireback/ui-core/sdk/navigation/AbacNavigation";
import { useEffect, useState } from "react";
import { RolePermissionTree } from "./RolePermissionTree";
import { useRoleGetActionQuery } from "@fireback/selfservice/sdk/abac/RoleGetAction";
import { usePageTitle } from "@fireback/ui-core/components/page-title/PageTitle";

export const RoleSingleScreen = () => {
  const router = useRouter();
  const uniqueId = router.query.uniqueId as string;
  const s = useS(strings);
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
          router.push(RoleNavigation.edit(uniqueId));
        }}
        getSingleHook={getSingleHook}
      >
        <GeneralEntityView
          entity={d}
          fields={[
            {
              label: s.role.name,
              elem: d?.name,
            },
          ]}
        />

        <PageSection title={s.role.permissions} className="mt-3">
          <RolePermissionTree value={value} />
        </PageSection>
      </CommonSingleManager>
    </>
  );
};
