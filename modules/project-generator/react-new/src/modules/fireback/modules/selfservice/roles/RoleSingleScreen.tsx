import { CommonSingleManager } from "../../../../fireback-ui/components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../../fireback-ui/components/general-entity-view/GeneralEntityView";
import { PageSection } from "../../../../fireback-ui/components/page-section/PageSection";
import { usePageTitle } from "../../../../fireback-ui/hooks/authContext";
import { useRouter } from "../../../../fireback-ui/hooks/useRouter";
import { useT } from "../../../../fireback-ui/hooks/useT";
import { RoleNavigation } from "../../../sdk/navigation/AbacNavigation";
import { useEffect, useState } from "react";
import { RolePermissionTree } from "./RolePermissionTree";
import { useRoleGetActionQuery } from "../../../sdk/abac/RoleGetAction";

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
          router.push(RoleNavigation.edit(uniqueId));
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
