import { CommonSingleManager } from "../../../components/entity-manager/CommonSingleManager";
import { GeneralEntityView } from "../../../components/general-entity-view/GeneralEntityView";
import { PageSection } from "../../../components/page-section/PageSection";
import { usePageTitle } from "../../../hooks/authContext";
import { useRouter } from "../../../hooks/useRouter";
import { useT } from "../../../hooks/useT";
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
