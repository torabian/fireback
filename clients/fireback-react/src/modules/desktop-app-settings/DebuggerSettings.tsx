import { FormCheckbox } from "@/components/forms/form-switch/FormSwitch";
import Link from "@/components/link/Link";
import { PageSection } from "@/components/page-section/PageSection";
import { useT } from "@/hooks/useT";
import { RemoteQueryContext as FirebackContext } from "src/sdk/fireback/core/react-tools";
import { RemoteQueryContext } from "src/sdk/fireback/core/react-tools";
import { useGetUserRoleWorkspaces } from "src/sdk/fireback/modules/workspaces/useGetUserRoleWorkspaces";
import { useContext, useState } from "react";
import { useQueryClient } from "react-query";

function UserRoleWorkspaceDebug() {
  const queryClient = useQueryClient();
  const { query: queryWorkspaces } = useGetUserRoleWorkspaces({
    queryClient,
    query: {},
    queryOptions: {
      cacheTime: 0,
    },
  });

  return (
    <>
      <h2>User Role Workspaces</h2>
      <p>Data:</p>
      <pre>{JSON.stringify(queryWorkspaces.data, null, 2)}</pre>
      <p>Error:</p>
      <pre>{JSON.stringify(queryWorkspaces.error, null, 2)}</pre>
    </>
  );
}

function SessionDebug() {
  const fireback = useContext(RemoteQueryContext);

  return (
    <>
      <h2>Fireback context:</h2>
      <pre>{JSON.stringify(fireback, null, 2)}</pre>
    </>
  );
}

export function DebuggerSettings({}: {}) {
  const [debugVisible, setDebugVisible] = useState(false);
  const firebackContext = useContext(FirebackContext);

  const t = useT();

  return (
    <PageSection title={t.generalSettings.debugSettings.title}>
      <p>{t.generalSettings.debugSettings.description}</p>

      <FormCheckbox
        value={debugVisible}
        label={t.debugInfo}
        onChange={() => setDebugVisible((m) => !m)}
      />
      {debugVisible && (
        <>
          <pre></pre>
          <Link href="/lalaland">Go to Lalaland</Link>
          <Link href="/view3d">View 3D</Link>
          <UserRoleWorkspaceDebug />
          <SessionDebug />
        </>
      )}
    </PageSection>
  );
}
